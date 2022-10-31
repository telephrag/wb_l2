package pattern

import "fmt"

// facade
type PassTestsFacade struct {
	tests    *Tests
	shaman   *Shaman
	dancable *Dancable
}

// f := (&PassTestsFacade).New(...)
func (f *PassTestsFacade) New(tests *Tests, shaman *Shaman, dancable Dancable) (self *PassTestsFacade) {
	f.tests = tests
	f.shaman = shaman
	f.dancable = &dancable
	return f
}

func (f *PassTestsFacade) Do() {
	if f.tests == nil || f.shaman == nil || f.dancable == nil {
		return
	}

	f.shaman.DanceWith(f.dancable)
	f.tests.Test()
}

// internal structs abstracted by facade
type Tests struct{}

func (t *Tests) Test() { fmt.Println("testing the programm for errors") }

type Shaman struct{}

func (s *Shaman) DanceWith(d *Dancable) { fmt.Printf("dancing with %s\n", (*d).UseForDancing()) }

type Dancable interface {
	UseForDancing() string
}

type Tambourine struct{}

func (t Tambourine) UseForDancing() string { return "tambourine" }

// demonstrate work on package import
func init() {
	fmt.Println("Demonstrating Facade pattern:")
	defer fmt.Println()

	t := &Tests{}
	s := &Shaman{}
	d := &Tambourine{}
	f := (&PassTestsFacade{}).New(t, s, d)

	f.Do()
}
