package cmd

import (
	"fmt"
	"os"

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

	fmt.Println("OS scenario :", os.Getenv("SCENARIO"))
	fmt.Println("OS output   :", os.Getenv("OUTPUT_DIR"))

	scenario, _ := cmd.Flags().GetString("scenario")
	output, _ := cmd.Flags().GetString("output")
	//fmt.Println("get called")
	if scenario != "" {
		fmt.Println("Flag scenario is", scenario)
	}
	if os.Getenv("SCENARIO") != "" {
		if scenario == "default" {
			scenario = os.Getenv("SCENARIO")
			fmt.Println("Scenario is:", scenario)
		}
	}

	if output != "" {
		fmt.Println("Flag output dir is", output)
	}
	if os.Getenv("OUTPUT_DIR") != "" {
		if output == "outputs" {
			output = os.Getenv("OUTPUT_DIR")
			fmt.Println("Output Directory:", output)
		}
	}

	if len(args) > 0 {
		fmt.Println("args are", args)
	} else {
		fmt.Println("no args")
	}
}
