package expenses

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	ID       int       `json:"id" gorm:"primaryKey"`
	Amount   float64   `json:"amount"`
	Category string    `json:"category"`
	Date     time.Time `json:"date"`
}

// var nextID int = 1

// AddExpense menambahkan pengeluaran baru dan mengembalikan struct Expense
func AddExpense(db *gorm.DB, amount float64, category string, date time.Time) Expense {
	expense := Expense{
		Amount:   amount,
		Category: category,
		Date:     date,
	}
	db.Create(&expense)
	return expense
}

// GetExpense mengembalikan semua pengeluaran yang ada
func GetExpenses(db *gorm.DB) []Expense {
	var expenses []Expense
	db.Find(&expenses)
	return expenses
}

// UpdateExpense memmperbarui pengeluaran berdasarkan ID dan mengembalikan error jika ID tidak ditemukan
func UpdateExpense(db *gorm.DB, id int, amount float64, category string, date time.Time) error {
	var expense Expense
	result := db.First(&expense, id)
	if result.Error != nil {
		return errors.New("expense not found")
	}
	expense.Amount = amount
	expense.Category = category
	expense.Date = date
	db.Save(&expense)
	return nil
}

// DeleteExpense menghapus pengeluaran berdasarkan ID dan mengebalikan error jika ID tidak ditemukan
func DeleteExpense(db *gorm.DB, id int) error {
	result := db.Delete(&Expense{}, id)
	if result.Error != nil {
		return errors.New("expense not found")
	}
	return nil
}
