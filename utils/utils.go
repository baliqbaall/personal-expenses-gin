package utils

import (
	"encoding/json"
	"io/ioutil"
	"personal_expanses/expenses"
)

func SaveExpenseToFile(filename string, expenses []expenses.Expense) error {
	data, err := json.Marshal(expenses)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func LoadExpensesFromFile(filename string) ([]expenses.Expense, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var expenses []expenses.Expense
	json.Unmarshal(data, &expenses)
	return expenses, err
}
