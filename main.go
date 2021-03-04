package main

import "alvinhtml.com/go-upload/router"

func main() {
	r := router.Init()

	r.Run(":8007")
}
