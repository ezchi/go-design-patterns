package printer

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestPrintAPI1(t *testing.T) {
	api1 := PrinterImpl1{}

	err := api1.PrintMessage("Hello")

	if err != nil {
		t.Errorf("can not use the API1 implementation: %v\n", err)
	}
}

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}

	err = errors.New("content received on Writer was empty")

	return
}

func TestPrintAPI2(t *testing.T) {
	api2 := PrinterImpl2{}

	err := api2.PrintMessage("Hello")

	if err != nil {
		expectedErrorMessage := "You need to pass an io.Writer to PrinterImpl2"

		if !strings.Contains(err.Error(), expectedErrorMessage) {
			t.Errorf("Error message was not corrent.\nActual: %s\nExpected: %s\n", err.Error(), expectedErrorMessage)
		}
	}

	testWriter := TestWriter{}

	api2 = PrinterImpl2{Writer: &testWriter}

	expectedMessage := "Hello"
	err = api2.PrintMessage(expectedMessage)

	if err != nil {
		t.Errorf("can not print message: %v\n", err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Fatalf("API2 can not write correctly on the io.Writer.\nActual: %s\nExpected: %s\n", testWriter.Msg, expectedMessage)

	}
}

func TestNormalPrinterPrint(t *testing.T) {
	expectedMessage := "Hello io.Writer"

	testWriter := TestWriter{}

	normal := NormalPrinter{
		Msg:     expectedMessage,
		Printer: &PrinterImpl2{Writer: &testWriter},
	}

	err := normal.Print()
	if err != nil {
		t.Error(err)
	}

	if testWriter.Msg != expectedMessage {
		t.Errorf("can not write to the io.Writer correctly:\n Got: %s\n Expected: %s\n", testWriter.Msg, expectedMessage)
	}
}

func TestPacktPrinterPrint(t *testing.T) {
	passedMessage := "Hello io.Writer"
	expectedMessage := fmt.Sprintf("Message from Packt: %s", passedMessage)

	packt := PacktPrinter{
		Msg:     passedMessage,
		Printer: &PrinterImpl1{},
	}

	err := packt.Print()
	if err != nil {
		t.Error(err)
	}

	testWriter := TestWriter{}

	packt = PacktPrinter{
		Msg:     passedMessage,
		Printer: &PrinterImpl2{Writer: &testWriter},
	}

	err = packt.Print()

	if err != nil {
		t.Error(err)
	}

	if testWriter.Msg != expectedMessage {
		t.Errorf("can not write to the io.Writer correctly:\n Got: %s\n Expected: %s\n", testWriter.Msg, expectedMessage)
	}
}
