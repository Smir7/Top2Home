package main

import "fmt"

func printNumber(ptrToNumber interface{}) {
	if ptrToNumber != nil { // проверка входящего парметра
		if numberPtr, ok := ptrToNumber.(*int); ok { // проверка если ptrToNumber не указатель инт
			if numberPtr != nil {
				fmt.Println(*numberPtr) //разименовали указатель
			} else {
				fmt.Println("nil") // проверили на приведение типа
			}
		} else {
			fmt.Println("not an int")
		}
	} else {
		fmt.Println("nil input")
	}
}

func main() {
	v := 10
	printNumber(&v) // ссылка на v вывод 10
	var pv *int     //ссылка на инт
	printNumber(pv) //вывод нил
	pv = &v
	printNumber(pv) // вывод 9
}
