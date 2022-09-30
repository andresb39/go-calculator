package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Calc struct{}

func (c Calc) Sum(a, b int) int {
	return a + b
}

func (c Calc) Subtract(a, b int) int {
	return a - b
}

func (c Calc) Multiply(a, b int) int {
	return a * b
}

func (c Calc) Divide(a, b int) int {
	return a / b
}

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func ParseInput(x string) (int, error) {
	result, err := strconv.Atoi(x)
	return result, err
}

func GetNumber(x string) (int, int, error) {
	fmt.Printf("Enter First Number: ")
	firstNumbre, err := ParseInput(ReadInput())
	if err != nil {
		fmt.Println("Error is not a number... please try again")
		os.Exit(2)
	}
	fmt.Printf("Enter Second Number: ")
	secondNumber, err := ParseInput(ReadInput())
	if err != nil {
		fmt.Println("Error is not a number ... please try again")
		os.Exit(2)
	}
	return firstNumbre, secondNumber, err
}

func Response(result int) {
	fmt.Printf("The result is: %d \n", result)
}

func (c Calc) Calculate(option string) {
	switch option {
	case "1":
		a, b, _ := GetNumber("sum")
		Response(c.Sum(a, b))
	case "2":
		a, b, _ := GetNumber("subtract")
		Response(c.Subtract(a, b))
	case "3":
		a, b, _ := GetNumber("multiply")
		Response(c.Multiply(a, b))
	case "4":
		a, b, _ := GetNumber("divide")
		Response(c.Divide(a, b))
	default:
		fmt.Println("¡Bye!")
		os.Exit(0)
	}

}
