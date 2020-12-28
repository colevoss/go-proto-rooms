package roomServer

import (
	nats "github.com/nats-io/nats.go"
	natsp "github.com/nats-io/nats.go/encoders/protobuf"
)

type RoomEvents struct {
	nc *nats.Conn
	ec *nats.EncodedConn
}

func (re *RoomEvents) Close() {
	re.ec.Close()
}

func NewRoomEvents() *RoomEvents {
	// nc, _ := nats.Connect(nats.DefaultURL)
	nc, _ := nats.Connect("nats://nats:4222")
	ec, _ := nats.NewEncodedConn(nc, natsp.PROTOBUF_ENCODER)

	re := &RoomEvents{
		nc: nc,
		ec: ec,
	}

	return re
}
