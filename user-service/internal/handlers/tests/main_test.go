package tests

import (
	"context"
	"os"
	"testing"

	"github.com/aurindo10/internal/app"
	"github.com/aurindo10/internal/repositories"
)

func TestMain(m *testing.M) {
	ctx, cancel := context.WithCancel(context.Background())
	isRead := make(chan bool)
	dbName := "dbtest"
	dbtest := repositories.NewTestDb(dbName)
	go app.Run(ctx, dbtest, isRead)
	<-isRead
	code := m.Run()
	defer repositories.ClearDb(dbName)
	cancel()
	os.Exit(code)
}
