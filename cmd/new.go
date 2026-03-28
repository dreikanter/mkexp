package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/dreikanter/mkexp/internal/experiment"
	"github.com/dreikanter/mkexp/internal/slug"
	"github.com/spf13/cobra"
)

var (
	flagReadme bool
	flagGit    bool
)

var newCmd = &cobra.Command{
	Use:   "new [text...]",
	Short: "Create a new experiment directory",
	Long: `Create a new date-prefixed experiment directory.

If text is provided, it is slugified to form the directory name.
If no text is provided, a random adjective-noun name is generated.

The directory is created under $MKEXP_PATH (default: current directory).
The absolute path of the new directory is printed to stdout.`,
	RunE: runNew,
}

func init() {
	newCmd.Flags().BoolVar(&flagReadme, "readme", false, "create a README.md in the new directory")
	newCmd.Flags().BoolVar(&flagGit, "git", false, "run git init in the new directory")
	rootCmd.AddCommand(newCmd)
}

func runNew(cmd *cobra.Command, args []string) error {
	var name string
	if len(args) > 0 {
		name = slug.Slugify(strings.Join(args, " "))
		if name == "" {
			return fmt.Errorf("could not generate a valid slug from: %q", strings.Join(args, " "))
		}
	}

	path, err := experiment.Create(experiment.Options{
		Name:   name,
		Path:   os.Getenv("MKEXP_PATH"),
		Readme: flagReadme,
		Git:    flagGit,
	})
	if err != nil {
		return err
	}

	fmt.Fprintln(cmd.OutOrStdout(), path)
	return nil
}
