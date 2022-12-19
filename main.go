package main

import (
	"exemple.com/beckand-2taskmi/database"

	_ "github.com/lib/pq"
)

func main() {
	database.Init()
}
