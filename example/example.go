package main

import (
	"github.com/liutaihua/xrpc"
	log "github.com/thinkboy/log4go"
)


var rpcClient *xrpc.Clients

var CallTestRPC = "TEST.CallTestRPC"

func main() {
	network := "tcp"
	addr := "localhost:1234"
	options := xrpc.ClientOptions{
		Proto: network,
		Addr:  addr,
	}
	rpcOptions := []xrpc.ClientOptions{option}
	rpcClient := xrpc.Dials(rpcOptions)

	// ping & reconnect
	rpcClient.Ping(CometServicePing)

	arg := struct {}{}
	reply := struct {}{}

	err = rpcClient.Call(CallTestRPC, &arg, &reply)
	if err != nil {
		log.Error("rpcClient.Call(%s, %v, reply) error(%v)", CallTestRPC, arg, err)
	}
}
