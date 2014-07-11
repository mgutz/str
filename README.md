# str
--
    import "github.com/mgutz/str"

Package str is a comprehensive set of string functions to build more Go
awesomeness. Str is a port of the JavaScript [string.js](http://stringjs.com)
including my contributions to the project.

Str does not duplicate functionality found in `strings` or `strconv`. Str may
add filter versions of functions found in those packages for use with Pipe.

Str is based on plain functions instead of object-based methods to be more
consistent with Go standard libraries.

    str.Between("<a>foo</a>", "<a>", "</a>") == "foo"

Str is designed to be pipelined.

    s := str.Pipe("\nabcdef\n", Clean, BetweenF("a", "f"), ChompLeftF("bc"))
    s == "de"

User-defined filters can be added to the pipeline by creating a function or
closure that returns a function with this signature

    func(string) string

## Usage

```go
var Verbose = false
```
Verbose flag enables console output for those functions that have counterparts
in Go's excellent stadard packages.

#### func  Between

```go
func Between(s, left, right string) string
```
Between extracts a string between left and right strings.

#### func  BetweenF

```go
func BetweenF(left, right string) func(string) string
```
BetweenF is the filter form for Between.

#### func  Camelize

```go
func Camelize(s string) string
```
Camelize return new string which removes any underscores or dashes and convert a
string into camel casing.

#### func  Capitalize

```go
func Capitalize(s string) string
```
Capitalize uppercases the first char of s and lowercases the rest.

#### func  CharAt

```go
func CharAt(s string, index int) string
```
CharAt returns a string from the character at the specified position.

#### func  CharAtF

```go
func CharAtF(index int) func(string) string
```
CharAtF is the filter form of CharAt.

#### func  ChompLeft

```go
func ChompLeft(s, prefix string) string
```
ChompLeft removes prefix at the start of a string.

#### func  ChompLeftF

```go
func ChompLeftF(prefix string) func(string) string
```
ChompLeftF is the filter form of ChompLeft.

#### func  ChompRight

```go
func ChompRight(s, suffix string) string
```
ChompRight removes suffix from end of s.

#### func  ChompRightF

```go
func ChompRightF(suffix string) func(string) string
```
ChompRightF is the filter form of ChompRight.

#### func  Clean

```go
func Clean(s string) string
```
Clean compresses all adjacent whitespace to a single space and trims s.

#### func  Dasherize

```go
func Dasherize(s string) string
```
Dasherize converts a camel cased string into a string delimited by dashes.

#### func  DecodeHTMLEntities

```go
func DecodeHTMLEntities(s string) string
```
DecodeHTMLEntities decodes HTML entities into their proper string
representation. DecodeHTMLEntities is an alias for html.UnescapeString

#### func  EnsurePrefix

```go
func EnsurePrefix(s, prefix string) string
```
EnsurePrefix ensures s starts with prefix.

#### func  EnsurePrefixF

```go
func EnsurePrefixF(prefix string) func(string) string
```
EnsurePrefixF is the filter form of EnsurePrefix.

#### func  EnsureSuffix

```go
func EnsureSuffix(s, suffix string) string
```
EnsureSuffix ensures s ends with suffix.

#### func  EnsureSuffixF

```go
func EnsureSuffixF(suffix string) func(string) string
```
EnsureSuffixF is the filter version of EnsureSuffix.

#### func  EscapeHTML

```go
func EscapeHTML(s string) string
```
EscapeHTML is alias for html.EscapeString.

#### func  Humanize

```go
func Humanize(s string) string
```
Humanize transforms s into a human friendly form.

#### func  IndexOf

```go
func IndexOf(s string, needle string, start int) int
```
IndexOf finds the index of needle in s starting from start.

#### func  IsAlpha

```go
func IsAlpha(s string) bool
```
IsAlpha returns true if a string contains only letters from ASCII (a-z,A-Z).
Other letters from other languages are not supported.

#### func  IsAlphaNumeric

```go
func IsAlphaNumeric(s string) bool
```
IsAlphaNumeric returns true if a string contains letters and digits.

#### func  IsEmpty

```go
func IsEmpty(s string) bool
```
IsEmpty returns true if the string is solely composed of whitespace.

#### func  IsLower

```go
func IsLower(s string) bool
```
IsLower returns true if s comprised of all lower case characters.

#### func  IsNumeric

```go
func IsNumeric(s string) bool
```
IsNumeric returns true if a string contains only digits from 0-9. Other digits
not in Latin (such as Arabic) are not currently supported.

#### func  IsUpper

```go
func IsUpper(s string) bool
```
IsUpper returns true if s contains all upper case chracters.

#### func  Left

```go
func Left(s string, n int) string
```
Left returns the left substring of length n.

#### func  LeftF

```go
func LeftF(n int) func(string) string
```
LeftF is the filter version of Left.

#### func  Lines

```go
func Lines(s string) []string
```
Lines convert windows newlines to unix newlines then convert to an Array of
lines.

#### func  Map

```go
func Map(arr []string, iterator func(string) string) []string
```
Map maps an array's iitem through an iterator.

#### func  Match

```go
func Match(s, pattern string) bool
```
Match returns true if patterns matches the string

#### func  Pad

```go
func Pad(s, c string, n int) string
```
Pad pads string s on both sides until it has length of n.

#### func  PadF

```go
func PadF(c string, n int) func(string) string
```
PadF is the filter version of Pad.

#### func  PadLeft

```go
func PadLeft(s, c string, n int) string
```
PadLeft pads string s on left side until it has length of n.

#### func  PadLeftF

```go
func PadLeftF(c string, n int) func(string) string
```
PadLeftF is the filter version of PadLeft.

#### func  PadRight

```go
func PadRight(s, c string, n int) string
```
PadRight pads string s on right side until it has length of n.

#### func  PadRightF

```go
func PadRightF(c string, n int) func(string) string
```
PadRightF is the filter version of Padright

#### func  Pipe

```go
func Pipe(s string, funcs ...func(string) string) string
```
Pipe pipes s through one or more string filters.

#### func  QuoteItems

```go
func QuoteItems(arr []string) []string
```
QuoteItems quotes all items in array, mostly for debugging.

#### func  ReplaceF

```go
func ReplaceF(old, new string, n int) func(string) string
```
ReplaceF is the filter version of strings.Replace.

#### func  ReplacePattern

```go
func ReplacePattern(s, pattern, repl string) string
```
ReplacePattern replaces string with regexp string. ReplacePattern returns a copy
of src, replacing matches of the Regexp with the replacement string repl. Inside
repl, $ signs are interpreted as in Expand, so for instance $1 represents the
text of the first submatch.

#### func  ReplacePatternF

```go
func ReplacePatternF(pattern, repl string) func(string) string
```
ReplacePatternF is the filter version of ReplaceRegexp.

#### func  Reverse

```go
func Reverse(s string) string
```
Reverse a string

#### func  Right

```go
func Right(s string, n int) string
```
Right returns the right substring of length n.

#### func  RightF

```go
func RightF(n int) func(string) string
```
RightF is the Filter version of Right.

#### func  SetTemplateDelimiters

```go
func SetTemplateDelimiters(opening, closing string)
```
SetTemplateDelimiters sets the delimiters for Template function. Defaults to
"{{" and "}}"

#### func  Slice

```go
func Slice(s string, start, end int) string
```
Slice slices a string. If end is negative then it is the from the end of the
string.

#### func  SliceF

```go
func SliceF(start, end int) func(string) string
```
SliceF is the filter for Slice.

#### func  Slugify

```go
func Slugify(s string) string
```
Slugify converts s into a dasherized string suitable for URL segment.

#### func  StripPunctuation

```go
func StripPunctuation(s string) string
```
StripPunctuation strips puncation from string.

#### func  StripTags

```go
func StripTags(s string, tags ...string) string
```
StripTags strips all of the html tags or tags specified by the parameters

#### func  Substr

```go
func Substr(s string, index int, n int) string
```
Substr returns a substring of s starting at index of length n.

#### func  SubstrF

```go
func SubstrF(index, n int) func(string) string
```
SubstrF is the filter version of Substr.

#### func  Template

```go
func Template(s string, values map[string]interface{}) string
```
Template is a string template which replaces template placeholders delimited by
"{{" and "}}" with values from map. The global delimiters may be set with
SetTemplateDelimiters.

#### func  TemplateDelimiters

```go
func TemplateDelimiters() (opening string, closing string)
```
TemplateDelimiters is the getter for the opening and closing delimiters for
Template.

#### func  TemplateWithDelimiters

```go
func TemplateWithDelimiters(s string, values map[string]interface{}, opening, closing string) string
```
TemplateWithDelimiters is string template with user-defineable opening and
closing delimiters.

#### func  ToArgv

```go
func ToArgv(s string) []string
```
ToArgv converts a s into an argv for exec.

#### func  ToBool

```go
func ToBool(s string) bool
```
ToBool fuzzily converts truthy values.

#### func  Underscore

```go
func Underscore(s string) string
```
Underscore returns converted camel cased string into a string delimited by
underscores.

#### func  UnescapeHTML

```go
func UnescapeHTML(s string) string
```
UnescapeHTML is an alias for html.UnescapeString.

#### func  WrapHTML

```go
func WrapHTML(s string, tag string, attrs map[string]string) string
```
WrapHTML wraps s within HTML tag having attributes attrs.

#### func  WrapHTMLF

```go
func WrapHTMLF(tag string, attrs map[string]string) func(string) string
```
WrapHTMLF is the filter version of WrapHTML
