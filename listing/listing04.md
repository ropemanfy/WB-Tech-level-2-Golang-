Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
...
программа выведет значения от 0 до 9, а после произойдет deadlock т.к. 
канал никогда не будет закрыт.
```