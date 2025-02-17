package main

import "fmt"
import "os"
import "strconv"
import "errors"

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	balanceText, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("Failed to find balance file")
	}

	var balance float64
	balance, err = strconv.ParseFloat(string(balanceText), 64)

	if err != nil {
		return 1000, errors.New("Failed to parse balance")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {

	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------------")
		// panic("Can't continue without balance")
	}
	fmt.Println("Welcome to GO Bank")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your account balance is:", accountBalance)
		case 2:
			fmt.Print("Your deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount!")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your new account balance is:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Your withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)
			if withdrawalAmount <= 0 {
				fmt.Println("Invalid withdrawal amount!")
				continue
			}
			if withdrawalAmount > accountBalance {
				fmt.Println("Insufficient funds!")
			} else {
				accountBalance -= withdrawalAmount
				fmt.Println("Your new account balance is:", accountBalance)
				writeBalanceToFile(accountBalance)
			}
		case 4:
			fmt.Println("Arrivederci!")
			fmt.Println("Thank you for banking with us!")
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}
