package chainofresponsability

import (
	"fmt"
	"io"
	"strings"
)

type ChainLogger interface {
	Next(string)
}

type FirstLogger struct {
	NextChain ChainLogger
}

func (f *FirstLogger) Next(s string) {
	fmt.Printf("First logger: %s\n", s)

	if f.NextChain != nil {
		f.NextChain.Next(s)
	}

}

type SecondLogger struct {
	NextChain ChainLogger
}

func (sl *SecondLogger) Next(s string) {

	if strings.Contains(strings.ToLower(s), "hello") {
		fmt.Printf("Second logger: %s\n", s)

		if sl.NextChain != nil {
			sl.NextChain.Next(s)
		}
		return
	}

	fmt.Printf("Finishing in second logging\n\n")
}

type WriterLogger struct {
	NextChain ChainLogger
	Writer    io.Writer
}

func (w *WriterLogger) Next(s string) {

	if w.Writer != nil {
		w.Writer.Write([]byte("WriterLogger: " + s))
	}

	if w.NextChain != nil {
		w.NextChain.Next(s)
	}

}

type myTestWriter struct {
	receivedMessage *string
}

func (m *myTestWriter) Write(p []byte) (int, error) {

	if m.receivedMessage == nil {
		m.receivedMessage = new(string)
	}

	tempMessage := fmt.Sprintf("%s%s", *m.receivedMessage, p)
	m.receivedMessage = &tempMessage

	return len(p), nil
}

func (m *myTestWriter) Next(s string) {
	m.Write([]byte(s))
}
