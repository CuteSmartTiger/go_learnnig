package main

import (
	"errors"
	"fmt"
)

var ErrDidNotWork = errors.New("did not work")

func DoTheThing(reallyDoIt bool) (err error) {
	if reallyDoIt {
		result, err := tryTheThing()
		fmt.Println(result, err)
		if err != nil || result != "it worked" {
			fmt.Println("step in to here")
			err = ErrDidNotWork
			fmt.Println("step after here err", err)

		}
	}
	fmt.Println("return err", err)
	return err
}

func tryTheThing() (string, error) {

	return "", ErrDidNotWork
}

func main() {
	fmt.Println("======starting========")
	fmt.Println(DoTheThing(true)) // A
	fmt.Println("==============") // B
	fmt.Println(DoTheThing(false))
	fmt.Println("======ending========")
}
