package implements

import "aCupOfGin/internal/wires"

var (
	UserController = wires.InitUserController(UserService)
)
