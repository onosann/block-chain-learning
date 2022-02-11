package main

import (
	json2 "encoding/json"
	"fmt"
	"testing"
)

func TestNewETHRPCClient(t *testing.T) {
	client2 := NewETHRPCClient("www.nihao.com").GetRpc()
	if client2 == nil {
		fmt.Println("初始化失败")
	}

	client := NewETHRPCClient("123://456").GetRpc()
	if client == nil {
		fmt.Println("初始化失败")
	}
}

func Test_GetTransactionByHash(t *testing.T) {
	nodeUrl := "https://mainnet.infura.io/v3/182344b8c3154851ad6a5544fc6b3b07"
	txHash := "0x90e78567f562ee38c8a6eaa1a09b12d6de02c8452b27f0da450d8fd0f07cb3c3"
	if txHash == "" || len(txHash) != 66 {
		fmt.Println("非法的交易哈希值")
		return
	}

	txInfo, err := NewETHRPCRequester(nodeUrl).GetTransactionByHash(txHash)
	if err != nil {
		fmt.Println("查询交易失败，信息是：", err.Error())
		return
	}
	json,_ :=json2.Marshal(txInfo)
	fmt.Println(string(json))
}

func Test_GetTransactions(t *testing.T) {
	nodeUrl := "https://mainnet.infura.io/v3/182344b8c3154851ad6a5544fc6b3b07"
	txHash1 := "0x8e43b2ffc11e4bf8484ee1c6c73a22575f03ec370ed2158059d573ebe4c42f9c"
	txHash2 := "0x90e78567f562ee38c8a6eaa1a09b12d6de02c8452b27f0da450d8fd0f07cb31d"
	txHash3 := "0xfc118d98562613065a84c0edb704527e60d1c1c3e2a61f8aef6deb351cede83d"

	txHashs :=[]string{}
	txHashs = append(txHashs,txHash1,txHash2,txHash3)

	if txHashs ==nil || len(txHashs)==0{
		fmt.Println("非法的交易哈希值数组")
		return
	}

	txInfos,err:= NewETHRPCRequester(nodeUrl).GetTransactions(txHashs)
	if err!=nil{
		fmt.Println("查询交易失败，信息是：", err.Error())
		return
	}

	json,_ :=json2.Marshal(txInfos)
	fmt.Println(string(json))
}

func Test_GetETHBalance(t *testing.T) {
	nodeUrl := "https://mainnet.infura.io/v3/182344b8c3154851ad6a5544fc6b3b07"
	address :="0xb873312eea2b9f07e73f3df5f628cf4e01026632"
	if address ==""|| len(address) !=42{
		fmt.Println("非法点交易地址值")
		return
	}
	balance,err := NewETHRPCRequester(nodeUrl).GetETHBalance(address)
	if err!=nil{
		fmt.Println("查询ETH余额失败，信息是：", err.Error())
		return
	}
	fmt.Println(balance)

}

func Test_GetETHBalances(t *testing.T) {
	nodeUrl := "https://mainnet.infura.io/v3/182344b8c3154851ad6a5544fc6b3b07"
	address1 :="0x11587ce064f95814E8c71D7cF1A5b6EB7a22bd83"
	address2 :="0x7a689a6bd3b8df172224a96e8f5c85d5f1dd8ace"
	address :=[]string{address1,address2}

	balance,err := NewETHRPCRequester(nodeUrl).GetETHBalances(address)
	if err!=nil{
		fmt.Println("查询ETH余额失败，信息是：", err.Error())
		return
	}
	fmt.Println(balance)
}

func Test_GetERC20Balances(t *testing.T) {
	nodeUrl := "https://mainnet.infura.io/v3/182344b8c3154851ad6a5544fc6b3b07"
	address := "0x11587ce064f95814E8c71D7cF1A5b6EB7a22bd83"
	contract1 := "0x78021ABD9b06f0456cB9DB95a846C302c34f8b8D"
	contract2 := "0xB8c77482e45F1F44dE1745F52C74426C631bDD52"

	params :=[]ERC20BalanceRpcReq{}
	item := ERC20BalanceRpcReq{}
	item.ContractAddres=contract1
	item.UserAddress=address
	item.ContractDecimal=18

	params =append(params,item)
	item.ContractAddres=contract2
	params =append(params,item)

	balance,err :=NewETHRPCRequester(nodeUrl).GetERC20Balances(params)
	if err!=nil{
		fmt.Println("查询eth余额失败，信息是：", err.Error())
		return
	}
	fmt.Println(balance)
}

func TestGetLatestBlockNumber(t *testing.T) {
	nodeUrl := "https://mainnet.infura.io/v3/182344b8c3154851ad6a5544fc6b3b07"
	number,err :=NewETHRPCRequester(nodeUrl).GetLatestBlockNumber()
	if err!=nil{
		fmt.Println("获取区块号失败，信息是", err.Error())
		return
	}
	fmt.Println("10进制: ",number.String())
}

// 测试区块号 14178047

func TestGetFullBlockInfo(t *testing.T){
	nodeUrl := "https://mainnet.infura.io/v3/182344b8c3154851ad6a5544fc6b3b07"
	requester := NewETHRPCRequester(nodeUrl)
	number,_ :=requester.GetLatestBlockNumber()
	fmt.Println("区块号是:\n",number)
	fullBlock,err := requester.GetBlockInfoByNumber(number)
	if err!=nil{
		fmt.Println("获取区块信息失败，信息是：",err.Error())
		return
	}
	json,_:=json2.Marshal(fullBlock)
	fmt.Println("根据区块号获取区块信息：\n",string(json))
}

func TestGetFullBlockByBlockHash(t *testing.T){
	nodeUrl := "https://mainnet.infura.io/v3/182344b8c3154851ad6a5544fc6b3b07"
	requester := NewETHRPCRequester(nodeUrl)
	blockHash := "0x5f266b7d1350a23cddf8b7d98859806f13909091698aeebc86a3aefcdcfd5d68"
	fullBlock,err := requester.GetBlockInfoByHash(blockHash)
	if err!=nil{
		fmt.Println("获取区块信息失败，信息是：",err.Error())
		return
	}
	json,_:=json2.Marshal(fullBlock)
	fmt.Println("根据区块号获取区块信息：\n",string(json))
}
func Test_MakeMethodId(t *testing.T){
	contractABI :=`[
		{ "constant": true,
          "inputs": [{"name" : "arg1", "type": "uint8"},{"name" : "arg2", "type": "uint8"}],
          "name" ："add","outputs":[{ "name":"","type":"uint8"}],
          "payable": false, "stateMutability":"pure", "type": "function"
		}]`
	methodName :="add"
	methodId,err := MakeMethoId(methodName,contractABI)
	if err!=nil{
		fmt.Println("生成 method 失败",err.Error())
		return
	}
	fmt.Println("生成 methodId 成功",methodId)
}

func Test_CreateETHWallet(t *testing.T) {
	nodeUrl := "https://mainnet.infura.io/v3/182344b8c3154851ad6a5544fc6b3b07"
	address1,err :=NewETHRPCRequester(nodeUrl).CreateETHWallet("12345")

	if err!=nil{
		fmt.Println("第一次，创建钱包失败",err.Error())
	}else{
		fmt.Println("第一次，创建钱包成功，以太坊地址是：",address1)
	}
	address2,err :=NewETHRPCRequester(nodeUrl).CreateETHWallet("123456a")
	if err!=nil{
		fmt.Println("第二次，创建钱包失败",err.Error())
	}else{
		fmt.Println("第二次，创建钱包成功，以太坊地址是：",address2)
	}
}