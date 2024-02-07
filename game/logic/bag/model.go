package bag

import "easy-game/game/logic/player"

type Item struct {
	ItemId int
	Num    int
}

var (
	BagChangeFields = []string{
		player.USER_ID,
		player.BAG,
	}
)
