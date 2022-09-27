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

var CharacterSetNumber int8
var CustomRunes string
var Length uint
var Strict bool
var StartAlpha bool
var Debug bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "runegen [-c integer | -r string]",
	Short: "Runegen generates random string of characters",
	Long:  "Runegen generates random string of characters based on the given options as parameters. Default to 6 characters alphanumeric string.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Args:    cobra.MatchAll(cobra.NoArgs, cobra.OnlyValidArgs),
	Version: "1.1.1",
	Run: func(cmd *cobra.Command, args []string) {
		switch CharacterSetNumber {
		case 1:
			fmt.Println(runegen.GetPolicyRandom(true, false, false, false, Strict, false, Length))
		case 2:
			fmt.Println(runegen.GetPolicyRandom(false, true, false, false, Strict, StartAlpha, Length))
		case 3:
			fmt.Println(runegen.GetPolicyRandom(true, true, false, false, Strict, StartAlpha, Length))
		case 4:
			fmt.Println(runegen.GetPolicyRandom(false, false, true, false, Strict, StartAlpha, Length))
		case 5:
			fmt.Println(runegen.GetPolicyRandom(true, false, true, false, Strict, StartAlpha, Length))
		case 6:
			fmt.Println(runegen.GetPolicyRandom(false, true, true, false, Strict, StartAlpha, Length))
		case 7:
			fmt.Println(runegen.GetPolicyRandom(true, true, true, false, Strict, StartAlpha, Length))
		case 8:
			fmt.Println(runegen.GetPolicyRandom(false, false, false, true, Strict, false, Length))
		case 9:
			fmt.Println(runegen.GetPolicyRandom(true, false, false, true, Strict, false, Length))
		case 10:
			fmt.Println(runegen.GetPolicyRandom(false, true, false, true, Strict, StartAlpha, Length))
		case 11:
			fmt.Println(runegen.GetPolicyRandom(true, true, false, true, Strict, StartAlpha, Length))
		case 12:
			fmt.Println(runegen.GetPolicyRandom(false, false, true, true, Strict, StartAlpha, Length))
		case 13:
			fmt.Println(runegen.GetPolicyRandom(true, false, true, true, Strict, StartAlpha, Length))
		case 14:
			fmt.Println(runegen.GetPolicyRandom(false, true, true, true, Strict, StartAlpha, Length))
		case 15:
			fmt.Println(runegen.GetPolicyRandom(true, true, true, true, Strict, StartAlpha, Length))
		case 16:
			fmt.Println(runegen.GetCustomRandom(CustomRunes, Length))
		default:
			fmt.Println(runegen.GetPolicyRandom(true, true, true, false, true, true, Length))
		}
	},
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

	rootCmd.Flags().StringVarP(&CustomRunes, "customrunes", "r",
		runegen.BigCaps+runegen.SmallCaps+runegen.Numerics, "Custom Rune Set to use. Note: use of -c and -r are mutually exclusive")

	rootCmd.Flags().UintVarP(&Length, "length", "l",
		6, "Length of the generated random string")

	rootCmd.Flags().BoolVarP(&Strict, "strict", "s",
		false, "Strict: make sure all rune set selected exist at least once")

	rootCmd.Flags().BoolVarP(&StartAlpha, "startalpha", "a",
		false, `Start Alpha: make sure generated string start with alphabet
This flag will be ignored if only symbol and/or numeric runes are selected`)

	rootCmd.MarkFlagsMutuallyExclusive("charset", "customrunes")

}
