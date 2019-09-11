
package indent_test

import (
	"github.com/reiver/go-indent"

	"math/rand"
	"strings"
	"time"

	"testing"
)

func TestHasIndentation_error(t *testing.T) {

	tests := []struct{
		Indentation interface{}
		Src interface{}

		IndentationString string
		SrcString string
	}{}

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	runes := []rune{
		'0',
		'1',
		'2',
		'3',
		'4',
		'5',
		'6',
		'7',
		'8',
		'9',



		'A',
		'B',
		'C',
		'D',
		'E',
		'F',
		'G',
		'H',
		'I',
		'J',
		'K',
		'L',
		'M',
		'N',
		'O',
		'P',
		'Q',
		'R',
		'S',
		'T',
		'U',
		'V',
		'W',
		'X',
		'Y',
		'Z',



		'a',
		'b',
		'c',
		'd',
		'e',
		'f',
		'g',
		'h',
		'i',
		'j',
		'k',
		'l',
		'm',
		'n',
		'o',
		'p',
		'q',
		'r',
		's',
		't',
		'u',
		'v',
		'w',
		'x',
		'y',
		'z',



		'۰',
		'۱',
		'۲',
		'۳',
		'۴',
		'۵',
		'۶',
		'۷',
		'۸',
		'۹',



		'ا',
		'ب',
		'پ',
		'ت',
		'ث',
		'ج',
		'چ',
		'ح',
		'خ',
		'د',
		'ذ',
		'ر',
		'ز',
		'ژ',
		'س',
		'ش',
		'ص',
		'ض',
		'ط',
		'ظ',
		'ع',
		'غ',
		'ف',
		'ق',
		'ک',
		'گ',
		'ل',
		'م',
		'ن',
		'و',
		'ه',
		'ی',



		'ㄱ',
		'ㄲ',
		'ㄳ',
		'ㄴ',
		'ㄵ',
		'ㄶ',
		'ㄷ',
		'ㄹ',
		'ㄺ',
		'ㄻ',
		'ㄼ',
		'ㄽ',
		'ㄾ',
		'ㄿ',
		'ㅀ',
		'ㅁ',
		'ㅂ',
		'ㅄ',
		'ㅅ',
		'ㅆ',
		'ㅇ',
		'ㅈ',
		'ㅊ',
		'ㅋ',
		'ㅌ',
		'ㅍ',
		'ㅎ',



		'👾',
		'👻',

		'😀',
		'😁',
		'😂',
		'😃',
		'😄',
		'😅',
		'😆',
		'😇',
		'😈',
		'😉',
		'😊',
		'😋',
		'😌',
		'😍',
		'😎',
		'😏',

		'😐',
		'😑',
		'😒',
		'😓',
		'😔',
		'😕',
		'😖',
		'😗',
		'😘',
		'😙',
		'😚',
		'😛',
		'😜',
		'😝',
		'😞',
		'😟',

		'😠',
		'😡',
		'😢',
		'😣',
		'😤',
		'😥',
		'😦',
		'😧',
		'😨',
		'😩',
		'😪',
		'😫',
		'😬',
		'😭',
		'😮',
		'😯',

		'😰',
		'😱',
		'😲',
		'😳',
		'😴',
		'😵',
		'😶',
		'😷',

		'🙁',
		'🙂',
		'🙃',
		'🙄',
	}

	indentations := []string{
		"\t",
		"\t\t",
		"\t\t\t",
		"\t\t\t\t",
		"\t\t\t\t\t",



		" ",
		"  ",
		"   ",
		"    ",
		"     ",



		"-",
		"--",
		"---",
		"----",
		"-----",



		" \t",
		"\t ",



		"    \t\t",
		"\t\t    ",

	}

	for _, indentation := range indentations {
		const limit = 20

		for i:=0; i<limit; i++ {

			var buffer strings.Builder
			{
				size := 1 + randomness.Intn(10)
				for i:=0; i<size; i++ {
					buffer.WriteRune( runes[randomness.Intn(len(runes))] )
				}
			}

			// indentation string
			// text        string
			{
				test := struct{
					Indentation interface{}
					Src interface{}

					IndentationString string
					SrcString string
				}{
					IndentationString: indentation,
					SrcString:         buffer.String(),
				}
				test.Indentation = test.IndentationString
				test.Src         = test.SrcString

				tests = append(tests, test)
			}

			// indentation string
			// text        []byte
			{
				test := struct{
					Indentation interface{}
					Src interface{}

					IndentationString string
					SrcString string
				}{
					IndentationString: indentation,
					SrcString:         buffer.String(),
				}
				test.Indentation = test.IndentationString
				test.Src         = []byte(test.SrcString)

				tests = append(tests, test)
			}

			// indentation string
			// text        io.ReaderAt
			{
				test := struct{
					Indentation interface{}
					Src interface{}

					IndentationString string
					SrcString string
				}{
					IndentationString: indentation,
					SrcString:         buffer.String(),
				}
				test.Indentation = test.IndentationString
				test.Src         = strings.NewReader(test.SrcString)

				tests = append(tests, test)
			}

			// indentation []byte
			// text        string
			{
				test := struct{
					Indentation interface{}
					Src interface{}

					IndentationString string
					SrcString string
				}{
					IndentationString: indentation,
					SrcString:         buffer.String(),
				}
				test.Indentation = []byte(test.IndentationString)
				test.Src         = test.SrcString

				tests = append(tests, test)
			}

			// indentation []byte
			// text        []byte
			{
				test := struct{
					Indentation interface{}
					Src interface{}

					IndentationString string
					SrcString string
				}{
					IndentationString: indentation,
					SrcString:         buffer.String(),
				}
				test.Indentation = []byte(test.IndentationString)
				test.Src         = []byte(test.SrcString)

				tests = append(tests, test)
			}

			// indentation []byte
			// text        io.ReaderAt
			{
				test := struct{
					Indentation interface{}
					Src interface{}

					IndentationString string
					SrcString string
				}{
					IndentationString: indentation,
					SrcString:         buffer.String(),
				}
				test.Indentation = []byte(test.IndentationString)
				test.Src         = strings.NewReader(test.SrcString)

				tests = append(tests, test)
			}

			// indentation io.ReaderAt
			// text        string
			{
				test := struct{
					Indentation interface{}
					Src interface{}

					IndentationString string
					SrcString string
				}{
					IndentationString: indentation,
					SrcString:         buffer.String(),
				}
				test.Indentation = strings.NewReader(test.IndentationString)
				test.Src         = test.SrcString

				tests = append(tests, test)
			}

			// indentation io.ReaderAt
			// text        []byte
			{
				test := struct{
					Indentation interface{}
					Src interface{}

					IndentationString string
					SrcString string
				}{
					IndentationString: indentation,
					SrcString:         buffer.String(),
				}
				test.Indentation = strings.NewReader(test.IndentationString)
				test.Src         = []byte(test.SrcString)

				tests = append(tests, test)
			}

			// indentation io.ReaderAt
			// text        io.ReaderAt
			{
				test := struct{
					Indentation interface{}
					Src interface{}

					IndentationString string
					SrcString string
				}{
					IndentationString: indentation,
					SrcString:         buffer.String(),
				}
				test.Indentation = strings.NewReader(test.IndentationString)
				test.Src         = strings.NewReader(test.SrcString)

				tests = append(tests, test)
			}
		}
	}

	for testNumber, test := range tests {

		var dst strings.Builder
		var src interface{} = test.Src

		swallowed, err := indent.Swallow(&dst, src, test.Indentation)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
			t.Logf("INDENTATION: %q (string)", test.IndentationString)
			t.Logf("SRC:         %q (string)", test.SrcString)
			t.Logf("INDENTATION TYPE: %T", test.Indentation)
			t.Logf("SRC:        TYPE  %T", test.Src)
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %q", err)
		}
		if expected, actual := false, swallowed; expected != actual {
			t.Errorf("For test #%d, expected to be told it did not have the indentation, but it actually did.", testNumber)
			t.Logf("INDENTATION: %q (string)", test.IndentationString)
			t.Logf("SRC:         %q (string)", test.SrcString)
			t.Logf("DST:         %q", dst.String())
			t.Logf("INDENTATION TYPE: %T", test.Indentation)
			t.Logf("SRC:        TYPE  %T", test.Src)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
		}
		if expected, actual := test.IndentationString, dst.String(); expected == actual {
			t.Errorf("For test #%d, expected to be told it did not have the indentation, but it actually did.", testNumber)
			t.Logf("INDENTATION: %q (string)", test.IndentationString)
			t.Logf("SRC:         %q (string)", test.SrcString)
			t.Logf("INDENTATION TYPE: %T", test.Indentation)
			t.Logf("SRC:        TYPE  %T", test.Src)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
		}
	}

}

func TestHasIndentation(t *testing.T) {

	tests := []struct{
		Indentation interface{}
		Src interface{}

		IndentationString string
		SrcString string
	}{}

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	runes := []rune{
		'0',
		'1',
		'2',
		'3',
		'4',
		'5',
		'6',
		'7',
		'8',
		'9',



		'A',
		'B',
		'C',
		'D',
		'E',
		'F',
		'G',
		'H',
		'I',
		'J',
		'K',
		'L',
		'M',
		'N',
		'O',
		'P',
		'Q',
		'R',
		'S',
		'T',
		'U',
		'V',
		'W',
		'X',
		'Y',
		'Z',



		'a',
		'b',
		'c',
		'd',
		'e',
		'f',
		'g',
		'h',
		'i',
		'j',
		'k',
		'l',
		'm',
		'n',
		'o',
		'p',
		'q',
		'r',
		's',
		't',
		'u',
		'v',
		'w',
		'x',
		'y',
		'z',



		'۰',
		'۱',
		'۲',
		'۳',
		'۴',
		'۵',
		'۶',
		'۷',
		'۸',
		'۹',



		'ا',
		'ب',
		'پ',
		'ت',
		'ث',
		'ج',
		'چ',
		'ح',
		'خ',
		'د',
		'ذ',
		'ر',
		'ز',
		'ژ',
		'س',
		'ش',
		'ص',
		'ض',
		'ط',
		'ظ',
		'ع',
		'غ',
		'ف',
		'ق',
		'ک',
		'گ',
		'ل',
		'م',
		'ن',
		'و',
		'ه',
		'ی',



		'ㄱ',
		'ㄲ',
		'ㄳ',
		'ㄴ',
		'ㄵ',
		'ㄶ',
		'ㄷ',
		'ㄹ',
		'ㄺ',
		'ㄻ',
		'ㄼ',
		'ㄽ',
		'ㄾ',
		'ㄿ',
		'ㅀ',
		'ㅁ',
		'ㅂ',
		'ㅄ',
		'ㅅ',
		'ㅆ',
		'ㅇ',
		'ㅈ',
		'ㅊ',
		'ㅋ',
		'ㅌ',
		'ㅍ',
		'ㅎ',



		'👾',
		'👻',

		'😀',
		'😁',
		'😂',
		'😃',
		'😄',
		'😅',
		'😆',
		'😇',
		'😈',
		'😉',
		'😊',
		'😋',
		'😌',
		'😍',
		'😎',
		'😏',

		'😐',
		'😑',
		'😒',
		'😓',
		'😔',
		'😕',
		'😖',
		'😗',
		'😘',
		'😙',
		'😚',
		'😛',
		'😜',
		'😝',
		'😞',
		'😟',

		'😠',
		'😡',
		'😢',
		'😣',
		'😤',
		'😥',
		'😦',
		'😧',
		'😨',
		'😩',
		'😪',
		'😫',
		'😬',
		'😭',
		'😮',
		'😯',

		'😰',
		'😱',
		'😲',
		'😳',
		'😴',
		'😵',
		'😶',
		'😷',

		'🙁',
		'🙂',
		'🙃',
		'🙄',
	}

	indentations := []string{
		"",



		"\t",
		"\t\t",
		"\t\t\t",
		"\t\t\t\t",
		"\t\t\t\t\t",



		" ",
		"  ",
		"   ",
		"    ",
		"     ",



		"-",
		"--",
		"---",
		"----",
		"-----",



		" \t",
		"\t ",



		"    \t\t",
		"\t\t    ",

	}

	for _, indentation := range indentations {
		const limit = 20

		for i:=0; i<limit; i++ {

			var buffer strings.Builder
			{
				size := 1 + randomness.Intn(10)
				for i:=0; i<size; i++ {
					buffer.WriteRune( runes[randomness.Intn(len(runes))] )
				}
			}

			// indentation string
			// text        string
			{
				test := struct{
					Indentation interface{}
					Src interface{}

					IndentationString string
					SrcString string
				}{
					IndentationString: indentation,
					SrcString:         indentation + buffer.String(),
				}
				test.Indentation = test.IndentationString
				test.Src         = test.SrcString

				tests = append(tests, test)
			}

			// indentation string
			// text        []byte
			{
				test := struct{
					Indentation interface{}
					Src interface{}

					IndentationString string
					SrcString string
				}{
					IndentationString: indentation,
					SrcString:         indentation + buffer.String(),
				}
				test.Indentation = test.IndentationString
				test.Src         = []byte(test.SrcString)

				tests = append(tests, test)
			}

			// indentation string
			// text        io.ReaderAt
			{
				test := struct{
					Indentation interface{}
					Src interface{}

					IndentationString string
					SrcString string
				}{
					IndentationString: indentation,
					SrcString:         indentation + buffer.String(),
				}
				test.Indentation = test.IndentationString
				test.Src         = strings.NewReader(test.SrcString)

				tests = append(tests, test)
			}

			// indentation []byte
			// text        string
			{
				test := struct{
					Indentation interface{}
					Src interface{}

					IndentationString string
					SrcString string
				}{
					IndentationString: indentation,
					SrcString:         indentation + buffer.String(),
				}
				test.Indentation = []byte(test.IndentationString)
				test.Src         = test.SrcString

				tests = append(tests, test)
			}

			// indentation []byte
			// text        []byte
			{
				test := struct{
					Indentation interface{}
					Src interface{}

					IndentationString string
					SrcString string
				}{
					IndentationString: indentation,
					SrcString:         indentation + buffer.String(),
				}
				test.Indentation = []byte(test.IndentationString)
				test.Src         = []byte(test.SrcString)

				tests = append(tests, test)
			}

			// indentation []byte
			// text        io.ReaderAt
			{
				test := struct{
					Indentation interface{}
					Src interface{}

					IndentationString string
					SrcString string
				}{
					IndentationString: indentation,
					SrcString:         indentation + buffer.String(),
				}
				test.Indentation = []byte(test.IndentationString)
				test.Src         = strings.NewReader(test.SrcString)

				tests = append(tests, test)
			}

			// indentation io.ReaderAt
			// text        string
			{
				test := struct{
					Indentation interface{}
					Src interface{}

					IndentationString string
					SrcString string
				}{
					IndentationString: indentation,
					SrcString:         indentation + buffer.String(),
				}
				test.Indentation = strings.NewReader(test.IndentationString)
				test.Src         = test.SrcString

				tests = append(tests, test)
			}

			// indentation io.ReaderAt
			// text        []byte
			{
				test := struct{
					Indentation interface{}
					Src interface{}

					IndentationString string
					SrcString string
				}{
					IndentationString: indentation,
					SrcString:         indentation + buffer.String(),
				}
				test.Indentation = strings.NewReader(test.IndentationString)
				test.Src         = []byte(test.SrcString)

				tests = append(tests, test)
			}

			// indentation io.ReaderAt
			// text        io.ReaderAt
			{
				test := struct{
					Indentation interface{}
					Src interface{}

					IndentationString string
					SrcString string
				}{
					IndentationString: indentation,
					SrcString:         indentation + buffer.String(),
				}
				test.Indentation = strings.NewReader(test.IndentationString)
				test.Src         = strings.NewReader(test.SrcString)

				tests = append(tests, test)
			}
		}
	}

	for testNumber, test := range tests {

		var dst strings.Builder
		var src interface{} = test.Src

		swallowed, err := indent.Swallow(&dst, src, test.Indentation)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
			t.Logf("INDENTATION: %q (string)", test.IndentationString)
			t.Logf("SRC:         %q (string)", test.SrcString)
			t.Logf("INDENTATION TYPE: %T", test.Indentation)
			t.Logf("SRC:        TYPE  %T", test.Src)
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %q", err)
		}
		if expected, actual := true, swallowed; expected != actual {
			t.Errorf("For test #%d, expected to be told it indeed had the indentation, but it actually didn't.", testNumber)
			t.Logf("INDENTATION: %q (string)", test.IndentationString)
			t.Logf("SRC:         %q (string)", test.SrcString)
			t.Logf("DST:         %q", dst.String())
			t.Logf("INDENTATION TYPE: %T", test.Indentation)
			t.Logf("SRC:        TYPE  %T", test.Src)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
		}
		if expected, actual := test.IndentationString, dst.String(); expected != actual {
			t.Errorf("For test #%d, expected to be told it indeed had the indentation, but it actually didn't.", testNumber)
			t.Logf("INDENTATION: %q (string)", test.IndentationString)
			t.Logf("SRC:         %q (string)", test.SrcString)
			t.Logf("INDENTATION TYPE: %T", test.Indentation)
			t.Logf("SRC:        TYPE  %T", test.Src)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
		}
	}
}

