package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/StephenButtolph/canoto/canoto/generate"
)

var errWrongArgCount = errors.New("expected 1 argument")

func init() {
	cobra.EnablePrefixMatching = true
}

func main() {
	cmd := &cobra.Command{
		Use:   "canoto",
		Short: "Generates a canoto file",
		RunE: func(_ *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("%w, got %d", errWrongArgCount, len(args))
			}
			return generate.File(args[0])
		},
	}

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "command failed %v\n", err)
		os.Exit(1)
	}
}
