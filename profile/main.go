// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/mebackend/profileapi/router"
// )

// func main() {
// 	fmt.Println("MongoDB API")
// 	r := router.Router()

// 	http.Handle("/", http.FileServer(http.Dir("./templates")))

// 	fmt.Println("Server is getting started...")
// 	log.Fatal(http.ListenAndServe(":4000", r))
// 	fmt.Println("Listening at port 4000 ...")
// }

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mebackend/profileapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()

	// Serve static HTML files
	fs := http.FileServer(http.Dir("./templates"))
	r.PathPrefix("/").Handler(fs)

	fmt.Println("Server is getting started on port 4000...")
	log.Fatal(http.ListenAndServe(":4000", r))
}

