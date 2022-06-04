package visibilityorder

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"github.com/samber/lo"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "visibilityorder",
	Doc:  "enforces public-first ordering of symbols",
	Run:  run,
}

// fileState maps from token type (const, func, type, or var) to unexported symbols of that type.
type fileState map[token.Token][]*ast.Ident

func run(pass *analysis.Pass) (res interface{}, err error) {
	for _, f := range pass.Files {
		state := fileState{
			token.CONST: {},
			token.FUNC:  {},
			token.TYPE:  {},
			token.VAR:   {},
		}
		for _, decl := range f.Decls {
			if err = state.addDeclaration(pass, decl); err != nil {
				return
			}
		}
	}

	return
}

func (s fileState) addDeclaration(pass *analysis.Pass, decl ast.Decl) error {
	switch d := decl.(type) {
	case *ast.GenDecl:
		s[d.Tok] = appendUnexportedSymbols(pass, s[d.Tok], genDeclSymbolIdentities(d))
	case *ast.FuncDecl:
		s[token.FUNC] = appendUnexportedSymbols(pass, s[token.FUNC], []*ast.Ident{d.Name})
	default:
		return fmt.Errorf("analyzers: unexpected declaration type %T", decl)
	}
	return nil
}

func appendUnexportedSymbols(pass *analysis.Pass, unexported, newIdents []*ast.Ident) []*ast.Ident {
	for _, id := range newIdents {
		if id.String() == "init" {
			continue // do not report incorrect placement of func init()
		}
		if !isExported(id) {
			unexported = append(unexported, id)
			continue
		}

		if len(unexported) != 0 {
			lastUnexportedID, _ := lo.Last(unexported)
			pass.Reportf(id.Pos(), "exported symbol %s appears after unexported symbol %s", id, lastUnexportedID)
		}
	}
	return unexported
}

func isExported(id *ast.Ident) bool {
	firstChar := string(id.String()[0])
	return firstChar == strings.ToUpper(firstChar)
}

func genDeclSymbolIdentities(d *ast.GenDecl) []*ast.Ident {
	switch d.Tok {
	case token.CONST:
		fallthrough
	case token.VAR:
		return lo.FlatMap(d.Specs, func(spec ast.Spec, _ int) []*ast.Ident { return spec.(*ast.ValueSpec).Names })
	case token.TYPE:
		return lo.Map(d.Specs, func(spec ast.Spec, _ int) *ast.Ident { return spec.(*ast.TypeSpec).Name })
	}
	// Ignores IMPORT tokens
	return []*ast.Ident{}
}
