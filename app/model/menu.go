package model

import (
	"time"
	"fmt"
)

type Menu struct {
	ID	int      `json:"id"`
	UserID int `gorm:"not null"`
	// User User
	Time_id	int	`json:"time_id" gorm:"not null"`
	Date	string	`json:"date" gorm:"not null"`
	Created_at	time.Time	`json:"created_at"`
	Updated_at	time.Time	`json:"updated_at"`
}

func CreateMenu(menu *Menu) {
	DB.Select("UserID", "Time_id", "Date").Create(menu)
}

func FindMenuById(id int) Menu {
	var menu Menu
	DB.Where("ID", id).First(&menu)
	return menu
}

func DeleteMenu(m *Menu) error {
	if rows := DB.Where(m).Delete(&Menu{}).RowsAffected; rows == 0 {
		return fmt.Errorf("Could not find Todo (%v) to delete", m)
	}
	return nil
}

func UpdateMenu(m *Menu) error {
	DB.Save(m)
	return nil
}