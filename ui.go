package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

type userInterface struct {
	tabWriter   *tabwriter.Writer
	stdinReader *bufio.Reader
}

func newUserInterface() *userInterface {
	w := tabwriter.Writer{}
	w.Init(os.Stdout, 0, 4, 0, '\t', 0)

	return &userInterface{
		tabWriter:   &w,
		stdinReader: bufio.NewReader(os.Stdin),
	}
}

func (u userInterface) write(msg string, args ...interface{}) {
	_, _ = fmt.Fprintln(u.tabWriter, fmt.Sprintf(msg, args...))
}

func (u userInterface) read() string {
	text, _ := u.stdinReader.ReadString('\n')
	return strings.TrimSuffix(text, "\n")
}

func (u userInterface) askUser(msg string, args ...interface{}) string {
	u.write(msg, args...)
	return u.read()
}

func (u userInterface) mustFlush() {
	if err := u.tabWriter.Flush(); err != nil {
		panic(err.Error())
	}
}
