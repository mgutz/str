## Testing

### Examples

Examples are great but having too many in an Example results in output
that is hard to correlate. Create an example function when you have
three or more results.

    func eg(index int, example interface{}) {
        output := fmt.Sprintf("%d: %v", index, example)
        fmt.Printf("%s\n", Clean(output))
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

### Regular Expressions

Use backticks \` for regular expression literals.

    var re = Regexp(`^abc$`)

Use `var` when creating `Regexp`. Regexp are safe to promote as package
variables.

    var re = Regexp(`^abc$`)


