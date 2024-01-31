package main

import (
	"app/api"
	"app/cron"
	infra_mongodb "app/infrastructure/mongodb"
	"app/infrastructure/repository"
	usecase_user "app/usecase/user"
	"log"
)

func main() {
	cron.StartCronJobs()

	conn := infra_mongodb.Connect()

	usecase := usecase_user.NewService(
		repository.NewUserPostgres(conn),
	)

	err := usecase.CreateAdminUser()
	if err != nil {
		log.Println("---------->     Error creating admin user     <----------")
		log.Println(err)
	}

	api.StartWebServer()
}
