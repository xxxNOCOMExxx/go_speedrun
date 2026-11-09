/*package main

import (
	"errors"
	"fmt"
)

func main() {
	// Проверка 1: Деление на ноль
	res, err := divide(10, 0)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", res)
	}

	// Проверка 2: Нормальное деление
	res2, err2 := divide(10, 2)
	if err2 != nil {
		fmt.Println("Ошибка:", err2)
	} else {
		fmt.Println("Результат:", res2)
	}
}

func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("На ноль нельзя")
	}
	return a / b, nil
}
*/
