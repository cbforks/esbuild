// +build !darwin
// +build !linux
// +build !windows

package logging

import "os"

const SupportsColorEscapes = false

func GetTerminalInfo(*os.File) TerminalInfo {
	return TerminalInfo{}
}

func writeStringWithColor(file *os.File, text string) {
	file.WriteString(text)
}
