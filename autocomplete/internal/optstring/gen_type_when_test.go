package optstring

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */


import (
	"testing"
)

func TestTypeWhenNone(t *testing.T) {
	const notChanged = "not changed"
	const changed    = "changed"

	{
		s := notChanged

		None().WhenNone(func(){
			s = changed
		})

		if expected, actual := changed, s; expected != actual {
			t.Errorf("Expected %v, but actually got %v.", expected, actual)
			return
		}
	}

	{
		s := notChanged

		var value string
		Some(value).WhenNone(func(){
			s = changed
		})

		if expected, actual := notChanged, s; expected != actual {
			t.Errorf("Expected %v, but actually got %v.", expected, actual)
			return
		}
	}
}

func TestTypeWhenSome(t *testing.T) {
	const notChanged = "not changed"
	const changed    = "changed"

	{
		s := notChanged

		None().WhenSome(func(datum string){
			s = changed
		})

		if expected, actual := notChanged, s; expected != actual {
			t.Errorf("Expected %v, but actually got %v.", expected, actual)
			return
		}
	}

	{
		s := notChanged

		var datum string
		Some(datum).WhenSome(func(datum string){
			s = changed
		})

		if expected, actual := changed, s; expected != actual {
			t.Errorf("Expected %v, but actually got %v.", expected, actual)
			return
		}
	}
}
