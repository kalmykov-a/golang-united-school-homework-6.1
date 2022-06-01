package golang_united_school_homework

import (
	"errors"
)

var (
	errorOutOfRange = errors.New("out of the shapesCapacity range")
	errorNotExist   = errors.New("shape by index doesn't exist or index went out of the range")
	errorNoCircles  = errors.New("circles are not exist in the list")
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) >= b.shapesCapacity {
		return errorOutOfRange
	} else {
		b.shapes = append(b.shapes, shape)
		return nil
	}
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= b.shapesCapacity-1 || len(b.shapes)-1 < i {
		return nil, errorOutOfRange
	} else {
		return b.shapes[i], nil
	}
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i >= b.shapesCapacity-1 || len(b.shapes)-1 < i {
		return nil, errorOutOfRange
	} else {
		res := b.shapes[i]
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		return res, nil
	}
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i >= b.shapesCapacity-1 || len(b.shapes) < i {
		return nil, errorNotExist
	} else {
		var res Shape
		res, b.shapes[i] = b.shapes[i], shape
		return res, nil
	}

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	sum := 0.0
	for _, v := range b.shapes {
		sum += v.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	sum := 0.0
	for _, v := range b.shapes {
		sum += v.CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	count := 0
	var indexes []int
	for i, sh := range b.shapes {
		if sh.(*Circle) != nil {
			count++
			indexes = append(indexes, i)
		}
	}
	if count == 0 {
		return errorNoCircles
	} else {
		for i, v := range indexes {
			b.shapes = append(b.shapes[:v-i], b.shapes[v-i+1:]...)
		}
		return nil
	}
}
