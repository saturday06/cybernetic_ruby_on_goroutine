package main

/*
#include "cybernetic_ruby_on_goroutine.h"
*/
import "C"

import (
	"fmt"

	"github.com/ruby-go-gem/go-gem-wrapper/ruby"
)

//export Init_cybernetic_ruby_on_goroutine
func Init_cybernetic_ruby_on_goroutine() {
	rb_mCyberneticRubyOnGoroutine := ruby.RbDefineModule("CyberneticRubyOnGoroutine")
	fmt.Printf("rb_mCyberneticRubyOnGoroutine: %v\n", rb_mCyberneticRubyOnGoroutine)
}

func main() {
}
