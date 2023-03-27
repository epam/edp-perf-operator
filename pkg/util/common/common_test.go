package common

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvertToStringArray(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input interface{}
		want  []string
	}{
		{
			name:  "should return slice with the same values",
			input: []string{"a", "b", "c"},
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "should return slice with string values if it asserts successfully, and empty string otherwise",
			input: []any{"a", 5, 12.4, "b", "c"},
			want:  []string{"a", "", "", "b", "c"},
		},
		{
			name:  "should fail due to type assertion error and return empty slice",
			input: []int{5, 4},
			want:  []string{},
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, ConvertToStringArray(tt.input))
		})
	}
}
