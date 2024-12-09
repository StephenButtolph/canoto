// Canoto is a code generator that generates Go code for reading and writing the
// canoto format.
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/StephenButtolph/canoto/canoto/internal/generate"
)

func init() {
	cobra.EnablePrefixMatching = true
}

func main() {
	cmd := &cobra.Command{
		Use:   "canoto",
		Short: "Generates the canoto file for all provided files",
		RunE: func(_ *cobra.Command, args []string) error {
			for _, arg := range args {
				if err := generate.File(arg); err != nil {
					return fmt.Errorf("failed to generate %q: %w", arg, err)
				}
			}
			return nil
		},
	}

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "command failed %v\n", err)
		os.Exit(1)
	}
}
