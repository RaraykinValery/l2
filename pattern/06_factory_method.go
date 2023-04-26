package pattern

import (
	"log"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

type action string

const (
	A action = "A"
	B action = "B"
	C action = "C"
)

type Creator interface {
	CreateProduct(action action) Product // Factory Method
}

type Product interface {
	Use() string // Every product should be usable
}

type ConcreteCreator struct{}

func NewCreator() Creator {
	return &ConcreteCreator{}
}

func (p *ConcreteCreator) CreateProduct(action action) Product {
	var product Product

	switch action {
	case A:
		product = &ConcreteProductA{string(action)}
	case B:
		product = &ConcreteProductB{string(action)}
	case C:
		product = &ConcreteProductC{string(action)}
	default:
		log.Fatalln("Unknown Action")
	}

	return product
}

type ConcreteProductA struct {
	action string
}

func (p *ConcreteProductA) Use() string {
	return p.action
}

type ConcreteProductB struct {
	action string
}

func (p *ConcreteProductB) Use() string {
	return p.action
}

type ConcreteProductC struct {
	action string
}

func (p *ConcreteProductC) Use() string {
	return p.action
}
