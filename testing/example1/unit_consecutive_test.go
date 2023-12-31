package example1

import "testing"

func TestConsecutive(t *testing.T) {
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
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.name)
		})
	}
}
