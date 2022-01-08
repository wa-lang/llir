// Package constant implements values representing immutable LLVM IR constants.
package constant

import (
	"github.com/wa-lang/llir/types"
	"github.com/wa-lang/llir/value"
)

// === [ Constants ] ===========================================================

// Convenience constants.
var (
	// None token constant.
	None = &NoneToken{} // none
	// Boolean constants.
	True  = NewInt(types.I1, 1) // true
	False = NewInt(types.I1, 0) // false
)

// Constant is an LLVM IR constant; a value that is immutable at runtime, such
// as an integer or floating-point literal, or the address of a function or
// global variable.
//
// A Constant has one of the following underlying types.
//
// Simple constants
//
// https://llvm.org/docs/LangRef.html#simple-constants
//
//    *constant.Int         // https://godoc.org/github.com/wa-lang/llir/constant#Int
//    *constant.Float       // https://godoc.org/github.com/wa-lang/llir/constant#Float
//    *constant.Null        // https://godoc.org/github.com/wa-lang/llir/constant#Null
//    *constant.NoneToken   // https://godoc.org/github.com/wa-lang/llir/constant#NoneToken
//
// Complex constants
//
// https://llvm.org/docs/LangRef.html#complex-constants
//
//    *constant.Struct            // https://godoc.org/github.com/wa-lang/llir/constant#Struct
//    *constant.Array             // https://godoc.org/github.com/wa-lang/llir/constant#Array
//    *constant.CharArray         // https://godoc.org/github.com/wa-lang/llir/constant#CharArray
//    *constant.Vector            // https://godoc.org/github.com/wa-lang/llir/constant#Vector
//    *constant.ZeroInitializer   // https://godoc.org/github.com/wa-lang/llir/constant#ZeroInitializer
//    TODO: include metadata node?
//
// Global variable and function addresses
//
// https://llvm.org/docs/LangRef.html#global-variable-and-function-addresses
//
//    *ir.Global   // https://godoc.org/github.com/wa-lang/llir#Global
//    *ir.Func     // https://godoc.org/github.com/wa-lang/llir#Func
//    *ir.Alias    // https://godoc.org/github.com/wa-lang/llir#Alias
//    *ir.IFunc    // https://godoc.org/github.com/wa-lang/llir#IFunc
//
// Undefined values
//
// https://llvm.org/docs/LangRef.html#undefined-values
//
//    *constant.Undef   // https://godoc.org/github.com/wa-lang/llir/constant#Undef
//
// Poison values
//
// https://llvm.org/docs/LangRef.html#poison-values
//
//    *constant.Poison   // https://godoc.org/github.com/wa-lang/llir/constant#Poison
//
// Addresses of basic blocks
//
// https://llvm.org/docs/LangRef.html#addresses-of-basic-blocks
//
//    *constant.BlockAddress   // https://godoc.org/github.com/wa-lang/llir/constant#BlockAddress
//
// Constant expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    constant.Expression   // https://godoc.org/github.com/wa-lang/llir/constant#Expression
type Constant interface {
	value.Value
	// IsConstant ensures that only constants can be assigned to the
	// constant.Constant interface.
	IsConstant()
}
