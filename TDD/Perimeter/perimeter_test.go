package main

import (
	"testing"
)

func TestPerimeter(t *testing.T){
	rectangle := Rectangle{10.0,10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}


// func TestArea(t *testing.T) {
// 	checkArea := func(t testing.TB, shape Shape, want float64){
// 		t.Helper()
// 		got := shape.Area()
// 		if got!= want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	}

// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{12, 6}
// 		want := 72.0

// 		checkArea(t, rectangle, want)
// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		want := 314.1592653589793
// 		checkArea(t, circle, want)
// 	})

// }

// table driven tests

func TestArea(t *testing.T){
	areaTests := []struct{
		name string
		shape Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12,6}, hasArea: 72.0},
		{name: "Circle",shape:Circle{10}, hasArea: 314.1592653589793},
	}

	for _,tt := range areaTests{
		got := tt.shape.Area()
		if got!= tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}
}