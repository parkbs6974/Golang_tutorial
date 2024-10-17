package main

import "fmt"

func main(){
	a := 3 // int - 64bit = int64
	var b float64 = 3.5
	var c int = int(b) //3.5 => 3 :: int로 변환이 되었기에 소수점이 사라진다
	d := float64(a) * b

	var e int64 = 7

	f := a * int(e) // int와 int64는 다른 타입이라고 간주가 되고 있다 (같은 사이즈라도 타입이 같아야 한다)

	fmt.Println(a,b,c,d,e,f)

}