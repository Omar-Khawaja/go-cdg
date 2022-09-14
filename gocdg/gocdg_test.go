package gocdg

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	cases := []struct {
		Name        string
		APIKey      string
		ExpectError bool
		ErrMessage  string
	}{
		{
			Name:        "API key is configured",
			APIKey:      "somevalidapikey",
			ExpectError: false,
		},
		{
			Name:        "API key is not configured",
			APIKey:      "",
			ExpectError: true,
			ErrMessage:  ErrAPIKeyNotConfigured,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			r := require.New(t)
			os.Setenv("CDG_API_KEY", c.APIKey)
			_, err := NewClient()
			if !c.ExpectError {
				r.NoError(err)
				return
			}
			r.Error(err)
			r.Equal(err.Error(), c.ErrMessage)
		})
	}
}
