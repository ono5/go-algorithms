package adapter

import "fmt"

// -----------------------------------------------
type LegacyPrinter interface {
	Print(s string) string
}

// -----------------------------------------------
type MyLegacyPrinter struct{}

func (l *MyLegacyPrinter) Print(s string) string {
	newMsg := fmt.Sprintf("Legacy Printer: %s\n", s)
	return newMsg
}

// -----------------------------------------------
type NewPrinter interface {
	PrintStored() string
}
type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrinterAdapter) PrintStored() string {
	var newMsg string
	if p.OldPrinter != nil {
		newMsg = p.OldPrinter.Print(p.Msg)
	} else {
		newMsg = p.Msg
	}
	return newMsg
}
