package connect

import (
	"sync"
)

type Room struct {
	Lock   sync.RWMutex
	RID    int64
	Next   *Channel
	Online int
	Empty  bool
}

func NewRoom(rid int64) *Room {
	return &Room{RID: rid}
}

func (r *Room) Join(c *Channel) {
	r.Lock.RLock()
	defer r.Lock.Unlock()
	if r.Next == nil {
		r.Next = c
	} else {
		r.Next.Prev = c
		c.Next = r.Next
		r.Next = c
	}
	r.Online++
}

func (r *Room) Leave(c *Channel) {
	r.Lock.RLock()
	defer r.Lock.RUnlock()
	if c.Next != nil {
		c.Next.Prev = c.Prev
	}
	if c.Prev != nil {
		c.Prev.Next = c.Next
	} else {
		r.Next = c.Next
	}
	r.Online--
	r.Empty = false
	if r.Online <= 0 {
		r.Empty = true
	}
}

func (r *Room) Broadcast(msg []byte) {
	r.Lock.RLock()
	defer r.Lock.Unlock()
	for ch := r.Next; ch != nil; ch = ch.Next {
		ch.SendM(msg)
	}
}
