package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/add", handleAdd)
	http.HandleFunc("/subtruct", handleSubtruct)
	http.HandleFunc("/multiply", handleMultiply)
	http.HandleFunc("/divide", handleDivide)
	http.ListenAndServe(":8080", nil)

}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	p1 := r.URL.Query().Get("a")
	p2 := r.URL.Query().Get("b")

	a, err := strconv.Atoi(p1)
	if err != nil {
		http.Error(w, "Параметр а должен быть числом", 400)
		return
	}

	b, err := strconv.Atoi(p2)
	if err != nil {
		http.Error(w, "Параметр b должен быть числом", 400)
		return
	}

	result := a + b

	fmt.Fprint(w, result)
}

func handleSubtruct(w http.ResponseWriter, r *http.Request) {
	p1 := r.URL.Query().Get("a")
	p2 := r.URL.Query().Get("b")

	a, err := strconv.Atoi(p1)
	if err != nil {
		http.Error(w, "Параметр a должен быть числом", 400)
		return
	}

	b, err := strconv.Atoi(p2)
	if err != nil {
		http.Error(w, "Параметр b должен быть числом", 400)
		return
	}

	result := a - b

	fmt.Fprint(w, result)

}

func handleMultiply(w http.ResponseWriter, r *http.Request) {
	p1 := r.URL.Query().Get("a")
	p2 := r.URL.Query().Get("b")

	a, err := strconv.Atoi(p1)
	if err != nil {
		http.Error(w, "Параметр a должен быть числом", 400)
		return
	}

	b, err := strconv.Atoi(p2)
	if err != nil {
		http.Error(w, "Параметр b должен быть числом", 400)
		return
	}

	result := a * b

	fmt.Fprint(w, result)

}

func handleDivide(w http.ResponseWriter, r *http.Request) {
	p1 := r.URL.Query().Get("a")
	p2 := r.URL.Query().Get("b")

	a, err := strconv.Atoi(p1)
	if err != nil {
		http.Error(w, "Параметр a должен быть числом", 400)
		return
	}

	b, err := strconv.Atoi(p2)
	if err != nil {
		http.Error(w, "Параметр b должен быть числом", 400)
		return
	}

	if b == 0 {
		http.Error(w, "Нельзя делить на ноль", 400)
		return
	}
	result := a / b

	fmt.Fprint(w, result)

}
