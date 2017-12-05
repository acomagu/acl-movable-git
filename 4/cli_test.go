package main

import (
	"bytes"
	"io/ioutil"
	"fmt"
	"testing"
)

var a = `a
b
c
`

var b = `d
e
f
`

func TestRun_Filenames(t *testing.T) {
	expected := `< a
< b
< c
> d
> e
> f
`

	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{
		outStream: outStream,
		errStream: errStream,
	}

	af, err := ioutil.TempFile("", "")
	if err != nil {
		fmt.Println(err)
	}
	af.WriteString(a)
	aName := af.Name()
	af.Close()

	bf, err := ioutil.TempFile("", "")
	if err != nil {
		fmt.Println(err)
	}
	bf.WriteString(b)
	bName := bf.Name()
	bf.Close()

	args := []string{"mdiff", aName, bName}

	status := cli.Run(args)
	if status != 0 {
		t.Errorf("ExitStatus=%d, want %d", status, 0)
	}

	if outStream.String() != expected {
		t.Errorf("Output=%q, want %q", outStream.String(), expected)
	}
}

func TestRun_U(t *testing.T) {
	expected := `- a
- b
- c
+ d
+ e
+ f
`
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{
		outStream: outStream,
		errStream: errStream,
	}

	af, err := ioutil.TempFile("", "")
	if err != nil {
		fmt.Println(err)
	}
	af.WriteString(a)
	aName := af.Name()
	af.Close()

	bf, err := ioutil.TempFile("", "")
	if err != nil {
		fmt.Println(err)
	}
	bf.WriteString(b)
	bName := bf.Name()
	bf.Close()

	args := []string{"mdiff", aName, bName, "-u"}

	status := cli.Run(args)
	if status != 0 {
		t.Errorf("ExitStatus=%d, want %d", status, 0)
	}

	if outStream.String() != expected {
		t.Errorf("Output=%q, want %q", outStream.String(), expected)
	}
}
