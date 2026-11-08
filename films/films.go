package main

import (
	"fmt"
)

func main() {
	catalog := map[string]float64{
		"The Shawshank Redemption": 9.3,
		"The Godfather":            9.2,
		"The Dark Knight":          9.0,
		"Pulp Fiction":             8.9,
		"Inception":                8.8,
		"Fight Club":               8.8,
		"The Matrix":               8.7,
		"Goodfellas":               8.7,
		"Seven":                    8.6,
		"Interstellar":             8.6,
		"Spirited Away":            8.6,
		"The Green Mile":           8.6,
		"Gladiator":                8.5,
		"The Departed":             8.5,
		"The Prestige":             8.5,
	}

	fmt.Println(addfilm(catalog))
	fmt.Println(mid_val(catalog))
	fmt.Println(most_cool(catalog))
	fmt.Println(most_cool2(catalog))
	fmt.Println(filterByRating(catalog, 8.9))

}

func addfilm(catalog map[string]float64) map[string]float64 {
	catalog["new film"] = 1.0
	return catalog
}

func mid_val(catalog map[string]float64) float64 {
	mid := 0.0
	for _, val := range catalog {
		mid += val
	}
	return mid / float64(len(catalog))
}
func most_cool(catalog map[string]float64) float64 {
	mid := 0.0
	for _, val := range catalog {
		if mid < val {
			mid = val
		}
	}
	return mid
}
func most_cool2(catalog map[string]float64) string {
	cool := most_cool(catalog)
	for name, num := range catalog {
		if num == cool {
			return name
		}
	}
	return ""
}
func filterByRating(catalog map[string]float64, limit float64) map[string]float64 {
	catalog_c := make(map[string]float64)
	for name, val := range catalog {
		if val > limit {
			catalog_c[name] = val
		}
	}
	return catalog_c
}
