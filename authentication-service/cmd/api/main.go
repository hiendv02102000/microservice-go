package main

import (
	"authentication/data/db"
	"fmt"
)

const webPort = ":80"

func main() {
	_, err := db.NewDB()
	fmt.Println(err)
}
