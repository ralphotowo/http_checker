package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const ZONE_ID
const 

func main() {
	response, err := http.Get("https://api.cloudflare.com/client/v4/zones/$ZONE_ID/dns_records")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

}
