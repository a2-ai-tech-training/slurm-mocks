package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `Get Stdout and Stderr and ExitCode of a scenario:
For example:\ngetCmd get -o ".../user-homes/<user>/<output>" -s "test1" -- sacct --json`,

	Run: cmd,
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringP("scenario", "s", "default", "$SCENARIO")
	getCmd.Flags().StringP("output", "o", "outputs", "$OUTPUT_DIR")
}

func cmd(cmd *cobra.Command, args []string) {
	// TODO: Work your own magic here
	scenario, _ := cmd.Flags().GetString("scenario")
	output, _ := cmd.Flags().GetString("output")
	fmt.Println("get called")

	if scenario == "default" {
		fmt.Println("scenario is default")
	}

	if output == "outputs" {
		fmt.Println("output is outputs")
	}
}
