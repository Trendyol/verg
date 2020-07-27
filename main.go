package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"runtime"
	"strings"
)

func main() {

	var semantic *Semantic
	var major, minor, patch, release, beta, alpha bool

	var (
		version = "dev"
		commit  = "none"
		date    = "unknown"
		builtBy = "unknown"
	)

	var info = fmt.Sprintf(
		"verg %s (%s, %s, %s) on %s (%s)",
		version,
		builtBy,
		date,
		commit,
		runtime.GOOS,
		runtime.GOARCH,
	)

	var cmd = &cobra.Command{
		Use:  "verg",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			s, err := New(args[0])

			if err != nil {
				log.Fatal(err)
				return
			}

			semantic = s

			semantic.Inc(major, minor, patch, release, beta, alpha)
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Println(semantic.String())
		},
	}

	var compareCmd = &cobra.Command{
		Use:     "compare",
		Example: "1.0.0 [< or <= or > or >= or ==] 2.0.0",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			items := strings.SplitN(args[0], " ", 3)

			if len(items) != 3 {
				log.Fatal("Command is not valid argumant. Ex: 1.0.0 [< or <= or > or >= or ==] 2.0.0")
				return
			}

			result, err := Compare(items[0], items[1], items[2])

			if err != nil {
				log.Fatal(err)
				return
			}

			fmt.Println(result)
		},
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of verg",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(info)
		},
	}

	cmd.Flags().BoolVarP(&major, "major", "x", false, "increment major version")
	cmd.Flags().BoolVarP(&minor, "minor", "y", false, "increment minor version")
	cmd.Flags().BoolVarP(&patch, "patch", "z", false, "increment patch version")
	cmd.Flags().BoolVarP(&release, "release", "r", false, "increment release version")
	cmd.Flags().BoolVarP(&beta, "beta", "b", false, "increment beta version")
	cmd.Flags().BoolVarP(&alpha, "alpha", "a", false, "increment alpha version")

	cmd.AddCommand(compareCmd, versionCmd)

	_ = cmd.Execute()
}
