package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/gustavohenriquess/go-intensive23/internal/Infra/database"
	usecase "github.com/gustavohenriquess/go-intensive23/internal/useCase"
	"github.com/gustavohenriquess/go-intensive23/pkg/rabbitmq"
	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close() // espera tudo rodar e depois executa o close
	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	err = rabbitmq.DeclareQueue(ch, []string{"order"})
	if err != nil {
		fmt.Println("Erro ao declarar fila")
		panic(err)
	}

	msgRabbitmqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgRabbitmqChannel) // escutando a fila // trava // T2
	rabbitmqWorker(msgRabbitmqChannel, uc)      // T1
}

func rabbitmqWorker(msgChan chan amqp.Delivery, uc *usecase.CalculateFinalPrice) {
	fmt.Println("Starting rabbitmq")
	for msg := range msgChan {
		var input usecase.OrderInput
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			panic(err)
		}
		output, err := uc.Execute(input)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		fmt.Println("Mensagem processada e salva no banco:", output)
	}
}
