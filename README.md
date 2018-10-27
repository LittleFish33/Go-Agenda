# Go-Agenda
服务计算课程作业——CLI 命令行实用程序开发实战 - Agenda

## 2018.10.26

### New features

- 用cobra创建项目结构
- 实现user和meeting数据json文件读写

### run

1. 根据自身文件目录，修改`testJson.go`文件中的`UFilepath`和`MFilepath`地址
2. 在`Go-Agenda$`目录下，执行命令`go run testJson.go`

### test

1. 如果终端有输出UJson、MJson文件内容，则读取成功
2. 如果UJson、MJson分别和UJsonO、MJsonO内容相同，写入成功
3. test后删除文件UJsonO、MJsonO

### File structure

```
Go-Agenda
├── cmd
│   └── root.go：
├── entity
│   ├── Data
│   │   ├── MJson: text格式，JsonIOForUM.go所有会议信息读取和写入的json文件
│   │   └── UJson: text格式，JsonIOForUM.go所有用户信息读取和写入的json文件
│   ├── Date.go: 将日期的年月日作为string封装到Data.go，用于Meeting.go中会议的日期类型
│   ├── JsonIOForUM.go: 实现对User和Meeting类型对象的json文件读写
│   ├── Meeting.go: 将Meeting的实体内容及相关方法封装
│   └── User.go: 将User的实体内容及相关方法封装
├── LICENSE
├── main.go
├── README.md
└── testJson.go: 测试json文件的读写，项目完成后删除
```





