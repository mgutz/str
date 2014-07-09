package str

//import "testing"
import "fmt"

func ExampleBetween() {
	eg(1, Between("<a>foo</a>", "<a>", "</a>"))
	eg(2, Between("<a>foo</a></a>", "<a>", "</a>"))
	eg(3, Between("<a><a>foo</a></a>", "<a>", "</a>"))
	eg(4, Between("<a><a>foo</a></a>", "<a>", "</a>"))
	eg(5, Between("<a>foo", "<a>", "</a>"))
	eg(6, Between("Some strings } are very {weird}, dont you think?", "{", "}"))
	eg(7, Between("This is ateststring", "", "test"))
	eg(8, Between("This is ateststring", "test", ""))
	// Output:
	// 1: foo
	// 2: foo
	// 3: <a>foo
	// 4: <a>foo
	// 5:
	// 6: weird
	// 7: This is a
	// 8: string
}

func ExampleBetweenF() {
	eg(1, Pipe("abc", BetweenF("a", "c")))
	// Output:
	// 1: b
}

func ExampleCamelize() {
	eg(1, Camelize("data_rate"))
	eg(2, Camelize("background-color"))
	eg(3, Camelize("-moz-something"))
	eg(4, Camelize("_car_speed_"))
	eg(5, Camelize("yes_we_can"))
	// Output:
	// 1: dataRate
	// 2: backgroundColor
	// 3: MozSomething
	// 4: CarSpeed
	// 5: yesWeCan
}

func ExampleCharAt() {
	eg(1, CharAt("abc", 1))
	eg(2, "foo"+CharAt("", -1)+CharAt("", 0)+CharAt("", 10)+CharAt("abc", -1)+CharAt("abc", 10)+"bar")
	// Output:
	// 1: b
	// 2: foobar
}

func ExampleCharAtF() {
	eg(1, Pipe("abc", CharAtF(1)))
	// Output:
	// 1: b
}

func ExampleChompLeft() {
	eg(1, ChompLeft("foobar", "foo"))
	eg(2, ChompLeft("foobar", "bar"))
	eg(3, ChompLeft("", "foo"))
	eg(4, ChompLeft("", ""))
	// Output:
	// 1: bar
	// 2: foobar
	// 3:
	// 4:
}

func ExampleChompLeftF() {
	eg(1, Pipe("abc", ChompLeftF("ab")))
	// Output:
	// 1: c
}

func ExampleChompRight() {
	eg(1, ChompRight("foobar", "foo"))
	eg(2, ChompRight("foobar", "bar"))
	eg(3, ChompRight("", "foo"))
	eg(4, ChompRight("", ""))
	// Output:
	// 1: foobar
	// 2: foo
	// 3:
	// 4:
}

func ExampleChompRightF() {
	eg(1, Pipe("abc", ChompRightF("bc")))
	// Output:
	// 1: a
}

func ExampleClean() {
	eg(1, Clean("clean"))
	eg(2, Clean(""))
	eg(3, Clean(" please\t    clean \t \n  me "))
	// Output:
	// 1: clean
	// 2:
	// 3: please clean me
}

func ExampleDasherize() {
	eg(1, Dasherize("dataRate"))
	eg(2, Dasherize("CarSpeed"))
	eg(3, Dasherize("yesWeCan"))
	eg(4, Dasherize(""))
	eg(5, Dasherize("ABC"))
	// Output:
	// 1: data-rate
	// 2: car-speed
	// 3: yes-we-can
	// 4:
	// 5: a-b-c
}

func ExampleEnsurePrefix() {
	eg(1, EnsurePrefix("foobar", "foo"))
	eg(2, EnsurePrefix("bar", "foo"))
	eg(3, EnsurePrefix("", ""))
	eg(4, EnsurePrefix("foo", ""))
	eg(5, EnsurePrefix("", "foo"))
	// Output:
	// 1: foobar
	// 2: foobar
	// 3:
	// 4: foo
	// 5: foo
}

func ExampleEnsurePrefixF() {
	eg(1, Pipe("dir", EnsurePrefixF("./")))
	// Output:
	// 1: ./dir
}

func ExampleEnsureSuffix() {
	eg(1, EnsureSuffix("foobar", "bar"))
	eg(2, EnsureSuffix("foo", "bar"))
	eg(3, EnsureSuffix("", ""))
	eg(4, EnsureSuffix("foo", ""))
	eg(5, EnsureSuffix("", "bar"))
	// Output:
	// 1: foobar
	// 2: foobar
	// 3:
	// 4: foo
	// 5: bar
}

func ExampleIndexOf() {
	eg(1, IndexOf("abcdef", "a", 0))
	eg(2, IndexOf("abcdef", "a", 3))
	eg(3, IndexOf("abcdef", "a", -2))
	eg(4, IndexOf("abcdef", "a", 10))
	eg(5, IndexOf("", "a", 0))
	eg(6, IndexOf("abcdef", "", 2))
	eg(7, IndexOf("abcdef", "", 1000))
	// Output:
	// 1: 0
	// 2: -1
	// 3: -1
	// 4: -1
	// 5: -1
	// 6: 2
	// 7: 6
}

func ExampleIsAlpha() {
	eg(1, IsAlpha("afaf"))
	eg(2, IsAlpha("FJslfjkasfs"))
	eg(3, IsAlpha("áéúóúÁÉÍÓÚãõÃÕàèìòùÀÈÌÒÙâêîôûÂÊÎÔÛäëïöüÄËÏÖÜçÇ"))
	eg(4, IsAlpha("adflj43faljsdf"))
	eg(5, IsAlpha("33"))
	eg(6, IsAlpha("TT....TTTafafetstYY"))
	eg(7, IsAlpha("-áéúóúÁÉÍÓÚãõÃÕàèìòùÀÈÌÒÙâêîôûÂÊÎÔÛäëïöüÄËÏÖÜçÇ"))
	// Output:
	// 1: true
	// 2: true
	// 3: true
	// 4: false
	// 5: false
	// 6: false
	// 7: false
}

func eg(index int, example interface{}) {
	output := fmt.Sprintf("%d: %v", index, example)
	fmt.Printf("%s\n", Clean(output))
}

func ExampleIsAlphaNumeric() {
	eg(1, IsAlphaNumeric("afaf35353afaf"))
	eg(2, IsAlphaNumeric("FFFF99fff"))
	eg(3, IsAlphaNumeric("99"))
	eg(4, IsAlphaNumeric("afff"))
	eg(5, IsAlphaNumeric("Infinity"))
	eg(6, IsAlphaNumeric("áéúóúÁÉÍÓÚãõÃÕàèìòùÀÈÌÒÙâêîôûÂÊÎÔÛäëïöüÄËÏÖÜçÇ1234567890"))
	eg(7, IsAlphaNumeric("-Infinity"))
	eg(8, IsAlphaNumeric("-33"))
	eg(9, IsAlphaNumeric("aaff.."))
	eg(10, IsAlphaNumeric(".áéúóúÁÉÍÓÚãõÃÕàèìòùÀÈÌÒÙâêîôûÂÊÎÔÛäëïöüÄËÏÖÜçÇ1234567890"))
	// Output:
	// 1: true
	// 2: true
	// 3: true
	// 4: true
	// 5: true
	// 6: true
	// 7: false
	// 8: false
	// 9: false
	// 10: false
}

// describe('- isEmpty()', function() {
//  it('should return true if the string is solely composed of whitespace or is null', function() {
//     T (S(' ').isEmpty());
//     T (S('\t\t\t    ').isEmpty());
//     T (S('\n\n ').isEmpty());
//     F (S('hey').isEmpty())
//     T (S(null).isEmpty())
//     T (S(null).isEmpty())
//   })
// })

func ExampleIsLower() {
	eg(1, IsLower("a"))
	eg(2, IsLower("A"))
	eg(3, IsLower("abc"))
	eg(4, IsLower("aBc"))
	eg(5, IsLower("áéúóúãõàèìòùâêîôûäëïöüç"))
	eg(6, IsLower("hi jp"))
	eg(7, IsLower("ÁÉÍÓÚÃÕÀÈÌÒÙÂÊÎÔÛÄËÏÖÜÇ"))
	eg(8, IsLower("áéúóúãõàèìòùâêîôûäëïöüçÁ"))
	eg(9, IsLower("áéúóúãõàèìòùâêîôû äëïöüç"))
	// Output:
	// 1: true
	// 2: false
	// 3: true
	// 4: false
	// 5: true
	// 6: false
	// 7: false
	// 8: false
	// 9: false
}

func ExampleIsNumeric() {
	eg(1, IsNumeric("3"))
	eg(2, IsNumeric("34.22"))
	eg(3, IsNumeric("-22.33"))
	eg(4, IsNumeric("NaN"))
	eg(5, IsNumeric("Infinity"))
	eg(6, IsNumeric("-Infinity"))
	eg(7, IsNumeric("JP"))
	eg(8, IsNumeric("-5"))
	eg(9, IsNumeric("00099242424"))
	// Output:
	// 1: true
	// 2: false
	// 3: false
	// 4: false
	// 5: false
	// 6: false
	// 7: false
	// 8: false
	// 9: true
}

func ExampleIsUpper() {
	eg(1, IsUpper("a"))
	eg(2, IsUpper("A"))
	eg(3, IsUpper("ABC"))
	eg(4, IsUpper("aBc"))
	eg(5, IsUpper("áéúóúãõàèìòùâêîôûäëïöüç"))
	eg(6, IsUpper("HI JP"))
	eg(7, IsUpper("ÁÉÍÓÚÃÕÀÈÌÒÙÂÊÎÔÛÄËÏÖÜÇ"))
	eg(8, IsUpper("áéúóúãõàèìòùâêîôûäëïöüçÁ"))
	eg(9, IsUpper("ÁÉÍÓÚÃÕÀÈÌÒÙÂÊÎ ÔÛÄËÏÖÜÇ"))
	// Output:
	// 1: false
	// 2: true
	// 3: true
	// 4: false
	// 5: false
	// 6: false
	// 7: true
	// 8: false
	// 9: false
}

// describe('- isNumeric()', function() {
//   it("should return true if the string only contains digits, this would not include Infinity or -Infinity", function() {
//     T (S("3").isNumeric());
//     F (S("34.22").isNumeric());
//     F (S("-22.33").isNumeric());
//     F (S("NaN").isNumeric());
//     F (S("Infinity").isNumeric());
//     F (S("-Infinity").isNumeric());
//     F (S("JP").isNumeric());
//     F (S("-5").isNumeric());
//     T (S("000992424242").isNumeric());
//   })
// })
//
// describe('- isUpper()', function() {
//   it('should return true if the character or string is uppercase', function() {
//     F (S('a').isUpper());
//     F (S('z').isUpper());
//     T (S('B').isUpper());
//     T (S('HIJP').isUpper());
//     T (S('ÁÉÍÓÚÃÕÀÈÌÒÙÂÊÎÔÛÄËÏÖÜÇ').isUpper());
//     F (S('HI JP').isUpper());
//     F (S('HelLO').isUpper());
//     F (S('áéúóúãõàèìòùâêîôûäëïöüç').isUpper());
//     F (S('áéúóúãõàèìòùâêîôûäëïöüçÁ').isUpper());
//     F (S('ÁÉÍÓÚÃÕÀÈÌÒÙÂÊÎÔÛÄËÏÖÜÇá').isUpper());
//   })
// })
//

func ExampleMatch() {
	eg(1, Match("foobar", `^fo.*r$`))
	eg(2, Match("foobar", `^fo.*x$`))
	eg(3, Match("", `^fo.*x$`))
	// Output:
	// 1: true
	// 2: false
	// 3: false
}

func ExamplePipe() {
	eg(1, Pipe("\nabcdef   \n", Clean, BetweenF("a", "f"), ChompLeftF("bc")))
	// Output:
	// 1: de
}
