package indent_test

import (
	"github.com/reiver/go-indent"

	"bytes"
	"io"

	"testing"
)

func TestWriterWrite(t *testing.T) {

	tests := []struct{
		Indentation string
		Data [][]byte
		Expected string
	}{
		{
			Indentation: "",
			Data: [][]byte{
				[]byte("apple\nbanana\ncherry"),
			},
			Expected: "apple\nbanana\ncherry",
		},
		{
			Indentation: "\t",
			Data: [][]byte{
				[]byte("apple\nbanana\ncherry"),
			},
			Expected: "\t"+"apple\n"+"\t"+"banana\n"+"\t"+"cherry",
		},
		{
			Indentation: "\t\t",
			Data: [][]byte{
				[]byte("apple\nbanana\ncherry"),
			},
			Expected: "\t\t"+"apple\n"+"\t\t"+"banana\n"+"\t\t"+"cherry",
		},



		{
			Indentation: "",
			Data: [][]byte{
				[]byte("apple\nbanana\ncherry\n"),
			},
			Expected: "apple\nbanana\ncherry\n",
		},
		{
			Indentation: "\t",
			Data: [][]byte{
				[]byte("apple\nbanana\ncherry\n"),
			},
			Expected: "\t"+"apple\n"+"\t"+"banana\n"+"\t"+"cherry\n",
		},
		{
			Indentation: "\t\t",
			Data: [][]byte{
				[]byte("apple\nbanana\ncherry\n"),
			},
			Expected: "\t\t"+"apple\n"+"\t\t"+"banana\n"+"\t\t"+"cherry\n",
		},









		{
			Indentation: "",
			Data: [][]byte{
				[]byte("apple\n"),
				[]byte("banana\n"),
				[]byte("cherry"),
			},
			Expected: "apple\n"+"banana\n"+"cherry",
		},
		{
			Indentation: "\t",
			Data: [][]byte{
				[]byte("apple\n"),
				[]byte("banana\n"),
				[]byte("cherry"),
			},
			Expected: "\t"+"apple\n"+"\t"+"banana\n"+"\t"+"cherry",
		},
		{
			Indentation: "\t\t",
			Data: [][]byte{
				[]byte("apple\n"),
				[]byte("banana\n"),
				[]byte("cherry"),
			},
			Expected: "\t\t"+"apple\n"+"\t\t"+"banana\n"+"\t\t"+"cherry",
		},



		{
			Indentation: "",
			Data: [][]byte{
				[]byte("apple\n"),
				[]byte("banana\n"),
				[]byte("cherry\n"),
			},
			Expected: "apple\n"+"banana\n"+"cherry\n",
		},
		{
			Indentation: "\t",
			Data: [][]byte{
				[]byte("apple\n"),
				[]byte("banana\n"),
				[]byte("cherry\n"),
			},
			Expected: "\t"+"apple\n"+"\t"+"banana\n"+"\t"+"cherry\n",
		},
		{
			Indentation: "\t\t",
			Data: [][]byte{
				[]byte("apple\n"),
				[]byte("banana\n"),
				[]byte("cherry\n"),
			},
			Expected: "\t\t"+"apple\n"+"\t\t"+"banana\n"+"\t\t"+"cherry\n",
		},









		{
			Indentation: "",
			Data: [][]byte{
				[]byte("apple\nban"),
				[]byte("ana\ncherry"),
			},
			Expected: "apple\nban"+"ana\ncherry",
		},
		{
			Indentation: "\t",
			Data: [][]byte{
				[]byte("apple\nban"),
				[]byte("ana\ncherry"),
			},
			Expected: "\tapple\n"+"\t"+"ban"+"ana\n"+"\t"+"cherry",
		},
		{
			Indentation: "\t\t",
			Data: [][]byte{
				[]byte("apple\nban"),
				[]byte("ana\ncherry"),
			},
			Expected: "\t\t"+"apple\n"+"\t\t"+"banana\n"+"\t\t"+"cherry",
		},



		{
			Indentation: "",
			Data: [][]byte{
				[]byte("apple\nban"),
				[]byte("ana\ncherry\n"),
			},
			Expected: "apple\nban"+"ana\ncherry\n",
		},
		{
			Indentation: "\t",
			Data: [][]byte{
				[]byte("apple\nban"),
				[]byte("ana\ncherry\n"),
			},
			Expected: "\tapple\n"+"\t"+"ban"+"ana\n"+"\t"+"cherry\n",
		},
		{
			Indentation: "\t\t",
			Data: [][]byte{
				[]byte("apple\nban"),
				[]byte("ana\ncherry\n"),
			},
			Expected: "\t\t"+"apple\n"+"\t\t"+"banana\n"+"\t\t"+"cherry\n",
		},









		{
			Indentation: "",
			Data: [][]byte{
				[]byte("apple\nban"),
				[]byte("ana\ncher"),
				[]byte("ry"),
			},
			Expected: "apple\nban"+"ana\ncher"+"ry",
		},
		{
			Indentation: "\t",
			Data: [][]byte{
				[]byte("apple\nban"),
				[]byte("ana\ncher"),
				[]byte("ry"),
			},
			Expected: "\t"+"apple\n"+"\t"+"ban"+"ana\n"+"\t"+"cher"+"ry",
		},
		{
			Indentation: "\t\t",
			Data: [][]byte{
				[]byte("apple\nban"),
				[]byte("ana\ncher"),
				[]byte("ry"),
			},
			Expected: "\t\t"+"apple\n"+"\t\t"+"ban"+"ana\n"+"\t\t"+"cher"+"ry",
		},



		{
			Indentation: "",
			Data: [][]byte{
				[]byte("apple\nban"),
				[]byte("ana\ncher"),
				[]byte("ry\n"),
			},
			Expected: "apple\nban"+"ana\ncher"+"ry\n",
		},
		{
			Indentation: "\t",
			Data: [][]byte{
				[]byte("apple\nban"),
				[]byte("ana\ncher"),
				[]byte("ry\n"),
			},
			Expected: "\t"+"apple\n"+"\t"+"ban"+"ana\n"+"\t"+"cher"+"ry\n",
		},
		{
			Indentation: "\t\t",
			Data: [][]byte{
				[]byte("apple\nban"),
				[]byte("ana\ncher"),
				[]byte("ry\n"),
			},
			Expected: "\t\t"+"apple\n"+"\t\t"+"ban"+"ana\n"+"\t\t"+"cher"+"ry\n",
		},









		{
			Indentation: "\t\t\t",
			Data: [][]byte{
				[]byte("name"),
				[]byte(" "),
				[]byte("{"),
				[]byte("\n"),

				[]byte("\t"),
				[]byte("one"),
				[]byte(":"),
				[]byte("\t"),
				[]byte("1st"),
				[]byte(";"),
				[]byte("\n"),

				[]byte("\t"),
				[]byte("two"),
				[]byte(":"),
				[]byte("\t"),
				[]byte("2nd"),
				[]byte(";"),
				[]byte("\n"),

				[]byte("}"),
				[]byte(""),
			},
			Expected: "\t\t\t"+"name {\n"+"\t\t\t"+"\tone:\t1st;\n"+"\t\t\t"+"\ttwo:\t2nd;\n"+"\t\t\t"+"}",
		},



		{
			Indentation: "\t\t\t",
			Data: [][]byte{
				[]byte("name"),
				[]byte(" "),
				[]byte("{"),
				[]byte("\n"),

				[]byte("\t"),
				[]byte("one"),
				[]byte(":"),
				[]byte("\t"),
				[]byte("1st"),
				[]byte(";"),
				[]byte("\n"),

				[]byte("\t"),
				[]byte("two"),
				[]byte(":"),
				[]byte("\t"),
				[]byte("2nd"),
				[]byte(";"),
				[]byte("\n"),

				[]byte("}"),
				[]byte("\n"),
			},
			Expected: "\t\t\t"+"name {\n"+"\t\t\t"+"\tone:\t1st;\n"+"\t\t\t"+"\ttwo:\t2nd;\n"+"\t\t\t"+"}\n",
		},









		{
			Indentation: "\t\t",
			Data: [][]byte{
				[]byte("apple"),
				[]byte(" "),
				[]byte("{"),
				[]byte("\n"),

				[]byte("\t"),
				[]byte("color"),
				[]byte(":"),
				[]byte("\t"),
				[]byte("red"),
				[]byte(";"),
				[]byte("\n"),

				[]byte("\t"),
				[]byte("size"),
				[]byte(":"),
				[]byte("\t"),
				[]byte("medium"),
				[]byte(";"),
				[]byte("\n"),

				[]byte("}"),
				[]byte("\n"),

				[]byte("\n"),

				[]byte("banana"),
				[]byte(" "),
				[]byte("{"),
				[]byte("\n"),

				[]byte("\t"),
				[]byte("color"),
				[]byte(":"),
				[]byte("\t"),
				[]byte("yellow"),
				[]byte(";"),
				[]byte("\n"),

				[]byte("\t"),
				[]byte("size"),
				[]byte(":"),
				[]byte("\t"),
				[]byte("medium"),
				[]byte(";"),
				[]byte("\n"),

				[]byte("}"),
				[]byte("\n"),

				[]byte("\n"),

				[]byte("cherry"),
				[]byte(" "),
				[]byte("{"),
				[]byte("\n"),

				[]byte("\t"),
				[]byte("color"),
				[]byte(":"),
				[]byte("\t"),
				[]byte("purple"),
				[]byte(";"),
				[]byte("\n"),

				[]byte("\t"),
				[]byte("size"),
				[]byte(":"),
				[]byte("\t"),
				[]byte("small"),
				[]byte(";"),
				[]byte("\n"),

				[]byte("}"),
				[]byte("\n"),
			},
			Expected:
`		apple {
			color:	red;
			size:	medium;
		}

		banana {
			color:	yellow;
			size:	medium;
		}

		cherry {
			color:	purple;
			size:	small;
		}
`,
		},



		{
			Indentation: "\t\t",
			Data: [][]byte{
				[]byte("\t"),

				[]byte("apple"),
				[]byte(" "),
				[]byte("{"),
				[]byte("\n"),

				[]byte("\t"),

				[]byte("\t"),
				[]byte("color"),
				[]byte(":"),
				[]byte("\t"),
				[]byte("red"),
				[]byte(";"),
				[]byte("\n"),

				[]byte("\t"),

				[]byte("\t"),
				[]byte("size"),
				[]byte(":"),
				[]byte("\t"),
				[]byte("medium"),
				[]byte(";"),
				[]byte("\n"),

				[]byte("\t"),

				[]byte("}"),
				[]byte("\n"),

				[]byte("\n"),

				[]byte("\t"),

				[]byte("banana"),
				[]byte(" "),
				[]byte("{"),
				[]byte("\n"),

				[]byte("\t"),

				[]byte("\t"),
				[]byte("color"),
				[]byte(":"),
				[]byte("\t"),
				[]byte("yellow"),
				[]byte(";"),
				[]byte("\n"),

				[]byte("\t"),

				[]byte("\t"),
				[]byte("size"),
				[]byte(":"),
				[]byte("\t"),
				[]byte("medium"),
				[]byte(";"),
				[]byte("\n"),

				[]byte("\t"),

				[]byte("}"),
				[]byte("\n"),

				[]byte("\n"),

				[]byte("\t"),

				[]byte("cherry"),
				[]byte(" "),
				[]byte("{"),
				[]byte("\n"),

				[]byte("\t"),

				[]byte("\t"),
				[]byte("color"),
				[]byte(":"),
				[]byte("\t"),
				[]byte("purple"),
				[]byte(";"),
				[]byte("\n"),

				[]byte("\t"),

				[]byte("\t"),
				[]byte("size"),
				[]byte(":"),
				[]byte("\t"),
				[]byte("small"),
				[]byte(";"),
				[]byte("\n"),

				[]byte("\t"),

				[]byte("}"),
				[]byte("\n"),
			},
			Expected:
`			apple {
				color:	red;
				size:	medium;
			}

			banana {
				color:	yellow;
				size:	medium;
			}

			cherry {
				color:	purple;
				size:	small;
			}
`,
		},



		{
			Indentation: "\t\t",
			Data: [][]byte{
				[]byte("\t\t"),

				[]byte("apple"),
				[]byte(" "),
				[]byte("{"),
				[]byte("\n"),

				[]byte("\t\t"),

				[]byte("\t"),
				[]byte("color"),
				[]byte(":"),
				[]byte("\t"),
				[]byte("red"),
				[]byte(";"),
				[]byte("\n"),

				[]byte("\t\t"),

				[]byte("\t"),
				[]byte("size"),
				[]byte(":"),
				[]byte("\t"),
				[]byte("medium"),
				[]byte(";"),
				[]byte("\n"),

				[]byte("\t\t"),

				[]byte("}"),
				[]byte("\n"),

				[]byte("\n"),

				[]byte("\t\t"),

				[]byte("banana"),
				[]byte(" "),
				[]byte("{"),
				[]byte("\n"),

				[]byte("\t\t"),

				[]byte("\t"),
				[]byte("color"),
				[]byte(":"),
				[]byte("\t"),
				[]byte("yellow"),
				[]byte(";"),
				[]byte("\n"),

				[]byte("\t\t"),

				[]byte("\t"),
				[]byte("size"),
				[]byte(":"),
				[]byte("\t"),
				[]byte("medium"),
				[]byte(";"),
				[]byte("\n"),

				[]byte("\t\t"),

				[]byte("}"),
				[]byte("\n"),

				[]byte("\n"),

				[]byte("\t\t"),

				[]byte("cherry"),
				[]byte(" "),
				[]byte("{"),
				[]byte("\n"),

				[]byte("\t\t"),

				[]byte("\t"),
				[]byte("color"),
				[]byte(":"),
				[]byte("\t"),
				[]byte("purple"),
				[]byte(";"),
				[]byte("\n"),

				[]byte("\t\t"),

				[]byte("\t"),
				[]byte("size"),
				[]byte(":"),
				[]byte("\t"),
				[]byte("small"),
				[]byte(";"),
				[]byte("\n"),

				[]byte("\t\t"),

				[]byte("}"),
				[]byte("\n"),
			},
			Expected:
`				apple {
					color:	red;
					size:	medium;
				}

				banana {
					color:	yellow;
					size:	medium;
				}

				cherry {
					color:	purple;
					size:	small;
				}
`,
		},









		{
			Indentation: "",
			Data: [][]byte{
				[]byte(
`apple {
	size:	3in 4in;
}

/*
 * This is a comment.
 */
apple/banana {
	one:	once   1st;
	two:	twice  2nd;
	three:	thrice 3rd;
	four:	fource 4th;
}

apple/banana/cherry {}

one/two/three[20<num][num<50] {
    title: "space man";
}
`),
			},
			Expected:
`apple {
	size:	3in 4in;
}

/*
 * This is a comment.
 */
apple/banana {
	one:	once   1st;
	two:	twice  2nd;
	three:	thrice 3rd;
	four:	fource 4th;
}

apple/banana/cherry {}

one/two/three[20<num][num<50] {
    title: "space man";
}
`,
		},



		{
			Indentation: "\t",
			Data: [][]byte{
				[]byte(
`apple {
	size:	3in 4in;
}

/*
 * This is a comment.
 */
apple/banana {
	one:	once   1st;
	two:	twice  2nd;
	three:	thrice 3rd;
	four:	fource 4th;
}

apple/banana/cherry {}

one/two/three[20<num][num<50] {
    title: "space man";
}
`),
			},
			Expected:
`	apple {
		size:	3in 4in;
	}

	/*
	 * This is a comment.
	 */
	apple/banana {
		one:	once   1st;
		two:	twice  2nd;
		three:	thrice 3rd;
		four:	fource 4th;
	}

	apple/banana/cherry {}

	one/two/three[20<num][num<50] {
	    title: "space man";
	}
`,
		},



		{
			Indentation: "\t\t",
			Data: [][]byte{
				[]byte(
`apple {
	size:	3in 4in;
}

/*
 * This is a comment.
 */
apple/banana {
	one:	once   1st;
	two:	twice  2nd;
	three:	thrice 3rd;
	four:	fource 4th;
}

apple/banana/cherry {}

one/two/three[20<num][num<50] {
    title: "space man";
}
`),
			},
			Expected:
`		apple {
			size:	3in 4in;
		}

		/*
		 * This is a comment.
		 */
		apple/banana {
			one:	once   1st;
			two:	twice  2nd;
			three:	thrice 3rd;
			four:	fource 4th;
		}

		apple/banana/cherry {}

		one/two/three[20<num][num<50] {
		    title: "space man";
		}
`,
		},
	}

	TestLoop: for testNumber, test := range tests {

		var buffer bytes.Buffer

		var writer io.Writer = &indent.Writer{
			Indentation: test.Indentation,
			Writer: &buffer,
		}

		var num int

		for datumNumber, datum := range test.Data {

			n, err := writer.Write(datum)
			if nil != err {
				t.Errorf("For test #%d and datum #%d, did not expect an error, but actually got one:(%T) %q", testNumber, datumNumber, err, err)
				t.Errorf("\tINDENTATION: %q", test.Indentation)
				t.Errorf("\tDATA: ...")
				for dN, d := range test.Data {
					t.Errorf("\t\t[%d] %q", dN, string(d))
				}
				continue TestLoop
			}
			num += n

		}

		if expected, actual := test.Expected, buffer.String(); expected != actual {
			t.Errorf("For test #%d ...", testNumber)
			t.Errorf("\tINDENTATION: %q", test.Indentation)
			t.Errorf("\tDATA: ...")
			for dN, d := range test.Data {
				t.Errorf("\t\t[%d] %q", dN, string(d))
			}
			t.Errorf("\tEXPECTED: %q", expected)
			t.Errorf("\tACTUAL:   %q", actual)
			continue
		}

		{
			var expected int
			for _, datum := range test.Data {
				expected += len(datum)
			}

			if actual := num; expected != actual {
				t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
				t.Errorf("\tINDENTATION: %q", test.Indentation)
				t.Errorf("\tDATA: ...")
				for dN, d := range test.Data {
					t.Errorf("\t\t[%d] %q", dN, string(d))
				}
				t.Errorf("\tEXPECTED: %q", expected)
				t.Errorf("\tACTUAL:   %q", actual)
				continue
			}
		}
	}
}
