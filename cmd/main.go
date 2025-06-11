package main

import (
	"log"

	"github.com/harshau007/listmanager/internal/listmanager"
	"github.com/harshau007/listmanager/internal/router"
)

func main() {
	lm := listmanager.New()
	r := router.New(lm)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
