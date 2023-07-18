package hw

// Вариант решения 2

import (
	"errors"
	"math"
)

// По условиям задачи, координаты не могут быть меньше 0.

type Point_v2 struct { //
	x, y float64 // Сделал точки приватные
}

func New(x, y float64) (Point_v2, error) {
	if x < 0 || y < 0 {
		return Point_v2{}, errors.New("координаты не могут быть меньше нуля")
	}
	return Point_v2{x: x, y: y}, nil
}

func (p *Point_v2) SetCoordinates(x, y float64) error {
	if x < 0 || y < 0 {
		return errors.New("координаты не могут быть меньше нуля")
	}
	p.x = x
	p.y = y
	return nil
}

func DBP_v2(p1, p2 Point_v2) (distance float64) { // Убрали условие и разнесли это проверку на методы создания и изменение точек
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.x-p1.x, 2))
}
