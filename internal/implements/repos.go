package implements

import "aCupOfGin/internal/wires"

var (
	UserRepo = wires.InitUserRepo(DBManager)
)
