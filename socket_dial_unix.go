// +build !windows

package main

import (
	"net"
	"github.com/go-sql-driver/mysql"
)

func RegisterSocketDial() {
	mysql.RegisterDial("socket", func(addr string) (net.Conn, error) {
		return net.Dial("unix",addr);
	})
}
