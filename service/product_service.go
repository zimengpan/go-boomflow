package service

import (
	"time"

	"github.com/zimengpan/go-boomflow/models"
	//"github.com/zimengpan/go-boomflow/models/mysql"
)

var mockProductDB1 = map[string]*models.Product{
	"1": &models.Product{
		"1",
		time.Date(2019, time.November, 10, 23, 0, 0, 0, time.UTC),
		time.Date(2019, time.November, 10, 23, 0, 0, 0, time.UTC),
		"A",
		"B",
	},
}

var mockProductDB2 = map[string]*models.Product{
	"AB": &models.Product{
		"1",
		time.Date(2019, time.November, 10, 23, 0, 0, 0, time.UTC),
		time.Date(2019, time.November, 10, 23, 0, 0, 0, time.UTC),
		"A",
		"B",
	},
}

func GetProductById(id string) (*models.Product, error) {
	//return mysql.SharedStore().GetProductById(id)
	return mockProductDB1[id], nil
}

func GetProductByAssetPair(assetA string, assetB string) (*models.Product, error) {
	//return mysql.SharedStore().GetProductById(id)
	aA, _ := GetAssetByAssetData(assetA)
	aB, _ := GetAssetByAssetData(assetB)

	symbolA := aA.Currency
	symbolB := aB.Currency

	key := symbolA + symbolB
	if symbolA > symbolB {
		key = symbolB + symbolA
	}

	return mockProductDB2[key], nil
}

func GetProducts() ([]*models.Product, error) {
	return []*models.Product{&models.Product{
		"1",
		time.Date(2019, time.November, 10, 23, 0, 0, 0, time.UTC),
		time.Date(2019, time.November, 10, 23, 0, 0, 0, time.UTC),
		"A",
		"B",
	}}, nil
	//return mysql.SharedStore().GetProducts()
}
