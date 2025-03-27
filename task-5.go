package main

import "fmt"

func main() {
	str := "Привет"
	runes := []rune(str)
	runes[2] = 'е'
	str = string(runes)
	fmt.Println(str)
}

//Невозможно изменить строку просто по индексу,
//сначала нужно превратить ее в срез рун и после этого
//заменить одну букву.
