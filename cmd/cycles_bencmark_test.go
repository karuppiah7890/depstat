package cmd

import "testing"

var err error

func BenchmarkCycles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err = runCycles(nil, nil)
		if err != nil {
			b.Errorf("error occurred while running cycles command: %v", err)
		}
	}
}
