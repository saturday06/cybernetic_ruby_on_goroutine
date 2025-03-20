package main

/*
#include "cybernetic_ruby_on_goroutine.h"

VALUE rb_cybernetic_ruby_on_goroutine(VALUE self, VALUE a, VALUE b);
*/
import "C"

import "github.com/ruby-go-gem/go-gem-wrapper/ruby"

//export Init_cybernetic_ruby_on_goroutine
func Init_cybernetic_ruby_on_goroutine() {
	rb_mCyberneticRubyOnGoroutine := ruby.RbDefineModule("CyberneticRubyOnGoroutine")
	ruby.RbDefineSingletonMethod(rb_mCyberneticRubyOnGoroutine, "sum", C.rb_cybernetic_ruby_on_goroutine, 2)
}

//export rb_cybernetic_ruby_on_goroutine
func rb_cybernetic_ruby_on_goroutine(_ C.VALUE, a C.VALUE, b C.VALUE) C.VALUE {
	aLong := ruby.NUM2LONG(ruby.VALUE(a))
	bLong := ruby.NUM2LONG(ruby.VALUE(b))

	sum := aLong + bLong

	return C.VALUE(ruby.LONG2NUM(sum))
}

func main() {
}
