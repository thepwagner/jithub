package gh_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thepwagner/jithub/gh"
)

func TestClient_PackageFiles(t *testing.T) {
	ctx, c := testClient(t)

	f, err := c.PackageFiles(ctx, "thepwagner", "test-java", "1.0")
	require.NoError(t, err)
	assert.Len(t, f, 6)
}

func TestClient_PackageFiles_VersionNotFound(t *testing.T) {
	ctx, c := testClient(t)

	f, err := c.PackageFiles(ctx, "thepwagner", "test-java", "eleventy.seven")
	require.NoError(t, err)
	assert.Empty(t, f)
}

func TestClient_PackageFiles_RepoNotFound(t *testing.T) {
	ctx, c := testClient(t)

	_, err := c.PackageFiles(ctx, "batman", "batcave-security-system", "1.0")
	assert.EqualError(t, err, "Could not resolve to a Repository with the name 'batcave-security-system'.")
}

func TestClient_TriggerPackageBuild(t *testing.T) {
	ctx, c := testClient(t)

	err := c.TriggerPackageBuild(ctx, "thepwagner", "test-java", "1f7d259eac8b700c24728e822279030b3b0748b1")
	require.NoError(t, err)
}

func testClient(t *testing.T) (context.Context, *gh.Client) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		t.Skip("GITHUB_TOKEN")
	}
	ctx := context.Background()
	c := gh.NewStaticTokenClient(ctx, token)
	return ctx, c
}
