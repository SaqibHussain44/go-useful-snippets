package main

import (
	"fmt"
	"time"
)

//Car a car model
type Car struct {
	Tankcapacity int
	Fuel         int
}

//FillGas a callback function structure
type FillGas func(litresRequired int, b *Car) (updated bool, litresFilled int, err error)

//Refill accepts a calback and determines Fuel required
func (b *Car) Refill(limit int, callback FillGas) error {
	fmt.Println("New Refil Request")
	fmt.Println("limit: ", limit)
	fmt.Println("sleeping for 3 seconds..")
	time.Sleep(3 * time.Second)
	litresRequired := b.Tankcapacity - b.Fuel
	callback(litresRequired, b)
	return nil
}

func main() {
	myb := Car{}
	myb.Fuel = 220
	myb.Tankcapacity = 250

	//with an anonymous function
	err := myb.Refill(100, func(litresRequired int, b *Car) (updated bool, filled int, err error) {
		fmt.Println("Refil Gas called.. ")
		fmt.Println("litresRequired: ", litresRequired)
		b.Fuel = b.Fuel + litresRequired
		return true, litresRequired, nil
	})

	fmt.Println("Error in Refiling: ", err)
	fmt.Println("Update Fuel: ", myb.Fuel)

	//defining the function first and passing it as an argument
	myfunc := func(litresRequired int, b *Car) (updated bool, filled int, err error) {
		fmt.Println("Refil Gas called.. ")
		fmt.Println("litresRequired: ", litresRequired)
		b.Fuel = b.Fuel + litresRequired
		return true, litresRequired, nil
	}

	myb.Fuel = 10

	myb.Refill(100, myfunc)

	fmt.Println("Error in Refiling: ", err)
	fmt.Println("Update Fuel: ", myb.Fuel)

}
