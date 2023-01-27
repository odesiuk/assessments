package treasures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindLocation(t *testing.T) {
	tests := map[string]struct {
		x, y   int
		exp    string
		expErr bool
	}{
		"any":    {x: 2, y: 5, exp: "20"},
		"(2,3)":  {x: 2, y: 3, exp: "9"},
		"(3,2)":  {x: 3, y: 2, exp: "8"},
		"(10,5)": {x: 10, y: 5, exp: "96"},
		"error":  {x: 100001, y: 5, expErr: true},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := FindLocation(tt.x, tt.y)

			assert.True(t, (err != nil) == tt.expErr)
			assert.Equal(t, tt.exp, got)
		})
	}
}
