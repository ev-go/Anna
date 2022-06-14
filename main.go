package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	fmt.Println("введите что-нибудь")
	var Pass string
	fmt.Scanf("%s\n", &Pass)
	fmt.Println("вы ввели:", Pass)

	client := http.Client{
		Timeout: time.Duration(6) * time.Second,
	}
	resp, err := client.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	// io.Copy(os.Stdout, resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// bodyString := string(body)

	fmt.Printf("\nAnswer for request : %s\t", body)

}
