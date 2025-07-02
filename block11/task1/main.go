package main

import "fmt"

type Students struct {
	name    string
	age     int
	course  int
	average float64
}

func NewStudent(Name string, Age int, Course int, Average float64) Students {
	return Students{
		name:    Name,
		age:     Age,
		course:  Course,
		average: Average,
	}
}

func (s Students) infoAboutStudent() {
	fmt.Printf("Студент: %s\n", s.name)
	fmt.Printf("Возраст: %d\n", s.age)
	fmt.Printf("Курс: %d\n", s.course)
	fmt.Printf("Средний балл: %.2f\n", s.average)
	fmt.Println()
}

func main() {
	//создаем 4 одарённых студента
	student1 := NewStudent("Антон", 19, 2, 4.2)
	student2 := NewStudent("Федор", 22, 4, 2.5)
	student3 := NewStudent("Паша", 20, 3, 3.2)
	student4 := NewStudent("Данил", 18, 1, 5)

	//узнаем информацию об одаренных учениках
	student1.infoAboutStudent()
	student2.infoAboutStudent()
	student3.infoAboutStudent()
	student4.infoAboutStudent()
}
