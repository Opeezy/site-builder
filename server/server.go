package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No env found")
	}
}

// type Child interface {
// 	())
// }

type Div struct {
	Classlist string `json:"class"`
	Children  interface {
	}
}

func embedHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	cmd := exec.Command("python", "pytl/pytl.py", string(b))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println(cmd.Run())
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
