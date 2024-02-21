package client

// 在线玩家迭代器
func IterationPlayerConn(f func(playerId string) bool) {
	WsConnTree.Lock.RLock()
	defer WsConnTree.Lock.RUnlock()
	for _, playerConn := range WsConnTree.Conns {
		if !f(playerConn.PlayerId) {
			break
		}
	}
}
