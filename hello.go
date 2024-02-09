package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintf(w, "Hello, World!")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func date(w http.ResponseWriter, r *http.Request) {
	fmt.Println("server: date handler started")
	defer fmt.Println("server: date handler ended")
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()

	if err != nil {
		http.Error(w, "date command errd", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "date: %s", string(dateOut))
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "favicon.ico")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/date", date)
	http.HandleFunc("/favicon.ico", faviconHandler)

	http.ListenAndServe(":8090", nil)
}
