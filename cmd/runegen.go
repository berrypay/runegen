/*
Runegen generates a random string of length.
When called without any parameters, it will generate a 6 characters long string using standard alphanumeric characters.

Usage:

	runegen [-c integer | -r string] [flags]

Flags:

	-c, --charset int8         rune set to use represented by binary representation of the set [bbbbb] where:
	                               0b00001 => numbers
	                               0b00010 => small letters
	                               0b00100 => capital letters
	                               0b01000 => symbols ._#@%&!~|$^*=
	                               0b10000 => use runes given by parameter "s"
	                           e.g -c 15 = 0b1111 => full case alphanumeric and symbol. (default 7)
	-r, --customrunes string   Custom Rune Set to use. Note: use of -c and -r are mutually exclusive (default "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	-h, --help                 help for runegen
	-l, --length uint          Length of the generated random string (default 6)
	-a, --startalpha           Start Alpha: make sure generated string start with alphabet
	                           This flag will be ignored if only symbol and/or numeric runes are selected
	-s, --strict               Strict: make sure all rune set selected exist at least once
	-v, --version              version for runegen

Copyright © 2022 Sallehuddin Abdul Latif sallehuddin@berrypay.com
*/
package main

import cmd "github.com/berrypay/runegen/cmd/root"

func main() {
	cmd.Execute()
}
