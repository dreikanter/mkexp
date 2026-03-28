package experiment_test

import (
	"os"
	"path/filepath"
	"testing"
)

// TestMain resolves TMPDIR to its real path to handle macOS symlinks
// (/var -> /private/var) so that t.TempDir() and os.Getwd() return
// consistent paths in tests.
func TestMain(m *testing.M) {
	if tmpdir := os.Getenv("TMPDIR"); tmpdir != "" {
		if real, err := filepath.EvalSymlinks(tmpdir); err == nil {
			os.Setenv("TMPDIR", real)
		}
	}
	os.Exit(m.Run())
}
