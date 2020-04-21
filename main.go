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

func compute(memSize, cpuSize int) float64 {
	// memory work
	arr := make([]float64, memSize)
	for i := range arr {
		arr[i] = rand.Float64()
	}

	// cpu work
	sum := 0.0
	for i := 0; i < cpuSize; i++ {
		sum += math.Sin(2.0 * math.Pi * arr[i])
	}
	return sum
}

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mem := vars["mem"]
	cpu := vars["cpu"]

	memSize, err := strconv.Atoi(mem)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cpuSize, err := strconv.Atoi(cpu)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	result := compute(memSize, cpuSize)
	w.Write([]byte(fmt.Sprintf("%f", result)))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Listening on :8081...")
	router := mux.NewRouter()
	router.
		Methods(http.MethodGet).
		Path("/work/{mem}/{cpu}").
		HandlerFunc(handler)
	if err := http.ListenAndServe("0.0.0.0:8081", router); err != nil {
		panic(err)
	}
}
