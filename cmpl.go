package otto

import (
	"github.com/BigtigerGG/otto/ast"
	"github.com/BigtigerGG/otto/file"
)

type _compiler struct {
	file    *file.File
	program *ast.Program
}

func (cmpl *_compiler) parse() *_nodeProgram {
	if cmpl.program != nil {
		cmpl.file = cmpl.program.File
	}
	return cmpl._parse(cmpl.program)
}
