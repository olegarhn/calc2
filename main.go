package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Считываем строку с выражением из консоли
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите выражение:")
	scanner.Scan()
	expression := scanner.Text()
	expression = strings.TrimSpace(expression)

	// Разделяем выражение на операнды и оператор
	str1, operator, str2 := parseExpression(expression)

	// Проверяем, что строки в кавычках и длина строки <= 10 символов
	if !isValidString(str1) || len(extractString(str1)) > 10 || (operator == "+" || operator == "-") && (!isValidString(str2) || len(extractString(str2)) > 10) {
		panic("Неверный формат строк")
	}

	// Проверяем число от 1 до 10 включительно для операций * и /
	number := extractNumber(str2)
	if (operator == "*" || operator == "/") && (number < 1 || number > 10) {
		panic("Неверный формат числа")
	}

	var result string
	switch operator {
	case "+":
		result = addStrings(str1, str2)
	case "-":
		result = subtractStrings(str1, str2)
	case "*":
		result = multiplyString(str1, str2)
	case "/":
		result = divideString(str1, str2)
	default:
		panic("Неподдерживаемая операция")
	}

	// Проверяем длину результата и обрезаем до 40 символов, если нужно
	if len(result) > 40 {
		result = result[:40] + "..."
	}

	// Выводим результат в кавычках
	fmt.Println("\"" + result + "\"")
}

func parseExpression(expression string) (string, string, string) {
	operators := []string{"+", "-", "*", "/"}
	for _, op := range operators {
		pos := strings.Index(expression, " "+op+" ")
		if pos != -1 {
			str1 := expression[:pos]
			str2 := expression[pos+len(op)+2:]
			return strings.TrimSpace(str1), op, strings.TrimSpace(str2)
		}
	}
	return "", "", ""
}

func isValidString(s string) bool {
	return len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"'
}

func extractString(s string) string {
	return s[1 : len(s)-1]
}

func extractNumber(s string) int {
	num := 0
	for i := 0; i < len(s); i++ {
		digit := s[i]
		if digit >= '0' && digit <= '9' {
			num = num*10 + int(digit-'0')
		}
	}
	return num
}

func addStrings(str1, str2 string) string {
	return str1[1:len(str1)-1] + str2[1:len(str2)-1]
}

func subtractStrings(str1, str2 string) string {
	innerStr1 := str1[1 : len(str1)-1]
	innerStr2 := str2[1 : len(str2)-1]
	if strings.Contains(innerStr1, innerStr2) {
		return strings.Replace(innerStr1, innerStr2, "", 1)
	}
	return innerStr1
}

func multiplyString(str, numStr string) string {
	innerStr := str[1 : len(str)-1]
	num := extractNumber(numStr)
	return strings.Repeat(innerStr, num)
}

func divideString(str, numStr string) string {
	innerStr := str[1 : len(str)-1]
	num := extractNumber(numStr)
	if num == 0 {
		return ""
	}
	return innerStr[:len(innerStr)/num]
}
