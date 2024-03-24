package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Сделайте ввод: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		fmt.Println(text)

		var operation string = ""
		if strings.Contains(text, "*") {
			operation = "*"

		} else if strings.Contains(text, "/") {
			operation = "/"

		} else if strings.Contains(text, "+") {
			operation = "+"

		} else if strings.Contains(text, "-") {
			operation = "-"

		}

		result := strings.Split(text, operation)

		num1 := result[0]
		num2 := result[1]

		if isRoman(num1) && isRoman(num2) {
			num11 := translate(num1)
			num12 := translate(num2)
			arabicResult := calculate(num11, num12, operation)
			if arabicResult <= 0 {
				panic("Невозможное число")
			}
			fmt.Println(intToRoman(arabicResult))

		} else if isArabic(num1) && isArabic(num2) {
			num11, _ := strconv.Atoi(num1)
			num12, _ := strconv.Atoi(num2)
			fmt.Println(calculate(num11, num12, operation))
		} else {
			panic("Числа должны быть целыми, арабскими или римскими")
		}
	}

}
func isRoman(num string) bool {
	romanNumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, romanNumeral := range romanNumerals {
		if num == romanNumeral {
			return true
		}
	}
	return false
}

func isArabic(num string) bool {
	arabicNumerals := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	for _, arabicNumeral := range arabicNumerals {
		if num == arabicNumeral {
			return true
		}
	}
	return false
}

func calculate(num1 int, num2 int, operation string) int {

	switch operation {

	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 == 0 {
			panic("Ошибка!Деление на ноль невозможно")
		}
		return num1 / num2
	}
	panic("Error")
}

func translate(num1 string) int {

	romanNumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	arabicNumerals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, romanNumeral := range romanNumerals {

		if romanNumeral == num1 {
			return arabicNumerals[i]

		}
	}
	panic("Error, разраб долбоеб")

}
func intToRoman(num int) string {
	var roman string = ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num {
			roman += romans[index]
			num -= numbers[index]
		}
		index -= 1
	}

	return roman
}
