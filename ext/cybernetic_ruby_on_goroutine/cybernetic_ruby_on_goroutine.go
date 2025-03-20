package main

/*
#include "cybernetic_ruby_on_goroutine.h"

VALUE rb_cybernetic_ruby_on_goroutine(VALUE self, VALUE a, VALUE b);
*/
import "C"

import (
	"fmt"
	"runtime"
	"syscall"
	"time"
	"unsafe"

	"github.com/ruby-go-gem/go-gem-wrapper/ruby"
)

//export Init_cybernetic_ruby_on_goroutine
func Init_cybernetic_ruby_on_goroutine() {
	rb_mCyberneticRubyOnGoroutine := ruby.RbDefineModule("CyberneticRubyOnGoroutine")
	ruby.RbDefineSingletonMethod(rb_mCyberneticRubyOnGoroutine, "sum", C.rb_cybernetic_ruby_on_goroutine, 2)
}

//export rb_cybernetic_ruby_on_goroutine
func rb_cybernetic_ruby_on_goroutine(_ C.VALUE, a C.VALUE, b C.VALUE) C.VALUE {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	fmt.Println("start", syscall.Gettid())

	done := func() <-chan struct{} {
		done := make(chan struct{})

		go func() {
			defer close(done)

			runtime.LockOSThread()
			defer runtime.UnlockOSThread()
			fmt.Println("sub goroutine", syscall.Gettid())

			time.Sleep(5 * time.Second)
			// state := 0
			script := "puts 'Hello, world!'"
			cscript := C.CString(script)
			defer C.free(unsafe.Pointer(cscript))
			fmt.Println("before eval")
			// C.rb_eval_string_protect(cscript, (*C.int)(unsafe.Pointer(&state)))
			fmt.Println("G3b")
		}()

		return done
	}()

	fmt.Println("before wait")
	<-done
	fmt.Println("done")

	return C.VALUE(ruby.LONG2NUM(3))
}

func main() {
}
