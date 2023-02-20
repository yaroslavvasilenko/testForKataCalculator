package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите значение")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		textSlice := strings.Split(text, " ")
		if len(textSlice) != 3 {
			log.Panic(fmt.Errorf(
				"%s строка не является математической операцией или не удовлетворяет условию два операнда и один оператор (+, -, /, *)", textSlice))
		}

		a, b, isRoman, err := handlerRomanAndInt(textSlice[0], textSlice[2])
		if err != nil {
			log.Panic(err)
		}
		if !((0 < a && a < 11) && (0 < b && b < 11)) {
			log.Panic(fmt.Errorf("a = %v, b = %v, не выполняется условие 0 < x < 11", a, b))
		}

		var result int

		switch textSlice[1] {
		case "+":
			result = a + b
		case "-":
			result = a - b
		case "*":
			result = a * b
		case "/":
			result = a / b
		default:
			log.Panic(fmt.Errorf("оператор - %v недоступен, доступые: +, -, /, *", textSlice[1]))
		}

		if isRoman {
			resultRoman, exception := romanToIntDefend(result)
			if exception != nil {
				fmt.Println(*exception)
				continue
			}
			fmt.Println(resultRoman)
		} else {
			fmt.Println(result)
		}

	}

}

// defines Arabic digits or Roman digits
// accepts two Roman or two Arabic digits, if one Roman and one Arabic digits - return error
func handlerRomanAndInt(a, b string) (int, int, bool, error) {
	var aInt, bInt, countRoman int
	var isRoman bool = false
	aInt, aIsInt := isArabic(a)
	if !aIsInt {
		aInt = romanToInt(a)
		countRoman++
	}
	bInt, bIsInt := isArabic(b)
	if !bIsInt {
		bInt = romanToInt(b)
		countRoman++
	}

	if countRoman == 1 {
		err := errors.New(
			"используются одновременно арабские и римский цифры")
		return 0, 0, false, err
	} else if countRoman == 2 {
		isRoman = true
	}

	return aInt, bInt, isRoman, nil
}

// define exception Roman numeral system and convert type to string
func romanToIntDefend(intRoman int) (string, *string) {
	if intRoman < 0 {
		exception := "в римской системе нет отрицательных чисел"
		return "", &exception
	} else if intRoman == 0 {
		exception := "в римской системе нет нуля"
		return "", &exception
	}

	return intToRoman(intRoman), nil

}
