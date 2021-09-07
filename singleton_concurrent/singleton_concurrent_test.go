package singletonconcurrent

import (
	"fmt"
	"testing"
	"time"
)

func TestStart_Instance(t *testing.T) {

	s1 := GetInstance()
	s2 := GetInstance()

	n := 5000

	for i := 0; i < n; i++ {
		go s1.AddOne()
		go s2.AddOne()
	}

	fmt.Printf("Before loop, current count is %d\n", s1.GetCount())

	var val int
	for val != n*2 {
		val = s1.GetCount()
		time.Sleep(10 * time.Millisecond)
	}
	s1.Stop()

}
