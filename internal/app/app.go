package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KadyrPoyraz/httplayout/config"
	helloworldhandler "github.com/KadyrPoyraz/httplayout/internal/handler/http/helloworld"
	"github.com/KadyrPoyraz/httplayout/internal/repository/db"
	"github.com/KadyrPoyraz/httplayout/internal/repository/helloworld"
	helloworldService "github.com/KadyrPoyraz/httplayout/internal/service/helloworld"
	"github.com/gorilla/mux"
)

type app struct{
    cnf config.Config
}

func New(cnf config.Config) App {
	return &app{
        cnf: cnf,
    }
}

func (a *app) Run() error {
    mainRouter := mux.NewRouter().PathPrefix("/api").Subrouter()

    h := helloworldhandler.New()
    h.Fill(mainRouter)

    dsnFormat := "postgresql://%s:%s@%s:%s/%s?sslmode=disable"
    dbCnf := a.cnf.DB
    dsn := fmt.Sprintf(dsnFormat, dbCnf.User, dbCnf.Password, dbCnf.Host, dbCnf.Port, dbCnf.Name)
    db, err := db.NewPostgresqlDB(dsn)
    if err != nil {
        return err
    }
    helloRepo := helloworld.NewPostgresqlRepo(db)

    helloService := helloworldService.New(helloRepo)

    ctx := context.Background()
    users, err := helloService.SayHello(ctx)
    if err != nil {
        panic(err)
    }

    fmt.Println(users)

    port := ":" + a.cnf.App.Port
    fmt.Printf("Server started on localhost%s\n", port)
    err = http.ListenAndServe(port, mainRouter)
    if err != nil {
        panic(err)
    }

    return nil
}

