package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ActionIndex(w http.ResponseWriter, r *http.Request) {
	data := []struct {
		Name string
		Age  int
	}{
		{"Dhandi", 24},
		{"Adam", 23},
		{"4522210147", 22},
		{"DhandiAdam", 21},
	}

	jsonInBytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonInBytes)
}

func main() {
	http.HandleFunc("/", ActionIndex)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
