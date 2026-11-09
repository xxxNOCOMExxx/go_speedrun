package main

import (
	"fmt"
	"strings"
)

func main() {
	// руны и то как их переводить в стринг
	s := "Hello world"
	z := "dead souls"
	r := "Привет"
	fmt.Println("Magic strating now")
	fmt.Println(translator(char_count(r)))
	fmt.Println(translator(char_count(z)))
	fmt.Println(translator(char_count(s)))

	//дубликаты
	num := make([]int, 0, 1250)
	num = append(num, 67, 67, 67, 67, 52)
	fmt.Println(booldublicate(num))

	// попытки понять что такое ok
	// 1
	names := []string{"Иван", "Петр", "Иван", "Ольга", "Петр", "Иван"}
	fmt.Println(count_voice(names))

	//2
	inst := map[string]bool{"git": true, "docker": true, "go": true}
	true_in := make([]string, 0, 5)
	true_in = append(true_in, "git", "docker", "postgresql")
	fmt.Println(view_true_inst(inst, true_in))
	//3
	slices := []int{4, 5, 9, 4, 5}
	fmt.Println(unique_num(slices))
	//4
	s1 := "God"
	s2 := "Dog"
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
	fmt.Println(s1, s2)

	fmt.Println(anogramma(s1, s2))
}

func char_count(str string) map[rune]int {
	mapa := make(map[rune]int)
	for _, sym := range str {
		mapa[sym]++
	}
	return mapa
}

func translator(result map[rune]int) string {
	var true_answer string
	for char, count := range result {
		answer := fmt.Sprintf("Буква '%c' встретилась: %d раз(а)\n", char, count)
		true_answer += answer
	}
	return true_answer
}
func booldublicate(num []int) bool {
	watch := make(map[int]bool)

	for _, v := range num {
		if watch[v] {
			return true
		}
		watch[v] = true
	}
	return false
}

func count_voice(names []string) map[string]int {
	vc := make(map[string]int)
	for _, cout := range names {
		vc[cout]++
	}
	return vc
}

func view_true_inst(inst map[string]bool, true_in []string) bool {
	for _, here := range true_in {
		if _, ok := inst[here]; !ok {
			return false
		}
	}
	return true
}

func unique_num(slice []int) int {
	ded_man := make(map[int]int)
	for _, ded := range slice {
		ded_man[ded]++
	}

	for num, ded := range ded_man {
		if ded == 1 {
			return num
		}
	}
	return 0
}
func anogramma(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	count := make(map[rune]int)
	for _, ch := range s1 {
		count[ch]++
	}
	for _, ch := range s2 {
		count[ch]--
	}
	for _, ch := range count {
		if ch != 0 {
			return false
		}
	}
	return true
}
