package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var typeHash = flag.String("s", "SHA256", "type of SHA Hash")

func main() {
	flag.Parse()
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
	inputToSHA(typeHash)
}

// Exercício 4.2
/* Escreva um programa que exiba o sha SHA256 de uma entrada-padrão por default, mas aceite uma flag de linha de comando para exibir o
hash SHA384 ou SHA512 em seu lugar*/
func inputToSHA(typeHash *string) {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		switch *typeHash {
		case "SHA256":
			{
				hash := sha256.Sum256([]byte(input.Text()))
				fmt.Printf("%s HASH %x\n", *typeHash, hash)
			}
		case "SHA384":
			{
				{
					hash := sha512.Sum384([]byte(input.Text()))
					fmt.Printf("%s HASH %x\n", *typeHash, hash)
				}
			}
		case "SHA512":
			{
				hash := sha512.Sum512([]byte(input.Text()))
				fmt.Printf("%s HASH %x\n", *typeHash, hash)
			}
		}

	}
}
