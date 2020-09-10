package main

import "strings"

const CreditCardType = "creditcard"

const (
	CardNumberLabel   = "NUMBER"
	ExpiryDateLabel   = "EXPIRY DATE"
	CardholderLabel   = "CARDHOLDER"
	PinLabel          = "PIN"
	PasswordCardLabel = "LOGIN PASSWORD"
	CvvLabel          = "CVC"
)

type CreditCard struct {
}

func (creditCard *CreditCard) Generate(items []EnpassItem) [][]string {
	records := make([][]string, 0)

	records = append(records, []string{"title", "card number", "expiry date (MM/YYYY)", "cardholder name", "PIN", "bank name", "CVV", "notes"})

	for _, item := range items {
		// build the map type -> slice of values
		fieldValuesByLabel := make(map[string][]string, 0)
		for _, field := range item.Fields {
			// skip the field that contains an empty value
			if field.Value == "" {
				continue
			}

			// set the uppercase to Label because "Enpass" exported values contain different values
			label := strings.ToUpper(field.Label)

			if fieldValuesByLabel[label] == nil {
				fieldValuesByLabel[label] = []string{field.Value}
			} else {
				fieldValuesByLabel[label] = append(fieldValuesByLabel[label], field.Value)
			}
		}

		cardNumber := joinValue(fieldValuesByLabel[CardNumberLabel])

		//  пропускаем те карты, где не известен номер карты
		if cardNumber == "" {
			continue
		}

		expiryDate := joinValue(fieldValuesByLabel[ExpiryDateLabel])
		cardHolderName := joinValue(fieldValuesByLabel[CardholderLabel])
		pin := joinValue(fieldValuesByLabel[PinLabel])
		cvv := joinValue(fieldValuesByLabel[CvvLabel])

		if len(pin) > 0 {
			pin = pin + ", password:" + joinValue(fieldValuesByLabel[PasswordCardLabel])
		} else {
			pin = "password:" + joinValue(fieldValuesByLabel[PasswordCardLabel])
		}

		var notes string
		if item.Note != "" {
			notes = notes + item.Note
		}

		//println(fieldValuesByLabel)
		records = append(records, []string{item.Title, cardNumber, expiryDate, cardHolderName, pin, item.Title, cvv, notes})
	}

	return records
}

func (creditCard *CreditCard) Type() string {
	return CreditCardType
}
