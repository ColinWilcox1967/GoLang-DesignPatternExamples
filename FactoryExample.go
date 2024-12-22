package main

import (
	"fmt"
	"sync"
)

type Singleton struct {}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()

	fmt.Println(s1 == s2) // Output: true
}

2. Factory Pattern
Purpose
Creates objects without specifying their exact class.

Example in Go
package main

import "fmt"

type Shape interface {
	Draw()
}

type Circle struct {}
func (c Circle) Draw() { fmt.Println("Drawing a Circle") }

type Square struct {}
func (s Square) Draw() { fmt.Println("Drawing a Square") }

func ShapeFactory(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return Circle{}
	case "square":
		return Square{}
	default:
		return nil
	}
}

func main() {
	shape := ShapeFactory("circle")
	shape.Draw()
}