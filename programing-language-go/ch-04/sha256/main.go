package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	
	
	fmt.Println(CheckBit(c1, c2))
}
func CheckBit(c1, c2 [32]byte) int {
	diff := 0

	for i := range c1 {
		b1, b2 := c1[i], c2[i]
		for j := 0; j < 8; j++ {
			mask := byte(1 << uint(j))
			if b1&mask != b2&mask {
				diff++
			}
			if i == 0 {
				fmt.Println("c1[i]:", c1[i])
				fmt.Println("c2[i]:",c2[i])
				fmt.Println("mask:", mask)
				fmt.Println("b1&mask:", b1&mask)
				fmt.Println("b2&mask:", b2&mask)
			}
		}
	}

	return diff
}