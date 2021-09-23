package composite

import "fmt"

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

func Swim() {
	fmt.Println("Swimming")
}

type CompositeSwimmmrA struct {
	MyAhlete Athlete
	MySwim   *func()
}

// -------------------------------

type Trainer interface {
	Train()
}

type Swimmer interface {
	Swim()
}

type SwimmerImplementor struct{}

func (s *SwimmerImplementor) Swim() {
	fmt.Println("Swimming!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

// -------------------------------

type Animal struct{}

func (a *Animal) Eat() {
	println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}

// -------------------------------

type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}
