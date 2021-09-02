package mediator

import (
	"fmt"
	"testing"
)

func TestMediator_Sum(t *testing.T) {

	result := Sum(One{}, Two{})

	fmt.Println(result)

	switch i := result.(type) {
	case interface{}:
		fmt.Println("Right")
	default:
		t.Error("Expected Three", i)
	}

}
