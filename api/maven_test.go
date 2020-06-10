package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thepwagner/jithub/api"
)

func TestMavenPackageFromURL(t *testing.T) {
	cases := map[string]*api.MavenPackage{
		"/":           nil,
		"/robots.txt": nil,
		"/com/github/thepwagner/test-java/1.0/test-java-1.0.pom": {
			GroupID:    "com.github.thepwagner",
			ArtifactID: "test-java",
			Version:    "1.0",
			Type:       "pom",
		},
		"/com/github/thepwagner/test-java/1.0/test-java-1.0.jar": {
			GroupID:    "com.github.thepwagner",
			ArtifactID: "test-java",
			Version:    "1.0",
			Type:       "jar",
		},
		"/com/github/thepwagner/test-java/1.0/unexpected-1.0.jar": nil,
	}

	for url, expected := range cases {
		t.Run(url, func(t *testing.T) {
			actual := api.MavenPackageFromURL(url)
			assert.Equal(t, expected, actual)
		})
	}
}
