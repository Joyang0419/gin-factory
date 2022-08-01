package implements

import "aCupOfGin/internal/wires"

var (
	GORMUserRepo = wires.InitGORMUserRepo(GORMDBManager)
	CSVUserRepo  = wires.InitCSVUserRepo(CSVDBManager)
)
