package main

import (
	"net"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/Microsoft/go-winio"
)

func RegisterSocketDial() {
	mysql.RegisterDial("socket", func(addr string) (net.Conn, error) {
		return winio.DialPipe(fmt.Sprintf("\\\\.\\pipe\\%s",addr), nil)
	})
}
