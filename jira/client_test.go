package jira

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasicAuthToken(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		want := "YUBiLmM6YWJj"

		c := &Client{
			Email: "a@b.c",
			Token: "abc",
		}

		res := c.basicAuthToken()
		require.Equal(t, want, res)
	})
}

func TestBaseURL(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		want := "https://a.atlassian.net/rest/api/3"
		tenant := "a"

		res := baseURL(tenant)
		require.Equal(t, want, res)
	})
}
