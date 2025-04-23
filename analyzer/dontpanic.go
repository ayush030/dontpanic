package analyzer

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

type operator struct {
	// TODO: exceptions and exclusions
}

func (op *operator) run(pass *analysis.Pass) (_ any, retErr error) {
	defer func() {
		if err := recover(); err != nil {
			switch err := err.(type) {
			case error:
				retErr = err
			default:
				retErr = fmt.Errorf("panic: %v", err)
			}
		}
	}()

	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{ // filter the nodes that are call expressions
		(*ast.CallExpr)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		callExp := node.(*ast.CallExpr)

		if callExp.Fun != nil {
			ident, ok := callExp.Fun.(*ast.Ident) // can be a selector expression as well, skipping those
			if ok && ident.Name == "panic" {
				pass.Reportf(node.Pos(), "avoid use of panic")
			}
		}
	})

	return nil, nil
}

func newAnalyser() *analysis.Analyzer {
	var op = operator{}

	return &analysis.Analyzer{
		Name: "dontpanic",
		Doc:  "prevents from throwing panics and promotes on handling them",
		Run:  op.run,
		Requires: []*analysis.Analyzer{
			inspect.Analyzer,
		},
	}
}

func GetAnalyser() *analysis.Analyzer {
	return newAnalyser()
}
