package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

type CheckRequest struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

type CheckResponse struct {
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	Status string `json:"status"`
}

func main() {
	ip := flag.String("ip", "", "IP to check")
	port := flag.Int("port", 0, "port to check")
	url := flag.String("url", "http://localhost:8080/check", "API URL")
	flag.Parse()

	if *ip == "" || *port == 0 {
		fmt.Println("usage: cli --ip 1.2.3.4 --port 80")
		os.Exit(1)
	}

	body, _ := json.Marshal(CheckRequest{IP: *ip, Port: *port})

	resp, err := http.Post(*url, "application/json", bytes.NewReader(body))
	if err != nil {
		fmt.Println("request error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("bad status:", resp.Status)
		os.Exit(1)
	}

	var out CheckResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		fmt.Println("decode error:", err)
		os.Exit(1)
	}

	data, _ := json.Marshal(out)
	fmt.Println(string(data))
}
