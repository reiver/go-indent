package indent

import (
	"github.com/reiver/go-oi"
	"github.com/reiver/go-utf8s"

	"bytes"
	"fmt"
	"io"
	"strings"
)

// Swallow returns whether ‘src’ begins with the value of ‘indentation’.
// Swallow also writes whatever it reads that matches the indentation to ‘dst’.
//
// Swallow is somewhat similar to strings.HasPrefix(), and bytes.HasPrefix.
//
// ‘src’ can be a rune, or a string, or a []byte, or an io.ReaderAt, or an io.RuneScanner.
//
// ‘indentation’ can be a rune, or a string, or a []byte, or an io.ReaderAt.
func Swallow(dst oi.RuneWriter, src interface{}, indentation interface{}) error {

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
			return fmt.Errorf("indent: indentation cannot be in type %T", indentation)
		}
	}

	var srcRuneScanner io.RuneScanner
	{
		switch casted := src.(type) {
		case io.RuneScanner:
			srcRuneScanner = casted
		case io.ReaderAt:
			srcRuneScanner = utf8s.NewRuneScanner(oi.ReadSeeker(casted))
		case []byte:
			srcRuneScanner = bytes.NewReader(casted)
		case string:
			srcRuneScanner = strings.NewReader(casted)
		case rune:
			srcRuneScanner = strings.NewReader(string(casted))
		default:
			return fmt.Errorf("indent: source cannot be in type %T", src)
		}
	}

	return swallow(dst, srcRuneScanner, indentationRuneScanner)
}

func swallow(dst oi.RuneWriter, src io.RuneScanner, indentation io.RuneScanner) error {

	for {
		rIndentation, _, err := indentation.ReadRune()
		if nil != err && io.EOF == err {
			return nil
		}
		if nil != err {
			return err
		}

		rSrc, _, err := src.ReadRune()
		if nil != err && io.EOF == err {
			return nil
		}
		if nil != err {
			return err
		}

		if expected, actual := rIndentation, rSrc; expected != actual {
			if err := src.UnreadRune(); nil != err {
				return err
			}
			return nil
		}

		if _, err := dst.WriteRune(rSrc); nil != err {
			return err
		}
	}
}
