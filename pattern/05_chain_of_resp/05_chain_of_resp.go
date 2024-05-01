package main

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
	Цепочка вызовов - это поведенческий паттерн, позволяющий передавать запросы последовательно
	по цепочке обработчиков. Каждый обработчик решает, может ли он обработать запрос,
	и передает его следующему обработчику в цепи, если не может.

	Плюсы:
	- Позволяет добавлять или изменять обработчики без изменения клиентского кода
	- Позволяет строить гибкие цепи обработки запросов
	- Убирает прямую зависимость между отправителем запроса и его получателем

	Минусы:
	- Запрос может быть не обработан, если не найден соответствующий обработчик
*/

type Handler interface {
	SendRequest(message int) string
}

type ConcreteHandlerA struct {
	next Handler
}

func (h *ConcreteHandlerA) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Im handler A!"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

type ConcreteHandlerB struct {
	next Handler
}

func (h *ConcreteHandlerB) SendRequest(message int) (result string) {
	if message == 2 {
		result = "Im handler B!"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}
