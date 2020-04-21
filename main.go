package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func compute(n int64) float64 {
	arr := make([]float64, n)
	for i := range arr {
		arr[i] = rand.Float64()
	}
	sum := 0.0
	sinSum := 0.0
	for _, a := range arr {
		sum += a
		sinSum += math.Sin(2.0 * math.Pi * a)
	}
	//fmt.Println("Sum is", sum)
	//fmt.Println("Sin sum is", sinSum)
	return sum / sinSum
}

func handler(w http.ResponseWriter, r *http.Request) {
	n := mux.Vars(r)["n"]
	//fmt.Println("received work", n)
	num, err := strconv.Atoi(n)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusNotFound)
		return
	}
	result := compute(int64(num))
	w.Write([]byte(fmt.Sprintf("%f", result)))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Listening on :8081...")
	router := mux.NewRouter()
	router.
		Methods(http.MethodGet).
		Path("/work/{n}").
		HandlerFunc(handler)
	if err := http.ListenAndServe("0.0.0.0:8081", router); err != nil {
		panic(err)
	}
}
