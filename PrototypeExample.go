package main

import "fmt"

type Prototype interface {
	Clone() Prototype
}package main

import "fmt"

type House struct {
	Walls   string
	Roof    string
	Windows int
}

type HouseBuilder struct {
	house House
}

func (b *HouseBuilder) SetWalls(walls string) *HouseBuilder {
	b.house.Walls = walls
	return b
}

func (b *HouseBuilder) SetRoof(roof string) *HouseBuilder {
	b.house.Roof = roof
	return b
}

func (b *HouseBuilder) SetWindows(windows int) *HouseBuilder {
	b.house.Windows = windows
	return b
}

func (b *HouseBuilder) Build() House {
	return b.house
}

func main() {
	builder := &HouseBuilder{}
	house := builder.SetWalls("Brick").SetRoof("Tile").SetWindows(4).Build()
	fmt.Println(house)
}