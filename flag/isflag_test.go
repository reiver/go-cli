package cliflag

import (
	"testing"
)

func TestIsFlagTrue(t *testing.T) {

	tests := []struct{
		Value string
	}{
		{
			Value: "--apple",
		},
		{
			Value: "--banana",
		},
		{
			Value: "--cherry",
		},



		{
			Value: "--apple=one",
		},
		{
			Value: "--banana=two",
		},
		{
			Value: "--cherry=three",
		},



		{
			Value: "--apple=o",
		},
		{
			Value: "--banana=t",
		},
		{
			Value: "--cherry=t",
		},



		{
			Value: "--apple=",
		},
		{
			Value: "--banana=",
		},
		{
			Value: "--cherry=",
		},



		{
			Value: "--apple:one",
		},
		{
			Value: "--banana:two",
		},
		{
			Value: "--cherry:three",
		},



		{
			Value: "--apple:o",
		},
		{
			Value: "--banana:t",
		},
		{
			Value: "--cherry:t",
		},



		{
			Value: "--apple:",
		},
		{
			Value: "--banana:",
		},
		{
			Value: "--cherry:",
		},



		{
			Value: "--a",
		},
		{
			Value: "--b",
		},
		{
			Value: "--c",
		},



		{
			Value: "--a=one",
		},
		{
			Value: "--b=two",
		},
		{
			Value: "--c=three",
		},



		{
			Value: "--a=o",
		},
		{
			Value: "--b=t",
		},
		{
			Value: "--c=t",
		},



		{
			Value: "--a=",
		},
		{
			Value: "--b=",
		},
		{
			Value: "--c=",
		},



		{
			Value: "--a:one",
		},
		{
			Value: "--b:two",
		},
		{
			Value: "--c:three",
		},



		{
			Value: "--a:o",
		},
		{
			Value: "--b:t",
		},
		{
			Value: "--c:t",
		},



		{
			Value: "--a:",
		},
		{
			Value: "--b:",
		},
		{
			Value: "--c:",
		},



		{
			Value: "+apple",
		},
		{
			Value: "+banana",
		},
		{
			Value: "+cherry",
		},



		{
			Value: "+apple=one",
		},
		{
			Value: "+banana=two",
		},
		{
			Value: "+cherry=three",
		},



		{
			Value: "+apple=o",
		},
		{
			Value: "+banana=t",
		},
		{
			Value: "+cherry=t",
		},



		{
			Value: "+apple=",
		},
		{
			Value: "+banana=",
		},
		{
			Value: "+cherry=",
		},



		{
			Value: "+apple:one",
		},
		{
			Value: "+banana:two",
		},
		{
			Value: "+cherry:three",
		},



		{
			Value: "+apple:o",
		},
		{
			Value: "+banana:t",
		},
		{
			Value: "+cherry:t",
		},



		{
			Value: "+apple:",
		},
		{
			Value: "+banana:",
		},
		{
			Value: "+cherry:",
		},



		{
			Value: "+a",
		},
		{
			Value: "+b",
		},
		{
			Value: "+c",
		},



		{
			Value: "+a=one",
		},
		{
			Value: "+b=two",
		},
		{
			Value: "+c=three",
		},



		{
			Value: "+a=o",
		},
		{
			Value: "+b=t",
		},
		{
			Value: "+c=t",
		},



		{
			Value: "+a=",
		},
		{
			Value: "+b=",
		},
		{
			Value: "+c=",
		},



		{
			Value: "+a:one",
		},
		{
			Value: "+b:two",
		},
		{
			Value: "+c:three",
		},



		{
			Value: "+a:o",
		},
		{
			Value: "+b:t",
		},
		{
			Value: "+c:t",
		},



		{
			Value: "+a:",
		},
		{
			Value: "+b:",
		},
		{
			Value: "+c:",
		},



		{
			Value: "-a",
		},
		{
			Value: "-b",
		},
		{
			Value: "-c",
		},



		{
			Value: "-a=one",
		},
		{
			Value: "-b=two",
		},
		{
			Value: "-c=three",
		},



		{
			Value: "-a=o",
		},
		{
			Value: "-b=t",
		},
		{
			Value: "-c=t",
		},



		{
			Value: "-a=",
		},
		{
			Value: "-b=",
		},
		{
			Value: "-c=",
		},



		{
			Value: "-a:one",
		},
		{
			Value: "-b:two",
		},
		{
			Value: "-c:three",
		},



		{
			Value: "-a:o",
		},
		{
			Value: "-b:t",
		},
		{
			Value: "-c:t",
		},



		{
			Value: "-a:",
		},
		{
			Value: "-b:",
		},
		{
			Value: "-c:",
		},



		{
			Value: "-aone",
		},
		{
			Value: "-btwo",
		},
		{
			Value: "-cthree",
		},



		{
			Value: "-ao",
		},
		{
			Value: "-bt",
		},
		{
			Value: "-ct",
		},
	}

	for testNumber, test := range tests {

		if expected, actual := true, isFlag(test.Value); expected != actual {
			t.Errorf("For test #%d, expected %t, but actually got %t.", testNumber, expected, actual)
			t.Errorf("\tVALUE: %q", test.Value)
			continue
		}
	}
}

func TestIsFlagFalse(t *testing.T) {

	tests := []struct{
		Value string
	}{
		{
			Value: "",
		},



		{
			Value: " ",
		},
		{
			Value: "  ",
		},
		{
			Value: "   ",
		},
		{
			Value: "    ",
		},









		{
			Value: "apple",
		},
		{
			Value: "banana",
		},
		{
			Value: "cherry",
		},



		{
			Value: "a",
		},
		{
			Value: "b",
		},
		{
			Value: "c",
		},









		{
			Value: "-",
		},
		{
			Value: "--",
		},
		{
			Value: "---",
		},
		{
			Value: "----",
		},



		{
			Value: "+",
		},
		{
			Value: "++",
		},
		{
			Value: "+++",
		},
		{
			Value: "++++",
		},









		{
			Value: "---a",
		},
		{
			Value: "---b",
		},
		{
			Value: "---c",
		},



		{
			Value: "---apple",
		},
		{
			Value: "---banana",
		},
		{
			Value: "---cherry",
		},




		{
			Value: "---a=",
		},
		{
			Value: "---b=",
		},
		{
			Value: "---c=",
		},



		{
			Value: "---apple=",
		},
		{
			Value: "---banana=",
		},
		{
			Value: "---cherry=",
		},



		{
			Value: "---a=o",
		},
		{
			Value: "---b=t",
		},
		{
			Value: "---c=t",
		},



		{
			Value: "---apple=o",
		},
		{
			Value: "---banana=t",
		},
		{
			Value: "---cherry=t",
		},



		{
			Value: "---a=one",
		},
		{
			Value: "---b=two",
		},
		{
			Value: "---c=three",
		},



		{
			Value: "---apple=one",
		},
		{
			Value: "---banana=two",
		},
		{
			Value: "---cherry=three",
		},



		{
			Value: "---a:",
		},
		{
			Value: "---b:",
		},
		{
			Value: "---c:",
		},



		{
			Value: "---apple:",
		},
		{
			Value: "---banana:",
		},
		{
			Value: "---cherry:",
		},



		{
			Value: "---a:o",
		},
		{
			Value: "---b:t",
		},
		{
			Value: "---c:t",
		},



		{
			Value: "---apple:o",
		},
		{
			Value: "---banana:t",
		},
		{
			Value: "---cherry:t",
		},



		{
			Value: "---a:one",
		},
		{
			Value: "---b:two",
		},
		{
			Value: "---c:three",
		},



		{
			Value: "---apple:one",
		},
		{
			Value: "---banana:two",
		},
		{
			Value: "---cherry:three",
		},









		{
			Value: "++a",
		},
		{
			Value: "++b",
		},
		{
			Value: "++c",
		},



		{
			Value: "++apple",
		},
		{
			Value: "++banana",
		},
		{
			Value: "++cherry",
		},




		{
			Value: "++a=",
		},
		{
			Value: "++b=",
		},
		{
			Value: "++c=",
		},



		{
			Value: "++apple=",
		},
		{
			Value: "++banana=",
		},
		{
			Value: "++cherry=",
		},



		{
			Value: "++a=o",
		},
		{
			Value: "++b=t",
		},
		{
			Value: "++c=t",
		},



		{
			Value: "++apple=o",
		},
		{
			Value: "++banana=t",
		},
		{
			Value: "++cherry=t",
		},



		{
			Value: "++a=one",
		},
		{
			Value: "++b=two",
		},
		{
			Value: "++c=three",
		},



		{
			Value: "++apple=one",
		},
		{
			Value: "++banana=two",
		},
		{
			Value: "++cherry=three",
		},



		{
			Value: "++a:",
		},
		{
			Value: "++b:",
		},
		{
			Value: "++c:",
		},



		{
			Value: "++apple:",
		},
		{
			Value: "++banana:",
		},
		{
			Value: "++cherry:",
		},



		{
			Value: "++a:o",
		},
		{
			Value: "++b:t",
		},
		{
			Value: "++c:t",
		},



		{
			Value: "++apple:o",
		},
		{
			Value: "++banana:t",
		},
		{
			Value: "++cherry:t",
		},



		{
			Value: "++a:one",
		},
		{
			Value: "++b:two",
		},
		{
			Value: "++c:three",
		},



		{
			Value: "++apple:one",
		},
		{
			Value: "++banana:two",
		},
		{
			Value: "++cherry:three",
		},



		{
			Value: `\-`,
		},
		{
			Value: `\--`,
		},
		{
			Value: `\---`,
		},
		{
			Value: `\----`,
		},



		{
			Value: `\+`,
		},
		{
			Value: `\++`,
		},
		{
			Value: `\+++`,
		},
		{
			Value: `\++++`,
		},









		{
			Value: `\-f`,
		},
		{
			Value: `\--f`,
		},
		{
			Value: `\---f`,
		},
		{
			Value: `\----f`,
		},



		{
			Value: `\+f`,
		},
		{
			Value: `\++f`,
		},
		{
			Value: `\+++f`,
		},
		{
			Value: `\++++f`,
		},









		{
			Value: `\-flag`,
		},
		{
			Value: `\--flag`,
		},
		{
			Value: `\---flag`,
		},
		{
			Value: `\----flag`,
		},



		{
			Value: `\+flag`,
		},
		{
			Value: `\++flag`,
		},
		{
			Value: `\+++flag`,
		},
		{
			Value: `\++++flag`,
		},









		{
			Value: `\-f=`,
		},
		{
			Value: `\--f=`,
		},
		{
			Value: `\---f=`,
		},
		{
			Value: `\----f=`,
		},



		{
			Value: `\+f=`,
		},
		{
			Value: `\++f=`,
		},
		{
			Value: `\+++f=`,
		},
		{
			Value: `\++++f=`,
		},









		{
			Value: `\-flag=`,
		},
		{
			Value: `\--flag=`,
		},
		{
			Value: `\---flag=`,
		},
		{
			Value: `\----flag=`,
		},



		{
			Value: `\+flag=`,
		},
		{
			Value: `\++flag=`,
		},
		{
			Value: `\+++flag=`,
		},
		{
			Value: `\++++flag=`,
		},









		{
			Value: `\-f:`,
		},
		{
			Value: `\--f:`,
		},
		{
			Value: `\---f:`,
		},
		{
			Value: `\----f:`,
		},



		{
			Value: `\+f:`,
		},
		{
			Value: `\++f:`,
		},
		{
			Value: `\+++f:`,
		},
		{
			Value: `\++++f:`,
		},









		{
			Value: `\-flag:`,
		},
		{
			Value: `\--flag:`,
		},
		{
			Value: `\---flag:`,
		},
		{
			Value: `\----flag:`,
		},



		{
			Value: `\+flag:`,
		},
		{
			Value: `\++flag:`,
		},
		{
			Value: `\+++flag:`,
		},
		{
			Value: `\++++flag:`,
		},
	}

	for testNumber, test := range tests {

		if expected, actual := false, isFlag(test.Value); expected != actual {
			t.Errorf("For test #%d, expected %t, but actually got %t.", testNumber, expected, actual)
			t.Errorf("\tVALUE: %q", test.Value)
			continue
		}
	}
}
