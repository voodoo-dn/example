package jsonassert

import (
	"github.com/kinbiko/jsonassert"
	"strings"
	"testing"
)

var replaces = map[string]string{
	"{{exists}}": "<<PRESENCE>>",
}

func AssertEqual(t *testing.T, expected string, actual string) {
	asserter := jsonassert.New(t)

	for r, v := range replaces {
		expected = strings.ReplaceAll(expected, r, v)
	}

	asserter.Assertf(actual, expected)
}
