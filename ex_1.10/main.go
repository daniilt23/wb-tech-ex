package main

import (
	"fmt"
	"math"
)

// сделаем функцию для округления в меньшую сторону числа
func floorByTen(value float64) float64 {
	if value > 0 {
		return math.Floor(value/10) * 10
	}
	return math.Ceil(value/10) * 10
}

func main() {
	mas := [8]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[float64][]float64)
	var key float64
	for _, val := range mas {
		key = floorByTen(val)              // находим ключ
		map_val := groups[key]             // получаем текущий слайл для данной группы
		groups[key] = append(map_val, val) // в конкретное значение мапы пишем данные из массива
	}

	fmt.Printf("%v", groups)
}
