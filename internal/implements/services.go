package implements

import "aCupOfGin/internal/wires"

var (
	UserService = wires.InitUserService(CSVUserRepo)
)
