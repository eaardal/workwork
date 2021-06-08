package gui

import "github.com/fatih/color"

var (
	FgHiRed     = color.New(color.FgHiRed).SprintFunc()
	FgHiGreen   = color.New(color.FgHiGreen).SprintFunc()
	FgHiMagenta = color.New(color.FgHiMagenta).SprintFunc()

	BoldFgHiRed    = color.New(color.Bold, color.FgHiRed).SprintFunc()
	BoldFgHiYellow = color.New(color.Bold, color.FgHiYellow).SprintFunc()
	BoldFgHiGreen  = color.New(color.Bold, color.FgHiGreen).SprintFunc()
)
