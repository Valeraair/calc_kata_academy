package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romanToAr(num string) int {
	romanmap := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	return romanmap[num]
}

func arToRoman(num int) string {
	values := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	rome := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var res string
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			res += rome[i]
			num -= values[i]
		}
	}
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Выдача паники, так как ввод некорректен!", err)
		return
	}

	parts := strings.Fields(input)
	if len(parts) > 3 {
		fmt.Println("Выдача паники, так как формат математической операции не удовлетворяет заданию — " +
			"два операнда и один оператор (+, -, /, *).")
		return
	} else if len(parts) < 3 {
		fmt.Println("Выдача паники, так как строка не является математической операцией.")
		return
	}

	var mode string
	var num1, num2 int

	if romanToAr(parts[0]) > 0 {
		if romanToAr(parts[2]) > 0 {
			num1, num2 = romanToAr(parts[0]), romanToAr(parts[2])
			mode = "rom"
		} else {
			fmt.Println("Выдача паники, так как используются одновременно разные системы счисления.")
			return
		}
	} else {
		if romanToAr(parts[2]) > 0 {
			fmt.Println("Выдача паники, так как используются одновременно разные системы счисления.")
			return
		}
	}

	if mode != "rom" {
		num1, err = strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Выдача паники, так как строка не является математической операцией.")
			return
		}

		num2, err = strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("Выдача паники, так как строка не является математической операцией.")
			return
		} else {
			mode = "ar"
		}
	}

	if mode == "ar" {
		if num1 > 10 || num2 > 10 || num1 <= 0 || num2 <= 0 {
			fmt.Println("Выдача паники, так как вводные данные выходят за условие задания (от 1 до 10)")
			return
		}
	}

	var (
		result  int
		resultR string
	)

	switch parts[1] {
	case "+":
		if mode == "rom" {
			resultR = arToRoman(num1 + num2)
		} else {
			result = num1 + num2
		}
	case "-":
		if mode == "rom" {
			if num1-num2 > 0 {
				resultR = arToRoman(num1 - num2)
			} else {
				fmt.Println("Выдача паники, так как в римской системе нет отрицательных чисел.")
				return
			}
		} else {
			result = num1 - num2
		}
	case "*":
		if mode == "rom" {
			resultR = arToRoman(num1 * num2)
		} else {
			result = num1 * num2
		}
	case "/":
		if num2 == 0 {
			fmt.Println("Выдача паники, так как деление на 0 в обычной арифметеке не имеет смысла")
			return
		}
		if mode == "rom" {
			resultR = arToRoman(num1 / num2)
		} else {
			result = num1 / num2
		}
	default:
		fmt.Println("Выдача паники, так как формат математической операции не удовлетворяет заданию — " +
			"два операнда и один оператор (+, -, /, *).")
		return
	}

	if mode == "rom" {
		fmt.Println(resultR)
	} else {
		fmt.Println(result)
	}
}
