package roomServer

import (
	"errors"
	"fmt"
	pb "proto-test/pkg/api"
	"sync"
)

type roomRecord struct {
	id                string
	members           map[string]*pb.RoomMember
	memberCount       int
	activeMemberCount int
}

func NewRoomRecord(id string) *roomRecord {
	return &roomRecord{
		id:      id,
		members: make(map[string]*pb.RoomMember),
	}
}

type roomStore struct {
	mutex sync.RWMutex
	rooms map[string]*roomRecord
}

func NewRoomStore() *roomStore {
	return &roomStore{
		rooms: make(map[string]*roomRecord),
	}
}

func (store *roomStore) AddRoom(id string) (*roomRecord, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	room, ok := store.rooms[id]

	if !ok {
		room = NewRoomRecord(id)
		store.rooms[id] = room
	}

	return room, nil
}

func (store *roomStore) AddMember(roomId string, member *pb.RoomConnect, present bool) (*pb.RoomMember, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	room := store.rooms[roomId]

	if room == nil {
		return nil, errors.Unwrap(fmt.Errorf("Room %s does not exist", roomId))
	}

	roomMember := &pb.RoomMember{
		UserId:  member.UserId,
		Name:    member.Name,
		Present: present,
	}

	room.members[member.UserId] = roomMember

	var members int
	var activeMembers int

	for _, member := range room.members {
		members += 1

		if member.Present {
			activeMembers += 1
		}
	}

	room.activeMemberCount = activeMembers
	room.memberCount = members

	return roomMember, nil
}

func (store *roomStore) GetRoom(roomId string) (*roomRecord, error) {
	room := store.rooms[roomId]

	if room == nil {
		return nil, errors.Unwrap(fmt.Errorf("Room %s does not exist", roomId))
	}

	return room, nil
}
