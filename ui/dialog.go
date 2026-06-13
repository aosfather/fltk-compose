package ui

import (
	"fltk"
)

func ShowFileDialog() []string {
	fd := fltk.NewNativeFileChooser()
	fd.Show()
	return fd.Filenames()
}

func ShowMessage(title, msg string) {
	fltk.MessageBox(title, msg)
}

func ShowMessageDialog(msg string, button string) {
	fltk.ChoiceDialog(msg, button)
}

func ShowConfirmDialog(msg string, yes string, cancel string) int {
	return fltk.ChoiceDialog(msg, yes, cancel)
}
