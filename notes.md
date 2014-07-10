

// EQ (S('the humanize_id string method_id').humanize().s, 'The humanize id string method')
// EQ (S('the  humanize string method  ').humanize().s, 'The humanize string method')
// EQ (S('   capitalize dash-CamelCase_underscore trim  ').humanize().s, 'Capitalize dash camel case underscore trim')
// EQ (S(123).humanize().s, '123')
// EQ (S('').humanize().s, '')
// EQ (S(null).humanize().s, '')
// EQ (S(undefined).humanize().s, '')

// func ExampleTruncate() {
// 	eg(1, Truncate("this is some long text", "...", 3))
// 	eg(2, Truncate("this is some long text", "...", 7))
// 	eg(3, Truncate("this is some long text", "...", 11))
// 	eg(4, Truncate("this is some long text", "...", 12))
// 	eg(5, Truncate("this is some long text", " read more", 12))
// 	eg(6, Truncate("this is some long text", "...", 120))
// 	// Output:
// 	// 1: ...
// 	// 2: this is...
// 	// 3: this is...
// 	// 4: this is some...
// 	// 5: this is some read more
// 	// 6: this is some long text
// }

