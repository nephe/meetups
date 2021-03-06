package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// START OMIT
func main() {
	s := string([]byte("dog"))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bsh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	bs := *((*[]byte)(unsafe.Pointer(&bsh)))
	fmt.Println(bs)
	bs[0] = 'h'
	fmt.Println(s)
}

// END OMIT
