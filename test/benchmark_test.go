package benchmark

import (
	"testing"

	"github.com/bjerkio/stackie/pkg/stackie"
)

func BenchmarkCmd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ = stackie.Setup()
	}
}
