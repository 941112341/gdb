package tree

import (
	"fmt"
	"testing"
)

type statusError struct {
}

func (statusError) Error() string {
	return "status error"
}

func TestName(t *testing.T) {
	i := []int(nil)
	fmt.Println(i == nil)
	fmt.Printf("%p\n", i)
	var x map[int]int = nil
	fmt.Printf("%p\n", x)
	//fmt.Println(x == nil)
	//var err error = Ret()
	var err error = Ret() // 指针 赋值给接口，type都会赋值进去，interface， ni的比较就一定为false，只限于接口和nil的比较
	fmt.Printf("%p\n", err)
	fmt.Println(err == nil)
}

func Ret() *statusError {
	return nil
}

func TestSync(t *testing.T) {
	//sync.Map{}
	//sync.RWMutex{}
	//sync.Cond{}
	//sync.Once{}
	//sync.WaitGroup{}
}
