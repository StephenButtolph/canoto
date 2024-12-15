// Canoto is command to generate code for reading and writing the canoto format.
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/StephenButtolph/canoto/generate"
)

const (
	canoto = "canoto"
	proto  = "proto"
)

func init() {
	cobra.EnablePrefixMatching = true
}

func main() {
	cmd := &cobra.Command{
		Use:   "canoto",
		Short: "Processes the provided files and generates the corresponding canoto and proto files",
		RunE: func(c *cobra.Command, args []string) error {
			flags := c.Flags()
			canoto, err := flags.GetBool(canoto)
			if err != nil {
				return fmt.Errorf("failed to get canoto flag: %w", err)
			}
			proto, err := flags.GetBool(proto)
			if err != nil {
				return fmt.Errorf("failed to get proto flag: %w", err)
			}

			for _, arg := range args {
				if canoto {
					if err := generate.Canoto(arg); err != nil {
						return fmt.Errorf("failed to generate canoto for %q: %w", arg, err)
					}
				}
				if proto {
					if err := generate.Proto(arg); err != nil {
						return fmt.Errorf("failed to generate proto for %q: %w", arg, err)
					}
				}
			}
			return nil
		},
	}

	flags := cmd.Flags()
	flags.Bool(canoto, true, "Generate canoto file")
	flags.Bool(proto, false, "Generate proto file")

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "command failed %v\n", err)
		os.Exit(1)
	}
}
