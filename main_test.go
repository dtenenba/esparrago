package main

import (
	"os"
	"reflect"
	"testing"
)

// reduce code duplication for a common pattern
func helper(t *testing.T, expectedErr interface{}, code string) {
	_, err := getExportedFunctions("", code)
	if err == nil {
		t.Error("Expected error, got nil!")
	}

	expectedErrType := reflect.TypeOf(expectedErr)
	actualErrType := reflect.TypeOf(err)

	if actualErrType != expectedErrType {
		t.Error("Expected error", expectedErrType.String(), "but got",
			actualErrType.String(), "with message", err)
	}

}

func Test_getExportedFunctions(t *testing.T) {
	t.Run("noFunctions", func(t *testing.T) {
		helper(t, badImportError{}, "package main")
	})
	t.Run("nonexistentfile", func(t *testing.T) {
		_, err := getExportedFunctions("thisfiledoesnotexist", nil)
		if err == nil {
			t.Error("Expected failure for nonexistent file.")
		}
	})
	t.Run("fromFile", func(t *testing.T) {
		res, err := getExportedFunctions("testdata/src0.go", nil)
		if err != nil {
			t.Error("Expected result (len 1), not error.")
		} else {
			if len(res) != 1 {
				t.Errorf("Expected length 1, not %d.", len(res))
			}
		}
	})
	t.Run("badExportComment", func(t *testing.T) {
		helper(t, badExportError{}, `
	package main

	import "C"

	//export funky
	func funko() {}
				`)
	})
	t.Run("hasReceiver", func(t *testing.T) {
		helper(t, receiverError{}, `
package main

import "C"

//export hasreceiver
func (m int) hasreceiver() {}
			`)

	})
	t.Run("multipleReturnValues", func(t *testing.T) {
		helper(t, multipleReturnValuesError{}, `
package main

import "C"

//export  multipleitemsreturned
func multipleitemsreturned() (int, error) {
	return 0, nil
}
			`)
	})
	t.Run("wrongPackageName", func(t *testing.T) {
		helper(t, wrongPackageError{}, `
package hello
			`)
	})
	t.Run("nonEmptyMainFunction", func(t *testing.T) {
		helper(t, nonEmptyMainFunctionError{}, `
package main

import "C"

func main() {
	println("Hello")
}
			`)
	})
	t.Run("noMainFunction", func(t *testing.T) {
		helper(t, noMainFunctionError{}, `
package main

import "C"

//export foo
func foo(i int) {

}
			`)
	})
	t.Run("noFunctionsToExport", func(t *testing.T) {
		helper(t, noFunctionsToExportError{}, `
package main

import "C"

func main(){}
			`)
	})
}

func Test_generateCcode(t *testing.T) {
	t.Run("someName", func(t *testing.T) {
	})
}

func Test_main(t *testing.T) {
	oldArgs := os.Args
	os.Setenv("TESTING_FOREIGN", "true")
	defer func() { os.Unsetenv("TESTING_FOREIGN") }()
	defer func() { os.Args = oldArgs }()

	t.Run("withValidArg", func(t *testing.T) {
		os.Args = []string{"cmd", "testdata/src0.go"}
		main()
	})

	t.Run("withNonexistentFile", func(t *testing.T) {
		os.Args = []string{"cmd", "a_file_that_does_not_exist"}
		main()
		if os.Getenv("FOREIGN_EXIT_CODE") != "1" {
			t.Error("Expected exit with code 1, got.",
				os.Getenv("FOREIGN_EXIT_CODE"))
		}
	})

	t.Run("withNoArguments", func(t *testing.T) {
		os.Args = []string{"cmd"}
		main()
		if os.Getenv("FOREIGN_EXIT_CODE") != "1" {
			t.Error("Expected exit with code 1, got.",
				os.Getenv("FOREIGN_EXIT_CODE"))
		}
	})
}
