package main

import (
	"fmt"
	"net/http"

	authcontroller "github.com/DhandiAdam/OnnlineFood/controllers"
)

func main() {
	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/logout", authcontroller.Logout)
	http.HandleFunc("/register", authcontroller.Register)
	http.HandleFunc("/Tampilan_Utama", authcontroller.Tampilan_Utama)
	http.HandleFunc("/send-data", authcontroller.SendDataToNodeJS)
	http.Handle("/static", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	fmt.Println("Server jalan di: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
