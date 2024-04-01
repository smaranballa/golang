package main

import (
	"fmt"
	"log"
	"net/http"
)

// func homePageHandler(res http.ResponseWriter, req *http.Request) {
// 	if req.URL.Path != "/home" {
// 		http.Error(res, "ERROR - 404 page not found", http.StatusNotFound)
// 		return
// 	}
// 	if req.Method != "GET" {
// 		http.Error(res, "Method not supported", http.StatusMethodNotAllowed)
// 		return
// 	}
// 	http.ServeFile(res, req, "./static/")
// }

func signUpHandler(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/sign-up" {
		http.Error(res, "ERROR - 404 page not found", http.StatusNotFound)
		return
	}
	if req.Method != http.MethodPost {
		http.Error(res, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	if err := req.ParseForm(); err != nil {
		http.Error(res, "Unable to parse the form", http.StatusExpectationFailed)
		return
	}
	name := req.FormValue("name")
	email := req.FormValue("email")
	password := req.FormValue("password")
	confirmPwd := req.FormValue("confirm_password")
	if password != confirmPwd {
		http.Error(res, fmt.Sprintf("Password not same, password: %s, confirm password:  %s", password, confirmPwd), http.StatusConflict)
	}
	fmt.Fprintf(res, "Name: %s\n", name)
	fmt.Fprintf(res, "Email: %s\n", email)
	fmt.Fprintf(res, "pwd: %s\n", password)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/sign-up", signUpHandler)
	fmt.Printf("Server started at port: %s", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil); err != nil {
		log.Fatal(err.Error())
	}
}
