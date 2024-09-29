package main

import "testing"

func BenchMarkBadFuncForMemory(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = BadFuncForMemory(1000)
	}
}
