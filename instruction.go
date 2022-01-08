package llir

// === [ Instructions ] ========================================================

// Instruction is an LLVM IR instruction. All instructions (except store and
// fence) implement the value.Named interface and may thus be used directly as
// values.
//
// An Instruction has one of the following underlying types.
//
// Unary instructions
//
// https://llvm.org/docs/LangRef.html#unary-operations
//
//    *ir.InstFNeg   // https://godoc.org/github.com/wa-lang/llir#InstFNeg
//
// Binary instructions
//
// https://llvm.org/docs/LangRef.html#binary-operations
//
//    *ir.InstAdd    // https://godoc.org/github.com/wa-lang/llir#InstAdd
//    *ir.InstFAdd   // https://godoc.org/github.com/wa-lang/llir#InstFAdd
//    *ir.InstSub    // https://godoc.org/github.com/wa-lang/llir#InstSub
//    *ir.InstFSub   // https://godoc.org/github.com/wa-lang/llir#InstFSub
//    *ir.InstMul    // https://godoc.org/github.com/wa-lang/llir#InstMul
//    *ir.InstFMul   // https://godoc.org/github.com/wa-lang/llir#InstFMul
//    *ir.InstUDiv   // https://godoc.org/github.com/wa-lang/llir#InstUDiv
//    *ir.InstSDiv   // https://godoc.org/github.com/wa-lang/llir#InstSDiv
//    *ir.InstFDiv   // https://godoc.org/github.com/wa-lang/llir#InstFDiv
//    *ir.InstURem   // https://godoc.org/github.com/wa-lang/llir#InstURem
//    *ir.InstSRem   // https://godoc.org/github.com/wa-lang/llir#InstSRem
//    *ir.InstFRem   // https://godoc.org/github.com/wa-lang/llir#InstFRem
//
// Bitwise instructions
//
// https://llvm.org/docs/LangRef.html#bitwise-binary-operations
//
//    *ir.InstShl    // https://godoc.org/github.com/wa-lang/llir#InstShl
//    *ir.InstLShr   // https://godoc.org/github.com/wa-lang/llir#InstLShr
//    *ir.InstAShr   // https://godoc.org/github.com/wa-lang/llir#InstAShr
//    *ir.InstAnd    // https://godoc.org/github.com/wa-lang/llir#InstAnd
//    *ir.InstOr     // https://godoc.org/github.com/wa-lang/llir#InstOr
//    *ir.InstXor    // https://godoc.org/github.com/wa-lang/llir#InstXor
//
// Vector instructions
//
// https://llvm.org/docs/LangRef.html#vector-operations
//
//    *ir.InstExtractElement   // https://godoc.org/github.com/wa-lang/llir#InstExtractElement
//    *ir.InstInsertElement    // https://godoc.org/github.com/wa-lang/llir#InstInsertElement
//    *ir.InstShuffleVector    // https://godoc.org/github.com/wa-lang/llir#InstShuffleVector
//
// Aggregate instructions
//
// https://llvm.org/docs/LangRef.html#aggregate-operations
//
//    *ir.InstExtractValue   // https://godoc.org/github.com/wa-lang/llir#InstExtractValue
//    *ir.InstInsertValue    // https://godoc.org/github.com/wa-lang/llir#InstInsertValue
//
// Memory instructions
//
// https://llvm.org/docs/LangRef.html#memory-access-and-addressing-operations
//
//    *ir.InstAlloca          // https://godoc.org/github.com/wa-lang/llir#InstAlloca
//    *ir.InstLoad            // https://godoc.org/github.com/wa-lang/llir#InstLoad
//    *ir.InstStore           // https://godoc.org/github.com/wa-lang/llir#InstStore
//    *ir.InstFence           // https://godoc.org/github.com/wa-lang/llir#InstFence
//    *ir.InstCmpXchg         // https://godoc.org/github.com/wa-lang/llir#InstCmpXchg
//    *ir.InstAtomicRMW       // https://godoc.org/github.com/wa-lang/llir#InstAtomicRMW
//    *ir.InstGetElementPtr   // https://godoc.org/github.com/wa-lang/llir#InstGetElementPtr
//
// Conversion instructions
//
// https://llvm.org/docs/LangRef.html#conversion-operations
//
//    *ir.InstTrunc           // https://godoc.org/github.com/wa-lang/llir#InstTrunc
//    *ir.InstZExt            // https://godoc.org/github.com/wa-lang/llir#InstZExt
//    *ir.InstSExt            // https://godoc.org/github.com/wa-lang/llir#InstSExt
//    *ir.InstFPTrunc         // https://godoc.org/github.com/wa-lang/llir#InstFPTrunc
//    *ir.InstFPExt           // https://godoc.org/github.com/wa-lang/llir#InstFPExt
//    *ir.InstFPToUI          // https://godoc.org/github.com/wa-lang/llir#InstFPToUI
//    *ir.InstFPToSI          // https://godoc.org/github.com/wa-lang/llir#InstFPToSI
//    *ir.InstUIToFP          // https://godoc.org/github.com/wa-lang/llir#InstUIToFP
//    *ir.InstSIToFP          // https://godoc.org/github.com/wa-lang/llir#InstSIToFP
//    *ir.InstPtrToInt        // https://godoc.org/github.com/wa-lang/llir#InstPtrToInt
//    *ir.InstIntToPtr        // https://godoc.org/github.com/wa-lang/llir#InstIntToPtr
//    *ir.InstBitCast         // https://godoc.org/github.com/wa-lang/llir#InstBitCast
//    *ir.InstAddrSpaceCast   // https://godoc.org/github.com/wa-lang/llir#InstAddrSpaceCast
//
// Other instructions
//
// https://llvm.org/docs/LangRef.html#other-operations
//
//    *ir.InstICmp         // https://godoc.org/github.com/wa-lang/llir#InstICmp
//    *ir.InstFCmp         // https://godoc.org/github.com/wa-lang/llir#InstFCmp
//    *ir.InstPhi          // https://godoc.org/github.com/wa-lang/llir#InstPhi
//    *ir.InstSelect       // https://godoc.org/github.com/wa-lang/llir#InstSelect
//    *ir.InstFreeze       // https://godoc.org/github.com/wa-lang/llir#InstFreeze
//    *ir.InstCall         // https://godoc.org/github.com/wa-lang/llir#InstCall
//    *ir.InstVAArg        // https://godoc.org/github.com/wa-lang/llir#InstVAArg
//    *ir.InstLandingPad   // https://godoc.org/github.com/wa-lang/llir#InstLandingPad
//    *ir.InstCatchPad     // https://godoc.org/github.com/wa-lang/llir#InstCatchPad
//    *ir.InstCleanupPad   // https://godoc.org/github.com/wa-lang/llir#InstCleanupPad
type Instruction interface {
	LLStringer
	// isInstruction ensures that only instructions can be assigned to the
	// instruction.Instruction interface.
	isInstruction()
}
