/*
Copyright © 2022 Sallehuddin Abdul Latif sallehuddin@berrypay.com

*/

// Package runegen provides libraries for generating random string from specified rune/runes.
package runegen

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

var SmallCaps = "abcdefghijklmnopqrstuvwxyz"
var BigCaps = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var Numerics = "0123456789"
var Symbols = "._#@%&!~|$^*="

// GetRandom returns a random string consists of runes from predefined character set/sets.
// Rune set to use are represented by binary representation of the set [bbbbb] where:
//
//	0b00001 => numbers
//	0b00010 => small letters
//	0b00100 => capital letters
//	0b01000 => symbols ._#@%&!~|$^*=
//
// It will use numbers, small letters and capital letters if the rune set specified not matched.
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

// GetPolicyRandom returns a random string of length n consists of runes from predefined or custom character set/sets
// complying with the complexity requirements.
//
//	num = true => use numeric set
//	sl = true => use small letters set
//	bl = true => use capital letters set
//	sym = true => use symbols set
//	strict = true => make sure all rune set selected exist in the resulting string
//	startAlpha = true => make sure the resulting string start with an alphabet character
func GetPolicyRandom(num bool, sl bool, bl bool, sym bool, strict bool, startAlpha bool, n uint) string {
	policyAlpha := startAlpha
	policyStrict := strict

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

	var alphaFirst bool

	if !sl && !bl {
		policyAlpha = false
	}

	var generated string

	// No strict checking requested
	if !policyStrict {
		if policyAlpha {
			for !alphaFirst {
				generated = generateRandom(curRune, n)
				alphaFirst = startedWithAlpha(generated)
				//fmt.Printf("Generated: %v and start with alpha: %v\n", generated, alphaFirst)
			}
			return generated
		}

		return generateRandom(curRune, n)
	}

	// We will reach this part when strict checking was requested
	hasNum, hasSL, hasBL, hasSym := false, false, false, false

	if policyAlpha {
		for !hasNum || !hasSL || !hasBL || !hasSym || !alphaFirst {
			generated = generateRandom(curRune, n)

			hasNum = numericExist(generated)
			//overwrite if numeric is not in the runes selection
			if !num {
				hasNum = true
			}

			hasSL = smallLetterExist(generated)
			//overwrite if small letters is not in the runes selection
			if !sl {
				hasSL = true
			}

			hasBL = capitalLetterExist(generated)
			//overwrite if small letters is not in the runes selection
			if !bl {
				hasBL = true
			}

			hasSym = symbolExist(generated)
			//overwrite if small letters is not in the runes selection
			if !sym {
				hasSym = true
			}

			alphaFirst = startedWithAlpha(generated)
			//overwrite if small letters or capital letters are not in the runes selection
			if !sl && !bl {
				alphaFirst = true
			}

			//fmt.Printf("Generated: %v and policy result: %v|%v|%v|%v|%v\n", generated, hasNum, hasSL, hasBL, hasSym, alphaFirst)
		}
	} else {
		for !hasNum || !hasSL || !hasBL || !hasSym {
			generated = generateRandom(curRune, n)

			hasNum = numericExist(generated)
			//overwrite if numeric is not in the runes selection
			if !num {
				hasNum = true
			}

			hasSL = smallLetterExist(generated)
			//overwrite if small letters is not in the runes selection
			if !sl {
				hasSL = true
			}

			hasBL = capitalLetterExist(generated)
			//overwrite if small letters is not in the runes selection
			if !bl {
				hasBL = true
			}

			hasSym = symbolExist(generated)
			//overwrite if small letters is not in the runes selection
			if !sym {
				hasSym = true
			}

			//fmt.Printf("Generated: %v and policy result: %v|%v|%v|%v\n", generated, hasNum, hasSL, hasBL, hasSym)
		}
	}

	return generated
}

// GetCustomRandom returns a random string of length n consists of runes from the specified string of characters s.
func GetCustomRandom(s string, n uint) string {
	return generateRandom([]rune(s), n)
}

// generateRandom returns a random string of length n consist of runes from the given set.
func generateRandom(r []rune, n uint) string {
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = r[randomizer.Intn(len(r))]
	}
	return string(b)
}

// numericExist returns a boolean indicating whether at least single numeric character exist in the string s
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

// smallLetterExist returns a boolean indicating whether at least single small letter character exist in the string s
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

// capitalLetterExist returns a boolean indicating whether at least single capital letter character exist in the string s
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

// symbolExist returns a boolean indicating whether at least single symbol character exist in the string s
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

// startedWithAlpha returns a boolean indicating whether the string s starts with an alphabet character
func startedWithAlpha(s string) bool {
	return strings.Contains(SmallCaps+BigCaps, s[0:1])
}

// Nonce Generator
// Generate random string Nonce of length n using crypto/rand
func GetNonceStr(n uint) string {
	runes := []rune(SmallCaps + BigCaps + Numerics)

	b := make([]rune, n)
	for i := range b {
		random, err := crand.Int(crand.Reader, big.NewInt(int64(len(runes))))
		if err != nil {
			//fallback to use math/rand
			randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
			random = big.NewInt(int64(randomizer.Intn(len(runes))))
		}
		b[i] = runes[random.Int64()]
	}
	return string(b)
}
