package indent

import (
	"github.com/reiver/go-utf8s"

	"io"
	"sync"
	"unicode/utf8"
)

type Writer struct {
	Indentation string
	Writer io.Writer

	mutex   sync.Mutex
	notPending bool
}

// Write makes indent.Writer fit the io.Writer interface.
//
// The returned integer returned from indent.Writer.Write() represents how many bytes of ‘bytes’ was written.
// I.e., it does not count the indentation that was written.
//
// I.e., so for this example:
//
//	var indenter indent.Writer = indent.Writer{
//		Indentation: "\t\t",
//		Writer:      writer,
//	}
//	
//	// ...
//	
//	var p []byte = []byte("Hello world!") // The length of this string is: 12
//
//	n, err := indenter.Write(p)
//
// If everything was successful, then ‘n’ would 12, not 14.
func (receiver *Writer) Write(bytes []byte) (int, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	var writer io.Writer = receiver.Writer
	if nil == writer {
		return 0, errNilWriter
	}

	if 0 == len(bytes) {
		return 0, nil
	}

	var n int

	var p []byte = bytes


	for 0 < len(p) {

		r, size := utf8.DecodeRune(p)

		p = p[size:]

		if !receiver.notPending {
			receiver.notPending = true

			if '\n' != r {
				_, err := receiver.writeIndent()
				if nil != err {
					return n, err
				}
			}
		}

		num, err := utf8s.WriteRune(writer, r)
		n += num
		if nil != err {
			return n, err
		}

		if '\n' == r {
			receiver.notPending = false
		}
	}

	return n, nil
}

func (receiver *Writer) writeIndent() (int, error) {
	if 0 >= len(receiver.Indentation) {
		return 0, nil
	}

	n, err := io.WriteString(receiver.Writer, receiver.Indentation)
	if nil != err {
		return n, err
	}

	return n, nil
}
