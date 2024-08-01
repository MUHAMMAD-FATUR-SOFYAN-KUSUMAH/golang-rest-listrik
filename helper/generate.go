package helper

import "math/rand"

func GenerateRandomToken() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	token := make([]rune, 10)
	for i := range token {
		token[i] = letters[rand.Intn(len(letters))]
	}
	return string(token)
}

func GenerateRandomNumeber() string {
	letters := []rune("0123456789")
	token := make([]rune, 10)
	for i := range token {
		token[i] = letters[rand.Intn(len(letters))]
	}
	return string(token)
}

func Generate_role(q int) string {
	switch q {
	case 1:
		return "admin"
	case 2:
		return "user"
	default:
		return "user"
	}
}
