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

func ParseInput(x string) int {
	operador, _ := strconv.Atoi(x)
	return operador
}

func GetNumber(x string) (int, int) {
	fmt.Printf("Enter First Number: ")
	first_numbre := ParseInput(ReadInput())
	fmt.Printf("Enter Second Number: ")
	second_number := ParseInput(ReadInput())
	return first_numbre, second_number
}

func Response(result int) {
	fmt.Printf("The result is: %d \n", result)
}

func (c Calc) Calculate(option string) {
	switch option {
	case "1":
		a, b := GetNumber("sum")
		Response(c.Sum(a, b))
	case "2":
		a, b := GetNumber("subtract")
		Response(c.Subtract(a, b))
	case "3":
		a, b := GetNumber("multiply")
		Response(c.Multiply(a, b))
	case "4":
		a, b := GetNumber("divide")
		Response(c.Divide(a, b))
	default:
		fmt.Println("¡Bye!")
		os.Exit(0)
	}
}
