package example1

import "testing"

func TestParallel(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "foo",
		},
		{
			name: "bar",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			t.Log(tt.name)
		})
	}
}
