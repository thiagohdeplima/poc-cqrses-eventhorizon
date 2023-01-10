package main

import (
	"errors"
	"fmt"
	"os"

	es "github.com/looplab/eventhorizon/eventstore/mongodb_v2"
)

func main() {
	a, e := getEventStore()

	fmt.Println(a, e)
}

func getEventStore() (*es.EventStore, error) {
	if os.Getenv("MONGODB_URL") == "" {
		return nil, errors.New("create event store")
	}

	return es.NewEventStore(os.Getenv("MONGODB_URL"), "eventstore")
}
