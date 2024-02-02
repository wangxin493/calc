package hello

import "fmt"

func Comp(num1 ,num2 int8, sign string) {
	switch sign {
	case "add":
		fmt.Printf("%d + %d = %d", num1, num2, num1 + num2);
	case "sub":
		fmt.Printf("%d + %d = %d", num1, num2, num1 - num2);
	}
}