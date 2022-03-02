package persistence

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	title                                    string `gorm:"primaryKey"`
	status, category, task_type, description string
	creation_date, completed_date            int64
}

func MigrateTables(db *gorm.DB) {
	db.AutoMigrate(&Task{})

}
