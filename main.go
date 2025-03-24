package main

// Объяснить, почему в первом вызове ошибки не выводится сообщение.
//Что происходит при передаче nil-значения интерфейса.

import "fmt"

type MyError struct{} // структура пустая

func (MyError) Error() string { // принимает пустую структуру и возвращвет ошибку строку
	return "MyError!"
}

func errorHandler(err error) { //принимает ошибку
	if err != nil { // если есть ошибка то печатает
		fmt.Println("Error:", err)
	}
}
func main() {
	var err *MyError //значение err присвоен указатель на структуру
	errorHandler(err)
	err = &MyError{}
	errorHandler(err)
}

// В первом вызове ошибки выводится nil потому что в main мы переменной err присвиваем указатель на пустую структуру.
// при передаче интерфейсу nil ничего не происходит
