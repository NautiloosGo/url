package app

import (
	"strings"
	"testing"
)

func TestGetRandomString(t *testing.T) {
	for n := 5; n <= 10000; n++ { // string length = n
		//Arrange
		list := []string{
			"1234567890",
			"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			"abcdefghijklmnopqrstuvwxyz",
			`!"#$%&'()*+,-./:;<=>?@"`,
			`[\]^_`,
			`{|}~`}
		stringlen := n
		// Act
		for _, lib := range list {
			answer := GetRandomString(n, lib)
			// Assert
			// check length
			if answerQty := len(answer); answerQty != stringlen {
				t.Fatalf("%q:\nWrong len, want %d, get %d", t.Name(), stringlen, answerQty)
			}
			// check answer srting for symbols not from list
			for _, s := range answer {
				if index := strings.Index(lib, string(s)); index >= 0 {

				} else {
					t.Fatalf("%q:\nsent: %s\nwith len = %d\nIllegal symbol: %s in %s", t.Name(), lib, n, string(s), answer)
				}
			}
		}
	}
}
