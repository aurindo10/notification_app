package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aurindo10/internal/app"
)

func main() {
	ctx := context.Background()

	if err := app.Run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
