package pattern

import "fmt"

// State interface and implementation
type ILightBulbLeverState interface {
	Flick()
}

type LeverUp struct {
	lever *Lever
}

func (s *LeverUp) Init(lever *Lever) (self *LeverUp) {
	s.lever = lever
	return s
}

func (s *LeverUp) Flick() {
	fmt.Println("Flicked the lever down and turned the light on.")
	s.lever.ChangeState(&LeverDown{s.lever})
}

type LeverDown struct {
	lever *Lever
}

func (s *LeverDown) Init(lever *Lever) (self *LeverDown) {
	s.lever = lever
	return s
}

func (s *LeverDown) Flick() {
	fmt.Println("Flicked the lever up and turned the light off.")
	s.lever.ChangeState(&LeverUp{s.lever})
}

// Statefull object
type Lever struct {
	state ILightBulbLeverState
}

func (l *Lever) Init() (self *Lever) {
	l.state = (&LeverUp{}).Init(l)
	return l
}

func (l *Lever) Flick() {
	l.state.Flick()
}

func (l *Lever) ChangeState(state ILightBulbLeverState) { // do I really need this one?
	l.state = state
}

// Demonstration
func init() {
	fmt.Println("Demonstrating State pattern:")
	defer fmt.Println()

	l := (&Lever{}).Init()
	l.ChangeState((&LeverDown{}).Init(l))
	l.Flick()
	l.Flick()
}
