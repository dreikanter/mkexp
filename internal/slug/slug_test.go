package slug_test

import (
	"strings"
	"testing"

	"github.com/dreikanter/mkexp/internal/slug"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSlugify(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello World", "hello-world"},
		{"  spaces  ", "spaces"},
		{"UPPER CASE", "upper-case"},
		{"foo--bar", "foo-bar"},
		{"foo bar baz", "foo-bar-baz"},
		{"", ""},
		{"!!!###", ""},
		{"foo 123 bar", "foo-123-bar"},
		{"-leading-dash", "leading-dash"},
		{"trailing-dash-", "trailing-dash"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.expected, slug.Slugify(tt.input))
		})
	}
}

func TestRandom(t *testing.T) {
	s := slug.Random()
	require.NotEmpty(t, s)
	parts := strings.Split(s, "-")
	require.Len(t, parts, 2, "random slug should be adjective-noun")
	assert.Regexp(t, `^[a-z]+$`, parts[0], "adjective should be lowercase alpha")
	assert.Regexp(t, `^[a-z]+$`, parts[1], "noun should be lowercase alpha")
}

func TestRandomIsRandom(t *testing.T) {
	seen := make(map[string]bool)
	for i := 0; i < 20; i++ {
		seen[slug.Random()] = true
	}
	assert.Greater(t, len(seen), 1, "20 random slugs should not all be identical")
}
