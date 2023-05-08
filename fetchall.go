package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // Запуск go-подпрограммы
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
		// Получение из канала ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // Отправка в канал ch
		return
	}
	time.Sleep(time.Second * 20)
	//nbytes, err := io.Copy(io.Discard, resp.Body)
	name := strings.Replace(strings.Replace(url, "/", "", 10), ":", "", 1)
	file, _ := os.Create(fmt.Sprintf("%s.html", name))
	io.Copy(file, resp.Body)
	file.Close()
	resp.Body.Close() // Исключение утечки ресурсов
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %s", secs, url)
}
