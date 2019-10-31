package main

const CreditCardType = "creditcard"

type CreditCard struct {
}

func (creditCard *CreditCard) Generate(items []EnpassItem) [][]string {
	records := make([][]string, 0)

	records = append(records, []string{"title", "card number", "expiry date (MM/YYYY)", "cardholder name", "PIN", "bank name", "CVV", "notes"})

	// TODO have to implement the body of function

	return records
}

func (creditCard *CreditCard) Type() string {
	return CreditCardType
}
