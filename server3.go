package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette2 = []color.Color{color.White, color.Black}

const (
	whiteIndex2 = 0 // Первый цвет палитры
	blackIndex2 = 1 // Следующий цвет палитры
)

func main() {
	//http.HandleFunc("/", handler3)
	//http.HandleFunc("/count", counter)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		param, err := strconv.ParseFloat(r.URL.Query().Get("circle"), 64)
		if err != nil {
			param = 0
		}
		lissajous2(w, param)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

}
func lissajous2(out io.Writer, cyclesCount float64) {
	const (
		cycles  = 5     // Количество полных колебаний x
		res     = 0.001 // Угловое разрешение
		size    = 100   // Канва изображения охватывает [size..+size]
		nframes = 64    // Количество кадров анимации
		delay   = 8     // Задержка между кадрами (единица - 10мс)
	)
	if cyclesCount == 0 {
		cyclesCount = cycles
	}
	fmt.Printf("cyclesCount: %.2f", cyclesCount)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0 // Относительная частота колебаний у
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Разность фаз
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette2)
		for t := 0.0; t < cyclesCount*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex2)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Примечание: игнорируем ошибки
}

//handler4 := func(w http.ResponseWriter, г *http.Request) {
//	lissajous(w)
//}
//http.HandleFuncCV", handler)
//или, что то же самое:
//http.HandleFuncCV", func(w http.ResponseWriter, r *http.Request) {
//lissajous(w)
//})
