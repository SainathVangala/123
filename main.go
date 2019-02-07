package main

import "encoding/json"
import "fmt"

type response1 struct {
	Version   string
        Lastcommitsha string
	Description string
       
}

type response2 struct {
	Myapplication []string
       
}

func main() {
 

   
 res2D := &response1{
        Version:   "1.0",
        Lastcommitsha: "asdfsdfasdf",
        Description: "pre-interview technical test"}
    strA, _ := json.Marshal(res2D)
  

  res3D := &response2{
        Myapplication: []string{(string(strA))}}
    fmt.Println(res3D)
 
           } 
 
