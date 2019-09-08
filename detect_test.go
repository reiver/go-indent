package indent_test

import (
	"github.com/reiver/go-indent"

	"io"
	"strings"

	"testing"
)

func TestDetect(t *testing.T) {

	tests := []struct{
		Src      string
		Expected string
	}{
		{
			Src:      "\tapple banana cherry",
			Expected: "\t",
		},
		{
			Src:      "\t\tapple banana cherry",
			Expected: "\t\t",
		},
		{
			Src:      "\t\t\tapple banana cherry",
			Expected: "\t\t\t",
		},
		{
			Src:      "\t\t\t\tapple banana cherry",
			Expected: "\t\t\t\t",
		},
		{
			Src:      "\t\t\t\t\tapple banana cherry",
			Expected: "\t\t\t\t\t",
		},



		{
			Src:      " apple banana cherry",
			Expected: " ",
		},
		{
			Src:      "  apple banana cherry",
			Expected: "  ",
		},
		{
			Src:      "   apple banana cherry",
			Expected: "   ",
		},
		{
			Src:      "    apple banana cherry",
			Expected: "    ",
		},
		{
			Src:      "     apple banana cherry",
			Expected: "     ",
		},



		{
			Src:      "\t apple banana cherry",
			Expected: "\t ",
		},
		{
			Src:      "\t \tapple banana cherry",
			Expected: "\t \t",
		},
		{
			Src:      "\t \t apple banana cherry",
			Expected: "\t \t ",
		},
		{
			Src:      "\t \t \tapple banana cherry",
			Expected: "\t \t \t",
		},
		{
			Src:      "\t \t \t apple banana cherry",
			Expected: "\t \t \t ",
		},
		{
			Src:      "\t \t \t \tapple banana cherry",
			Expected: "\t \t \t \t",
		},
		{
			Src:      "\t \t \t \t apple banana cherry",
			Expected: "\t \t \t \t ",
		},
		{
			Src:      "\t \t \t \t \tapple banana cherry",
			Expected: "\t \t \t \t \t",
		},



		{
			Src:      " \tapple banana cherry",
			Expected: " \t",
		},
		{
			Src:      " \t apple banana cherry",
			Expected: " \t ",
		},
		{
			Src:      " \t \tapple banana cherry",
			Expected: " \t \t",
		},
		{
			Src:      " \t \t apple banana cherry",
			Expected: " \t \t ",
		},
		{
			Src:      " \t \t \tapple banana cherry",
			Expected: " \t \t \t",
		},
		{
			Src:      " \t \t \t apple banana cherry",
			Expected: " \t \t \t ",
		},
		{
			Src:      " \t \t \t \tapple banana cherry",
			Expected: " \t \t \t \t",
		},
		{
			Src:      " \t \t \t \t apple banana cherry",
			Expected: " \t \t \t \t ",
		},



		{
			Src:      "  \t\t apple banana cherry",
			Expected: "  \t\t ",
		},
	}

	for testNumber, test := range tests {

		// dst: io.Writer
		// src: string
		{
			var dst strings.Builder

			err := indent.Detect(&dst, test.Src)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
				t.Logf("ERROR TYPE: %T", err)
				t.Logf("ERROR: %q", err)
				continue
			}

			if expected, actual := test.Expected, dst.String(); expected != actual {
				t.Errorf("For test #%d, did not actually get what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL: %q", actual)
				continue
			}
		}

		// dst: []byte
		// src: string
		{
			var dst []byte = make([]byte, len(test.Expected))

			err := indent.Detect(dst, test.Src)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
				t.Logf("ERROR TYPE: %T", err)
				t.Logf("ERROR: %q", err)
				continue
			}

			if expected, actual := test.Expected, string(dst); expected != actual {
				t.Errorf("For test #%d, did not actually get what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL: %q", actual)
				continue
			}
		}



		// dst: io.Writer
		// src: []byte
		{
			var dst strings.Builder
			var src []byte = []byte(test.Src)

			err := indent.Detect(&dst, src)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
				t.Logf("ERROR TYPE: %T", err)
				t.Logf("ERROR: %q", err)
				continue
			}

			if expected, actual := test.Expected, dst.String(); expected != actual {
				t.Errorf("For test #%d, did not actually get what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL: %q", actual)
				continue
			}
		}

		// dst: []byte
		// src: []byte
		{
			var dst []byte = make([]byte, len(test.Expected))
			var src []byte = []byte(test.Src)

			err := indent.Detect(dst, src)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
				t.Logf("ERROR TYPE: %T", err)
				t.Logf("ERROR: %q", err)
				continue
			}

			if expected, actual := test.Expected, string(dst); expected != actual {
				t.Errorf("For test #%d, did not actually get what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL: %q", actual)
				continue
			}
		}



		// dst: io.Writer
		// src: io.ReadSeeker
		{
			var dst strings.Builder
			var src io.ReadSeeker = strings.NewReader(test.Src)

			err := indent.Detect(&dst, src)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
				t.Logf("ERROR TYPE: %T", err)
				t.Logf("ERROR: %q", err)
				continue
			}

			if expected, actual := test.Expected, dst.String(); expected != actual {
				t.Errorf("For test #%d, did not actually get what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL: %q", actual)
				continue
			}
		}

		// dst: []byte
		// src: io.ReadSeeker
		{
			var dst []byte = make([]byte, len(test.Expected))
			var src io.ReadSeeker = strings.NewReader(test.Src)

			err := indent.Detect(dst, src)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
				t.Logf("ERROR TYPE: %T", err)
				t.Logf("ERROR: %q", err)
				continue
			}

			if expected, actual := test.Expected, string(dst); expected != actual {
				t.Errorf("For test #%d, did not actually get what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL: %q", actual)
				continue
			}
		}
	}
}
