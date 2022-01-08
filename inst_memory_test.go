package llir

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/wa-lang/llir/constant"
	"github.com/wa-lang/llir/types"
)

func TestTypeCheckStore(t *testing.T) {
	cases := []struct {
		fromTyp, toTyp types.Type
		panicMessage   string // "OK" if not panic'ing.
	}{
		{types.I8, types.I8Ptr,
			"OK"},

		{types.I64, types.I8Ptr,
			"store operands are not compatible: src=i64; dst=i8*"},
		{types.I8, types.I8,
			"invalid store dst operand type; expected *types.Pointer, got *types.IntType"},
	}

	errOK := errors.New("OK")

	for _, c := range cases {
		testName := fmt.Sprintf("%v to %v", c.fromTyp, c.toTyp)
		t.Run(testName, func(t *testing.T) {
			var panicErr error
			fromVal := constant.NewZeroInitializer(c.fromTyp)
			toVal := constant.NewZeroInitializer(c.toTyp)
			func() {
				defer func() { panicErr = recover().(error) }()
				store := NewStore(fromVal, toVal)
				_ = store.LLString()
				panic(errOK)
			}()
			got := panicErr.Error()
			if got != c.panicMessage {
				t.Errorf("expected %q, got %q", c.panicMessage, got)
			}
		})
	}
}
