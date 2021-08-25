package composite

import (
	"fmt"
	"testing"
)

func TestComposite(t *testing.T) {

	swimmer2 := CompositeSwimmerA{MySwim: Swim}
	swimmer2.MySwim()
	swimmer2.MyAthele.Train()

	fish := Shark{Swim: Swim}
	fish.Eat()
	fish.Swim()

	swimmer := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}

	swimmer.Train()
	swimmer.Swim()

	root := Tree{
		LeafValue: 0,
		Right: &Tree{
			LeafValue: 5,
			Right:     &Tree{6, nil, nil},
			Left:      nil,
		},
		Left: &Tree{4, nil, nil},
	}

	fmt.Println(root)

}
