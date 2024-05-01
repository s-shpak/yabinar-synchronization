package mutex

import "testing"

func BenchmarkFrequentReadsFrequentWrites(b *testing.B) {
	ctr := NewMuxCounter()
	for n := 0; n < b.N; n++ {
		ctr.FrequentReadsFrequentWrites(nil)
	}
}

func BenchmarkFrequentReadsFrequentWritesRW(b *testing.B) {
	ctr := NewRWMuxCounter()
	for n := 0; n < b.N; n++ {
		ctr.FrequentReadsFrequentWrites(nil)
	}
}

func BenchmarkFrequentReadsSpareWrites(b *testing.B) {
	ctr := NewMuxCounter()
	for n := 0; n < b.N; n++ {
		ctr.FrequentReadsSpareWrites(nil)
	}
}

func BenchmarkFrequentReadsSpareWritesRW(b *testing.B) {
	ctr := NewRWMuxCounter()
	for n := 0; n < b.N; n++ {
		ctr.FrequentReadsSpareWrites(nil)
	}
}
