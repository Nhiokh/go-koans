package goKoans

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
)

func aboutBasics() {
	assert(__bool__ == true)  // what is truth?
	assert(__bool__ != false) // in it there is nothing false

	var i int = __int__
	assert(i == 1.0000000000000000000000000000000000000) // precision is in the eye of the beholder

	k := __int__ // short assignment can be used, as well (using ':=' operator)
	assert(k == 1.0000000000000000000000000000000000000)

	assert(5%2 == __int__) // modulo operator is allowed...
	assert(5^2 == __int__) // ...as well as power operator

	var x int
	assert(x == __int__) // zero values are valued in Go

	var f float32
	assert(f == __float32__) // for types of all types

	var s string
	assert(s == __string__) // both typical or atypical types

	var c struct {
		x int
		f float32
		s string
	}
	assert(c.x == __int__)     // and types within composite types
	assert(c.f == __float32__) // which match the other types
	assert(c.s == __string__)  // in a typical way
}

// Initialization variables
const (
	__string__       string  = "impossibly lame value"
	__int__          int     = -1
	__positive_int__ int     = 42
	__byte__         byte    = 255
	__bool__         bool    = false // ugh
	__boolean__      bool    = true  // oh well
	__float32__      float32 = -1.0
	__delete_me__    bool    = false
)

var __runner__ runner = nil

func _getRecentLine() string {
	_, file, line, _ := runtime.Caller(2)
	buf, _ := ioutil.ReadFile(file)
	code := strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])
	return fmt.Sprintf("%v:%d\n%s", path.Base(file), line, code)
}

func assert(o bool) {
	if !o {
		fmt.Printf("\n%c[35m%s%c[0m\n\n", 27, _getRecentLine(), 27)
		os.Exit(1)
	}
}
