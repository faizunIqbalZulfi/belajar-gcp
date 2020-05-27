package main

import (
	"belajar-gcp/config"
	"belajar-gcp/databases/postgres"
	"belajar-gcp/domain/repository"
	"belajar-gcp/domain/usecase"
	"log"

	"github.com/jinzhu/gorm"
)

type Service struct {
	UseCase usecase.UseCase
	db      *gorm.DB
}

func MakeHandler() Service {
	psql := config.Conf.Db

	log.Printf(" ========= %+v ", psql)
	log.Printf(" ========= %+v ", config.Conf.HTTPServer)

	db, err := postgres.Open(psql)
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewRepository(db)
	usecase := usecase.NewUseCase(&repo)

	return Service{
		UseCase: usecase,
		db:      db,
	}
}
