package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strconv"
)

func main() {
	//	fmt.Println("Hello, World!")
	r := gin.Default()
	r.GET("/fibonacci", ginFibonacciHandler)
	port := "localhost:8080"

	fmt.Println("Server is running on address", port)
	r.Run(port)
	//на стандартном роутере
	//http.HandleFunc("/", fibonacciHandler)
	//log.Println("Starting server...")
	//log.Fatal(http.ListenAndServe(":8080", nil))
}

func fibonacciHandler(w http.ResponseWriter, r *http.Request) {
	result := calculateFibonacci(6)
	s := strconv.Itoa(int(result.Int64()))
	w.Write([]byte(s))
}

func ginFibonacciHandler(c *gin.Context) {
	n := c.DefaultQuery("n", "0")
	nInt, err := strconv.Atoi(n)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	}
	result := calculateFibonacci(nInt)
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func calculateFibonacci(n int) *big.Int {
	if n <= 0 {
		return big.NewInt(0)
	}
	if n <= 1 {
		return big.NewInt(1)
	}

	a := big.NewInt(0)
	b := big.NewInt(1)

	var result *big.Int

	for i := 2; i <= n; i++ {
		result = new(big.Int).Set(a)
		result.Add(result, b)
		a.Set(b)
		b.Set(result)
	}

	return result
}
