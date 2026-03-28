package experiment

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/dreikanter/mkexp/internal/slug"
)

// Options configures experiment directory creation.
type Options struct {
	Name   string // slug for the dir name suffix; if empty, a random name is generated
	Path   string // base directory; if empty, current working directory is used
	Readme bool   // create a README.md in the new directory
	Git    bool   // run git init in the new directory
}

// Create creates a new experiment directory and returns its absolute path.
func Create(opts Options) (string, error) {
	name := opts.Name
	if name == "" {
		name = slug.Random()
	}

	base := opts.Path
	if base == "" {
		var err error
		base, err = os.Getwd()
		if err != nil {
			return "", fmt.Errorf("getting working directory: %w", err)
		}
	}

	datePrefix := time.Now().Format("20060102")
	dirPath, err := filepath.Abs(filepath.Join(base, datePrefix+"_"+name))
	if err != nil {
		return "", fmt.Errorf("resolving path: %w", err)
	}

	if err := os.MkdirAll(dirPath, 0o755); err != nil {
		return "", fmt.Errorf("creating directory: %w", err)
	}

	if opts.Readme {
		if err := writeReadme(dirPath, name, time.Now().Format("2006-01-02")); err != nil {
			return "", err
		}
	}

	if opts.Git {
		if err := gitInit(dirPath); err != nil {
			return "", err
		}
	}

	return dirPath, nil
}

func writeReadme(dirPath, name, date string) error {
	content := fmt.Sprintf("# %s\n\nCreated: %s\n", name, date)
	if err := os.WriteFile(filepath.Join(dirPath, "README.md"), []byte(content), 0o644); err != nil {
		return fmt.Errorf("writing README: %w", err)
	}
	return nil
}

func gitInit(dirPath string) error {
	cmd := exec.Command("git", "init", dirPath)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git init: %w\n%s", err, out)
	}
	return nil
}
