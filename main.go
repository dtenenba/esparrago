package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

type badExportError struct{ error }
type receiverError struct{ error }
type multipleReturnValuesError struct{ error }

func getExportedFunctions(filename string, src interface{}) (exportedFuncs []*ast.FuncDecl, err error) {
	// FIXME make sure source has 'import "C"' in it....
	// FIXME make sure package is main and that there is an empty main function
	// src is the input for which we want to print the AST.
	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset

	f, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	// cmap := ast.NewCommentMap(fset, f, f.Comments)

	var keepers []*ast.FuncDecl

	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		fmt.Println("func name", fn.Name.Name)
		fmt.Println("doc", fn.Doc)
		if fn.Doc != nil {
			for _, comment := range fn.Doc.List {
				fmt.Println("comment", comment.Text)
				if strings.HasPrefix(comment.Text, "//export ") {
					segs := strings.Split(comment.Text, " ")
					exportedFunc := segs[len(segs)-1]
					if exportedFunc != fn.Name.Name {
						return nil,
							badExportError{fmt.Errorf("Function name in comment (%s) does not match function name in function (%s)\n",
								exportedFunc, fn.Name.Name)}
						// return nil, fmt.Errorf("Function name in comment (%s) does not match function name in function (%s)\n",
						// 	exportedFunc, fn.Name.Name)
					}
					if fn.Recv != nil {
						return nil, receiverError{errors.New("Can't export methods to foreign languages, only functions.")}
					}
					// TODO make sure function is either void or returns only one thing
					// using fn.Type.Results.NumFields()

					if fn.Type.Results != nil {
						fmt.Println("output is", fn.Type.Results.List)
						for _, field := range fn.Type.Results.List {
							fmt.Println("  type is ", field.Type)
						}
						if fn.Type.Results.NumFields() > 1 {
							return nil,
								multipleReturnValuesError{fmt.Errorf("Function %s must return 0 or 1 items, not %d\n",
									fn.Name.Name, fn.Type.Results.NumFields())}
						}
					}
					for _, param := range fn.Type.Params.List {
						fmt.Println("  param type is", param.Type)
						fmt.Println("  param names", param.Names)
					}

					// assume this func is ok
					keepers = append(keepers, fn)
				}
			}
		}
		// fmt.Println(fn)
	}
	return keepers, nil
}

func generateCcode(exportedFuncs []*ast.FuncDecl) {

}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Supply the name of a Go source file.")
		os.Exit(1)
	}
	exportedFuncs, err := getExportedFunctions(os.Args[1], nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	generateCcode(exportedFuncs)
}
