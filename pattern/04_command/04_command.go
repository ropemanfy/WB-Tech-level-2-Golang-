package main

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
	Команда — это поведенческий паттерн, позволяющий заворачивать запросы или простые операции в отдельные объекты.
	Это позволяет откладывать выполнение команд, выстраивать их в очереди, а также хранить историю и делать отмену.

	Плюсы:
	- Убирается прямая связь между отправителями и исполнителями запросов
	- Позволяет удобно реализовывать различные операции: отмена и повтор запросов,
	отложенный запуск запросов, выстраивание очереди запросов

	Минусы:
	- Усложняет код из-за необходимости реализации дополнительных классов
*/

type Command interface {
	execute()
}

type onCommand struct {
	device Device
}

func (o *onCommand) execute() {
	o.device.on()
}

type offCommand struct {
	device Device
}

func (o *offCommand) execute() {
	o.device.off()
}

type Device interface {
	on()
	off()
}

type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("TV on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("TV off")
}

type button struct {
	command Command
}

func (b *button) press() {
	b.command.execute()
}
