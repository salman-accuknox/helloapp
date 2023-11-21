// main.go

package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WORLD\n")
}
func whoamiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You\n")
}
func main() {
	http.HandleFunc("/api/hello", helloHandler)
   	http.HandleFunc("/api/whoami", whoamiHandler)
	// Start the server on port 8080
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
