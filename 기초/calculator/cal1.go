package main

import (
	"bufio" // 한줄을 입력 받기 위해 reader strim할때 필요
	"fmt"
	"os"      // 표준 입력 받기
	"strconv" // 문자열을 숫자로 타입 변환할때 필요함
	"strings" // string을 한줄 입력 받았을때 찌꺼기 제거 할때 필요함
)

func main2() {

	fmt.Println("숫자를 입력하세요")
	reader := bufio.NewReader(os.Stdin) //입력 받음 os.Stdin os에서 표준 입력 받음
	line, _ := reader.ReadString('\n')  //입력 받고"\n"문자가 나올때 까지 읽음  line변수에 한줄 읽은 값을 넣어줌 _ 에러는 처리 하지 않음 이름 없는 변수
	line = strings.TrimSpace(line)      //개행 제거 (찌거기 제거)

	n1, _ := strconv.Atoi(line) //문자를 숫자로 변환함

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n2, _ := strconv.Atoi(line)

	fmt.Printf("입력하신 숫자는 %d, %d 입니다.\n", n1, n2)

	fmt.Println("연산자를 입력하세요")

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	// if line == "+" {
	// 	fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	// } else if line == "-" {
	// 	fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	// } else if line == "*" {
	// 	fmt.Printf("%d * %d = %d", n1, n2, n1*n2)
	// } else if line == "/" {
	// 	fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
	// } else if line == "%" {
	// 	fmt.Print(n1, " % ")
	// 	fmt.Printf("%d = %d", n2, n1%n2)
	// } else {
	// 	fmt.Println("잘못 입력하셨습니다.")
	// }

	switch line {
	case "+":
		fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	case "-":
		fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	case "*":
		fmt.Printf("%d * %d = %d", n1, n2, n1*n2)
	case "/":
		fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
	case "%":
		fmt.Print(n1, " % ")
		fmt.Printf("%d = %d", n2, n1%n2)
	default:
		fmt.Println("잘못 입력하셨습니다.")
	}
}
