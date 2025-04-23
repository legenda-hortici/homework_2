package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func generatePassword(n int) (string, error) {
	if n <= 0 {
		return "", fmt.Errorf("Пароль должен быть больше 0 символов")
	}

	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	var password strings.Builder

	for i := 0; i < n; i++ {
		password.WriteRune(chars[rand.Intn(len(chars))])
	}

	return password.String(), nil
}

func main() {
	/*
		Теоретически можно угадать пароль если он имеет короткую длину символов
	*/
	password, err := generatePassword(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(password)
}
