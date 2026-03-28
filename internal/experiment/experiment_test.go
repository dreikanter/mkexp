package experiment_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/dreikanter/mkexp/internal/experiment"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreate_BasicDir(t *testing.T) {
	base := t.TempDir()
	path, err := experiment.Create(experiment.Options{
		Name: "test-project",
		Path: base,
	})
	require.NoError(t, err)

	today := time.Now().Format("20060102")
	expected := filepath.Join(base, today+"_test-project")
	assert.Equal(t, expected, path)
	assert.DirExists(t, path)
}

func TestCreate_RandomName(t *testing.T) {
	base := t.TempDir()
	path, err := experiment.Create(experiment.Options{
		Path: base,
	})
	require.NoError(t, err)
	require.NotEmpty(t, path)

	today := time.Now().Format("20060102")
	assert.True(t, strings.HasPrefix(filepath.Base(path), today+"_"),
		"dir name should start with today's date")
	assert.DirExists(t, path)
}

func TestCreate_WithReadme(t *testing.T) {
	base := t.TempDir()
	path, err := experiment.Create(experiment.Options{
		Name:   "readme-test",
		Path:   base,
		Readme: true,
	})
	require.NoError(t, err)
	assert.FileExists(t, filepath.Join(path, "README.md"))
}

func TestCreate_WithGit(t *testing.T) {
	base := t.TempDir()
	path, err := experiment.Create(experiment.Options{
		Name: "git-test",
		Path: base,
		Git:  true,
	})
	require.NoError(t, err)
	assert.DirExists(t, filepath.Join(path, ".git"))
}

func TestCreate_DefaultPath(t *testing.T) {
	base := t.TempDir()
	orig, err := os.Getwd()
	require.NoError(t, err)
	require.NoError(t, os.Chdir(base))
	defer func() { _ = os.Chdir(orig) }()

	path, err := experiment.Create(experiment.Options{Name: "cwd-test"})
	require.NoError(t, err)
	assert.True(t, strings.HasPrefix(path, base))
	assert.DirExists(t, path)
}
