package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"flag"
)

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) Run(args []string) int {
	af, err := os.Open(args[1])
	if err != nil {
		fmt.Sprintln(c.errStream, err)
		return 1
	}
	defer af.Close()
	a, err := ioutil.ReadAll(af)
	if err != nil {
		fmt.Sprintln(c.errStream, err)
		return 1
	}

	bf, err := os.Open(args[2])
	if err != nil {
		fmt.Sprintln(c.errStream, err)
		return 1
	}
	defer bf.Close()
	b, err := ioutil.ReadAll(bf)
	if err != nil {
		fmt.Sprintln(c.errStream, err)
		return 1
	}

	var unified bool
	flags := flag.NewFlagSet("mdiff", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.BoolVar(&unified, "u", false, "output as unified context")

	if err := flags.Parse(args[1:]); err != nil {
		fmt.Sprintln(c.errStream, err)
		return 1
	}
	for 0 < flags.NArg() {
		err := flags.Parse(flags.Args()[1:])
		if err != nil {
			fmt.Sprintln(c.errStream, err)
			return 1
		}
	}

	fmt.Fprint(c.outStream, diff(string(a), string(b), unified))

	return 0
}
