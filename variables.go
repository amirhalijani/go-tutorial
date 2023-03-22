// Ways to define variables
var example_1 = 3

var example_2 int = 3

var example_3 int
example_3 = 3

var x,y,z = 1, 2, "example"

var (
	a int = 1
	b int = 2
	c = "example"
)

example_4 := 3

d, e := 4, "example"

// Default Values
var firstName string // -> ""
var age int // -> 0
// -> other default: nil (null) 

const UserAge = 21
const UserName = "Kurt"