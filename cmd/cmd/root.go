/*
Copyright © 2022 Sallehuddin Abdul Latif sallehuddin@berrypay.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "runegen [-c integer | -c 16 -s string] [-n integer]",
	Short: "Generates random string of characters. Default to 6 characters alphanumeric string",
	Long: `runegen [-c integer] [-n integer]
Generates random string of characters based on the given options as parameters.

Optional parameters:
  -c #: rune set to use represented by binary representation of the set [bbbbb] where
        0b00001 => numbers
        0b00010 => small letters
        0b00100 => capital letters
        0b01000 => symbols ._#@%&!~|$^*=
				0b10000 => use runes given by parameter "s"
        e.g -c 15 = 0b1111 => full case alphanumeric and symbol. Default to 7

  -n #: an integer denoting the length of random string to be generated. Default to 6

  -s string: user defined runes to use e.g "123abc&$"`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cmd.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
