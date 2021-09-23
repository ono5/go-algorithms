package composite

import (
	"fmt"
	"testing"
)

func TestAthlete_Train(t *testing.T) {
	Athlete := Athlete{}
	Athlete.Train()
}

func TestSwimmer_Swim(t *testing.T) {
	localSwim := Swim
	swimmer := CompositeSwimmmrA{
		MySwim: &localSwim,
	}
	swimmer.MyAhlete.Train()
	(*swimmer.MySwim)()
}

func TestSwimmer_Swim2(t *testing.T) {
	swimmer := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImplementor{},
	}
	swimmer.Train()
	swimmer.Swim()
}

func TestAnimal_Swim(t *testing.T) {
	fish := Shark{
		Swim: Swim,
	}
	fish.Eat()
	fish.Swim()
}

func TestTree(t *testing.T) {
	root := Tree{
		LeafValue: 0,
		Right: &Tree{
			LeafValue: 5,
			Right:     &Tree{6, nil, nil},
		},
		Left: &Tree{4, nil, nil},
	}
	fmt.Println(root.Right.Right.LeafValue)
}
