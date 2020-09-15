package rabbitmq_go_connect_demo

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Configuration struct {
	AMQPConnectionURL string
}

type AddTask struct {
	Number1 int
	Number2 int
}

var Config = Configuration{
	AMQPConnectionURL: "",
}

func init() {
	Config.AMQPConnectionURL = os.Getenv("AMQPConnectionURL")
}
