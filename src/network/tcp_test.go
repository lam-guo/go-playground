package network

import (
	"fmt"
	"net"
	"net/http"
	"testing"
)

func TestTcp(t *testing.T) {
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8080")
	if err != nil {
		t.Fatalf("add err:%s", err)
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		t.Fatalf("listen err:%s", err)
	}
	con, err := listener.Accept()
	if err != nil {
		t.Fatalf("accept err:%s", err)
	}
	body := []byte{}
	n, err := con.Read(body)
	if err != nil {
		t.Fatalf("read err:%s", err)
	}
	fmt.Println("read:=", n)
}

func TestServe(t *testing.T) {
	h := &Handler{}
	http.ListenAndServe("0.0.0.0:8080", h)
}

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	println(w, r)
	header := w.Header()
	fmt.Printf("header:%v\n", header.Values("1"))
}

func Serve(l net.Listener, handler http.Handler) error {
	srv := &http.Server{Handler: handler}
	return srv.Serve(l)
}
