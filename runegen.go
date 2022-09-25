/*
Copyright © 2022 Sallehuddin Abdul Latif sallehuddin@berrypay.com

*/

package runegen

import (
	"math/rand"
	"time"
)

var smallCaps = "abcdefghijklmnopqrstuvwxyz"
var bigCaps = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var numeric = "0123456789"
var symbol = "._#@%&!~|$^*="

func GetRandom(charSet uint8, n int) string {
	switch charSet {
	case 1:
		return generateRandom([]rune(numeric), n)
	case 2:
		return generateRandom([]rune(smallCaps), n)
	case 3:
		return generateRandom([]rune(smallCaps+numeric), n)
	case 4:
		return generateRandom([]rune(bigCaps), n)
	case 5:
		return generateRandom([]rune(bigCaps+numeric), n)
	case 6:
		return generateRandom([]rune(bigCaps+smallCaps), n)
	case 7:
		return generateRandom([]rune(bigCaps+smallCaps+numeric), n)
	case 8:
		return generateRandom([]rune(symbol), n)
	case 10:
		return generateRandom([]rune(smallCaps+symbol), n)
	case 11:
		return generateRandom([]rune(bigCaps+smallCaps+symbol), n)
	case 12:
		return generateRandom([]rune(bigCaps+symbol), n)
	case 13:
		return generateRandom([]rune(bigCaps+numeric+symbol), n)
	case 14:
		return generateRandom([]rune(bigCaps+smallCaps+symbol), n)
	case 15:
		return generateRandom([]rune(bigCaps+smallCaps+numeric+symbol), n)
	default:
		return generateRandom([]rune(bigCaps+smallCaps+numeric), n)
	}
}

func GetCustomRandom(s string, n int) string {
	return generateRandom([]rune(s), n)
}

func generateRandom(r []rune, n int) string {
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = r[randomizer.Intn(len(r))]
	}
	return string(b)
}
