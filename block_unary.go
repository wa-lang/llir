package llir

import (
	"github.com/wa-lang/llir/value"
)

// --- [ Unary instructions ] --------------------------------------------------

// ~~~ [ fneg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFNeg appends a new fneg instruction to the basic block based on the given
// operand.
func (block *Block) NewFNeg(x value.Value) *InstFNeg {
	inst := NewFNeg(x)
	block.Insts = append(block.Insts, inst)
	return inst
}
