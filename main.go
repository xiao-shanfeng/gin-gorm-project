package main

import "example.com/m/v2/router"

func main() {
	r := router.Routers()
	r.Run("localhost:8000")
}
