// Created by: Charlotte Jhu
// Created on: March 2023
//
// This program calculates income tax

package main

import "fmt"

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
	roundedPay := fmt.Sprintf("%.2f\n", pay)
	governmentRoundedPay := fmt.Sprintf("%.2f", governmentPay)

	// output
	fmt.Println()
	fmt.Println("Your pay after tax is: $", roundedPay)
	fmt.Println("The government's portion is: $", governmentRoundedPay)
	fmt.Println("\nDone.")
}
