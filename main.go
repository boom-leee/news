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
	//connectionString := "postgres://koyeb-adm:D0ZGrelqfRI6@ep-empty-meadow-a15erppx.ap-southeast-1.pg.koyeb.app/koyebdb"
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		"postgres",
		"password",
		"`localhost`",
		"5432",
		"postgres",
	)
	pool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		log.Fatalln("Can not connect to sql")
	}
	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalln("Can not connect to sql")
	}
	defer pool.Close()
	//init adapter
	dummyAdapter := outAdapter.NewDummyAdapter(pool)
	userAdapter := outAdapter.NewUserAdapter(pool)
	categoryAdapter := outAdapter.NewCategoryAdapter(pool)
	//init Use case
	dummyUseCase := service.NewDummyService(dummyAdapter)
	getAllUserUseCase := service.NewGetAllUsersService(userAdapter)
	getAllCategoryUseCase := service.NewCategoriesService(categoryAdapter)

	//init handler
	dummyHandler := rest.NewDummyHandler(dummyUseCase)
	userHandler := rest.NewUserHandlers(getAllUserUseCase)
	categoryHandler := rest.NewCategoryHandlers(getAllCategoryUseCase)

	router := rest.AppRouter(dummyHandler, userHandler, categoryHandler)
	http.ListenAndServe(":3000", router)
}
