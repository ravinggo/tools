package output_tests

import (
	"io"
	_ "net/http/pprof"
	"testing"

	"github.com/ravinggo/zerolog"

	"github.com/ravinggo/tools/zerolog-gen/output_tests/output"
)

func BenchmarkLog(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	user := &output.User{
		ID:   12312,
		Name: "sdasdasdas",
		Age:  99,
	}
	log := zerolog.New(io.Discard)
	for i := 0; i < b.N; i++ {
		log.Debug().Any("user", user).Send()
	}
}
