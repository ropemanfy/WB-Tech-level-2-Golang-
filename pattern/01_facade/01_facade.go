package main

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
	Фасад - структурный шаблон проектирования, позволяющий скрыть сложность системы путём сведения всех возможных внешних
	вызовов к одному объекту, делегирующему их соответствующим объектам системы.

	Плюсы:
	- Упрощает работу клиента со сложной системой
	- Уменьшает связность кода

	Минусы:
	- Ограничивает доступ к подсистеме
	- Может стать god object'ом
*/

type SubsystemA struct {
}

func (s *SubsystemA) methodA() {
	fmt.Println("some logic in SubsystemA")
}

type SubsystemB struct {
}

func (s *SubsystemB) methodB() {
	fmt.Println("some logic in SubsystemB")
}

type SubsystemC struct {
}

func (s *SubsystemC) methodC() {
	fmt.Println("some logic in SubsystemC")
}

type Facade struct {
	a *SubsystemA
	b *SubsystemB
	c *SubsystemC
}

func (f *Facade) FacadeMethodOne() {
	f.a.methodA()
	f.b.methodB()
	f.c.methodC()
}

func (f Facade) FacadeMethodTwo() {
	f.a.methodA()
	f.b.methodB()
}
