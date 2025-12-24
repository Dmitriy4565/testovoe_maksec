package checker

import (
	"context"
	"fmt"
	"net"
	"time"
)

type Result struct {
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	Status string `json:"status"`
}

func CheckPort(ctx context.Context, ip string, port int, timeout time.Duration) Result {
	dialer := net.Dialer{}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	addr := net.JoinHostPort(ip, fmt.Sprint(port))
	conn, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		return Result{IP: ip, Port: port, Status: "closed"}
	}
	conn.Close()
	return Result{IP: ip, Port: port, Status: "open"}
}
