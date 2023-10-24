package config

import (
	"fmt"
	"os"
	"strconv"
)

var PORT int
var VER int

func init() {
	port := os.Getenv("PORT")
	version := os.Getenv("VERSION")
	if port == "" || version == "" {
		panic("invalid port and version")
	}
	var err error
	PORT, err = strconv.Atoi(port)
	if err != nil {
		panic(fmt.Sprintf("parse port err: %v", err))
	}
	VER, err = strconv.Atoi(version)
	if err != nil {
		panic(fmt.Sprintf("parse version error: %v", err))
	}
}
