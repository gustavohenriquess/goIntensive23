package main

import (
	"encoding/json"
	"fmt"
	"os"

	usecase "github.com/gustavohenriquess/go-intensive23/internal/useCase"
	"github.com/gustavohenriquess/go-intensive23/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
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

	for {
		Menu(ch)
	}
}

func Menu(ch *amqp.Channel) {
	fmt.Println("1- Send Message")
	fmt.Println("0- Exit")

	var option int
	fmt.Scan(&option)
	fmt.Println("")

	switch option {
	case 1:
		SendMessage(ch)
	case 0:
		os.Exit(0)
	}
}

func SendMessage(ch *amqp.Channel) {
	price, tax := GetInfos()

	msg := usecase.OrderInput{Price: price, Tax: tax}
	jsonBody, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	err = rabbitmq.Publish(ch, "order", jsonBody)

	if err != nil {
		panic(err)
	}
}

func GetInfos() (float64, float64) {
	var price, tax float64

	fmt.Println("Type the Price: ")
	fmt.Scan(&price)

	fmt.Println("Type the Tax: ")
	fmt.Scan(&tax)
	fmt.Println("")

	return price, tax
}
