package entity

import (
	"bufio"
	"os"
)

<<<<<<< HEAD
var UFilepath = "F:/Users/HP/GOPATH/src/Go-Agenda/entity/Data/UJson"
var ULoginFilepath = "F:/Users/HP/GOPATH/src/Go-Agenda/entity/Data/ULoginJson"
var MFilepath = "F:/Users/HP/GOPATH/src/Go-Agenda/entity/Data/MJson"
=======
var UFilepath = "F:/GOProject/src/Go-Agenda/entity/Data/UJson"
var ULoginFilepath = "F:/GOProject/src/Go-Agenda/entity/Data/ULoginJson"
var MFilepath = "F:/GOProject/src/Go-Agenda/entity/Data/MJson"
>>>>>>> d7679333bec7ebada26c28715d286fad80b5ae38
var Users []User
var Meetings []Meeting
var Loginuser User

// 读所有用户和会议的数据
func ReadJson() {
	Users = UReadFromJsonFile(UFilepath)
	Meetings = MReadFromJsonFile(MFilepath)
}

<<<<<<< HEAD
】// 读登录信息
=======
// 读登录信息
>>>>>>> d7679333bec7ebada26c28715d286fad80b5ae38
func ReadLoginJson() {
	inputFile, _ := os.Open(ULoginFilepath)
	inputReader := bufio.NewReader(inputFile)
	line, _ := inputReader.ReadString('\n')
	if line != "" {
		Loginuser = UserJsonDecode([]byte(line))
	}
}

// 写登录信息
func WriteLoginJson() {
	outputFile, _ := os.OpenFile(ULoginFilepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	var userByte = UserJsonEncode(Loginuser)
	userByte = append(userByte, '\n')
	outputFile.Write(userByte)
	outputFile.Close()
}

// 写所有用户和会议的数据
func WriteJson() {
	UWriteToJsonFile(Users, UFilepath)
	MWriteToJsonFile(Meetings, MFilepath)
}

<<<<<<< HEAD

=======
>>>>>>> d7679333bec7ebada26c28715d286fad80b5ae38
func GetUsers() []User {
	return Users
}

func GetMeetings() []Meeting {
	return Meetings
}

func AddUser(u User) string {
	Users = append(Users, u)
	WriteJson()
	return u.Name
}

// 判断是否登录
func IsLogin() bool {
	return Loginuser.Name != ""
}
<<<<<<< HEAD
=======

func AddMeeting(meet Meeting) {
	Meetings = append(Meetings, meet)
	WriteJson()
}

func GetMeetingByTitle(title string) Meeting {
	var meeting Meeting
	for _, meet := range Meetings {
		if title == meet.Title {
			meeting = meet
		}
	}
	return meeting
}

func IsUser(username string) bool {
	for _, user := range Users {
		if user.Name == username {
			return true
		}
	}
	return false
}
>>>>>>> d7679333bec7ebada26c28715d286fad80b5ae38
