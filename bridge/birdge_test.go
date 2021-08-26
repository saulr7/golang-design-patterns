package bridge

import (
	"testing"
)

func TestPrintAPI(t *testing.T) {

	ap1 := PrinterImpl1{}

	err := ap1.PrintMessage("Hello")

	if err != nil {
		t.Errorf("Error trying to use the API1 implementation: Message: %s\n", err.Error())
	}
}

func TestPrintapi2(t *testing.T) {

	testWriter := TestWriter{}

	api2 := PrinterImpl2{
		Writer: &testWriter,
	}

	expectedMessage := "Hello"

	err := api2.PrintMessage(expectedMessage)

	if err != nil {
		t.Errorf("Error trying to use the API2 implementation: %s\n", err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Fatalf("API2 did not write correctly on the io.Writer. \n  Actual: %s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}

}

func TestNormalPrinter(t *testing.T) {

	expectedMessage := "Hello io.Writer"

	normal := NormalPrinter{
		Msg:     expectedMessage,
		Printer: &PrinterImpl1{},
	}

	err := normal.Print()

	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestPacktPrinter(t *testing.T) {

	passedMsg := "Hello io.Writer"
	expectedMsg := "Message from Packt: Hello io.Writer"

	packt := PacktPrinter{
		Msg:     passedMsg,
		Printer: &PrinterImpl1{},
	}

	err := packt.Print()

	if err != nil {
		t.Errorf(err.Error())
	}

	testWriter := TestWriter{}
	packt = PacktPrinter{
		Msg: passedMsg,
		Printer: &PrinterImpl2{
			Writer: &testWriter,
		},
	}

	err = packt.Print()

	if err != nil {
		t.Errorf(err.Error())
	}
	if testWriter.Msg != expectedMsg {
		t.Errorf("The expected message on the io.Writer doesn't match actual.\n  Actual: %s\nExpected: %s\n", testWriter.Msg, expectedMsg)
	}
}
