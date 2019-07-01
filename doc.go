/*
Package indent provides tools for indentation of UTF-8 text, for the Go programming language.

For example, if you take some text, package indent can make it so it puts tabs, or spaces for indentation
in front of each non-empty line.

So, for example, this:

	.
	Hello world!

	I came; I saw; I conquered.

Can becomes this:

	.
	        Hello world!

	        I came; I saw; I conquered.
Or:
	.
	-----Hello world!

	-----I came; I saw; I conquered.

You can specify how much indentation you want, and what you want the indentation to be
(ex: tabs, spaces, or any other symbol, or string).


Example

Here is how one might use indent.Writer:

	import "github.com/reiver/go-indent"
	
	// ...
	
	var writer io.Writer
	
	// ...
	
	var indentationWriter indent.Writer
	
	indentationWriter.Indentation = "\t\t"
	indentationWriter.Writer      = writer
	
	// ...
	
	var s string =
	"This is a sentence.\n"+
	"This is another sentence.\n"+
	"\n"+
	"This is a new paragraph.\n"
	
	// ...
	
	io.WriteString(&indentationWriter, s) // <---- Note we write to the ‘indentationWriter’, and not ‘writer’.
	
	// Result will be:
	//
	// "\t\tThis is a sentence.\n"+
	// "\t\tThis is another sentence.\n"+
	// "\n"+
	// "\t\tThis is a new paragraph.\n"
	//
	// Note the "\t\t" indentation in front of each non-empty line.
	//
	// Note the empty line (i.e., "\n") did not get indentation.
*/
package indent
