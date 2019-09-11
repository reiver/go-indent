package indent

import (
	"github.com/reiver/go-oi"
	"github.com/reiver/go-utf8s"

	"bytes"
	"fmt"
	"io"
	"strings"
)

// HasIndentation returns whether ‘value’ begins with the value of ‘indentation’.
//
// HasIndentation is somewhat similar to strings.HasPrefix(), and bytes.HasPrefix,
// except the prefix and value (using with HasIndentation) can be of more types,
// AND if ‘value’ is io.RuneScanner, then it reads (and discards) the indentation in ‘value’.
//
// ‘value’ can be a rune, or a string, or a []byte, or an io.ReaderAt, or an io.RuneScanner.
//
// ‘indentation’ can be a rune, or a string, or a []byte, or an io.ReaderAt.
func HasIndentation(value interface{}, indentation interface{}) (bool, error) {

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
			return false, fmt.Errorf("indent: indentation cannot be in type %T", indentation)
		}
	}

	var valueRuneScanner io.RuneScanner
	{
		switch casted := value.(type) {
		case io.RuneScanner:
			valueRuneScanner = casted
		case io.ReaderAt:
			valueRuneScanner = utf8s.NewRuneScanner(oi.ReadSeeker(casted))
		case []byte:
			valueRuneScanner = bytes.NewReader(casted)
		case string:
			valueRuneScanner = strings.NewReader(casted)
		case rune:
			valueRuneScanner = strings.NewReader(string(casted))
		default:
			return false, fmt.Errorf("indent: value cannot be in type %T", value)
		}
	}

	return hasIndentation(valueRuneScanner, indentationRuneScanner)
}

func hasIndentation(value io.RuneScanner, indentation io.RuneScanner) (bool, error) {

	for {
		rIndentation, _, err := indentation.ReadRune()
		if nil != err && io.EOF == err {
			return true, nil
		}
		if nil != err {
			return false, err
		}

		rValue, _, err := value.ReadRune()
		if nil != err && io.EOF == err {
			return false, nil
		}
		if nil != err {
			return false, err
		}

		if expected, actual := rIndentation, rValue; expected != actual {
			if err := value.UnreadRune(); nil != err {
				return false, err
			}
			return false, nil
		}
	}
}
