package main

import "fmt"

func main() {
	transactions := []float64{}

	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)	
	}
	ballance := calculateBallance(transactions)
	fmt.Printf("Your ballance: %.2f", ballance)
}

func scanTransaction() float64 {
	fmt.Print("Enter You Transaction: ")
	var transaction float64
	fmt.Scan(&transaction)
	return transaction
}

func calculateBallance(transactions []float64) float64 {
		ballance := 0.0	

		for _, value := range transactions {
			ballance += value
		}	
		return ballance
	}


