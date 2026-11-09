package main

import (
	"errors"
	"fmt"
)

func main() {
	users := map[string]string{
		"admin": "supersecret",
		"alex":  "12345",
	}

	// Попытка 1: Несуществующий юзер
	if err := loginAdmin(users, "hacker", "123"); err != nil {
		fmt.Println("Попытка 1:", err) // Выведет: доступ запрещен: пользователь не найден
	}

	// Попытка 2: Неверный пароль
	if err := loginAdmin(users, "alex", "00000"); err != nil {
		fmt.Println("Попытка 2:", err) // Выведет: доступ запрещен: неверный пароль
	}

	// Попытка 3: Успешный вход
	if err := loginAdmin(users, "alex", "12345"); err == nil {
		fmt.Println("Попытка 3: Успешный вход в систему!")
	}
}

func checkPassword(users map[string]string, login string, pass string) error {
	realPass, ok := users[login]
	if !ok {
		return errors.New("пользователь не найден")
	}
	if realPass != pass {
		return errors.New("неверный пароль")
	}

	return nil
}

func loginAdmin(users map[string]string, login string, pass string) error {
	err := checkPassword(users, login, pass)
	if err != nil {
		return fmt.Errorf("доступ запрещен: %w", err)
	}

	return nil
}
