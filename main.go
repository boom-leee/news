package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"news-api/adapter/in/rest"
	outAdapter "news-api/adapter/out"
	"news-api/application/domain/service"
)

func main() {
	ctx := context.Background()
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		"postgres",
		"password",
		"localhost",
		"5432",
		"postgres",
	)
	d, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		log.Fatalln("Can not connect to sql")
	}
	err = d.Ping(ctx)
	if err != nil {
		log.Fatalln("Can not connect to sql")
	}
	defer d.Close()
	//init adapter
	dummyAdapter := outAdapter.NewDummyAdapter()

	//init Use case
	dummyUseCase := service.NewDummyService(dummyAdapter)

	//init handler
	dummyHandler := rest.NewDummyHandler(dummyUseCase)
	router := rest.AppRouter(dummyHandler)
	http.ListenAndServe(":3000", router)
}
