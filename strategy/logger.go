package strategy

import "fmt"

// 实现一个日志记录器（相当于 Context）
type LogManager struct {
	Logging
}

func NewLogManager(logging Logging) *LogManager {
	return &LogManager{logging}
}

// 抽象的日志
type Logging interface {
	Info()
	Error()
}

// 实现具体的日志：文件方式日志
type FileLogging struct {
}

func (fl *FileLogging) Info() {
	fmt.Println("文件记录 Info")
}

func (fl *FileLogging) Error() {
	fmt.Println("文件记录 Error")
}

// 实现具体的日志：数据库方式日志
type DBLogging struct {
}

func (dl *DBLogging) Info() {
	fmt.Println("数据库记录 Info")
}

func (dl *DBLogging) Error() {
	fmt.Println("数据库记录 Error")
}
