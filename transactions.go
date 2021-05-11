package go_demo

import (
	"strconv"
	"time"
)

type TransactionId string
type CustomerId string

type Request struct {
	Id          string
	Customer_id string
	Load_amount string
	Time        string
}

type Transaction struct {
	Id          TransactionId
	Customer_id CustomerId
	Amount      float64
	Time        string
}

/**
	Create transaction from Request struct
**/
func createTransaction(request Request) (Transaction, error) {
	amount, err := strconv.ParseFloat(request.Load_amount[1:], 64)
	transaction := Transaction{
		Id:          TransactionId(request.Id),
		Customer_id: CustomerId(request.Customer_id),
		Amount:      amount,
		Time:        request.Time,
	}

	return transaction, err
}

/**
	This function check whether the system should accept the request of loading fund
        return:
                true : accept the request
                false: reject the request
**/
func shouldLoadFund(transaction Transaction, customers map[CustomerId][]Transaction) bool {
	loadFundTime, err := time.Parse(time.RFC3339, transaction.Time)
	if err != nil {
		return false
	}

	year, month, day := loadFundTime.UTC().Date()
	_, week := loadFundTime.UTC().ISOWeek()

	loadedAmoutOnTheDay := 0.0
	loadedTimesOnTheDay := 0
	loadedAmountOnTheWeek := 0.0

	if transaction.Amount > MaxLoadPerDay {
		return false
	}

	for _, loadedTransaction := range customers[transaction.Customer_id] {
		loadedTime, err := time.Parse(time.RFC3339, loadedTransaction.Time)
		// skip for invalid time
		if err != nil {
			continue
		}

		loadedYear, loadedMonth, loadedDay := loadedTime.UTC().Date()
		_, loadedWeek := loadedTime.UTC().ISOWeek()

		if year == loadedYear && month == loadedMonth && day == loadedDay {
			loadedAmoutOnTheDay += loadedTransaction.Amount
			loadedTimesOnTheDay += 1

			if loadedAmoutOnTheDay+transaction.Amount > MaxLoadPerDay || loadedTimesOnTheDay >= MaxLoadTimesPerDay {
				return false
			}
		}

		if year == loadedYear && week == loadedWeek {
			loadedAmountOnTheWeek += loadedTransaction.Amount

			if loadedAmountOnTheWeek+transaction.Amount > MaxLoadPerWeek {
				return false
			}
		}
	}

	return true
}
