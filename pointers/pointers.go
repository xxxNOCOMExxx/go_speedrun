package main

import "fmt"

/*func main() {
	health := 50
	fmt.Println("Здоровье до зелья:", health)

	p_h := 50

	heal(&health, p_h)

	fmt.Println("Здоровье после зелья:", health)
	playerScore := 150
	resetScore(&playerScore)
	fmt.Println("Счет игрока:", playerScore) // Должно вывести 0

	var emptyScore *int    // Пустой указатель (равен nil)
	resetScore(emptyScore) // Программа НЕ должна упасть здесь!
	fmt.Println("Программа успешно завершилась!")

	x := 10
	y := 99
	fmt.Printf("До обмена: x = %d, y = %d\n", x, y)

	swap(&x, &y)

	fmt.Printf("После обмена: x = %d, y = %d\n", x, y) // Должно выве
}

func heal(hp *int, amount int) {
	*hp += amount
}

func resetScore(score *int) {
	// Твой код здесь:
	if score == nil {
		return
	} else {
		*score = 0
	}
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}
*/

type Server struct {
	Name     string
	RAM      int
	IsOnline bool
}

func main() {
	myServer := Server{
		Name:     "Prod-DB-01",
		RAM:      16,
		IsOnline: false,
	}

	fmt.Printf("До апгрейда: %+v\n", myServer)

	upgradeServer(&myServer, 64)
	fmt.Printf("После апгрейда: %+v\n", myServer) // Должно быть RAM: 64, IsOnline: true
}

func upgradeServer(srv *Server, newRAM int) {
	srv.IsOnline = true
	srv.RAM = newRAM
}
