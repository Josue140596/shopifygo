package main

import (
	"github.com/Josue140596/shopifygo/internal/db"
)

func main() {
	// rootCtx := context.Background()
	db.NewConnection("host=localhost port=5432 user=root password=secret dbname=shopify_db sslmode=disable")
}
