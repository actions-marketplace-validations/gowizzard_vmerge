package convert_test

import (
	"github.com/gowizzard/vmerge/convert"
	"reflect"
	"testing"
)

// TestInteger is to test the convert integer function
// We convert a number it to an integer and check the result
func TestInteger(t *testing.T) {

	tests := []struct {
		number   string
		expected int
	}{
		{
			number:   "5",
			expected: 5,
		},
		{
			number:   "265",
			expected: 265,
		},
		{
			number:   "78",
			expected: 78,
		},
		{
			number:   "132",
			expected: 132,
		},
		{
			number:   "25",
			expected: 25,
		},
	}

	for _, value := range tests {

		integer := convert.Integer(value.number)

		if !reflect.DeepEqual(value.expected, integer) {
			t.Fatalf("expected: %d, got %d", value.expected, integer)
		}

	}

}