package main

import (
	"./entity"
)

func main() {
	const UFilepath string = "/home/wjh/gowork/src/Go-Agenda/entity/Data/UJson"
	const MFilepath string = "/home/wjh/gowork/src/Go-Agenda/entity/Data/MJson"
	
	var users []entity.User = entity.UReadFromJsonFile(UFilepath)
	entity.UWriteToJsonFile(users, UFilepath)
	var meetings []entity.Meeting = entity.MReadFromJsonFile(MFilepath)
	entity.MWriteToJsonFile(meetings, MFilepath)
}