package merror

import (
	"errors"
	"fmt"
	"testing"

	pkgError "github.com/pkg/errors"
)

// run benchmark go test -bench=.
func BenchmarkPKGError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = pkgError.New("dadada")
	}
}

// run benchmark go test -bench=.
func BenchmarkNewAppError(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewAppError("dadada")
	}
}

func BenchmarkAppError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// start := time.Now()
		_ = AppError(errors.New("text"), "bench")
		// fmt.Println(time.Since(start))
	}
}

func BenchmarkPkgErrors(b *testing.B) {
	for i := 0; i < b.N; i++ {

		_ = errors.New("dadada ")
	}
}

func BenchmarkPkgFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {

		_ = fmt.Errorf("dadada %d", i)
	}
}
