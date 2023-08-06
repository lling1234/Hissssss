package key

import "fmt"

const (
	userMappingServer = "uid:server:%d"
)

func GenUidMappingServer(uid int64) string {
	return fmt.Sprintf(userMappingServer, uid)
}
