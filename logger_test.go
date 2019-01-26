package logger

import (
	"bytes"
	"log"
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_constructor(t *testing.T) {
	l := New("prefix")
	assert := asserter.New(t)
	assert(l != nil).Fail()
}

func Test_output_of_logger(t *testing.T) {
	assert := asserter.New(t)
	w := bytes.NewBufferString("")
	l := Adapt(log.New(w, "", log.LstdFlags))

	exp := "emotions assign value to things"
	l.Log(exp)
	assert().Contains(w.Bytes(), exp)

	w.Reset()
	l.Logf("%s", exp)
	assert().Contains(w.Bytes(), exp)
}

func Test_silent(t *testing.T) {
	l := Silent
	l.Log("nada")
	l.Logf("%s", "nada")
	//t.Fail()
}
