package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/post", postHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что метод запроса POST.
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Читаем данные из тела запроса.
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Выводим данные на экран.
	fmt.Fprintln(w, "Data received:")
	fmt.Fprintln(w, string(body))
}

// Обработчик, возвращающий компонент пути запрашиваемого URL.
func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprint(w, "<body>")
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	fmt.Fprintf(w, "Headers = %q\n", r.Header)
	fmt.Fprint(w, "</body>")

}

// Счетчик, возвращающий количество сделанных запросов,
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
