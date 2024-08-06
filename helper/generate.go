package helper

import (
	"math/rand"
)

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

func Int_StatusToString(q int) string {
	switch q {
	case 1:
		return "belum lunas"
	case 2:
		return "menunggu konfirmasi"
	case 3:
		return "lunas"
	case 4:
		return "dibatalkan"
	default:
		return "ngacok"
	}
}
