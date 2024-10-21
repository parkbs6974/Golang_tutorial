package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main(){
	str1 := "hello Wold"
	str2 := str1

	stringData1 := unsafe.StringData(str1)
	stringData2 := unsafe.Pointer(&str1)
	// 	•	unsafe.StringData(str1): 문자열 데이터(실제 문자들이 저장된 메모리)의 주소를 반환.
	// •	unsafe.Pointer(&str1): 문자열 변수(str1 자체)의 메모리 주소를 반환.
	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	fmt.Println(stringData1)
	fmt.Println(stringData2)
	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)


}