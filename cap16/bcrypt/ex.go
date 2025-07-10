package main

import "fmt"
import "golang.org/x/crypto/bcrypt"

func main() {
	senha := "minha senha"
	senhaErrada := "errada"
	
	// func GenerateFromPassword(password []byte, cost int) ([]byte, error)

	senhaHash, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(senhaHash)) // gerou um hash(código aleatório)


	// func CompareHashAndPassword(hashedPassword, password []byte) error

	fmt.Println("senha certa:",bcrypt.CompareHashAndPassword(senhaHash, []byte(senha)))
	fmt.Println("senha errada:" ,bcrypt.CompareHashAndPassword(senhaHash, []byte(senhaErrada)))

	
}