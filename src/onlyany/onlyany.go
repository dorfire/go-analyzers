package onlyany

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer implements a Go static analyzer that reports uses of `interface{}`.
var Analyzer = &analysis.Analyzer{
	Name: "onlyany",
	Doc: "prefers any over interface{} - see https://go-review.googlesource.com/c/gofrontend/+/382248/",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (res any, err error) {
	astInspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	astInspector.Preorder([]ast.Node{(*ast.Expr)(nil)}, func(node ast.Node) {
		switch e := n.(type) {
		case *ast.InterfaceType:
			if len(e.Methods.List) == 0 {
				pass.Reportf(e.Pos(), "use any instead of an empty interface")
			}
		}
	})
	return
}
