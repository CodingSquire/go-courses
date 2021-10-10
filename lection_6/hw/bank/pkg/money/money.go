package money

import "fmt"

type Money int

func Validate(balance Money) (err error) {
	if balance < 0 {
		err = fmt.Errorf("balance should be positive")
	}
	return
}
