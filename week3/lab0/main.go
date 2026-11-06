package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"fmt"
)

func proofMe(txt1, txt2 string) {
	fmt.Println("\nProof the Hash Program")

	// MD5
	a1 := md5.Sum([]byte(txt1))
	b1 := md5.Sum([]byte(txt2))
	fmt.Printf("\nHash (MD5):\nOutput A = %x\nOutput B = %x\n", a1, b1)
	if a1 == b1 {
		fmt.Println("=> Match!")
	} else {
		fmt.Println("=> No Match!")
	}

	// SHA1
	a2 := sha1.Sum([]byte(txt1))
	b2 := sha1.Sum([]byte(txt2))
	fmt.Printf("\nHash (SHA1):\nOutput A = %x\nOutput B = %x\n", a2, b2)
	if a2 == b2 {
		fmt.Println("=> Match!")
	} else {
		fmt.Println("=> No Match!")
	}

	// SHA256
	a3 := sha256.Sum256([]byte(txt1))
	b3 := sha256.Sum256([]byte(txt2))
	fmt.Printf("\nHash (SHA256):\nOutput A = %x\nOutput B = %x\n", a3, b3)
	if a3 == b3 {
		fmt.Println("=> Match!")
	} else {
		fmt.Println("=> No Match!")
	}

	// SHA512
	a4 := sha512.Sum512([]byte(txt1))
	b4 := sha512.Sum512([]byte(txt2))
	fmt.Printf("\nHash (SHA512):\nOutput A = %x\nOutput B = %x\n", a4, b4)
	if a4 == b4 {
		fmt.Println("=> Match!")
	} else {
		fmt.Println("=> No Match!")
	}

	// SHA3-256
	a5 := sha3.Sum256([]byte(txt1))
	b5 := sha3.Sum256([]byte(txt2))
	fmt.Printf("\nHash (SHA3-256):\nOutput A = %x\nOutput B = %x\n", a5, b5)
	if a5 == b5 {
		fmt.Println("=> Match!")
	} else {
		fmt.Println("=> No Match!")
	}
}

func main() {
	var txt1, txt2 string

	fmt.Println("======== Name + Hashing Program ========")

	fmt.Print("Please input value 1: ")
	fmt.Scanln(&txt1)

	fmt.Print("Please input value 2: ")
	fmt.Scanln(&txt2)

	proofMe(txt1, txt2)
}
