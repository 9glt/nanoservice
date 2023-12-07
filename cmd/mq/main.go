package main

import (
	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	_ = nc

	//TODO: implement

	// nc.Subscribe("records.list", func(m *nats.Msg)
	// nc.Subscribe("records.read", func(m *nats.Msg)
	// nc.Subscribe("records.create", func(m *nats.Msg)
	// nc.Subscribe("records.update", func(m *nats.Msg)
	// nc.Subscribe("records.delete", func(m *nats.Msg)
}
