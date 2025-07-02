package main

import "fmt"

type engine struct {
	power  int     //мощность
	volume float64 //объём
}
type auto struct {
	color      string
	Model      string
	Year       int
	EngineInfo engine
}

func Info(c auto) {
	fmt.Printf("модель:%s\n", c.Model)
	fmt.Printf("год выпуска: %d\n", c.Year)
	fmt.Printf("цвет: %s\n", c.color)
	fmt.Printf("мощность двигателя: %d\n", c.EngineInfo.power)
	fmt.Printf("объём: %.2f\n", c.EngineInfo.volume)
	fmt.Println()
}
func main() {
	//создаем двигатель
	superEngine := engine{300, 3}
	poorEngine := engine{20, 1}
	//создаем автомобили
	car1 := auto{"red", "ferrari", 2000, superEngine}
	car2 := auto{"yellow", "dodge", 1990, superEngine}
	car3 := auto{"white", "lada", 1986, poorEngine}

	Info(car1)
	Info(car2)
	Info(car3)

}
