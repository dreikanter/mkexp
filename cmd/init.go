package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Print shell integration snippet",
	Long:  `Print a shell snippet for .zshrc. Add 'eval "$(mkexp init)"' to your .zshrc to enable the mkx function.`,
	RunE:  runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, _ []string) error {
	snippet := `# mkexp shell integration
if command -v mkexp &>/dev/null; then
  mkx() { cd "$(mkexp new "$@")" }
fi`
	fmt.Fprintln(cmd.OutOrStdout(), snippet)
	return nil
}
