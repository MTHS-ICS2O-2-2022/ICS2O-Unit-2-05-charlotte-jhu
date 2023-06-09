// Copyright (c) 2023 Charlotte Jhu All rights reserved
//
// Created by: Charlotte Jhu
// Created on: March 2023
// This file contains the JS functions for index.html

'use strict'
/**
 * This function calculates how much income tax you have to pay
 */
function myButtonClicked() {
  // input
  const hoursWorked = parseFloat(document.getElementById("hours-worked").value)
  const hourlyWage = parseFloat(document.getElementById("hourly-wage").value)
  const TAX_RATE = 0.18

  // process
  const pay = (hoursWorked * hourlyWage) * (1 - TAX_RATE)
  const governmentPay = (hoursWorked * hourlyWage) * TAX_RATE

  //output
  document.getElementById("your-pay").innerHTML = "You get $ " + pay.toFixed(2)
  document.getElementById("government-pay").innerHTML = "The government gets $" + governmentPay.toFixed(2)
}
