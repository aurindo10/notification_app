package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aurindo10/internal/app"
	"github.com/aurindo10/internal/repositories"
)

func main() {
	ctx := context.Background()
	db := repositories.NewDb()
	isRead := make(chan bool)
	if err := app.Run(ctx, db, isRead); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
