package main

import (
	"fmt"
	"log"

	"github.com/kyeett/wasm-adventure/preferences"
)

func main() {

	pref, err := preferences.New("test")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pref.GetString("name"))
	fmt.Println(pref.GetInt("age"))
	fmt.Println(pref.GetBool("toggle"))

	//pref.SetItem("name", "magnus")
	//pref.SetItem("age", 18)
	//pref.SetItem("toggle", true)
}
