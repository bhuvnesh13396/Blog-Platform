package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	cost = 12
)

func main() {
	pass := "qweqweqwe"
	err := do(pass)
	fmt.Println("err", err)
}

func do(pass string) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	if err != nil {
		return err
	}
	fmt.Println("hash:", string(hash))

	c, err := bcrypt.Cost(hash)
	if err != nil {
		return
	}
	fmt.Println("calculated cost:", c)

	return bcrypt.CompareHashAndPassword(hash, []byte(pass))
}
