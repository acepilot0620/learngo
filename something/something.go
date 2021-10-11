package something

import "fmt"

func SayHello(){ // 함수의 첫글자가 대문자면 Public 함수
	fmt.Println("Hello World")
}

func sayBye(){ // 함수의 첫글자가 소문자면 Private 함수
	fmt.Println("Bye")
}