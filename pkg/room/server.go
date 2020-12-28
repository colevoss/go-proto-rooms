package roomServer

import (
	"log"
	pb "proto-test/pkg/api"
)

type RoomServer struct {
	store  *roomStore
	events *RoomEvents
}

func NewRoomServer() *RoomServer {
	return &RoomServer{
		store:  NewRoomStore(),
		events: NewRoomEvents(),
	}
}

/**
 */
func (server *RoomServer) ConnectToRoom(connect *pb.RoomConnect, stream pb.RoomService_ConnectToRoomServer) error {
	log.Printf("Received connection to room")
	context := stream.Context()

	log.Printf("server events %v", server.events)

	defer func(connection *pb.RoomConnect) {
		if connection == nil {
			return
		}

		member, _ := server.store.AddMember(connection.RoomId, connection, false)
		server.events.ec.Publish("rooms."+connection.RoomId+".user.left", member)

		room, _ := server.store.GetRoom(connection.RoomId)

		res := &pb.Room{
			Id:                room.id,
			ActiveMemberCount: int32(room.activeMemberCount),
			MemberCount:       int32(room.memberCount),
		}
		server.events.ec.Publish("rooms."+res.Id+".updated", res)
	}(connect)

	log.Printf("Received stream from client user: %s room: %s", connect.UserId, connect.RoomId)

	room, _ := server.store.AddRoom(connect.RoomId)
	member, _ := server.store.AddMember(connect.RoomId, connect, true)

	server.events.ec.Publish("rooms."+connect.RoomId+".user.joined", member)

	room, _ = server.store.GetRoom(connect.RoomId)

	res := &pb.Room{
		Id:                room.id,
		ActiveMemberCount: int32(room.activeMemberCount),
		MemberCount:       int32(room.memberCount),
	}

	// done := make(chan bool, 1)
	// go func() {
	// 	// Handle room updates
	// 	for
	// 		select {
	// 		case <-done:
	// 			return
	// 		}
	// 	}
	// }()

	stream.Send(res)

	server.events.ec.Publish("rooms."+res.Id+".updated", res)

	roomUpdateChan := make(chan *pb.Room)
	server.events.ec.BindRecvChan("rooms."+res.Id+".updated", roomUpdateChan)

	for {
		select {
		case room := <-roomUpdateChan:
			stream.Send(room)
		case <-context.Done():
			log.Printf("Disconnected")
			// done <- true
			return context.Err()
		}
	}
}

/*
 */
func (server *RoomServer) Members(connect *pb.RoomConnect, stream pb.RoomService_MembersServer) error {
	room, _ := server.store.GetRoom(connect.RoomId)
	context := stream.Context()

	for _, member := range room.members {
		stream.Send(member)
	}

	memberJoinedChan := make(chan *pb.RoomMember)
	memberLeftChan := make(chan *pb.RoomMember)

	server.events.ec.BindRecvChan("rooms."+connect.RoomId+".user.joined", memberJoinedChan)
	server.events.ec.BindRecvChan("rooms."+connect.RoomId+".user.left", memberLeftChan)

	// done := make(chan bool, 1)

	// go func() {
	// 	for {
	// 		select {
	// 		case <-done:
	// 			close(memberJoinedChan)
	// 			close(memberLeftChan)
	// 			return
	// 		}
	// 	}
	// }()

	for {
		select {
		case member := <-memberJoinedChan:
			log.Printf("Received Member joined event %s", member.UserId)
			stream.Send(member)

		case member := <-memberLeftChan:
			log.Printf("Received Member left room %s", member.UserId)
			stream.Send(member)

		case <-context.Done():
			close(memberJoinedChan)
			close(memberLeftChan)
			// done <- true
			return context.Err()
		}
	}
}
