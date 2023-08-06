package connect

import (
	"sync"
)

type Bucket struct {
	Lock     sync.RWMutex
	Channels map[int64]*Channel
	Rooms    map[int64]*Room
}

type Option struct {
	ChannelSize int
	RoomSize    int
}

func NewBucket(option Option) *Bucket {
	b := &Bucket{
		Channels: make(map[int64]*Channel, option.ChannelSize),
		Rooms:    make(map[int64]*Room, option.RoomSize),
	}
	return b
}

func (b *Bucket) ChannelX(uid int64) *Channel {
	b.Lock.RLock()
	c, _ := b.Channels[uid]
	b.Lock.RUnlock()
	return c
}

func (b *Bucket) RoomX(rid int64) *Room {
	b.Lock.RLock()
	r, _ := b.Rooms[rid]
	b.Lock.RUnlock()
	return r
}

func (b *Bucket) RoomN() int {
	return len(b.Rooms)
}

func (b *Bucket) ChannelN() int {
	return len(b.Channels)
}

func (b *Bucket) JoinC(uid int64, ch *Channel) {
	b.Lock.RLock()
	b.Channels[uid] = ch
	b.Lock.RUnlock()
}

func (b *Bucket) JoinR(rid int64, ch *Channel) {
	b.Lock.RLock()
	defer b.Lock.RUnlock()
	r, ok := b.Rooms[rid]
	if !ok {
		r = NewRoom(rid)
		b.Rooms[rid] = r
	}
	r.Join(ch)
	ch.Room = r
}

func (b *Bucket) LeaveC(ch *Channel) {
	b.Lock.RLock()
	defer b.Lock.RUnlock()
	r := ch.Room
	delete(b.Channels, ch.UID)
	if r != nil {
		r.Leave(ch)
		if r.Empty {
			delete(b.Rooms, r.RID)
		}
	}
}

// BroadcastToAllChannel 广播到所有用户
func (b *Bucket) BroadcastToAllChannel(msg []byte) {
	b.Lock.RLock()
	for _, ch := range b.Channels {
		ch.SendM(msg)
	}
	b.Lock.RUnlock()
}

// BroadcastToAllRoom 广播到所有房间
func (b *Bucket) BroadcastToAllRoom(msg []byte) {
	b.Lock.RLock()
	for _, room := range b.Rooms {
		room.Broadcast(msg)
	}
	b.Lock.RUnlock()
}
