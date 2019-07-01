package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/globalsign/mgo"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Print("stating sever")
	a := make(chan int)
	go func() {
		for {
			select {
			case <-a:
				return
			}
			fmt.Printf("-->")
			time.Sleep(1 * time.Second)
		}
	}()
	r := mux.NewRouter()
	r.HandleFunc("/dirif", GetPathHandler)
	srv := &http.Server{
		Handler: r,
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	con, err := mgo.Dial("172.17.0.2:27017")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("connection is success")
	defer con.Close()
	log.Print(srv.Addr)
	a <- 1
	log.Fatal(srv.ListenAndServe())
}
func GetPathHandler(w http.ResponseWriter, r *http.Request) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("have error when get os information %v \n ", err)
	}

	if err != nil {
		fmt.Printf("have error when get os information %v \n ", err)
	}

	js, err := json.Marshal(map[string]string{
		"data":   dir,
		"status": "200",
	})
	fmt.Fprint(w, string(js))
	return

}
