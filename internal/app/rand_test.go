package app

import (
	"strings"
	"testing"
)

func TestGetRandomString(t *testing.T) {
	//Arrange
	for n := 5; n <= 10000; n++ {
		list := []string{
			"1234567890",
			"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			"abcdefghijklmnopqrstuvwxyz",
			`!"#$%&'()*+,-./:;<=>?@"`,
			`[\]^_`,
			`{|}~`}
		for _, lib := range list {
			// Act
			answer := GetRandomString(n, lib)
			// Assert
			if answerQty := len(answer); answerQty != n {
				t.Fatalf("%q:\nWrong len, want %d, get %d", t.Name(), n, answerQty)
			}
			for _, s := range answer {
				if index := strings.Index(lib, string(s)); index >= 0 {

				} else {
					t.Fatalf("%q:\nsent: %s\nwith len = %d\nIllegal symbol: %s in %s", t.Name(), lib, n, string(s), answer)
				}
			}
		}
	}
}
