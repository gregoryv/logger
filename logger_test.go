package logger

import (
	"bytes"
	"log"
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_output_of_logger(t *testing.T) {
	buf := bytes.NewBufferString("")
	log := New(buf, "", log.LstdFlags)
	log.Print("hello")
	assert := asserter.New(t)
	assert().Contains(buf.Bytes(), "hello")

	log.Printf("%s is ok", "world")
	assert().Contains(buf.Bytes(), "world is ok")
}

func Test_working_with_tAdapter(t *testing.T) {
	log := NewTadapter(t)
	log.Print("working")
	log.Printf("It's %s", "working")
	//t.Fail()
}

func Test_silent(t *testing.T) {
	log := NewSilent()
	log.Print("nada")
	log.Printf("%s", "nada")
	//t.Fail()
}
