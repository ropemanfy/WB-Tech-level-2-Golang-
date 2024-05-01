package main

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
	Посетитель — это поведенческий паттерн проектирования, который позволяет добавлять в программу новые операции,
	не изменяя классы объектов, над которыми эти операции могут выполняться.

	Плюсы:
	- Упрощает добавление операций, работающих со сложными структурами объектов
	- Объединяет родственные операции в одном классе

	Минусы:
	- Может привести к нарушению инкапсуляции элементов
	- Паттерн не оправдан, если иерархия элементов часто меняется
*/

type Shape interface {
	GetType() string
	Accept(visitor Visitor)
}

type Rectangle struct {
}

func (r *Rectangle) GetType() string {
	return "Rectangle"
}

func (r *Rectangle) Accept(visitor Visitor) {
	visitor.VisitForRectangle(r)
}

type Circle struct {
}

func (c *Circle) GetType() string {
	return "Circle"
}

func (c *Circle) Accept(visitor Visitor) {
	visitor.VisitForCircle(c)
}

type Visitor interface {
	VisitForRectangle(r *Rectangle)
	VisitForCircle(c *Circle)
}

type Area struct {
}

func (a *Area) VisitForRectangle(r *Rectangle) {
	fmt.Println("rectangle area")
}

func (a *Area) VisitForCircle(c *Circle) {
	fmt.Println("circle area")
}
