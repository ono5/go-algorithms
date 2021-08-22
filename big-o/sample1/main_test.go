package main

import "testing"

func BenchmarkFindCookie(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findCookie(snacks)
	}
}
