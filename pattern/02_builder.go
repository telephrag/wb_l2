package pattern

import "fmt"

// struct that will be constructed by builders
type HouseWithWalls interface {
	SetWalls(walls string)
}

type HouseWithRoof interface {
	SetRoof(roof string)
}

type House struct {
	walls, roof string
}

func (h *House) String() string {
	return fmt.Sprintf("This is a house with %s roof and %s walls.", h.roof, h.walls)
}

func (h *House) SetWalls(walls string) { h.walls = walls }

func (h *House) SetRoof(roof string) { h.roof = roof }

// builders
type ConcreteWallsBuilder struct{}

func (b *ConcreteWallsBuilder) BuildConcreteWalls(h HouseWithWalls) {
	h.SetWalls("concrete")
}

type BrickWallsBuilder struct{}

func (b *BrickWallsBuilder) BuildBrickWalls(h HouseWithWalls) {
	h.SetWalls("brick")
}

type MetalRoofBuilder struct{}

func (b *MetalRoofBuilder) BuildMetalRoof(h HouseWithRoof) {
	h.SetRoof("metal")
}

// demonstrate work on package import
func init() {
	fmt.Println("Demonstrating Builder pattern:")
	defer fmt.Println()

	h := &House{}
	(&ConcreteWallsBuilder{}).BuildConcreteWalls(h)
	(&MetalRoofBuilder{}).BuildMetalRoof(h)
	fmt.Println(h.String())
}
