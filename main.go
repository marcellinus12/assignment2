package main

import (
	"fmt"
	"web-server/routers"
)

func main() {
	const PORT = ":8088"

	fmt.Sprintln("server start at", PORT)
	routers.StartServer().Run(PORT)
}

// var PORT = ":8088"

// func main() {
// 	http.HandleFunc("/", greet)

// 	fmt.Println("web server run")
// 	http.ListenAndServe(PORT, nil)
// }

// func greet(w http.ResponseWriter, r *http.Request) {
// 	msg := "hello world"
// 	fmt.Fprintln(w, msg)
// }
