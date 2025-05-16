//class -> struct + method receivers

type Rectangle struct {
    Length float64
    Width  float64
}

// Method using value receiver
func (r Rectangle) Area() float64 {
    return r.Length * r.Width
}

// Method using pointer receiver (to modify)
func (r *Rectangle) Scale(factor float64) {
    r.Length *= factor
    r.Width *= factor
}

//inheritance -> composition via embedded structs

type Rectangle struct {
    Length float64
    Width  float64
}

// Method using value receiver
func (r Rectangle) Area() float64 {
    return r.Length * r.Width
}

// Method using pointer receiver (to modify)
func (r *Rectangle) Scale(factor float64) {
    r.Length *= factor
    r.Width *= factor
}

//Interfaces → Native interfaces (duck typing)
type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func PrintArea(s Shape) {
    fmt.Println("Area:", s.Area())
}

//Constructors → Factory functions like NewType()
type User struct {
    Name string
    Age  int
}

func NewUser(name string, age int) *User {
    return &User{
        Name: name,
        Age:  age,
    }
}

// Access control → Exported vs Unexported

type Person struct {
    Name string // Exported
    age  int    // Unexported
}



