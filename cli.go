package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/spf13/cobra"
)

const long_description = "lcdl is a cli tool to download and execute leetcode problems offline and in your own IDE. use -h for usage help"

var rootCmd = &cobra.Command{
	Use:   "lcdl",
	Short: "lcdl is a cli tool to download leetcode problems.",
	Long:  long_description,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(long_description)
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes lcdl in the current directory",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("LCDL Initialized")
	},
}

var checkAll bool
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Locally tests solutions with base test cases.",
	Long:  "Locally tests solutions with base test cases. Checks latest edited solution by default, [problem numbers] for specific problems and --all to check all.",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if checkAll {
			fmt.Println("Not Implemented Yet")
		} else {
			_, err := ConvertArgsToInts(args)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error converting arguments: %v\n", err)
				os.Exit(1)
			}

		}
	},
}

var getCmd = &cobra.Command{
	Use: "get",
	Short: "Downloads solutions for offline use.",
	Long: "Download solutions for offline use with an arbitrary amount of problems i.e lcdl get 1 2 3",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := ConvertArgsToInts(args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error converting arguments: %v\n", err)
            os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(getCmd)

	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().BoolVarP(&checkAll, "all", "a", false, "Check all solutions")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while executing lcdl: %s \n", err)
	}
}
