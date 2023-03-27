// Created by: Charlotte Jhu
// Created on: March 2023
//
// This program calculates income tax

package main

import (
	"fmt"

	"github.com/leekchan/accounting"
)

func main() {
	// This function calculates income tax
	// variables
	var hoursWorked float64
	var hourlyWage float64

	// input
	fmt.Println("This program calculates income tax.")
	fmt.Println()
	fmt.Print("Enter the hours worked: ")
	fmt.Scanln(&hoursWorked)
	fmt.Print("Enter the hourly wage: ")
	fmt.Scanln(&hourlyWage)

	// process
	pay := (hoursWorked * hourlyWage) * (1 - 0.18)
	governmentPay := (hoursWorked * hourlyWage) * 0.18
	accountingFormater := accounting.Accounting{Symbol: "$", Precision: 2}
	fmt.Println(accountingFormater.FormatMoney(123456789.213123)) 

	// output
	fmt.Println()
	fmt.Println("Your pay after tax is: $", pay)
	fmt.Println("The government's portion is: $", governmentPay)
	fmt.Println("\nDone.")
}
