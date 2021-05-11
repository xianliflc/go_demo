package go_demo

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Id          string `json:"id"`
	Customer_id string `json:"customer_id"`
	Accepted    bool   `json:"accepted"`
}

/**
        Build and return the response string
**/
func getResponse(transactionId TransactionId, customerId CustomerId, result bool) string {
	response := Response{
		Id:          string(transactionId),
		Customer_id: string(customerId),
		Accepted:    result,
	}

	jsonString, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(jsonString)
}
