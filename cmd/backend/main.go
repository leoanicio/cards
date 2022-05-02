package main

import (
	"fmt"

	"github.com/leoanicio/cards/pkg/models/bill"
)

func main() {
	myBill := bill.NewBill("Leonardo")

	myBill.addItem("Carne", 20.0)
	myBill.addItem("Arroz", 6.75)
	myBill.computeTotal()

	fmt.Println(myBill)
}
