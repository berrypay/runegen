/*
Copyright © 2022 Sallehuddin Abdul Latif sallehuddin@berrypay.com
*/
package root

import (
	"fmt"
	"os"

	"github.com/berrypay/runegen"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "runegen",
	Short: "Runegen generates random string of characters",
	Long:  "Runegen generates random string of characters based on the given options as parameters. Default to 6 characters alphanumeric string.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Version: "1.1.0a",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Strict Alpha:" + runegen.GetPolicyRandom(true, true, true, true, true, true, 16))
		fmt.Println("Strict:" + runegen.GetPolicyRandom(true, true, true, true, true, false, 16))
		fmt.Println("Not Strict Alpha:" + runegen.GetPolicyRandom(true, true, true, true, false, true, 16))
		fmt.Println("Not Strict:" + runegen.GetPolicyRandom(true, true, true, true, false, false, 16))
	},
}

var CharacterSetNumber int8
var CustomRunes string
var Length int64
var Strict bool

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().Int8VarP(
		&CharacterSetNumber,
		"charset",
		"c",
		7,
		`rune set to use represented by binary representation of the set [bbbbb] where:
    0b00001 => numbers
    0b00010 => small letters
    0b00100 => capital letters
    0b01000 => symbols ._#@%&!~|$^*=
    0b10000 => use runes given by parameter "s"
e.g -c 15 = 0b1111 => full case alphanumeric and symbol.`)

	rootCmd.Flags().StringVarP(
		&CustomRunes,
		"customrunes",
		"r",
		runegen.BigCaps+runegen.SmallCaps+runegen.Numerics,
		"Custom Rune Set to use")

	rootCmd.Flags().Int64VarP(&Length, "length", "l", 6, "Length of the generated random string")

	rootCmd.Flags().BoolVarP(&Strict, "strict", "s", false, "Strict: make sure all rune set selected exist at least once")

	rootCmd.MarkFlagsMutuallyExclusive("charset", "customrunes")

}
