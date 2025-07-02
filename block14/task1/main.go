package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name      string
	BirthDate time.Time
	Grades    []float64
}

func (s *Student) Age() int {
	now := time.Now()
	age := now.Year() - s.BirthDate.Year()
	if now.Month() < s.BirthDate.Month() || (now.Month() == s.BirthDate.Month() && now.Day() < s.BirthDate.Day()) {
		age--
	}
	return age
}

func (s *Student) Status() string {
	var sum float64
	for _, grade := range s.Grades {
		sum += grade
	}
	averageGrade := sum / float64(len(s.Grades))
	switch {
	case averageGrade >= 4.5:
		return "отличник"
	case averageGrade >= 3.5:
		return "хорошист"
	default:
		return "троечник"
	}
}

func main() {
	student := Student{
		Name:      "Павел Иванов",
		BirthDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
		Grades:    []float64{5, 4, 5, 3, 4},
	}
	fmt.Println("Возраст:", student.Age())
	fmt.Println("Статус:", student.Status())
}
