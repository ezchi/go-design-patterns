package prototype

import (
	"errors"
	"fmt"
)

type ShirtColor byte

const (
	White ShirtColor = iota
	Black
	Blue
)

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtCloner interface {
	GetClone(s ShirtColor) (ItemInfoGetter, error)
}

func GetCloner() ShirtCloner {
	return nil
}

type ShirtCache struct{}

func (c ShirtCache) GetClone(s ShirtColor) (ItemInfoGetter, error) {
	switch s {
	case White:
		newItem := *whitePrototype
		return &newItem, nil

	case Black:
		newItem := *blackPrototype
		return &newItem, nil

	case Blue:
		newItem := *bluePrototype
		return &newItem, nil

	}

	return nil, errors.New("not implemented yet")
}

type Shirt struct {
	Color ShirtColor
	Price float32
	SKU   string
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and color id %d that cost %f\n", s.SKU, s.Color, s.Price)
}

func (s *Shirt) GetPrice() float32 {
	return s.Price
}

func GetShirtCloner() ShirtCloner {
	return ShirtCache{}
}

var whitePrototype *Shirt = &Shirt{
	Price: 15.0,
	SKU:   "empty",
	Color: White,
}

var blackPrototype *Shirt = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}

var bluePrototype *Shirt = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}
