// Canoto is command to generate code for reading and writing the canoto format.
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/StephenButtolph/canoto"
	"github.com/StephenButtolph/canoto/generate"
)

const (
	canotoFlag   = "canoto"
	libraryFlag  = "library"
	protoFlag    = "proto"
	versionFlag  = "version"
	importFlag   = "import"
	internalFlag = "internal"
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
			showVersion, err := flags.GetBool(versionFlag)
			if err != nil {
				return fmt.Errorf("failed to get version flag: %w", err)
			}
			if showVersion {
				fmt.Println("canoto/" + canoto.Version)
				return nil
			}

			library, err := flags.GetString(libraryFlag)
			if err != nil {
				return fmt.Errorf("failed to get library flag: %w", err)
			}
			if library != "" {
				if err := generate.Library(library); err != nil {
					return fmt.Errorf("failed to generate library in %q: %w", library, err)
				}
			}

			canoto, err := flags.GetBool(canotoFlag)
			if err != nil {
				return fmt.Errorf("failed to get canoto flag: %w", err)
			}
			proto, err := flags.GetBool(protoFlag)
			if err != nil {
				return fmt.Errorf("failed to get proto flag: %w", err)
			}

			canotoImport, err := flags.GetString(importFlag)
			if err != nil {
				return fmt.Errorf("failed to get import flag: %w", err)
			}
			canotoImport = `"` + canotoImport + `"`
			internal, err := flags.GetBool(internalFlag)
			if err != nil {
				return fmt.Errorf("failed to get internal flag: %w", err)
			}

			for _, arg := range args {
				if canoto {
					if err := generate.Canoto(arg, canotoImport, internal); err != nil {
						return fmt.Errorf("failed to generate canoto for %q: %w", arg, err)
					}
				}
				if proto {
					if err := generate.Proto(arg, canotoImport, internal); err != nil {
						return fmt.Errorf("failed to generate proto for %q: %w", arg, err)
					}
				}
			}
			return nil
		},
	}

	flags := cmd.Flags()
	flags.Bool(versionFlag, false, "Display the version and exit")
	flags.Bool(canotoFlag, true, "Generate canoto file")
	flags.String(libraryFlag, "", "Generate the canoto library in the specified directory")
	flags.Bool(protoFlag, false, "Generate proto file")
	flags.String(importFlag, "github.com/StephenButtolph/canoto", "Package to depend on for canoto serialization primitives")
	flags.Bool(internalFlag, false, "Generate a file that assumes the canoto functional does not need to be imported")

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "command failed %v\n", err)
		os.Exit(1)
	}
}
