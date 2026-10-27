package main

import (
	"fmt"
	"net/http"
)

func main() {
	var host string = "127.0.0.1"
	var port string = ":80"
	fmt.Printf("Запуск сервера по адрему: http://%s%s \n", host, port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	srv := &http.Server{Addr: port, Handler: mux}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Errorf("Ошибка запуска сервера: %w", err)
	}

}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
	w.Write([]byte("\n"))
}
