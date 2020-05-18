package main

import (
	"fmt"
	/*"io/ioutil"*/
	"net/http"
)

func main() {

	client_id, client_secret := getclientinfo()
	url := "https://accounts.spotify.com/authorization?client_id=" + client_id
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Invalid client_id passed: %d", err)
	}
	fmt.Println(resp)
	fmt.Println(client_secret)
}

func getclientinfo() (string, string) {
	client_id, client_secret := client_info()
	return client_id, client_secret
}
