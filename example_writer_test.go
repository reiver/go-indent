package indent_test

import (
	"github.com/reiver/go-indent"

	"bytes"
	"fmt"
	"io"
)

func ExampleWriter() {

	var buffer bytes.Buffer

	var writer io.Writer = &buffer


	var indentationWriter indent.Writer

	indentationWriter.Indentation = "\t\t"
	indentationWriter.Writer      = writer


	var s string =
	"This is a sentence.\n"+
	"This is another sentence.\n"+
	"\n"+
	"This is a new paragraph.\n"


	io.WriteString(&indentationWriter, s) // <---- Note we write to the ‘indentationWriter’, and not ‘writer’.


	fmt.Printf("%q", buffer.String())

	// Output:
	// "\t\tThis is a sentence.\n\t\tThis is another sentence.\n\n\t\tThis is a new paragraph.\n"
}
