package pattern

import (
	"errors"
	"fmt"
)

// interface we will fabricate

type Car interface {
	Drive()
}

// concrete implementations of interface above
type Impreza struct{}

func NewImpreza() Car {
	fmt.Println("Importing Impreza from Japan.")
	return &Impreza{}
}

func (c *Impreza) Drive() {
	fmt.Println("haha impreza goes wroooooo stutututtutu, wrooooom")
}

type E30 struct{}

func NewE30() Car {
	fmt.Println("Importing E30 from near abroad.")
	return &E30{}
}

func (c *E30) Drive() {
	fmt.Println("haha e30 goes wrorororooroooom, wroooooooo, wroooooooooooooooooo")
}

// factory
type CarImporter struct {
	importMethods map[string]func() Car
}

func (i *CarImporter) Init(importMethods map[string]func() Car) (self *CarImporter) {
	i.importMethods = importMethods
	return i
}

func (i *CarImporter) Import(name string) (Car, error) {
	if im, ok := i.importMethods[name]; ok {
		return im(), nil
	} else {
		return nil, errors.New("can't import such car")
	}
}

// demonstration
func init() {
	fmt.Println("Demonstrating Factory Method pattern:")
	defer fmt.Println()

	im := map[string]func() Car{"impreza": NewImpreza, "e30": NewE30}
	ci := (&CarImporter{}).Init(im)

	if car, err := ci.Import("e30"); err == nil {
		car.Drive()
	}

	if car, err := ci.Import("impreza"); err == nil {
		car.Drive()
	}

	if car, err := ci.Import("civic"); err == nil {
		car.Drive()
	}
}
