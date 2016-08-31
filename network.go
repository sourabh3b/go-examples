package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}
 
func main() {
	person := Person{
			Name: Name {Family:"abc",Personal:"xyx"},
			Email:[]Email{Email{Kind:"home",Address:"home@gmail.com"},Email{Kind:"work",Address:"work@gmail.com"}}
	}
	saveJson
}
