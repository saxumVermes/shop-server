package shop

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type ShoeServer interface {
	// Add returns the shoe that was successfully created, and nil
	Add(brand string, model string, price float32, colors string) (shoe Shoe, ok bool)
	// List lists all shoes in the shop
	List() (list []Shoe)
	DeleteById(id string) (success bool)
	// Find returns a shoe, and a boolean if it was found
	Find(id string) (shoe Shoe, found bool)
}

type Address struct {
	City string
	Zip  int
}

type Shoe struct {
	gorm.Model
	SID    string  `json:"id" gorm:"primary_key"`
	SModel string  `json:"model"`
	Brand  string  `json:"brand"`
	Price  float32 `json:"price"`
	Colors string  `json:"colors"`
}

func New(name string, city string, zip int) ShoeServer {
	if os.Getenv("SHOE_TEST_ENV") != "" {
		logrus.Infoln("Running without persistent database")
		return newInMem(name, city, zip)
	}
	logrus.Infoln("Running with persistent database, dialect:", Dialect)
	return newORM(name, city, zip)
}
