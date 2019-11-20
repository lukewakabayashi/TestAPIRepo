package main

import (
	"fmt"

	"github.com/frazercomputing/owner-verification-portal-back-end/auth"
)

func main() {
	a, err := auth.NewDefault()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a.Sign(18292))
	fmt.Println(auth.New([]byte("2347a")).Sign(18292))
}
