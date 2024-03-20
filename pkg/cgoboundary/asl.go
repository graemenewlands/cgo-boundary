package cgoboundary

/*
#cgo CFLAGS: -I/home/graeme/git/cgo-boundary/pkg/cgoboundary
#cgo LDFLAGS: -L/home/graeme/git/cgo-boundary/pkg/cgoboundary -ltest
#include "test.h"
*/
import "C"

func FunctionToIterate(byref func(), iterations int) {
    for i := 0; i < iterations; i++ {
        byref()
    }
}


func CGOOverhead() {
    C.trivialFunction()
	return
}

func NoCGOOverhead() {
    return 
}
