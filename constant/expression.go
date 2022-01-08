package constant

// === [ Expressions ] =========================================================

// Expression is an LLVM IR constant expression.
//
// An Expression has one of the following underlying types.
//
// Unary expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *constant.ExprFNeg   // https://godoc.org/github.com/wa-lang/llir/constant#ExprFNeg
//
// Binary expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *constant.ExprAdd    // https://godoc.org/github.com/wa-lang/llir/constant#ExprAdd
//    *constant.ExprFAdd   // https://godoc.org/github.com/wa-lang/llir/constant#ExprFAdd
//    *constant.ExprSub    // https://godoc.org/github.com/wa-lang/llir/constant#ExprSub
//    *constant.ExprFSub   // https://godoc.org/github.com/wa-lang/llir/constant#ExprFSub
//    *constant.ExprMul    // https://godoc.org/github.com/wa-lang/llir/constant#ExprMul
//    *constant.ExprFMul   // https://godoc.org/github.com/wa-lang/llir/constant#ExprFMul
//    *constant.ExprUDiv   // https://godoc.org/github.com/wa-lang/llir/constant#ExprUDiv
//    *constant.ExprSDiv   // https://godoc.org/github.com/wa-lang/llir/constant#ExprSDiv
//    *constant.ExprFDiv   // https://godoc.org/github.com/wa-lang/llir/constant#ExprFDiv
//    *constant.ExprURem   // https://godoc.org/github.com/wa-lang/llir/constant#ExprURem
//    *constant.ExprSRem   // https://godoc.org/github.com/wa-lang/llir/constant#ExprSRem
//    *constant.ExprFRem   // https://godoc.org/github.com/wa-lang/llir/constant#ExprFRem
//
// Bitwise expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *constant.ExprShl    // https://godoc.org/github.com/wa-lang/llir/constant#ExprShl
//    *constant.ExprLShr   // https://godoc.org/github.com/wa-lang/llir/constant#ExprLShr
//    *constant.ExprAShr   // https://godoc.org/github.com/wa-lang/llir/constant#ExprAShr
//    *constant.ExprAnd    // https://godoc.org/github.com/wa-lang/llir/constant#ExprAnd
//    *constant.ExprOr     // https://godoc.org/github.com/wa-lang/llir/constant#ExprOr
//    *constant.ExprXor    // https://godoc.org/github.com/wa-lang/llir/constant#ExprXor
//
// Vector expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *constant.ExprExtractElement   // https://godoc.org/github.com/wa-lang/llir/constant#ExprExtractElement
//    *constant.ExprInsertElement    // https://godoc.org/github.com/wa-lang/llir/constant#ExprInsertElement
//    *constant.ExprShuffleVector    // https://godoc.org/github.com/wa-lang/llir/constant#ExprShuffleVector
//
// Aggregate expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *constant.ExprExtractValue   // https://godoc.org/github.com/wa-lang/llir/constant#ExprExtractValue
//    *constant.ExprInsertValue    // https://godoc.org/github.com/wa-lang/llir/constant#ExprInsertValue
//
// Memory expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *constant.ExprGetElementPtr   // https://godoc.org/github.com/wa-lang/llir/constant#ExprGetElementPtr
//
// Conversion expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *constant.ExprTrunc           // https://godoc.org/github.com/wa-lang/llir/constant#ExprTrunc
//    *constant.ExprZExt            // https://godoc.org/github.com/wa-lang/llir/constant#ExprZExt
//    *constant.ExprSExt            // https://godoc.org/github.com/wa-lang/llir/constant#ExprSExt
//    *constant.ExprFPTrunc         // https://godoc.org/github.com/wa-lang/llir/constant#ExprFPTrunc
//    *constant.ExprFPExt           // https://godoc.org/github.com/wa-lang/llir/constant#ExprFPExt
//    *constant.ExprFPToUI          // https://godoc.org/github.com/wa-lang/llir/constant#ExprFPToUI
//    *constant.ExprFPToSI          // https://godoc.org/github.com/wa-lang/llir/constant#ExprFPToSI
//    *constant.ExprUIToFP          // https://godoc.org/github.com/wa-lang/llir/constant#ExprUIToFP
//    *constant.ExprSIToFP          // https://godoc.org/github.com/wa-lang/llir/constant#ExprSIToFP
//    *constant.ExprPtrToInt        // https://godoc.org/github.com/wa-lang/llir/constant#ExprPtrToInt
//    *constant.ExprIntToPtr        // https://godoc.org/github.com/wa-lang/llir/constant#ExprIntToPtr
//    *constant.ExprBitCast         // https://godoc.org/github.com/wa-lang/llir/constant#ExprBitCast
//    *constant.ExprAddrSpaceCast   // https://godoc.org/github.com/wa-lang/llir/constant#ExprAddrSpaceCast
//
// Other expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *constant.ExprICmp     // https://godoc.org/github.com/wa-lang/llir/constant#ExprICmp
//    *constant.ExprFCmp     // https://godoc.org/github.com/wa-lang/llir/constant#ExprFCmp
//    *constant.ExprSelect   // https://godoc.org/github.com/wa-lang/llir/constant#ExprSelect
type Expression interface {
	Constant
	// IsExpression ensures that only constants expressions can be assigned to
	// the constant.Expression interface.
	IsExpression()
}
