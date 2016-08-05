package efficient

import (
	"fmt"
	"reflect"
	"unsafe"
)

// BytesToString 将bs转换为string,此方法利用[]byte和string
// 头结构“部分相同”，以非安全的指针类型转换来实现，能有效改善
// 执行效率。
// Note:修改 bs 数据自身，返回的string 内容也会随之改变
func BytesToString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

// PrintDataPointer 提供通用的打印输出数据指针的方法
func PrintDataPointer(format string, ptr interface{}) {
	p := reflect.ValueOf(ptr).Pointer()
	h := (*uintptr)(unsafe.Pointer(p))
	fmt.Printf(format, *h)
}
