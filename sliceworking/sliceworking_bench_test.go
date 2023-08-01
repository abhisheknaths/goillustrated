package sliceworking

import "testing"

func Benchmark_AppendWithCapacity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendWithCapacity(1000)
	}
}

func Benchmark_AppendWithoutCapacity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendWithoutCapacity(1000)
	}
}

func Benchmark_Compare(b *testing.B) {
	tests := map[string]struct {
		testFunc func(b *testing.B)
	}{
		"AppendWithCapacity": {
			testFunc: Benchmark_AppendWithCapacity,
		},
		"AppendWithoutCapacity": {
			testFunc: Benchmark_AppendWithoutCapacity,
		},
	}
	for name, test := range tests {
		b.Run(name, test.testFunc)
	}
}
