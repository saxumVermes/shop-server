package shop

import (
	"fmt"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

type InMemShoeServer struct {
	Products map[string]Shoe
	Name     string
	Address
	sync.RWMutex
}

var counter = 1

func (s *InMemShoeServer) Add(brand, model string, price float32, colors string) (Shoe, bool) {
	s.Lock()
	defer s.Unlock()
	id := fmt.Sprintf("%3.3s-%03d", strings.ToLower(strings.TrimSpace(brand)), counter)
	if _, ok := s.Products[id]; ok {
		logrus.Warnf("Product already exists with id %s", id)
		return Shoe{}, false
	}
	shoe := Shoe{
		SID:    id,
		SModel: model,
		Brand:  brand,
		Price:  price,
		Colors: colors,
	}
	s.Products[id] = shoe
	counter++
	return shoe, true
}

func (s *InMemShoeServer) List() []Shoe {
	p := []Shoe{}
	s.Lock()
	for _, v := range s.Products {
		p = append(p, v)
	}
	s.Unlock()
	return p
}

func (s *InMemShoeServer) DeleteById(id string) bool {
	if _, ok := s.Products[id]; !ok {
		return false
	}
	delete(s.Products, id)
	return true
}

func (s *InMemShoeServer) Find(id string) (Shoe, bool) {
	shoe, ok := s.Products[id]
	if !ok {
		return shoe, false
	}
	return shoe, true
}

func newInMem(name string, city string, zip int) *InMemShoeServer {
	return &InMemShoeServer{
		Products: map[string]Shoe{},
		Name:     name,
		Address: Address{
			City: city,
			Zip:  zip,
		},
	}
}
