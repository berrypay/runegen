/*
Copyright © 2022 Sallehuddin Abdul Latif sallehuddin@berrypay.com

*/

package runegen

import (
	"math/rand"
	"strings"
	"time"
)

var SmallCaps = "abcdefghijklmnopqrstuvwxyz"
var BigCaps = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var Numerics = "0123456789"
var Symbols = "._#@%&!~|$^*="

func GetRandom(charSet uint8, n uint) string {
	switch charSet {
	case 1:
		return generateRandom([]rune(Numerics), n)
	case 2:
		return generateRandom([]rune(SmallCaps), n)
	case 3:
		return generateRandom([]rune(SmallCaps+Numerics), n)
	case 4:
		return generateRandom([]rune(BigCaps), n)
	case 5:
		return generateRandom([]rune(BigCaps+Numerics), n)
	case 6:
		return generateRandom([]rune(BigCaps+SmallCaps), n)
	case 7:
		return generateRandom([]rune(BigCaps+SmallCaps+Numerics), n)
	case 8:
		return generateRandom([]rune(Symbols), n)
	case 10:
		return generateRandom([]rune(SmallCaps+Symbols), n)
	case 11:
		return generateRandom([]rune(BigCaps+SmallCaps+Symbols), n)
	case 12:
		return generateRandom([]rune(BigCaps+Symbols), n)
	case 13:
		return generateRandom([]rune(BigCaps+Numerics+Symbols), n)
	case 14:
		return generateRandom([]rune(BigCaps+SmallCaps+Symbols), n)
	case 15:
		return generateRandom([]rune(BigCaps+SmallCaps+Numerics+Symbols), n)
	default:
		return generateRandom([]rune(BigCaps+SmallCaps+Numerics), n)
	}
}

func GetPolicyRandom(num bool, sl bool, bl bool, sym bool, strict bool, startAlpha bool, n uint) string {
	var curRune []rune
	if num {
		curRune = []rune(string(curRune) + Numerics)
	}

	if sl {
		curRune = []rune(string(curRune) + SmallCaps)
	}

	if bl {
		curRune = []rune(string(curRune) + BigCaps)
	}

	if sym {
		curRune = []rune(string(curRune) + Symbols)
	}

	if !num && !sl && !bl && !sym {
		// we won't allow empty runes... if all not set, then use all runes
		curRune = []rune(Numerics + SmallCaps + BigCaps + Symbols)
	}

	alphaFirst := false
	var generated string

	if !strict {
		if startAlpha {
			for !alphaFirst {
				generated = generateRandom(curRune, n)
				alphaFirst = startedWithAlpha(generated)
			}
			return generated
		}

		return generateRandom(curRune, n)
	}

	hasNum, hasSL, hasBL, hasSym := false, false, false, false

	if startAlpha {
		for !hasNum || !hasSL || !hasBL || !hasSym || !alphaFirst {
			generated = generateRandom(curRune, n)
			hasNum = numericExist(generated)
			hasSL = smallLetterExist(generated)
			hasBL = capitalLetterExist(generated)
			hasSym = symbolExist(generated)
			alphaFirst = startedWithAlpha(generated)
		}
	} else {
		for !hasNum || !hasSL || !hasBL || !hasSym {
			generated = generateRandom(curRune, n)
			hasNum = numericExist(generated)
			hasSL = smallLetterExist(generated)
			hasBL = capitalLetterExist(generated)
			hasSym = symbolExist(generated)
		}
	}

	return generated
}

func GetCustomRandom(s string, n uint) string {
	return generateRandom([]rune(s), n)
}

func generateRandom(r []rune, n uint) string {
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = r[randomizer.Intn(len(r))]
	}
	return string(b)
}

func numericExist(s string) bool {
	isExist := false
	testRunes := []rune(Numerics)
	for i := 0; i < len(testRunes); i++ {
		char := string(testRunes[i])
		if strings.Contains(s, char) {
			isExist = true
		}
	}
	return isExist
}

func smallLetterExist(s string) bool {
	isExist := false
	testRunes := []rune(SmallCaps)
	for i := 0; i < len(testRunes); i++ {
		char := string(testRunes[i])
		if strings.Contains(s, char) {
			isExist = true
		}
	}
	return isExist
}

func capitalLetterExist(s string) bool {
	isExist := false
	testRunes := []rune(BigCaps)
	for i := 0; i < len(testRunes); i++ {
		char := string(testRunes[i])
		if strings.Contains(s, char) {
			isExist = true
		}
	}
	return isExist
}

func symbolExist(s string) bool {
	isExist := false
	testRunes := []rune(Symbols)
	for i := 0; i < len(testRunes); i++ {
		char := string(testRunes[i])
		if strings.Contains(s, char) {
			isExist = true
		}
	}
	return isExist
}

func startedWithAlpha(s string) bool {
	return strings.Contains(SmallCaps+BigCaps, s[0:1])
}
