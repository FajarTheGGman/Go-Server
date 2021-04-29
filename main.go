/*
Copyright© 2019 By Fajar Firdaus

*/

package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main(){
	fmt.Print("{\n")
	fmt.Print("  Coder : Fajar Firdaus\n")
	fmt.Print("  Fb : Ace.of.spades729\n")
	fmt.Print("  IG : fajar_firdaus_7\n")
	fmt.Print("  YT : iTech7732\n")
	fmt.Print("}\n\n")

	files, err := ioutil.ReadDir("./htdocs")
    if err != nil {
        fmt.Print("[!] File Empty")
	}
	
	http.Handle("/", http.FileServer(http.Dir("./htdocs")))
	
	fmt.Print("[Server Content]\n")
    for _, f := range files {
		fmt.Print("- " + f.Name() + "\n")
	}

	fmt.Print("\n[+] Server Running At Port 3000")
	http.ListenAndServe(":3000", nil)
	
}
