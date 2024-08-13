/*
Copyright © 2022 Sallehuddin Abdul Latif sallehuddin@berrypay.com

*/

package runegen

import (
	"fmt"
	"testing"
)

func TestGetRandom(t *testing.T) {
	for i := 1; i < 10; i++ {
		got := GetRandom(1, 6)
		if len(got) != 6 {
			t.Error("Not enough length generated, test failed.")
		} else {
			fmt.Println("Random runes generated: " + got)
		}
	}
}

func TestGetNonceStr(t *testing.T) {
	for i := 1; i < 10; i++ {
		got := GetNonceStr(64)
		if len(got) < 64 {
			t.Error("Not enough length generated, test failed.")
		} else {
			fmt.Println("Nonce generated: " + got)
		}
	}
}
