package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
		fmt.Println("------ BODY MASS INDEX ------")
	for {
		userHeight, userWeight := getUserInput()
		IMT, err := calculateIMT(userHeight, userWeight)	
		if err != nil {
			fmt.Println(err)
			continue
		}
		outputResult(IMT)
		isRepeatCalc := checkRepeatCalc()
		if !isRepeatCalc {
			break
			}
		}
	}

func getUserInput() (float64, float64) {
	var userHeight float64 
	var userWeight float64
	fmt.Print("Enter your height, sm: ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your weight, kg: ")
	fmt.Scan(&userWeight)
	return userHeight, userWeight
}

func calculateIMT(userHeight, userWeight float64) (float64, error) {
	if userHeight <= 0 || userWeight <= 0 {
		return 0, errors.New("Wrong data!")	
	}
	const	IMTPower = 2
	IMT := userWeight / math.Pow(userHeight / 100, IMTPower)
	return IMT, nil
}

func outputResult(IMT float64) {
	fmt.Printf("Your Body Mass Index is %.0f \n", IMT)
	
	switch {
		case IMT < 16:
			println("You are severely underweigh")
		case IMT < 18.5:
			println("You are underweight")
		case IMT < 25:
			println("You have a normal body weight")	
		case IMT < 30:
			println("You are overweight")		
		default:
			println("You have obesity")
		}
}
		
func checkRepeatCalc() bool {
	var userChoise string
	fmt.Println("Do you want to do the calculation again? (Y/n)")
	fmt.Scan(&userChoise)

	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}
