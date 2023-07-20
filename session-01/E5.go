package main

import (
	"fmt"
	"math/rand"
)

func raiseError() error {

	return nil
}

func getRandom() (int, error) {
	randomInteger := rand.Int()
	err := raiseError()
	if err != nil {
		return 0, err
	}

	return randomInteger, nil
}

func main() {
	fmt.Println(getRandom())
}
