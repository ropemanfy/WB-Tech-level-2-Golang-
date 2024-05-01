package main

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
	Строитель - порождающий шаблон проектирования предоставляет способ создания составного объекта.

	Плюсы:
	- Позволяет создавать тяжелые объекты пошагово
	- Переиспользование кода

	Минусы:
	- Усложняет код из-за введения дополнительных классов
	- Клиент привязан к конкретным классам билдера
*/

type Car struct {
	bodyStyle string
	brand     string
	color     string
}

type CarBuilder struct {
	car Car
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{car: Car{}}
}

func (c *CarBuilder) SetBodyStyle(body string) {
	c.car.bodyStyle = body
}

func (c *CarBuilder) SetBrand(brand string) {
	c.car.brand = brand
}

func (c *CarBuilder) SetColor(color string) {
	c.car.color = color
}

func (c *CarBuilder) GetCar() Car {
	return c.car
}
