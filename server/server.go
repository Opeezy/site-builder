package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No env found")
	}
}

func embedHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))
}

func main() {
	portKey, exists := os.LookupEnv("PORT")
	if !exists {
		log.Fatal("No port key found")
	}
	homeKey, exists := os.LookupEnv("HOME_FILES")
	if !exists {
		log.Fatal("No home files key found")
	}
	port := fmt.Sprintf(":%v", portKey)

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(homeKey))))
	http.HandleFunc("/save", embedHandler)

	log.Println("Starting server")
	err := http.ListenAndServe(port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
