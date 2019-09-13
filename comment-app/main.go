package main

import (
	"os"
)

func main() {
	if os.Getenv("GO_ENV") == "production" {
		startInProd()
	} else {
		startInDev()
	}
}
