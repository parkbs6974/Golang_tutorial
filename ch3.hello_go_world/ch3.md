# 3강 Hello World

## G0의 역사

- 2009년에 발표된 언어
- 켄 톰슨(c언어의 모체가 되는 b언어 개발), 롭 파이크(UTF-8 개발)
- 컴파일 언어
  => 코드를 컴파일러(변환 프로그램)를 통해 기계가 읽을 수 있는 언어로 변환

### OPCODE

- 명령 (add) => 숫자(0011)로 변환

### ANSI code

- Only English / 1 byte (0~255)

### UNICODE-16

- 1문자 2byte (65535개를 표현 가능) / 모든 문자 표현 가능

### UTF-8 (UNICODE의 일부)

- Go에서의 default
- 1~3byte

## package main

### Go의 시작은 package로 시작이 되어야 한다

- code가 속한 package를 표시하기 위해서
- package는 코드를 묶는 단위이다

### main

- main은 프로그램 시작점을 가진 함수 포함하는 package 의미 이다

프로그램 --(load)--> Memory --(read code)--> CPU

- read code 부분에서 시적점(starting point)이 어디인지 알려 주는 역활
- starting point를 가지고 있는 package는 하나여야 한다

<b>1개의 main package와 여러 다양한 package로 구성이 되어 있다.</b>

## import "fmt" => 기능을 가지고 있다(해당 기능을 사용함)

- fmt라는 package를 사용하겠다

### import main은 가능한가?

- 불가능 하다

## func main(){}

- func : 함수선언
- main : 함수의 이름, keyword 역활을 가지고 있다(starting point)
