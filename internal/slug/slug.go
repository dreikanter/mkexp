package slug

import (
	"math/rand"
	"regexp"
	"strings"
)

var nonAlphanumeric = regexp.MustCompile(`[^a-z0-9]+`)

// Slugify converts arbitrary text into a URL-safe slug.
// Returns empty string if the input produces no valid characters.
func Slugify(text string) string {
	s := strings.ToLower(text)
	s = nonAlphanumeric.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}

// Random returns a random adjective-noun slug, e.g. "silent-river".
func Random() string {
	adj := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]
	return adj + "-" + noun
}
