package reflection

import (
	"fmt"
	"testing"
	"unsafe"
)

type User struct {
	age byte
	id  byte
}

func TestUnsafe(t *testing.T) {
	user := &User{1, 0}
	hi := (*int16)(unsafe.Pointer(&user))
	if *hi == int16(1) {
		fmt.Println("低位在前")
	} else {
		fmt.Println(*hi, "高位在前")
	}
}
