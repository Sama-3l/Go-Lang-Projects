package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web request")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("REsponse is of type %T\n", res)

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", string(body))
}
