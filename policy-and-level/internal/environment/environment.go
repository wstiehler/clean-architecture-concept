package environment

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var lock = &sync.Mutex{}

type Single struct {
	DATABASE_NAME     string
	DATABASE_HOST     string
	DATABASE_PORT     int64
	DATABASE_USER     string
	DATABASE_PASSWORD string
}

func init() {
	env := GetInstance()
	env.Setup()
}

func (e *Single) Setup() {
	e.DATABASE_HOST = getenv("DATABASE_HOST", "localhost")
	e.DATABASE_NAME = getenv("DATABASE_NAME", "postgres")
	e.DATABASE_PORT = getenvInt64("DATABASE_PORT", 5432)
	e.DATABASE_USER = getenv("DATABASE_USER", "postgres")
	e.DATABASE_PASSWORD = getenv("DATABASE_PASSWORD", "postgres")

}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func getenvInt64(key string, fallback int64) int64 {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	valueInt, _ := strconv.ParseInt(value, 10, 64)
	return valueInt
}

var singleInstance *Single

func GetInstance() *Single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now")
			singleInstance = &Single{}
			singleInstance.Setup()
		} else {
			fmt.Println("Single instance already created")
		}

	}
	return singleInstance
}
