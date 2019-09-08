package indent

import (
	"github.com/reiver/go-buffers"
	"github.com/reiver/go-oi"
	"github.com/reiver/go-utf8s"

	"bytes"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
)

// Detect detects the indentation in ‘src’ and stores what it detected in ‘dst’.
//
// ‘dst’ can be a io.ReadSeeker, io.ReaderAt, []byte, or a string.
//
// ‘src’ can be a io.Writer, or a []bte.
func Detect(dst interface{}, src interface{}) (error) {
	if nil == src {
		return errNilSource
	}
	if nil == dst {
		return errNilDestination
	}

	var readSeeker io.ReadSeeker
	{
		switch casted := src.(type) {
		case io.ReadSeeker:
			readSeeker = casted
		case io.ReaderAt:
			readSeeker = oi.ReadSeeker(casted)
		case string:
			readSeeker = strings.NewReader(casted)
		case []byte:
			readSeeker = bytes.NewReader(casted)
		default:
			return fmt.Errorf("indent: Unsupported Source Type: %T", src)
		}
	}
	if nil == readSeeker {
		return errInternalError
	}

	var writer io.Writer
	{
		switch casted := dst.(type) {
		case io.Writer:
			writer = casted
		case []byte:
			writer = buffers.NewWriter(casted)
		default:
			return fmt.Errorf("indent: Unsupported Destination Type: %T", dst)
		}
	}
	if nil == writer {
		return errInternalError
	}

	return detect(writer, readSeeker)
}


func detect(writer io.Writer, readSeeker io.ReadSeeker) error {
	if nil == readSeeker {
		return errNilSource
	}
	if nil == writer {
		return errNilDestination
	}

	var original int64
	{
		var err error

		original, err = readSeeker.Seek(0, io.SeekCurrent)
		if nil != err {
			return err
		}
	}

	var buffer []byte
	Loop: for {
		r, size, err := utf8s.ReadRune(readSeeker)
		if nil != err && io.EOF == err {
			readSeeker.Seek(int64(-size), io.SeekCurrent)
			return fmt.Errorf("indent: End Of File (io.EOF) received before seeing a non-indentation character; indentation so for: %q", buffer)
		}
		if nil != err {
			readSeeker.Seek(int64(-size), io.SeekCurrent)
			return err
		}
		if utf8.RuneError == r && 0 == size {
			break Loop
		}
		if utf8.RuneError == r {
			readSeeker.Seek(int64(-size), io.SeekCurrent)
			return errUTF8RuneError
		}

		switch r {
		case '\t',' ': // <--- These are the characters that a permitted to make up the indentation.
			// Nothing here.
		default:
			break Loop
		}

		buffer = append(buffer, string(r)...)
	}

	{
		offset, err := readSeeker.Seek(original, io.SeekStart)
		if nil != err {
			return err
		}
		if expected, actual := original, offset; expected != actual {
			return fmt.Errorf("indent: Internal Error: could not seek back to original offset of %d", original)
		}
	}

	{
		n, err := writer.Write(buffer)
		if nil != err {
			return err
		}
		if expected, actual := len(buffer), n; expected != actual {
			return fmt.Errorf("indent: Internal Error: expected to write %d bytes, but actually wrote %d bytes", expected, actual)
		}
	}

	return nil
}
