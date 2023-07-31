package main

import (
	"fmt"
	"golang.org/x/crypto/chacha20"
	"log"
)

func main() {

	key := make([]byte, 32)

	nonce := make([]byte, 24)

	plaintext := []byte("Hello, World! it is pair 5 here")

	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		log.Fatal(err)
	}

	ciphertext := make([]byte, len(plaintext))
	cipher.XORKeyStream(ciphertext, plaintext)

	fmt.Printf("Ciphertext: %x\n", ciphertext)

	decrypted := make([]byte, len(ciphertext))
	cipher, err = chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		log.Fatal(err)
	}
	cipher.XORKeyStream(decrypted, ciphertext)

	fmt.Printf("Decrypted: %s\n", decrypted)
}
