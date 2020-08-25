package client

import (
	"testing"
)

func BenchmarkInitESClient(b *testing.B) {
	for n := 0; n < b.N; n++ {
		InitESClient()
	}
}
