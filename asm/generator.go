package asm

import (
	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
)

// generator keeps track of top-level entities when translating from AST to IR
// representation.
type generator struct {
	// LLVM IR module being generated.
	m *ir.Module
	// index of AST top-level entities.
	old oldIndex
	// index of IR top-level entities.
	new newIndex

	// TODO: add rw mutex to gen.todo for access to blockaddress constant.

	// Fix dummy basic blocks after translation of function bodies and assignment
	// of local IDs.
	todo []*constant.BlockAddress
}

// newGenerator returns a new generator for translating an LLVM IR module from
// AST to IR representation.
func newGenerator() *generator {
	return &generator{
		m: ir.NewModule(),
		old: oldIndex{
			typeDefs:          make(map[string]*ast.TypeDef),
			comdatDefs:        make(map[string]*ast.ComdatDef),
			globals:           make(map[ir.GlobalIdent]ast.LlvmNode),
			attrGroupDefs:     make(map[int64][]*ast.AttrGroupDef),
			namedMetadataDefs: make(map[string][]*ast.NamedMetadataDef),
			metadataDefs:      make(map[int64]*ast.MetadataDef),
		},
		new: newIndex{
			typeDefs:          make(map[string]types.Type),
			comdatDefs:        make(map[string]*ir.ComdatDef),
			globals:           make(map[ir.GlobalIdent]constant.Constant),
			attrGroupDefs:     make(map[int64]*ir.AttrGroupDef),
			namedMetadataDefs: make(map[string]*metadata.NamedDef),
			metadataDefs:      make(map[int64]metadata.Definition),
		},
	}
}

// oldIndex is an index of AST top-level entities.
type oldIndex struct {
	// typeDefs maps from type identifier (without '%' prefix) to type
	// definition.
	typeDefs map[string]*ast.TypeDef
	// comdatDefs maps from comdat name (without '$' prefix) to comdat
	// definition.
	comdatDefs map[string]*ast.ComdatDef
	// globals maps from global identifier (without '@' prefix) to global
	// declarations and defintions, indirect symbol definitions, function
	// declarations and definitions.
	//
	// The value has one of the following types.
	//    *ast.GlobalDecl
	//    *ast.GlobalDef
	//    *ast.AliasDef
	//    *ast.IFuncDef
	//    *ast.FuncDecl
	//    *ast.FuncDef
	globals map[ir.GlobalIdent]ast.LlvmNode
	// attrGroupDefs maps from attribute group ID (without '#' prefix) to
	// attribute group definitions. Each ID maps to one or more attribute group
	// definitions which will be merged into a single attribute group definition
	// during translation.
	attrGroupDefs map[int64][]*ast.AttrGroupDef
	// namedMetadataDefs maps from metadata name (without '!' prefix) to named
	// metadata definitions with the same name.
	namedMetadataDefs map[string][]*ast.NamedMetadataDef
	// metadataDefs maps from metadata ID (without '!' prefix) to metadata
	// definition.
	metadataDefs map[int64]*ast.MetadataDef
	// useListOrders is a slice of use-list orders in their order of occurrence
	// in the input.
	useListOrders []*ast.UseListOrder
	// useListOrderBBs is a slice of basic block specific use-list orders in
	// their order of occurrence in the input.
	useListOrderBBs []*ast.UseListOrderBB

	// globalOrder records the global identifier of global declarations and
	// definitions, indirect symbol definitions, and function declarations and
	// definitions in their order of occurrence in the input.
	globalOrder []ir.GlobalIdent
}

// newIndex is an index of IR top-level entities.
type newIndex struct {
	// typeDefs maps from type identifier (without '%' prefix) to type
	// definition.
	typeDefs map[string]types.Type
	// comdatDefs maps from comdat name (without '$' prefix) to comdat
	// definition.
	comdatDefs map[string]*ir.ComdatDef
	// globals maps from global identifier (without '@' prefix) to global
	// declarations and defintions, indirect symbol definitions, function
	// declarations and definitions.
	//
	// The value has one of the following types.
	//    *ir.Global
	//    *ir.Alias
	//    *ir.IFunc
	//    *ir.Func
	globals map[ir.GlobalIdent]constant.Constant
	// attrGroupDefs maps from attribute group ID (without '#' prefix) to
	// attribute group definition.
	attrGroupDefs map[int64]*ir.AttrGroupDef
	// namedMetadataDefs maps from metadata name (without '!' prefix) to named
	// metadata definition.
	namedMetadataDefs map[string]*metadata.NamedDef
	// metadataDefs maps from metadata ID (without '!' prefix) to metadata
	// definition.
	metadataDefs map[int64]metadata.Definition
}
