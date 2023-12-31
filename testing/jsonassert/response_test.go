package jsonassert

import "testing"

func TestResponse(t *testing.T) {
	tests := []struct {
		name             string
		actualResponse   string
		expectedResponse string
	}{
		{
			name: "Successful comparison 1",
			actualResponse: `[
				{"username": "foo", "createdAt": "2023-12-31T14:30:17+00:00"},
				{"username": "bar", "createdAt": "2023-12-31T15:30:17+00:00"}
			]`,
			expectedResponse: `[
				{"username": "foo", "createdAt": "{{exists}}"},
				{"username": "bar", "createdAt": "{{exists}}"}
			]`,
		},
		{
			name:             "Failed comparison",
			actualResponse:   `{"username": "foo"}`,
			expectedResponse: `{"username": "foo", "createdAt": "{{exists}}"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertEqual(t, tt.expectedResponse, tt.actualResponse)
		})
	}
}
