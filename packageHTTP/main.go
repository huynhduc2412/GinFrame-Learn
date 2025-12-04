package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/demo", demoHandler)
	log.Println("Server is starting ...")
	err := http.ListenAndServe(":8080" , nil) // localhost:8080
	if err != nil {
		log.Fatal("Server error: " , err)
	}
}

func demoHandler(w http.ResponseWriter, r *http.Request) {
	// log.Printf("%+v" , r)
	if r.Method != http.MethodGet {
		http.Error(w , "Not support method" , http.StatusMethodNotAllowed)
		return
	}
	response := map[string]string {
		"message" : "Hello World!",
		"myName" : "Duc",
	}
	w.Header().Set("Content-Type" , "application/json")
	
	// data , err := json.Marshal(response)
	// if err != nil {
	// 	http.Error(w , "Error Json" , http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(data)

	json.NewEncoder(w).Encode(response)
}