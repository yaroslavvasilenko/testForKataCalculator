package main

import (
	"strconv"
	"strings"
)

// define Arabic digits and convert type to int
func isArabic(a string) (int, bool) {
	var isArDec = true
	aInt, err := strconv.Atoi(a)
	if err != nil {
		isArDec = false
	}

	return aInt, isArDec
}

// Roman digits switch to Arabic digits
func intToRoman(num int) string {
	changerToRoman := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, val := range changerToRoman {
		for num >= val.value {
			roman.WriteString(val.digit)
			num -= val.value
		}
	}

	return roman.String()
}

// Arabic december switch to Roman december
func romanToInt(roman string) int {
	translateRoman := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var decNum, tmpNum int
	for i := len(roman) - 1; i >= 0; i-- {
		romanDigit := roman[i]
		decDigit := translateRoman[romanDigit]
		if decDigit < tmpNum {
			decNum -= decDigit
		} else {
			decNum += decDigit
			tmpNum = decDigit
		}
	}
	return decNum
}
