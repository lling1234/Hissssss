package key

import "fmt"

const (
	msgSeqUser        = "msg:user_seq:%d_%d" // "msg:user_seq:1001_1002"
	msgSeqRoom        = "msg:room_seq:%d"
	userMappingServer = "uid:server:%d"
)

func GenUidMappingServer(uid int64) string {
	return fmt.Sprintf(userMappingServer, uid)
}

func GenMsgSeqUserToUser(from, to int64) string {
	if from > to {
		return fmt.Sprintf(msgSeqUser, from, to)
	}
	return fmt.Sprintf(msgSeqUser, to, from)
}

func GenRoomMsgSeq(rid int64) string {
	return fmt.Sprintf(msgSeqRoom, rid)
}
