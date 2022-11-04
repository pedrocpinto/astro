package constants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	testCases := []struct {
		name         string
		input        string
		expectedName string
		expectedErr  error
	}{
		{
			name:         "When 'light year' is provided, the light year struct is returned",
			input:        LighYear,
			expectedName: "Light Year",
			expectedErr:  nil,
		},
	}

	for _, tc := range testCases {
		constant, err := GetConstant(tc.input)
		assert.Equal(t, tc.expectedName, constant.name)
		assert.Nil(t, err)
	}
}
