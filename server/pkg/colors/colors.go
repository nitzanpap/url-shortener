package colors

import (
	"github.com/fatih/color"
)

// Error prints a red error message
func Error(msg string) string {
	return color.RedString(msg)
}

// Warning prints a yellow warning message
func Warning(msg string) string {
	return color.YellowString(msg)
}

// Info prints a cyan info message
func Info(msg string) string {
	return color.CyanString(msg)
}

// Success prints a green success message
func Success(msg string) string {
	return color.GreenString(msg)
}
