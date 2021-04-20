package gui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

type UserInterface interface {
	Write(msg string, args ...interface{})
	Read() string
	Ask(msg string, args ...interface{}) string
	MustFlush()
}

type userInterface struct {
	tabWriter   *tabwriter.Writer
	stdinReader *bufio.Reader
}

func NewUserInterface() *userInterface {
	w := tabwriter.Writer{}
	w.Init(os.Stdout, 0, 8, 2, '\t', 0) //, tabwriter.Debug)

	return &userInterface{
		tabWriter:   &w,
		stdinReader: bufio.NewReader(os.Stdin),
	}
}

func (u userInterface) Write(msg string, args ...interface{}) {
	_, _ = fmt.Fprintln(u.tabWriter, fmt.Sprintf(msg, args...))
}

func (u userInterface) Read() string {
	text, _ := u.stdinReader.ReadString('\n')
	return strings.TrimSuffix(text, "\n")
}

func (u userInterface) Ask(msg string, args ...interface{}) string {
	u.Write(msg, args...)
	return u.Read()
}

func (u userInterface) MustFlush() {
	if err := u.tabWriter.Flush(); err != nil {
		panic(err.Error())
	}
}

func debugStr(str string) {
	fmt.Printf("plain string: ")
	fmt.Printf("%s", str)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", str)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
	fmt.Printf("\n")
}
