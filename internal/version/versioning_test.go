package version_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Hasuzawa/nspredicate-parser/internal/version"
)

func TestVersioning(t *testing.T) {
	for _, subtest := range []struct {
		name       string
		versioning version.Versioning
		expected   string
	}{
		{
			name: "0.0.0",
			versioning: version.Versioning{
				Major: 0,
				Minor: 0,
				Patch: 0,
			},
			expected: "0.0.0",
		},
		{
			name: "1.2.3",
			versioning: version.Versioning{
				Major: 1,
				Minor: 2,
				Patch: 3,
			},
			expected: "1.2.3",
		},
		{
			name: "3.56.3694",
			versioning: version.Versioning{
				Major: 3,
				Minor: 56,
				Patch: 3_694,
			},
			expected: "3.56.3694",
		},
		{
			name: "0.0.0",
			versioning: version.Versioning{
				Major: 0_000,
				Minor: 00_000,
				Patch: 0000_0000,
			},
			expected: "0.0.0",
		},
	} {
		t.Run(subtest.name, func(t *testing.T) {
			assert.Equal(t, subtest.expected, subtest.versioning.String())
		})
	}
}
