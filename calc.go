package main

import (
 "bufio"
 "fmt"
 "os"
 "strings"
)

func convertToArabic(input string, arabicNumbers map[string]int, romanNumbers map[string]int) (int, error) {
 if val, ok := arabicNumbers[input]; ok {
  return val, nil
 } else if val, ok := romanNumbers[input]; ok {
  return val, nil
 }
 return 0, fmt.Errorf("неверное число: %s", input)
}

func convertToRoman(num int) string {
  if num <= 0 || num > 100 {
   return ""
  }
 
  thousands := []string{"", "M", "MM", "MMM"}
  hundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
  tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
  ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
 
  return thousands[num/1000] + hundreds[(num%1000)/100] + tens[(num%100)/10] + ones[num%10]
 }

func calculate(a, op, b string) (interface{}, error) {
 arabicNumbers := map[string]int{
  "1": 1, "2": 2, "3": 3, "4": 4, "5": 5,
  "6": 6, "7": 7, "8": 8, "9": 9, "10": 10,
 }
 romanNumbers := map[string]int{
  "I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
  "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
 }

 isArabic := false
 for _, char := range a {
  if char >= '0' && char <= '9' {
   isArabic = true
   break
  }
 }

 numA, err := convertToArabic(a, arabicNumbers, romanNumbers)
 if err != nil {
  return nil, err
 }

 numB, err := convertToArabic(b, arabicNumbers, romanNumbers)
 if err != nil {
  return nil, err
 }

 var result int
 switch op {
 case "+":
  result = numA + numB
 case "-":
  result = numA - numB
 case "*":
  result = numA * numB
 case "/":
  result = numA / numB
 default:
  return nil, fmt.Errorf("неверная операция: %s", op)
 }

 if !isArabic {
  if result <= 0 {
   return nil, fmt.Errorf("результат должен быть положительным числом или ноль")
  }
  if result <= 10{
	return convertToRoman(result), nil
  }
 }

 return result, nil
}

func main() {
 var input string
 fmt.Println("Введите математическую операцию (например, 1 + 2):")
 reader := bufio.NewReader(os.Stdin)
 input, _ = reader.ReadString('\n')
 input = strings.TrimSpace(input)

 parts := strings.Split(input, " ")
 if len(parts) != 3 {
  fmt.Println("Неверный формат математической операции")
  os.Exit(0)
 }

 a, op, b := parts[0], parts[1], parts[2]
 result, err := calculate(a, op, b)
 if err != nil {
  fmt.Println("Ошибка расчета:", err)
  os.Exit(0)
 }

 fmt.Println("Результат:", result)
}
