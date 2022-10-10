// get request is made to api and all the fields are being read
// post request is made and id, createdAt fields are read from that response from post request 
package main

import (
	"bytes"
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

type PostResponse struct {
	Name string
	Job string 
	Id string 
	CreatedAt string 
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

	// posting a json object
	newData , _ := json.Marshal(map[string]string {
		"name" : "teja p",
		"job" : "engineer",
	})
	responseBody := bytes.NewBuffer(newData)
	
	resp,err := http.Post("https://reqres.in/api/users","application/json",responseBody)

	if err!=nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err !=nil {
		log.Fatal(err)
	}

	var res PostResponse;
	
	json.Unmarshal(body,&res)
	
	fmt.Println(res.Id,res.CreatedAt,res.Name,res.Job)
}