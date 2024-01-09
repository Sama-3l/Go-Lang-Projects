package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const thisUrl = "https://lco.dev"
const checkUrl = "https://lco.dev:3000/local?coursename=reactjs&paymentid=uebfuifbbis"

func main() {
	fmt.Println("LCO web request ")

	res, err := http.Get(thisUrl)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T\n", res)

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", string(body))

	//parsing URL
	resultUrl, _ := url.Parse(checkUrl)

	fmt.Println(resultUrl.Scheme)
	fmt.Println(resultUrl.Host)
	fmt.Println(resultUrl.Port())
	fmt.Println(resultUrl.Path)
	fmt.Println(resultUrl.RawQuery)

	qparams := resultUrl.Query()

	fmt.Printf("%T\n", qparams)
	fmt.Println(qparams)
	fmt.Printf("%T\n", qparams["coursename"])
}
