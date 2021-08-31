package commander

import "testing"

func TestCommand(t *testing.T) {

	queue := CommandQueue{}

	queue.AddCommand(CreateCommand("First message"))
	queue.AddCommand(CreateCommand("Second message"))
	queue.AddCommand(CreateCommand("Third message"))

	queue.AddCommand(CreateCommand("Fourth message"))

	if len(queue.queue) != 4 {
		t.Error("must be 0", len(queue.queue))
	}

	queue.AddCommand(CreateCommand("Fifth message"))

}
