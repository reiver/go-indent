# go-indent

Package indent provides tools for indentation of UTF-8 text, for the Go programming language.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-indent

[![GoDoc](https://godoc.org/github.com/reiver/go-indent?status.svg)](https://godoc.org/github.com/reiver/go-indent)


## Example
```go
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
```
