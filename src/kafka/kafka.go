package kafka

import (
	infra_mongodb "app/infrastructure/mongodb"
	"app/infrastructure/repository"
	kafka_hanlders "app/kafka/hanlders"
	usecase_user "app/usecase/user"

	"github.com/segmentio/kafka-go"
)

func StartKafka() {

	db := infra_mongodb.Connect()

	repositoryUser := repository.NewUserPostgres(db)
	usecaseUser := usecase_user.NewService(repositoryUser)

	var topicParams []KafkaReadTopicsParams

	topicParams = append(topicParams, KafkaReadTopicsParams{
		Topic: "user",
		Handler: func(m kafka.Message) error {
			return kafka_hanlders.CreateUser(m, usecaseUser)
		},
	})

	startKafkaConnection(topicParams)
	readTopics()
}
