package main

import (
	"math/rand"
	"strings"
	"time"
)

func GeneratePassword(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var password strings.Builder
	password.Grow(n)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		password.WriteByte(charset[rand.Intn(len(charset))])
	}

	return password.String()
}
func main() {
	password := GeneratePassword(35)
	println(password)
}

// Теоритически и если задаться такой целью, то да подобрать пароль можно в данном случае, потому что это функция псевдо генерации.
// Тестировать можно на количество символов в паролей(на пример n > 0), также можно задать ограничение по макисмальному их
//колличеству. По содержащимся символам и их регистру.
