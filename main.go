package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Shopify/sarama"
	"github.com/couchbase/gocb/v2"
)

type Message struct {
	Id        int
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type CouchbaseClient struct {
	c *gocb.Cluster
	b *gocb.Bucket
}

var (
	couchbaseClient *CouchbaseClient
)

func main() {
	topic, envExist := os.LookupEnv("KF_TOPIC")

	if !envExist {
		log.Fatal("Topic name is not exist")
	}

	worker, err := connectConsumer([]string{"localhost:9093"})
	if err != nil {
		panic(err)
	}

	// Calling ConsumePartition. It will open one connection per broker
	// and share it for all partitions that live on it.
	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)

	if err != nil {
		panic(err)
	}

	couchbaseClient, err = NewCouchbaseClient()
	if err != nil {
		log.Fatal(err)
	}

	defer couchbaseClient.Close()

	fmt.Println("Consumer has been started")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Count how many message processed
	msgCount := 0

	// Get signal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgCount++

				var message Message
				if err := json.Unmarshal(msg.Value, &message); err != nil {
					panic(err)
				}

				fmt.Printf("Received message Count %d: | Topic(%s) \n", msgCount, string(msg.Topic))
				fmt.Println("Message:", message)

				saveToCouchbase(message)
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCount, "messages")

	if err := worker.Close(); err != nil {
		panic(err)
	}
}

func connectConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create new consumer
	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewCouchbaseClient() (*CouchbaseClient, error) {
	options := gocb.ClusterOptions{
		Username: os.Getenv("CB_USERNAME"),
		Password: os.Getenv("CB_PASSWORD"),
	}

	cluster, err := gocb.Connect(os.Getenv("CB_HOST"), options)

	if err != nil {
		return nil, err
	}

	err = cluster.WaitUntilReady(5*time.Second, nil)

	if err != nil {
		return nil, err
	}

	bucket := cluster.Bucket(os.Getenv("CB_BUCKET"))

	return &CouchbaseClient{
		c: cluster,
		b: bucket,
	}, nil
}

func (c *CouchbaseClient) Close() error {
	return c.c.Close(&gocb.ClusterCloseOptions{})
}

func saveToCouchbase(message Message) {
	id := uuid.New()

	collection := couchbaseClient.b.Collection("")

	_, err := collection.Insert(id.String(), message, &gocb.InsertOptions{
		Context: context.Background(),
	})

	if err != nil {
		fmt.Printf("insert error: %err\n", err)
	}
}
