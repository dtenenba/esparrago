package main

import "testing"

func Test_getExportedFunctions(t *testing.T) {
	t.Run("noFunctions", func(t *testing.T) {
		//FIXME this should eventually fail when sanity checks are added
		res, err := getExportedFunctions("", "package foo")
		if err != nil {
			t.Error("Expected [[], nil], got error")
		} else {
			t.Log("res is", res)
		}
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
		_, err := getExportedFunctions("", `
package main

//export funky
func funko() {}
			`)
		if err == nil {
			t.Error("Should have gotten an error")
		} else {
			if _, ok := err.(badExportError); !ok {
				t.Error("got the wrong error!")
			}
		}
	})
	t.Run("hasReceiver", func(t *testing.T) {
		_, err := getExportedFunctions("", `
package main

import "C"

//export hasreceiver
func (m int) hasreceiver() {}
			`)

		if err == nil {
			t.Error("Should have gotten an error!")
		}
		if _, ok := err.(receiverError); !ok {
			t.Error("got the wrong error!")
		}
	})
	t.Run("multipleReturnValues", func(t *testing.T) {
		_, err := getExportedFunctions("", `
package main

import "C"

//export  multipleitemsreturned
func multipleitemsreturned() (int, error) {
	return 0, nil
}
			`)
		if err == nil {
			t.Error("Should have gotten an error!")
		}
		if _, ok := err.(multipleReturnValuesError); !ok {
			t.Error("got the wrong error!")
		}
	})
}

func Test_generateCcode(t *testing.T) {
	t.Run("someName", func(t *testing.T) {
	})
}
