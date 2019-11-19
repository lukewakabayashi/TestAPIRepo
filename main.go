package main

import (
	"fmt"

	"github.com/frazercomputing/owner-verification-portal-back-end/auth"
)

func main() {
	fmt.Println(auth.NewDefault().Sign(18292))
	fmt.Println(auth.New("2347").Sign(18292))
}
