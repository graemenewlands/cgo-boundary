package cgoboundary

import (
    "testing"
)

func benchmarkTestCGOOverhead(b *testing.B, n int) {
    for i:=0; i < b.N; i++ {
        FunctionToIterate(CGOOverhead, n)
    }
}

func benchmarkTestNoCGOOverhead(b *testing.B, n int) {
    for i:=0; i < b.N; i++ {
        FunctionToIterate(NoCGOOverhead, n)
    }
}


func BenchmarkCGOOverhead_1(b *testing.B) { benchmarkTestCGOOverhead(b, 1) }
func BenchmarkCGOOverhead_10(b *testing.B) { benchmarkTestCGOOverhead(b, 10) }
func BenchmarkCGOOverhead_100(b *testing.B) { benchmarkTestCGOOverhead(b, 100) }
func BenchmarkCGOOverhead_1000(b *testing.B) { benchmarkTestCGOOverhead(b, 1000) }
func BenchmarkCGOOverhead_10000(b *testing.B) { benchmarkTestCGOOverhead(b, 10000) }
func BenchmarkCGOOverhead_100000(b *testing.B) { benchmarkTestCGOOverhead(b, 100000) }

func BenchmarkNoCGOOverhead_1(b *testing.B) { benchmarkTestNoCGOOverhead(b, 1) }
func BenchmarkNoCGOOverhead_10(b *testing.B) { benchmarkTestNoCGOOverhead(b, 10) }
func BenchmarkNoCGOOverhead_100(b *testing.B) { benchmarkTestNoCGOOverhead(b, 100) }
func BenchmarkNoCGOOverhead_1000(b *testing.B) { benchmarkTestNoCGOOverhead(b, 1000) }
func BenchmarkNoCGOOverhead_10000(b *testing.B) { benchmarkTestNoCGOOverhead(b, 10000) }
func BenchmarkNoCGOOverhead_100000(b *testing.B) { benchmarkTestNoCGOOverhead(b, 100000) }

