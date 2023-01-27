package solution

import (
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Predefined errors. DO NOT CHANGE.
var (
	ErrDocumentAlreadyExists = errors.New("error document already exists")
	ErrDocumentNotFound      = errors.New("error document not found")
)

const (
	AthletesCollection   = "athletes"
	ActivitiesCollection = "activities"
)

// AthletesStore which stores athletes and their activities.
type AthletesStore struct {
	client *firestore.Client
}

// NewAthletesStore instantiates a new AthletesStore.
func NewAthletesStore(client *firestore.Client) *AthletesStore {
	return &AthletesStore{client: client}
}

// GetAthleteByDocID get a single athlete by `DocID`; return a `ErrDocumentNotFound` error
func (ats *AthletesStore) GetAthleteByDocID(ctx context.Context, docID string) (map[string]interface{}, error) {
	doc, err := ats.client.Collection(AthletesCollection).Doc(docID).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, ErrDocumentNotFound
		}

		return nil, fmt.Errorf("get error: %w", err)
	}

	return doc.Data(), nil
}

// CreateAthlete Create a new athlete with the specific document `id` (which should be equal to `DocID`)
// if a document with the specified identifier already exists, return a `ErrDocumentAlreadyExists` error.
func (ats *AthletesStore) CreateAthlete(ctx context.Context, docID string, firstName, lastName, country string) error {
	doc, err := ats.client.Collection(AthletesCollection).Doc(docID).Get(ctx)
	if err == nil && doc.Exists() {
		return ErrDocumentAlreadyExists
	}

	if status.Code(err) != codes.NotFound {
		return fmt.Errorf("get error: %w", err)
	}

	_, err = ats.client.Doc(AthletesCollection+"/"+docID).Create(ctx, map[string]interface{}{
		"first_name": firstName,
		"last_name":  lastName,
		"country":    country,
	})

	if err != nil {
		return fmt.Errorf("firestore Doc Create error:%w", err)
	}

	return nil
}

// GetAthletesByCountry get all athletes in the specified `country`.
func (ats *AthletesStore) GetAthletesByCountry(ctx context.Context, country string) ([]map[string]interface{}, error) {

	iter := ats.client.Collection(AthletesCollection).Where("country", "==", country).Documents(ctx)
	defer iter.Stop()

	list, err := iter.GetAll()
	if err != nil {
		return nil, fmt.Errorf("get all error: %w", err)
	}

	if len(list) == 0 {
		return nil, ErrDocumentNotFound
	}

	result := make([]map[string]interface{}, 0)

	for _, doc := range list {
		result = append(result, doc.Data())
	}

	return result, nil
}

// DeleteActivitiesByType delete all activities with the specified activity `type`.
func (ats *AthletesStore) DeleteActivitiesByType(ctx context.Context, activityType string) error {
	iter := ats.client.Collection(ActivitiesCollection).Where("type", "==", activityType).Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}

			return fmt.Errorf("get next error: %w", err)
		}

		_, err = doc.Ref.Delete(ctx)
		if err != nil {
			return fmt.Errorf("delete %s error: %w", doc.Ref.ID, err)
		}
	}

	return nil
}

// UpdateActivityMap update an activity's `processedCoordinates` field with every second coordinate from the activity's `coordinates` field.
// Treat `coordinates` field as type `[]interface{}`.
//
func (ats *AthletesStore) UpdateActivityMap(ctx context.Context, docID string) error {
	doc, err := ats.client.Collection(ActivitiesCollection).Doc(docID).Get(ctx)
	if err != nil {
		return fmt.Errorf("get error: %w", err)
	}

	processedCoordinates := make([]interface{}, 0)

	coordinates := doc.Data()["coordinates"].([]interface{})
	for i, coordinate := range coordinates {
		if i%2 == 0 {
			processedCoordinates = append(processedCoordinates, coordinate)
		}
	}

	_, err = doc.Ref.Update(ctx, []firestore.Update{
		{
			Path:  "processed_coordinates",
			Value: processedCoordinates,
		},
	})
	if err != nil {
		return fmt.Errorf("update error: %w", err)
	}

	return nil
}
