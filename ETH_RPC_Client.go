package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
)
type ETHRPCClient struct {
	NodeUrl string
	client *rpc.Client
}

func NewETHRPCClient(nodeUrl string) *ETHRPCClient{
	// &代表取的是指针
	client := &ETHRPCClient{
		NodeUrl: nodeUrl,
	}
	client.initRpc()
	return client
}

func (erc *ETHRPCClient) initRpc(){
	rpcClient,err :=rpc.DialHTTP(erc.NodeUrl)
	if err !=nil {
		errInfo:=fmt.Errorf("初始化 rpc client 失败%s",err.Error()).Error()
		panic(errInfo)
	}
	erc.client=rpcClient
}

func (erc *ETHRPCClient) GetRpc() *rpc.Client{
	if erc.client ==nil {
		erc.initRpc()
	}
	return erc.client
}