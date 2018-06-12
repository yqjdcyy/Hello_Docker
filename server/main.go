package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	// Port 服务端口
	Port int
	// Path 文件保存目录
	Path string
)

func init() {
	flag.IntVar(&Port, "port", 8888, "server start port")
}

func main() {
	
	flag.Parse()

	addr := fmt.Sprintf(":%d", Port)
	fmt.Printf("Server start at %s", addr)
	r := mux.NewRouter()
	r.HandleFunc("/call", CallHandler).Methods("GET", "POST")
	http.ListenAndServe(addr, r)
}

// CallHandler 回调
func CallHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "CALL.BACK")
}
