package printer

import (
	"errors"
	"fmt"
	"io"
)

type PrinterAPI interface {
	PrintMessage(string) error
}

type PrinterImpl1 struct{}

func (p *PrinterImpl1) PrintMessage(msg string) error {
	_, err := fmt.Print(msg)

	return err
}

type PrinterImpl2 struct {
	Writer io.Writer
}

func (p *PrinterImpl2) PrintMessage(msg string) error {
	if p.Writer == nil {
		return errors.New("You need to pass an io.Writer to PrinterImpl2")
	}

	_, err := p.Writer.Write([]byte(msg))

	return err
}

type PrinterAbstraction interface {
	Print() error
}

type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (n *NormalPrinter) Print() error {
	if n.Printer == nil {
		return errors.New("Printer is not assigned yet")
	}

	return n.Printer.PrintMessage(n.Msg)
}

type PacktPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (p *PacktPrinter) Print() error {
	return p.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s", p.Msg))
}
