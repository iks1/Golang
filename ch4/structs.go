/*type Employee struct{
	ID int
	Name string
	Adress string
	DoB time.Time 
	Position string 
	Salary int 
	ManagerId int 
}

var dilbert Employee*/

type tree struct{
	value int 
	left, right *tree
}

func sort(values []int){
	var root *tree 
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t!= nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t=new(tree)
		t.value = value
		return t
	}
	if value < t.value {
	    t.left = add(t.left, value)
	} else {
	    t.right = add(t.right, value)
	}
	return t
}

//struct literals 

type Point struct { X, Y int}
p := Point{1, 2}

//return from function 
func scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

// 
type Movie struct{
	Title string 
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

