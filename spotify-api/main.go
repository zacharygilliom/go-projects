package main

import (
	"fmt"
	/*"io/ioutil"*/
	b64 "encoding/base64"
	//"net/http"
)

func main() {
	encodedString := encodeIdAndSecret()
	fmt.Println(encodedString)
}

func getClientInfo() (string, string) {
	client_id, client_secret := client_info()
	return client_id, client_secret
}

/*func authUser() {

	client_id, client_secret := getclientinfo()
	url := "https://accounts.spotify.com/authorize?client_id=" + client_id + "response_type=code&scope=playlist-read-private"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(client_secret)
		fmt.Println("Invalid client_id passed: %d", err)
	} else {
		fmt.Println("Authorization Successful")
		fmt.Println(resp)

	}
}
*/
func encodeIdAndSecret() string {
	client_id, client_secret := getClientInfo()
	data := client_id + ":" + client_secret
	encString := b64.StdEncoding.EncodeToString([]byte(data))
	return encString
}
