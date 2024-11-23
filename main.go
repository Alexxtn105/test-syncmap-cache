package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var cache = &sync.Map{}

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

	// проверяем, есть ли результат в кэше
	if cacheResult, ok := cache.Load(nInt); ok {
		fmt.Println("Данные в кэше!")
		// возвращаем результат из кэша
		c.JSON(http.StatusOK, gin.H{"result": cacheResult})
		return
	}
	fmt.Println("Данных в кэше нет!")
	// расчет числа Фибоначчи
	result := calculateFibonacci(nInt)

	// Сохраняем результат в кэш со временем жизни (TTL). Например, на 10 секунд
	cache.Store(nInt, result)

	// запускаем горутину для удаления данных из кэша через заданное время
	go func() {
		time.Sleep(10 * time.Second)
		cache.Delete(nInt)
	}()

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
