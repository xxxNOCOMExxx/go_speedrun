package main

import "fmt"

func main() {
	num := make([]int, 0, 1250)
	num = append(num, 67, 67, 67, 67, 52)
	fmt.Println(symmator(num))
	fmt.Println(removeDuplicates(num))

	num2 := make([]float64, 0, 1250)
	num2 = append(num2, 67.67, 67.68, 52.14232, 512.0, 15.1, 18.15)
	fmt.Println(filterFloat64(num2))
}

func symmator(mass []int) int {
	total := 0

	for _, zxc := range mass {
		total += zxc
	}
	return total
}

func removeDuplicates(mass []int) []int {
	copy_mass := make([]int, 0, 1250)

	for _, zxc1 := range mass {
		i_here := false
		for _, zxc2 := range copy_mass {
			if zxc2 == zxc1 {
				i_here = true
				break
			}
		}
		if !i_here {
			copy_mass = append(copy_mass, zxc1)
		}
	}
	return copy_mass
}
func filterFloat64(mass []float64) ([]float64, string) {
	min_rating := 67.67
	copy_mass := make([]float64, 0, len(mass))
	for _, zxc := range mass {
		if zxc > min_rating {
			copy_mass = append(copy_mass, zxc)
		}
	}
	cursed := fmt.Sprintf("Ваш рейтинг состоит из %d элементов", len(copy_mass))
	return copy_mass, cursed
}
