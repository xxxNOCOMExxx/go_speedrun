/*package main

import (
	"errors"
	"fmt"
)

func main() {
	// 3. Главная функция получает ошибку и решает, что делать (пишет в консоль)
	err := buyCoffee(300)
	if err != nil {
		fmt.Println("Не удалось купить кофе:", err)
		return
	}
	fmt.Println("Кофе успешно куплен!")
}

func buyCoffee(money int) error {
	// 2. Вызываем оплату. Если она вернула ошибку — пробрасываем её выше
	err := pay(money, 250)
	if err != nil {
		return fmt.Errorf("ошибка оплаты: %w", err)
	}
	return nil
}

func pay(balance int, price int) error {
	// 1. Самый нижний уровень. Проверяем условие и создаем ошибку
	if balance < price {
		return errors.New("недостаточно средств на счете")
	}
	return nil // Ошибок нет
}
*/
