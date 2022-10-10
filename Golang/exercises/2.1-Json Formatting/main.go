// get request is made to api and all the fields are being read

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct{ // Structure of json response expected from api hit
	Data map[string]string
	Support map[string]string 
}

func main(){
	fmt.Println("hello, going to hit https://reqres.in/api/users/3")
	
	response, err := http.Get("https://reqres.in/api/users/3")

	if err != nil {
		fmt.Println("Error after hitting the api")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData , err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var m Response;

	json.Unmarshal([]byte(string(responseData)),&m)

	fmt.Println("iterating over Data Map")
	for k,v := range m.Data {
		fmt.Println("key is ",k," value is ",v)
	}

	fmt.Println("iterating over Support Map")
	for k,v := range m.Support {
		fmt.Println("key is ",k," value is ",v)
	}	
}