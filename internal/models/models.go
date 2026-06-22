package models

import (
	"fmt"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBModel struct {
	Order OrderModel
	User  UserModel
	DB    *gorm.DB
}

func InitDB(dataSourceName string) (*DBModel, error) {
	// First connect to default 'postgres' database to create 'pizza_shop' if it doesn't exist
	defaultDSN := strings.Replace(dataSourceName, "dbname=pizza_shop", "dbname=postgres", 1)
	
	defaultDB, err := gorm.Open(postgres.Open(defaultDSN), &gorm.Config{})
	if err == nil {
		var count int64
		defaultDB.Raw("SELECT count(*) FROM pg_database WHERE datname = 'pizza_shop'").Scan(&count)
		if count == 0 {
			defaultDB.Exec("CREATE DATABASE pizza_shop;")
		}
		sqlDB, _ := defaultDB.DB()
		if sqlDB != nil {
			sqlDB.Close()
		}
	}

	// Now connect to the actual pizza_shop database
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&Order{}, &OrderItem{}, &User{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database %v", err)
	}

	dbModel := &DBModel{
		DB:    db,
		Order: OrderModel{DB: db},
		User:  UserModel{DB: db},
	}

	return dbModel, nil
}
