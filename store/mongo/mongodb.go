package mongo

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

//Config db config
type Config struct {
	Host     string
	Port     int
	DB       string
	User     string
	Password string
	Socket   string
}

// MongoInstance : MongoInstance Struct
type Client struct {
	Client *mongo.Client
	DB     *mongo.Database
	models []interface{}
}

func (cfg *Config) URI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/?compressors=disabled&gssapiServiceName=mongodb",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)
}
// MI : An instance of MongoInstance Struct
var MI Client
// ConnectDB - database connection
func ConnectDB(c *Config, models []interface{}) *Client {
	fmt.Println(c.URI())
	client, err := mongo.NewClient(options.Client().ApplyURI(c.URI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected!")

	MI = Client{
		Client: client,
		DB:     client.Database(os.Getenv("DATABASE_NAME")),
		models: models,
	}
	return &MI
}

// Teardown teardown all db tables
func (c *Client) Teardown() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return errors.Wrap(c.DB.Drop(ctx), "DropTableIfExists")
}
