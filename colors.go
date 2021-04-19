package main

import "github.com/fatih/color"

var (
	hiRed   = color.New(color.FgHiRed).SprintFunc()
	hiWhite = color.New(color.FgHiWhite).SprintFunc()
	hiGreen = color.New(color.FgHiGreen).SprintFunc()

	boldHiRed    = color.New(color.Bold, color.FgHiRed).SprintFunc()
	boldHiYellow = color.New(color.Bold, color.FgHiYellow).SprintFunc()
	boldHiGreen  = color.New(color.Bold, color.FgHiGreen).SprintFunc()
)
