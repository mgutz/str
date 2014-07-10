// Credits:
// https://github.com/jprichardson/string.js
// https://github.com/anbinh/dna/blob/master/string.go

// Package str is a set of useful string helpers which can be composed
// into a pipeline.
package str

import (
	"fmt"
	"html"
	"log"
	"math"
	"regexp"
	"strings"
	"unicode/utf8"
)

// Verbose flag enables console output for those functions that have
// counterparts in Go's excellent stadard packages.
var Verbose = false
var templateOpen = "{{"
var templateClose = "}}"

var beginEndSpacesRe = regexp.MustCompile("^\\s+|\\s+$")
var camelizeRe = regexp.MustCompile(`(\-|_|\s)+(.)?`)
var camelizeRe2 = regexp.MustCompile(`(\-|_|\s)+`)
var capitalsRe = regexp.MustCompile("([A-Z])")
var dashSpaceRe = regexp.MustCompile(`[-\s]+`)
var dashesRe = regexp.MustCompile("-+")
var isAlphaNumericRe = regexp.MustCompile(`[^0-9a-z\xC0-\xFF]`)
var isAlphaRe = regexp.MustCompile(`[^a-z\xC0-\xFF]`)
var nWhitespaceRe = regexp.MustCompile(`\s+`)
var notDigitsRe = regexp.MustCompile(`[^0-9]`)
var slugifyRe = regexp.MustCompile(`[^\w\s\-]`)
var spaceUnderscoreRe = regexp.MustCompile("[_\\s]+")
var spacesRe = regexp.MustCompile("[\\s\\xA0]+")
var stripPuncRe = regexp.MustCompile(`[^\w\s]|_`)
var templateRe = regexp.MustCompile(`([\-\[\]()*\s])`)
var templateRe2 = regexp.MustCompile(`\$`)
var underscoreRe = regexp.MustCompile(`([a-z\d])([A-Z]+)`)
var whitespaceRe = regexp.MustCompile(`^[\s\xa0]*$`)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Between extracts a string between left and right strings.
func Between(s, left, right string) string {
	l := len(left)
	startPos := strings.Index(s, left)
	endPos := IndexOf(s, right, startPos+l)
	//log.Printf("%s: left %s right %s start %d end %d", s, left, right, startPos+l, endPos)
	if endPos < 0 {
		return ""
	} else if right == "" {
		return s[endPos:]
	} else {
		return s[startPos+l : endPos]
	}
}

// BetweenF is the filter form for Between.
func BetweenF(left, right string) func(string) string {
	return func(s string) string {
		return Between(s, left, right)
	}
}

// Camelize return new string which removes any underscores or dashes and convert a string into camel casing.
func Camelize(s string) string {
	return camelizeRe.ReplaceAllStringFunc(s, func(val string) string {
		val = strings.ToUpper(val)
		val = camelizeRe2.ReplaceAllString(val, "")
		return val
	})
}

// Capitalize uppercases the first char of s and lowercases the rest.
func Capitalize(s string) string {
	return strings.ToUpper(s[0:1]) + strings.ToLower(s[1:])
}

// CharAt returns a string from the character at the specified position.
func CharAt(s string, index int) string {
	l := len(s)
	shortcut := index < 0 || index > l-1 || l == 0
	if shortcut {
		return ""
	}
	return s[index : index+1]
}

// CharAtF is the filter form of CharAt.
func CharAtF(index int) func(string) string {
	return func(s string) string {
		return CharAt(s, index)
	}
}

// ChompLeft removes prefix at the start of a string.
func ChompLeft(s, prefix string) string {
	if strings.HasPrefix(s, prefix) {
		return s[len(prefix):]
	}
	return s
}

// ChompLeftF is the filter form of ChompLeft.
func ChompLeftF(prefix string) func(string) string {
	return func(s string) string {
		return ChompLeft(s, prefix)
	}
}

// ChompRight removes suffix from end of s.
func ChompRight(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		return s[:len(s)-len(suffix)]
	}
	return s
}

// ChompRightF is the filter form of ChompRight.
func ChompRightF(suffix string) func(string) string {
	return func(s string) string {
		return ChompRight(s, suffix)
	}
}

// Clean compresses all adjacent whitespace to a single space and trims s.
func Clean(s string) string {
	s = spacesRe.ReplaceAllString(s, " ")
	s = beginEndSpacesRe.ReplaceAllString(s, "")
	return s
}

// Dasherize  converts a camel cased string into a string delimited by dashes.
func Dasherize(s string) string {
	s = strings.TrimSpace(s)
	s = spaceUnderscoreRe.ReplaceAllString(s, "-")
	s = capitalsRe.ReplaceAllString(s, "-$1")
	s = dashesRe.ReplaceAllString(s, "-")
	s = strings.ToLower(s)
	return s
}

// EscapeHTML is alias for html.EscapeString.
func EscapeHTML(s string) string {
	if Verbose {
		fmt.Println("Use html.EscapeString instead of EscapeHTML")
	}
	return html.EscapeString(s)
}

// DecodeHTMLEntities decodes HTML entities into their proper string representation.
// DecodeHTMLEntities is an alias for html.UnescapeString
func DecodeHTMLEntities(s string) string {
	if Verbose {
		fmt.Println("Use html.UnescapeString instead of DecodeHTMLEntities")
	}
	return html.UnescapeString(s)
}

// EnsurePrefix ensures s starts with prefix.
func EnsurePrefix(s, prefix string) string {
	if strings.HasPrefix(s, prefix) {
		return s
	}
	return prefix + s
}

// EnsurePrefixF is the filter form of EnsurePrefix.
func EnsurePrefixF(prefix string) func(string) string {
	return func(s string) string {
		return EnsurePrefix(s, prefix)
	}
}

// EnsureSuffix ensures s ends with suffix.
func EnsureSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		return s
	}
	return s + suffix
}

// EnsureSuffixF is the filter version of EnsureSuffix.
func EnsureSuffixF(suffix string) func(string) string {
	return func(s string) string {
		return EnsureSuffix(s, suffix)
	}
}

// Humanize transforms s into a human friendly form.
func Humanize(s string) string {
	if s == "" {
		return s
	}
	s = Underscore(s)
	log.Printf("hum %s\n", s)

	var humanizeRe = regexp.MustCompile(`_id$`)
	s = humanizeRe.ReplaceAllString(s, "")
	log.Printf("hum %s\n", s)

	s = strings.Replace(s, "_", " ", -1)
	s = strings.TrimSpace(s)
	s = Capitalize(s)
	return s
}

// IndexOf finds the index of needle in s starting from start.
func IndexOf(s string, needle string, start int) int {
	l := len(s)
	if needle == "" {
		if start < 0 {
			return 0
		} else if start < l {
			return start
		} else {
			return l
		}
	}
	if start < 0 || start > l-1 {
		return -1
	}
	pos := strings.Index(s[start:], needle)
	if pos == -1 {
		return -1
	}
	return start + pos
}

// IsAlpha returns true if a string contains only letters from ASCII (a-z,A-Z). Other letters from other languages are not supported.
func IsAlpha(s string) bool {
	return !isAlphaRe.MatchString(strings.ToLower(s))
}

// IsAlphaNumeric returns true if a string contains letters and digits.
func IsAlphaNumeric(s string) bool {
	return !isAlphaNumericRe.MatchString(strings.ToLower(s))
}

// IsLower returns true if s comprised of all lower case characters.
func IsLower(s string) bool {
	return IsAlpha(s) && s == strings.ToLower(s)
}

// IsNumeric returns true if a string contains only digits from 0-9. Other digits not in Latin (such as Arabic) are not currently supported.
func IsNumeric(s string) bool {
	return !notDigitsRe.MatchString(s)
}

// IsUpper returns true if s contains all upper case chracters.
func IsUpper(s string) bool {
	return IsAlpha(s) && s == strings.ToUpper(s)
}

// IsEmpty returns true if the string is solely composed of whitespace.
func IsEmpty(s string) bool {
	if s == "" {
		return true
	}
	return whitespaceRe.MatchString(s)
}

// Left returns the left substring of length n.
func Left(s string, n int) string {
	if n < 0 {
		return Right(s, -n)
	}
	return Substr(s, 0, n)
}

// LeftF is the filter version of Left.
func LeftF(n int) func(string) string {
	return func(s string) string {
		return Left(s, n)
	}
}

// Lines convert windows newlines to unix newlines then convert to an Array of lines.
func Lines(s string) []string {
	s = strings.Replace(s, "\r\n", "\n", -1)
	return strings.Split(s, "\n")
}

// Match returns true if patterns matches the string
func Match(s, pattern string) bool {
	r := regexp.MustCompile(pattern)
	return r.MatchString(s)
}

// Pad pads string s on both sides until it has length of n.
func Pad(s, c string, n int) string {
	L := len(s)
	if L >= n {
		return s
	}
	n -= L

	left := strings.Repeat(c, int(math.Ceil(float64(n)/2)))
	right := strings.Repeat(c, int(math.Floor(float64(n)/2)))
	return left + s + right
}

// PadF is the filter version of Pad.
func PadF(c string, n int) func(string) string {
	return func(s string) string {
		return Pad(s, c, n)
	}
}

// PadLeft pads string s on left side until it has length of n.
func PadLeft(s, c string, n int) string {
	L := len(s)
	if L > n {
		return s
	}
	return strings.Repeat(c, (n-L)) + s
}

// PadLeftF is the filter version of PadLeft.
func PadLeftF(c string, n int) func(string) string {
	return func(s string) string {
		return PadLeft(s, c, n)
	}
}

// PadRight pads string s on right side until it has length of n.
func PadRight(s, c string, n int) string {
	L := len(s)
	if L > n {
		return s
	}
	return s + strings.Repeat(c, (n-L))
}

// PadRightF is the filter version of Padright
func PadRightF(c string, n int) func(string) string {
	return func(s string) string {
		return PadRight(s, c, n)
	}
}

// Pipe pipes s through one or more string filters.
func Pipe(s string, funcs ...func(string) string) string {
	for _, fn := range funcs {
		s = fn(s)
	}
	return s
}

// ReplaceF is the filter version of strings.Replace.
func ReplaceF(old, new string, n int) func(string) string {
	return func(s string) string {
		return strings.Replace(s, old, new, n)
	}
}

// ReplacePattern replaces string with regexp string.
// ReplacePattern returns a copy of src, replacing matches of the Regexp with the replacement string repl. Inside repl, $ signs are interpreted as in Expand, so for instance $1 represents the text of the first submatch.
func ReplacePattern(s, pattern, repl string) string {
	r := regexp.MustCompile(pattern)
	return r.ReplaceAllString(s, repl)
}

// ReplacePatternF is the filter version of ReplaceRegexp.
func ReplacePatternF(pattern, repl string) func(string) string {
	return func(s string) string {
		return ReplacePattern(s, pattern, repl)
	}
}

// Reverse a string
func Reverse(s string) string {
	cs := make([]rune, utf8.RuneCountInString(s))
	i := len(cs)
	for _, c := range s {
		i--
		cs[i] = c
	}
	return string(cs)
}

// Right returns the right substring of length n.
func Right(s string, n int) string {
	if n < 0 {
		return Left(s, -n)
	}
	return Substr(s, len(s)-n, n)
}

// RightF is the Filter version of Right.
func RightF(n int) func(string) string {
	return func(s string) string {
		return Right(s, n)
	}
}

// Slice slices a string. If end is negative then it is the from the end
// of the string.
func Slice(s string, start, end int) string {
	if end > -1 {
		return s[start:end]
	}
	L := len(s)
	if L+end > 0 {
		return s[start : L-end]
	}
	return s[start:]
}

// SliceF is the filter for Slice.
func SliceF(start, end int) func(string) string {
	return func(s string) string {
		return Slice(s, start, end)
	}
}

// Substr returns a substring of s starting at index of length n.
func Substr(s string, index int, n int) string {
	L := len(s)
	if index < 0 || index >= L || s == "" {
		return ""
	}
	end := index + n
	if end >= L {
		end = L
	}
	if end <= index {
		return ""
	}
	return s[index:end]
}

// SubstrF is the filter version of Substr.
func SubstrF(index, n int) func(string) string {
	return func(s string) string {
		return Substr(s, index, n)
	}
}

// Slugify converts s into a dasherized string suitable for URL segment.
func Slugify(s string) string {
	sl := slugifyRe.ReplaceAllString(s, "")
	sl = strings.ToLower(sl)
	sl = Dasherize(sl)
	return sl
}

// StripPunctuation strips puncation from string.
func StripPunctuation(s string) string {
	s = stripPuncRe.ReplaceAllString(s, "")
	s = nWhitespaceRe.ReplaceAllString(s, " ")
	return s
}

// StripTags strips all of the html tags or tags specified by the parameters
func StripTags(s string, tags ...string) string {
	if len(tags) == 0 {
		tags = append(tags, "")
	}
	for _, tag := range tags {
		stripTagsRe := regexp.MustCompile(`(?i)<\/?` + tag + `[^<>]*>`)
		s = stripTagsRe.ReplaceAllString(s, "")
	}
	return s
}

// TemplateWithDelimiters is string template with user-defineable opening and closing delimiters.
func TemplateWithDelimiters(s string, values map[string]interface{}, opening, closing string) string {
	openDelim := templateRe.ReplaceAllString(opening, "\\$1")
	openDelim = templateRe2.ReplaceAllString(openDelim, "\\$")
	closingDelim := templateRe.ReplaceAllString(closing, "\\$1")
	closingDelim = templateRe2.ReplaceAllString(closingDelim, "\\$")

	r := regexp.MustCompile(openDelim + `(.+?)` + closingDelim)
	matches := r.FindAllStringSubmatch(s, -1)
	for _, submatches := range matches {
		match := submatches[0]
		key := submatches[1]
		//log.Printf("match %s key %s\n", match, key)
		if values[key] != nil {
			v := fmt.Sprintf("%v", values[key])
			s = strings.Replace(s, match, v, -1)
		}
	}

	return s
}

// SetTemplateDelimiters sets the delimiters for Template function. Defaults to "{{" and "}}"
func SetTemplateDelimiters(opening, closing string) {
	templateOpen = opening
	templateClose = closing
}

// TemplateDelimiters is the getter for the opening and closing delimiters for Template.
func TemplateDelimiters() (opening string, closing string) {
	return templateOpen, templateClose
}

// Template is a string template which replaces template placeholders delimited
// by "{{" and "}}" with values from map. The global delimiters may be set with
// SetTemplateDelimiters.
func Template(s string, values map[string]interface{}) string {
	return TemplateWithDelimiters(s, values, templateOpen, templateClose)
}

// ToBool fuzzily converts truthy values.
func ToBool(s string) bool {
	s = strings.ToLower(s)
	return s == "true" || s == "yes" || s == "on" || s == "1"
}

// Truncate truncates the string, accounting for word placement and chars count
// adding a morestr (defaults to ellipsis)
func Truncate(s, morestr string, n int) string {
	L := len(s)
	if L <= n {
		return s
	}

	if morestr == "" {
		morestr = "..."
	}

	tmpl := func(c string) string {
		if strings.ToUpper(c) != strings.ToLower(c) {
			return "A"
		}
		return " "
	}
	template := s[0 : n+1]
	var truncateRe = regexp.MustCompile(`.(?=\W*\w*$)`)
	truncateRe.ReplaceAllStringFunc(template, tmpl) // 'Hello, world' -> 'HellAA AAAAA'
	var wwRe = regexp.MustCompile(`\w\w`)
	var whitespaceRe2 = regexp.MustCompile(`\s*\S+$`)
	if wwRe.MatchString(template[len(template)-2:]) {
		template = whitespaceRe2.ReplaceAllString(template, "")
	} else {
		template = strings.TrimRight(template, " \t\n")
	}

	if len(template+morestr) > L {
		return s
	}
	return s[0:len(template)] + morestr
}

//     truncate: function(length, pruneStr) { //from underscore.string, author: github.com/rwz
//       var str = this.s;
//
//       length = ~~length;
//       pruneStr = pruneStr || '...';
//
//       if (str.length <= length) return new this.constructor(str);
//
//       var tmpl = function(c){ return c.toUpperCase() !== c.toLowerCase() ? 'A' : ' '; },
//         template = str.slice(0, length+1).replace(/.(?=\W*\w*$)/g, tmpl); // 'Hello, world' -> 'HellAA AAAAA'
//
//       if (template.slice(template.length-2).match(/\w\w/))
//         template = template.replace(/\s*\S+$/, '');
//       else
//         template = new S(template.slice(0, template.length-1)).trimRight().s;
//
//       return (template+pruneStr).length > str.length ? new S(str) : new S(str.slice(0, template.length)+pruneStr);
//     },

// Underscore returns converted camel cased string into a string delimited by underscores.
func Underscore(s string) string {
	if s == "" {
		return ""
	}
	u := strings.TrimSpace(s)

	u = underscoreRe.ReplaceAllString(u, "${1}_$2")
	log.Printf("trim %s\n", u)
	u = dashSpaceRe.ReplaceAllString(u, "_")
	log.Printf("trim2 %s\n", u)
	u = strings.ToLower(u)
	log.Printf("trim3 %s\n", u)
	log.Printf("trimx %s\n", s[0:1])
	if IsUpper(s[0:1]) {
		return "_" + u
	}
	return u
}

// UnescapeHTML is an alias for html.UnescapeString.
func UnescapeHTML(s string) string {
	if Verbose {
		fmt.Println("Use html.UnescapeString instead of UnescapeHTML")
	}
	return html.UnescapeString(s)
}

// WrapHTML wraps s within HTML tag having attributes attrs.
func WrapHTML(s string, tag string, attrs map[string]string) string {
	escapeHTMLAttributeQuotes := func(v string) string {
		v = strings.Replace(v, "<", "&lt;", -1)
		v = strings.Replace(v, "&", "&amp;", -1)
		v = strings.Replace(v, "\"", "&quot;", -1)
		return v
	}
	if tag == "" {
		tag = "div"
	}
	el := "<" + tag
	for name, val := range attrs {
		el += " " + name + "=\"" + escapeHTMLAttributeQuotes(val) + "\""
	}
	el += ">" + s + "</" + tag + ">"
	return el
}

// WrapHTMLF is the filter version of WrapHTML
func WrapHTMLF(tag string, attrs map[string]string) func(string) string {
	return func(s string) string {
		return WrapHTML(s, tag, attrs)
	}
}
