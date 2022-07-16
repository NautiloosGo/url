package app

import (
	"testing"
)

const letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"

// compare random func
func benchGet(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRandomString(n, letters)
	}
}

func BenchmarkGetRandomString3(b *testing.B)  { benchGet(3, b) }
func BenchmarkGetRandomString5(b *testing.B)  { benchGet(5, b) }
func BenchmarkGetRandomString7(b *testing.B)  { benchGet(7, b) }
func BenchmarkGetRandomString10(b *testing.B) { benchGet(10, b) }
func BenchmarkGetRandomString15(b *testing.B) { benchGet(15, b) }
func BenchmarkGetRandomString20(b *testing.B) { benchGet(20, b) }

func benchGetFaster(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRandomStringFaster(n, letters)
	}
}

func BenchmarkGetRandomStringFaster3(b *testing.B)  { benchGetFaster(3, b) }
func BenchmarkGetRandomStringFaster5(b *testing.B)  { benchGetFaster(5, b) }
func BenchmarkGetRandomStringFaster7(b *testing.B)  { benchGetFaster(7, b) }
func BenchmarkGetRandomStringFaster10(b *testing.B) { benchGetFaster(10, b) }
func BenchmarkGetRandomStringFaster15(b *testing.B) { benchGetFaster(15, b) }
func BenchmarkGetRandomStringFaster20(b *testing.B) { benchGetFaster(20, b) }
