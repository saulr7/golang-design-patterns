package memento

import "fmt"

//Struct to save as Memento
type State struct {
	Description string
}

type memento struct {
	state State
}

type originator struct {
	state State
}

func (o *originator) NewMemento() memento {
	return memento{state: o.state}
}

func (o *originator) ExtractAndStoreState(m memento) {
	o.state = m.state
}

type careTaker struct {
	mementoList []memento
}

func (c *careTaker) Add(m memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *careTaker) memento(i int) (memento, error) {

	if len(c.mementoList) < i || i < 0 {
		return memento{}, fmt.Errorf("index not found")

	}
	return c.mementoList[i], nil
}
