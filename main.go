package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strconv"
	"strings"
)

const (
	template1 = `
{{if .Define_USE_RINTERNALS }}#define USE_RINTERNALS{{end}}
{{if .IncludeRHeader }}#include <R.h>{{end}}
{{if .IncludeRinternalsHeader }}#include <Rinternals.h>{{end}}
#include "_cgo_export.h"

{{range .ExportedFunctions}}
For now, just print a string: {{.}}
{{end}}


	`
)

type templateData struct {
	DefineUseRinternals     bool
	IncludeRHeader          bool
	IncludeRinternalsHeader bool
	ExportedFunctions       []string // FIXME change type to something else
}

type badExportError struct{ error }
type receiverError struct{ error }
type multipleReturnValuesError struct{ error }
type badImportError struct{ error }
type wrongPackageError struct{ error }
type noMainFunctionError struct{ error }
type nonEmptyMainFunctionError struct{ error }
type noFunctionsToExportError struct{ error }

func getExportedFunctions(filename string, src interface{}) (
	exportedFuncs []*ast.FuncDecl, err error) {
	// FIXME make this work with `go generate`; be sure and handle the
	// `go generate` comment line in source files we process, don't associate it with
	// a function (like an export comment).
	fset := token.NewFileSet() // positions are relative to fset

	f, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	if f.Name.Name != "main" {
		return nil,
			wrongPackageError{fmt.Errorf("Package name '%s' invalid, must be 'main'.",
				f.Name.Name)}
	}

	foundC := false
	for _, anImport := range f.Imports {
		importPath := anImport.Path.Value
		for _, quoter := range []string{"\"", "'", "`"} {
			importPath = strings.Replace(importPath, quoter, "", -1)
			if importPath == "C" {
				foundC = true
				break
			}
		}
	}
	if !foundC {
		return nil, badImportError{errors.New("You must import \"C\"!")}
	}

	// cmap := ast.NewCommentMap(fset, f, f.Comments)

	var keepers []*ast.FuncDecl
	foundMainFunction := false

	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		fmt.Println("func name", fn.Name.Name)
		fmt.Println("doc", fn.Doc)
		if fn.Name.Name == "main" {
			if len(fn.Body.List) > 0 {
				return nil, nonEmptyMainFunctionError{errors.New("main function must be empty")}
			}
			foundMainFunction = true
		}
		if fn.Doc != nil {
			for _, comment := range fn.Doc.List {
				fmt.Println("comment", comment.Text)
				if strings.HasPrefix(comment.Text, "//export ") {
					segs := strings.Split(comment.Text, " ")
					exportedFunc := segs[len(segs)-1]
					if exportedFunc != fn.Name.Name {
						return nil,
							badExportError{
								fmt.Errorf("Function name in comment (%s) does not match function name in function (%s)\n",
									exportedFunc, fn.Name.Name)}
						// return nil, fmt.Errorf("Function name in comment (%s) does not match function name in function (%s)\n",
						// 	exportedFunc, fn.Name.Name)
					}
					if fn.Recv != nil {
						return nil, receiverError{
							errors.New("Can't export methods to foreign languages, only functions.")}
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
								multipleReturnValuesError{
									fmt.Errorf("Function %s must return 0 or 1 items, not %d\n",
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
	if !foundMainFunction {
		return nil, noMainFunctionError{errors.New("No main() function")}
	}

	if len(keepers) == 0 {
		return nil, noFunctionsToExportError{errors.New("No functions to export")}
	}

	return keepers, nil
}

func generateCcode(exportedFuncs []*ast.FuncDecl) {

}

func exit(exitCode int) int {
	if _, ok := os.LookupEnv("TESTING_FOREIGN"); ok {
		os.Setenv("FOREIGN_EXIT_CODE", strconv.Itoa(exitCode))
	} else { // These two lines can't be covered by tests
		os.Exit(exitCode) // without a lot of hassle.
	}
	return exitCode
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Supply the name of a Go source file.")
		exit(1)
		return
	}
	exportedFuncs, err := getExportedFunctions(os.Args[1], nil)
	if err != nil {
		os.Stderr.WriteString(err.Error()) // errors to stderr, save stdout for correct output
		os.Stderr.WriteString("\n")
		exit(1)
		return
	}
	generateCcode(exportedFuncs)
}
