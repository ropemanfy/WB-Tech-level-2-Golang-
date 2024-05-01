package main

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
	Фабричный метод — это порождающий паттерн проектирования, который решает проблему создания
	различных продуктов, без указания конкретных классов продуктов.

	Плюсы:
	- Выделяет код производства продуктов в одно место, упрощая поддержку кода
	- Избавляет класс от привязки к конкретным классам продуктов
	- Реализует принцип открытости/закрытости

	Минусы:
	- Может привести к созданию больших параллельных иерархий классов,
	так как для каждого класса продукта надо создать свой подкласс создателя
*/

type ITransport interface {
	setName(name string)
	getName() string
	setColor(color string)
	getColor() string
}

type transport struct {
	name  string
	color string
}

func (t *transport) setName(name string) {
	t.name = name
}

func (t *transport) getName() string {
	return t.name
}

func (t *transport) setColor(color string) {
	t.color = color
}

func (t *transport) getColor() string {
	return t.color
}

func getTransport(t string) (ITransport, error) {
	switch t {
	case "car":
		return newCar(), nil
	case "motorcycle":
		return newMotorcycle(), nil
	}
	return nil, fmt.Errorf("wrong type")
}

type car struct {
	transport
}

func newCar() ITransport {
	return &car{transport: transport{
		name:  "zhigul",
		color: "crimson",
	},
	}
}

type motorcycle struct {
	transport
}

func newMotorcycle() ITransport {
	return &car{transport: transport{
		name:  "kawasaki",
		color: "green",
	},
	}
}
