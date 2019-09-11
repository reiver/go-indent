package indent

import (
	"github.com/reiver/go-oi"
	"github.com/reiver/go-utf8s"

	"bytes"
	"io"
	"strings"
)

// HasIndentation returns whether ‘value’ begins with the value of ‘indentation’.
//
// HasIndentation is somewhat similar to strings.HasPrefix(), and bytes.HasPrefix,
// except the prefix and value (using with HasIndentation) can be of more types.
//
// ‘value’ can be a rune, or a string, or a []byte, or an io.ReaderAt.
//
// ‘indentation’ can be a rune, or a string, or a []byte, or an io.ReaderAt.
func HasIndentation(value interface{}, indentation interface{}) bool {

	var indentationRuneScanner io.RuneScanner
	{
		switch casted := indentation.(type) {
		case io.ReaderAt:
			indentationRuneScanner = utf8s.NewRuneScanner(oi.ReadSeeker(casted))
		case []byte:
			indentationRuneScanner = bytes.NewReader(casted)
		case string:
			indentationRuneScanner = strings.NewReader(casted)
		case rune:
			indentationRuneScanner = strings.NewReader(string(casted))
		default:
			return false
		}
	}

	var valueRuneScanner io.RuneScanner
	{
		switch casted := value.(type) {
		case io.ReaderAt:
			valueRuneScanner = utf8s.NewRuneScanner(oi.ReadSeeker(casted))
		case []byte:
			valueRuneScanner = bytes.NewReader(casted)
		case string:
			valueRuneScanner = strings.NewReader(casted)
		case rune:
			valueRuneScanner = strings.NewReader(string(casted))
		default:
			return false
		}
	}

	return hasIndentation(valueRuneScanner, indentationRuneScanner)
}

func hasIndentation(value io.RuneScanner, indentation io.RuneScanner) bool {

	for {
		rIndentation, _, err := indentation.ReadRune()
		if nil != err && io.EOF == err {
			return true
		}
		if nil != err {
			return false
		}

		rValue, _, err := value.ReadRune()
		if nil != err && io.EOF == err {
			return false
		}
		if nil != err {
			return false
		}

		if expected, actual := rIndentation, rValue; expected != actual {
			return false
		}
	}
}
