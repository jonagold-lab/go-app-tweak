package apptweak

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAuthClient(t *testing.T) {
	input := "x12345x"
	output := NewAuthClient(input, &http.Client{})
	expectedOutput := &Client{token: input, httpClient: &http.Client{}}
	assert.Equal(t, expectedOutput, output, "Auth Client not working as expected")
}
