package main

import (
	"fmt"
	"net/http"
	"personal_expanses/expenses"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Aplikasi Pencatatan Pengeluaran Pribadi")

	// Konfigurasi koneksi database MySQL
	dsn := "root:@tcp(127.0.0.1:3306)/personal_expenses?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrasi skema database
	db.AutoMigrate(&expenses.Expense{})

	// Inisialisasi Gin router
	r := gin.Default()

	// Define the router and their handlers
	r.GET("/expenses", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, expenses.GetExpenses(db))
	})

	r.POST("/expenses", func(ctx *gin.Context) {
		var expense expenses.Expense
		if err := ctx.ShouldBindJSON(&expense); err == nil {
			expense.Date = time.Now() // set tanggal saat ini
			expenses.AddExpense(db, expense.Amount, expense.Category, expense.Date)
			ctx.JSON(http.StatusOK, gin.H{"status": "success"})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	r.PUT("/expenses/:id", func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expense ID"})
			return
		}

		var expense expenses.Expense
		if err := ctx.ShouldBindJSON(&expense); err == nil {
			err := expenses.UpdateExpense(db, id, expense.Amount, expense.Category, time.Now())
			if err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"status": "success"})
			}
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	r.DELETE("/expenses/:id", func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expense ID"})
			return
		}

		err = expenses.DeleteExpense(db, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"status": "success"})
		}
	})

	r.Run() // Menjalankan server
}

// formatRupiah memformat jumlah uang dalam format rupiah
func formatRupiah(amount float64) string {
	p := message.NewPrinter(language.Indonesian)
	return p.Sprintf("Rp%.0f", amount)
}
