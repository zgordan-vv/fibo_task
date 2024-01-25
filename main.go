// Fibonacci sequence API
// Endpoints (all of them are GET requests):
// - /previous: returns the previous number of the Fibonacci sequence, the sequence is not changed.
// - /current: returns the current number of the Fibonacci sequence, the sequence is not changed.
// - /next: returns the next number of the Fibonacci sequence and adds a new number to the sequence.
// - /reset: resets the sequence to the beginning and returns the curent value.
// - /make_trouble: causes a panic for testing reasons to check recovery.
//
// Example:
// http://127.0.0.1:8080/current returns 0
// http://127.0.0.1:8080/next returns 1
// http://127.0.0.1:8080/next returns 1
// http://127.0.0.1:8080/next returns 2
// http://127.0.0.1:8080/previous returns 1
// http://127.0.0.1:8080/current returns 2
// http://127.0.0.1:8080/reset returns 0

package main

import (
	"github.com/gin-gonic/gin"
)

var previous, current uint64

func main() {
	router := gin.Default()
	router.GET("/previous", getPrevious)
	router.GET("/current", getCurrent)
	router.GET("/next", getNext)
	router.GET("/reset", reset)
	router.GET("/make_trouble", func(c *gin.Context) {
		c.String(200, "%s", "Server error, try again later")
		panic("Oops, something went wrong")
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}

// getPrevious returns the previous number of the Fibonacci sequence.
// No recalculation is done.
func getPrevious(c *gin.Context) {
	c.String(200, "%d", previous)
}

// getCurrent returns the current number of the Fibonacci sequence.
// No recalculation is done.
func getCurrent(c *gin.Context) {
	c.String(200, "%d", current)
}

// getNext returns the next number of the Fibonacci sequence
// and recalculates the previous and current numbers.
func getNext(c *gin.Context) {
	var next uint64
	if current == 0 {
		next = 1
	} else {
		next = previous + current
	}
	previous = current
	current = next
	c.String(200, "%d", next)
}

func reset(c *gin.Context) {
	previous = 0
	current = 0
	c.String(200, "%d", current)
}
