package main

import (
	"github.com/liutaihua/xrpc"
	log "github.com/thinkboy/log4go"
)


var (
	rpcClient *xrpc.Clients
	CallTestRPC = "TEST.CallTestRPC"
	RPCPing = "TEST.Ping"
)

func main() {
	network := "tcp"
	addr := "localhost:1234"
	option := xrpc.ClientOptions{
		Proto: network,
		Addr:  addr,
	}
	rpcOptions := []xrpc.ClientOptions{option}
	rpcClient := xrpc.Dials(rpcOptions)

	// ping & reconnect
	rpcClient.Ping(RPCPing)

	arg := struct {}{}
	reply := struct {}{}

	err := rpcClient.Call(CallTestRPC, &arg, &reply)
	if err != nil {
		log.Error("rpcClient.Call(%s, %v, reply) error(%v)", CallTestRPC, arg, err)
	}
}
