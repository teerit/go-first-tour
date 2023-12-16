package main

import (
	"errors"
	"fmt"
)

func main() {
	_, err := someF()
	if err != nil {
		// Handle error ...
		return
	}

	_, err = someOtherF()
	if err != nil {
		// Handle error again ..
		return
	}
	fmt.Println()
}

func someF() (int, error) {
	// Some operations ...

	return 0, someE()
}

func someOtherF() (int, error) {
	return 0, someE()
}

func someE() error {
	return errors.New("same type of error")
}
