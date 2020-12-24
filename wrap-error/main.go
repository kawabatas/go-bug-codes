package main

import (
	"errors"
	"fmt"
)

func main() {
	baseErr := errors.New("base err")
	wrapErr := fmt.Errorf("wrap err %w", baseErr)

	err := fmt.Errorf("wrap err %w", baseErr)
	if err.Error() == "base err" {
		fmt.Println("string")
	}
	if errors.Is(err, baseErr) {
		fmt.Println("base")
	}
	if errors.Is(err, wrapErr) {
		fmt.Println("wrap")
	}
}
