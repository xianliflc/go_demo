// demo project demo.go
package go_demo

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

func processRequest(customers map[CustomerId][]Transaction, request Request) string {
	var customerId CustomerId = CustomerId(request.Customer_id)
	transaction, err := createTransaction(request)

	// invalid amount
	if err != nil {
		return getResponse(transaction.Id, transaction.Customer_id, false) + "\n"
	}

	_, ok := customers[customerId]
	if !ok {
		customers[customerId] = []Transaction{}
	}

	if shouldLoadFund(transaction, customers) {
		customers[customerId] = append(customers[customerId], transaction)
		return getResponse(transaction.Id, transaction.Customer_id, true) + "\n"
	} else {
		return getResponse(transaction.Id, transaction.Customer_id, false) + "\n"
	}
}

func main() {
	file, err := os.Open(InputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Each customer is a list of transactions
	customers := make(map[CustomerId][]Transaction)
	var response string = ""

	for scanner.Scan() {
		var temp Request
		_ = json.Unmarshal([]byte(scanner.Text()), &temp)

		response += processRequest(customers, temp)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if err := writeData(response); err != nil {
		log.Fatal(err)
	}
}
