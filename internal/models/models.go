package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"

)

type DBModel struct {
	Order OrderModel
}

func InitDB(dataSourceName string) (*DBModel, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed to migrate database: %v", err)
	}
	err = db.AutoMigrate(&Order{}, &OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("Failed to migrate database %v", err)
	}
	dbModel := &DBModel{
		Order: OrderModel{DB: db},
	}
	 return dbModel, nil
}
