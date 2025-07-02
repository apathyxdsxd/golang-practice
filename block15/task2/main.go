package main

import "fmt"

type Transport interface {
	Move() string
	Stop() string
}
type Car struct{}

func (car Car) Move() string { return "автомобиль едет" }
func (car Car) Stop() string { return "автомобилб остановился" }

type Bicycle struct{}

func (bic Bicycle) Move() string { return "велосипед едет" }
func (bic Bicycle) Stop() string { return "велосипед остановился" }

type Boat struct{}

func (bo Boat) Move() string { return "лодка плывёт" }
func (bo Boat) Stop() string { return "лодка остановилась" }
func main() {
	var car Car
	var bike Bicycle
	var boat Boat

	fmt.Println("автомобиль:")
	fmt.Printf("движение: %s\n", car.Move())
	fmt.Printf("остановка: %s\n\n", car.Stop())

	fmt.Println("велосипед:")
	fmt.Printf("движение: %s\n", bike.Move())
	fmt.Printf("остановка: %s\n\n", bike.Stop())

	fmt.Println("лодка:")
	fmt.Printf("движение: %s\n", boat.Move())
	fmt.Printf("остановка: %s\n\n", boat.Stop())
}
