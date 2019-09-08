package indent_test

import (
	"github.com/reiver/go-indent"

	"fmt"
	"strings"
)

func ExampleDetect() {

	var src string =
`		if 5 = x then
			publish "Hello world!"
		end
`

	var dst strings.Builder

	err := indent.Detect(&dst, src)

	if nil != err {
		fmt.Printf("ERROR: something went wrong when trying to detect the indentation: %s\n", err)
		return
	}

	var indentation string = dst.String()

	fmt.Printf("Indentation: %q\n", indentation)

	// Output:
	// Indentation: "\t\t"
}
