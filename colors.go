package main

import "github.com/fatih/color"

var (
	fgHiRed     = color.New(color.FgHiRed).SprintFunc()
	fgHiWhite   = color.New(color.FgHiWhite).SprintFunc()
	fgHiGreen   = color.New(color.FgHiGreen).SprintFunc()
	fgHiMagenta = color.New(color.FgHiMagenta).SprintFunc()

	boldFgHiRed    = color.New(color.Bold, color.FgHiRed).SprintFunc()
	boldFgHiYellow = color.New(color.Bold, color.FgHiYellow).SprintFunc()
	boldFgHiGreen  = color.New(color.Bold, color.FgHiGreen).SprintFunc()
)
