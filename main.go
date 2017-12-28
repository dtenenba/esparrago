package foreign

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

type FuncVisitor struct {
}

func twohey() {
	// src is the input for which we want to print the AST.

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "testdata/src0.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	// cmap := ast.NewCommentMap(fset, f, f.Comments)

	// var keepers []*ast.FuncDecl

	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		fmt.Println("hoza", fn.Name.Name)
		fmt.Println("doc", fn.Doc)
		if fn.Doc != nil {
			for _, comment := range fn.Doc.List {
				fmt.Println("comment", comment.Text)
				if strings.HasPrefix(comment.Text, "//export ") {
					segs := strings.Split(comment.Text, " ")
					exportedFunc := segs[len(segs)-1]
					if exportedFunc != fn.Name.Name {
						fmt.Printf("Function name in comment (%s) does not match function name in function (%s)\n",
							exportedFunc, fn.Name.Name)
						os.Exit(1)
					}
					if fn.Recv != nil {
						fmt.Println("Can't export methods to foreign languages, only functions.")
						os.Exit(1)
					}
					fmt.Println("output is", fn.Type.Results)
				}
			}
		}
		fmt.Println(fn)
	}

}

func (v *FuncVisitor) Visit(node ast.Node) (w ast.Visitor) {
	fmt.Println("visited node:", node)
	return v
}
