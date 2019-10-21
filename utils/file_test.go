package utils

import (
	"io/ioutil"
	"testing"
)

func TestReadFileContent(t *testing.T) {

	b, err := ioutil.ReadFile("static/important.md")
	if err != nil {
		println(err)
	}
	println(string(b))
}
