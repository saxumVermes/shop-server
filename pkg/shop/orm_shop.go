package shop

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/mattn/go-sqlite3"
)

var ErrShoeExists = fmt.Errorf("shoe already exists")
var Dialect string
var DBUri string

type ORMShoeServer struct {
	Name string
	Address
}

func newORM(name string, city string, zip int) ORMShoeServer {
	return ORMShoeServer{
		Name: name,
		Address: Address{
			City: city,
			Zip:  zip,
		},
	}
}

func (s ORMShoeServer) Add(brand, model string, price float32, colors string) (shoe Shoe, ok bool) {
	db, err := gorm.Open(Dialect, DBUri)
	if err != nil {
		panic("failed to connect to database")
	}
	defer func() {
		err = db.Close()
		if err != nil {
			fmt.Errorf("can not close database: %v", err)
		}
	}()

	var lastShoe Shoe
	db.Last(&lastShoe)
	if lastShoe.ID == 0 {
		lastShoe.ID = 1
	}
	id := fmt.Sprintf("%3.3s-%03d", strings.ToLower(strings.TrimSpace(brand)), lastShoe.ID+1)
	shoe = Shoe{
		SID:    id,
		SModel: model,
		Brand:  brand,
		Price:  price,
		Colors: colors,
	}

	db.AutoMigrate(&Shoe{})
	if !db.NewRecord(shoe) {
		return Shoe{}, false
	}
	db.Create(&shoe)

	return shoe, true
}

func (s ORMShoeServer) List() []Shoe {
	db, err := gorm.Open(Dialect, DBUri)
	if err != nil {
		panic("failed to connect to database")
	}
	defer func() {
		err = db.Close()
		if err != nil {
			fmt.Errorf("listing: can not close database: %v", err)
		}
	}()

	shoe := Shoe{}
	shoes := []Shoe{}
	db.Model(&shoe).Find(&shoes)

	return shoes
}

func (s ORMShoeServer) DeleteById(id string) (success bool) {
	db, err := gorm.Open(Dialect, DBUri)
	if err != nil {
		panic("failed to connect to database")
	}
	defer func() {
		err = db.Close()
		if err != nil {
			fmt.Errorf("deletion: can not close database: %v", err)
		}
	}()

	shoe := Shoe{}
	if db.Find(&shoe).Where("s_id = ?", id).Delete(&shoe).Error != nil {
		fmt.Println(shoe)
		return false
	}
	return true
}

func (s ORMShoeServer) Find(id string) (Shoe, bool) {
	db, err := gorm.Open(Dialect, DBUri)
	if err != nil {
		panic("failed to connect to database")
	}
	defer func() {
		err = db.Close()
		if err != nil {
			fmt.Errorf("deletion: can not close database: %v", err)
		}
	}()

	shoe := Shoe{}
	if db.Find(&shoe).Where("s_id = ?", id).Error != nil {
		return shoe, false
	}
	return shoe, true
}
