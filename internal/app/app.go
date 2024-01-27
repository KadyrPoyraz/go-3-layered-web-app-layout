package app

import (
	"context"
	"fmt"
	"net/http"

	helloworldhandler "github.com/KadyrPoyraz/httplayout/internal/handler/http/helloworld"
	helloworldService "github.com/KadyrPoyraz/httplayout/internal/service/helloworld"
	"github.com/KadyrPoyraz/httplayout/internal/repository/db"
	"github.com/KadyrPoyraz/httplayout/internal/repository/helloworld"
	"github.com/gorilla/mux"
)

type app struct{}

func New() App {
	return &app{}
}

func (*app) Run() {
    mainRouter := mux.NewRouter().PathPrefix("/api").Subrouter()

    h := helloworldhandler.New()
    h.Fill(mainRouter)

    POSTGRES_USER := "user"
    POSTGRES_PASSWORD := "228"
    POSTGRES_HOST := "127.0.0.1"
    POSTGRES_PORT := "1337"
    POSTGRES_DB := "db"

    dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_HOST, POSTGRES_PORT, POSTGRES_DB)
    db, err := db.NewPostgresqlDB(dsn)
    if err != nil {
        fmt.Println(err)
        return
    }
    helloRepo := helloworld.NewPostgresqlRepo(db)

    helloService := helloworldService.New(helloRepo)

    ctx := context.Background()
    users, err := helloService.SayHello(ctx)
    if err != nil {
        panic(err)
    }

    fmt.Println(users)

    port := ":42069"
    fmt.Printf("Server started on localhost%s\n", port)
    http.ListenAndServe(":42069", mainRouter)
}

