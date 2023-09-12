package snowid

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	epoch = int64(1694519782066)

	nodeIDBits = uint(5)
	maxNodeID  = -1 ^ (int64(-1) << nodeIDBits)

	dataCenterIDBits = uint(5)
	maxDatacenterID  = -1 ^ (-1 << dataCenterIDBits)

	sequenceBits    = uint(12)
	maxSequenceBits = -1 ^ (-1 << sequenceBits)

	nodeIDShift     = sequenceBits
	dataCenterShift = sequenceBits + nodeIDBits
	timestampShift  = sequenceBits + nodeIDBits + dataCenterIDBits
)

type Snow struct {
	sync.Mutex
	timestamp    int64
	nodeID       int64
	dataCenterID int64
	sequence     int64
	epoch        time.Time
}

func New(nodeID int64, dataCenterID int64) (*Snow, error) {
	if dataCenterID < 0 || dataCenterID > maxDatacenterID {
		return nil, errors.New(fmt.Sprintf("dataCenterID must between 0 and %d", maxDatacenterID-1))
	}
	if nodeID < 0 || nodeID > maxNodeID {
		return nil, errors.New(fmt.Sprintf("nodeID must between 0 and %d", maxNodeID-1))
	}
	now := time.Now()
	return &Snow{
		timestamp:    0,
		nodeID:       nodeID,
		dataCenterID: dataCenterID,
		sequence:     0,
		epoch:        now.Add(time.Unix(epoch/1000, (epoch%1000)*1000000).Sub(now)),
	}, nil
}

func (s *Snow) ID() int64 {
	s.Lock()
	defer s.Unlock()

	now := time.Since(s.epoch).Milliseconds()

	if s.timestamp == now {
		s.sequence = (s.sequence + 1) & maxSequenceBits
		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.sequence = 0
	}

	s.timestamp = now
	r := (now << timestampShift) | (s.dataCenterID << dataCenterShift) | (s.nodeID << nodeIDShift) | (s.sequence)
	return r
}
