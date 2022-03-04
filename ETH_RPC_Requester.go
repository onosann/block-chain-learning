package main

import (
	"ETHtest/model"
	"ETHtest/tool"
	"errors"
	"fmt"
	abi2 "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

type ETHRPCRequester struct {
	NonceManager *NonceManager
	client *ETHRPCClient
}

type ERC20BalanceRpcReq struct {
	ContractAddres string
	UserAddress string
	ContractDecimal int
}

func NewETHRPCRequester(nodeUrl string) *ETHRPCRequester {
	requester :=&ETHRPCRequester{}
	//实例化 nonce 管理器
	requester.NonceManager =NewNonceManager()
	//实例化 rpc 客户端
	requester.client=NewETHRPCClient(nodeUrl)
	return requester
}

func (r *ETHRPCRequester) GetTransactionByHash(txHash string)(model.Transaction,error){
	methodName :="eth_getTransactionByHash"
	result :=model.Transaction{}
	err:= r.client.GetRpc().Call(&result,methodName,txHash)
	return result,err
}

func (r *ETHRPCRequester) GetTransactions(txHashs []string) ([]*model.Transaction,error) {
	name :="eth_getTransactionByHash"
	rets :=[]*model.Transaction{}
	size :=len(txHashs)

	reqs :=[]rpc.BatchElem{}
	for i:=0;i<size;i++{
		ret :=model.Transaction{}

		req:= rpc.BatchElem{
			Method: name,
			Args:[]interface{}{txHashs[i]},
			// 这里不太清楚要干嘛
			Result: &ret,
		}
		reqs =append(reqs,req)
		rets =append(rets,&ret)
	}
	err :=r.client.GetRpc().BatchCall(reqs)
	return rets,err
}

func (r *ETHRPCRequester) GetETHBalance(address string) (string,error){
	name :="eth_getBalance"
	result :=""

	err := r.client.GetRpc().Call(&result,name,address,"latest")
	if err!= nil{
		return "",err
	}

	if result== ""{
		return "",errors.New("eth balance is null")
	}

	ten,_ :=new(big.Int).SetString(result[2:],16)
	return ten.String(),nil
}

func (r *ETHRPCRequester) GetETHBalances(address []string) ([]string,error) {
	name := "eth_getBalance"
	rets := []*string{}

	size := len(address)
	reqs := []rpc.BatchElem{}
	for i := 0; i < size; i++ {
		ret := ""
		req := rpc.BatchElem{
			Method: name,
			Args:   []interface{}{address[i], "latest"},
			//&ret 传入单个请求的结果引用，保证它在函数内部被修改值后，回到函数外时仍然有效
			Result: &ret,
		}
		reqs = append(reqs, req)
		rets = append(rets, &ret)
	}
	err := r.client.GetRpc().BatchCall(reqs)
	if err != nil {
		return nil, err
	}

	for _, req := range reqs {
		if req.Error != nil {
			return nil, req.Error
		}
	}
	finalRet := []string{}
	for _, item := range rets {
		ten, _ := new(big.Int).SetString((*item)[2:], 16)
		finalRet = append(finalRet, ten.String())
	}
	return finalRet,err
}

func (r *ETHRPCRequester) GetERC20Balances(paramArr []ERC20BalanceRpcReq)([]string,error){
	name :="eth_call"
	methodID :="0x70a08231"

	rets :=[]*string{}
	size :=len(paramArr)
	reqs :=[]rpc.BatchElem{}
	for i:=0;i<size;i++{
		ret :=""
		arg := &model.CallArg{}
		userAddress :=paramArr[i].UserAddress

		arg.To =common.HexToAddress(paramArr[i].ContractAddres)

		arg.Data =methodID+"000000000000000000000000"+userAddress[2:]

		req := rpc.BatchElem{
			Method: name,
			Args: []interface{}{arg,"latest"},
			Result: &ret,
		}
		reqs =append(reqs,req)
		rets =append(rets,&ret)
	}
	err := r.client.GetRpc().BatchCall(reqs)
	if err != nil {
		return nil,err
	}

	for _,req := range reqs{
		if req.Error != nil {
			return nil,req.Error
		}
	}
	finalRet := []string{}
	for _,item := range rets{
		if *item == ""{
			continue
		}
		ten,_ :=new(big.Int).SetString((*item)[2:],16)
		finalRet =append(finalRet,ten.String())
	}
	return finalRet,err
}

func (r *ETHRPCRequester) GetLatestBlockNumber() (*big.Int,error)  {
	methodName :="eth_blockNumber"
	number :=""
	err := r.client.client.Call(&number,methodName)
	if err != nil{
		return  nil,fmt.Errorf("获取最新区块号失败！ %s",err.Error())
	}
	ten,_ :=new(big.Int).SetString(number[2:],16)
	return ten,nil
}

func (r *ETHRPCRequester) GetBlockInfoByNumber(blockNumber *big.Int) (*model.FullBLock, error){
	number :=fmt.Sprintf("%#x", blockNumber)
	methodName := "eth_getBlockByNumber"
	fullBlock := model.FullBLock{}

	err := r.client.client.Call(&fullBlock, methodName, number, true)
	if err!=nil {
		return nil,fmt.Errorf("get block info failed! %s",err.Error())
	}
	if fullBlock.Number == ""{
		return nil, fmt.Errorf("block info is empty %s",blockNumber.String())
	}
	return &fullBlock,nil
}

func (r *ETHRPCRequester) GetBlockInfoByHash(blockHash string) (*model.FullBLock, error){
	methodName := "eth_getBlockByHash"
	fullBlock := model.FullBLock{}

	err := r.client.client.Call(&fullBlock, methodName, blockHash, true)
	if err!=nil {
		return nil,fmt.Errorf("get block info failed! %s",err.Error())
	}
	if fullBlock.Number == ""{
		return nil, fmt.Errorf("block info is empty %s",blockHash)
	}
	return &fullBlock,nil
}

func MakeMethoId(methodName string,abiStr string) (string,error) {
	abi := &abi2.ABI{}
	err := abi.UnmarshalJSON([]byte(abiStr))
	if err != nil{
		return  "",err
	}
	method := abi.Methods[methodName]
	methodIdBytes := method.ID
	methodId := "0x" +common.Bytes2Hex(methodIdBytes)
	return methodId,nil
}

func (r *ETHRPCRequester) CreateETHWallet(password string) (string,error){
	if password ==""{
		return "",errors.New("password cant empty")
	}
	if len(password) <6 {
		return "", errors.New("password's len must more than 6")
	}
	keydir :="./keystores"

	ks := keystore.NewKeyStore(keydir,keystore.StandardScryptN,keystore.StandardScryptP)
	wallet,err := ks.NewAccount(password)
	if err!= nil{
		return "0x",err
	}
	return  wallet.Address.String(),nil
}

func (r *ETHRPCRequester) SendTransaction(address string,transaction *types.Transaction) (string,error)  {
	signTx,err := tool.SignETHTransaction(address,transaction)
	if err!= nil{
		return "",fmt.Errorf("签名失败！ %s", err.Error())
	}
	txRlpData,err := rlp.EncodeToBytes(signTx)
	if err!=nil {
		return "",fmt.Errorf("rlp 序列化失败！ %s",err.Error())
	}
	txHash :=""
	methodName := "eth_sendRawTransaction"
	err = r.client.client.Call(&txHash,methodName,common.BytesToHash(txRlpData))
	if err!=nil{
		return "", fmt.Errorf("发送交易失败！ %s",err.Error())
	}
	oldNonce :=r.NonceManager.GetNonce(address)
	if oldNonce== nil{
		r.NonceManager.SetNonce(address,new(big.Int).SetUint64(transaction.Nonce()))
	}
	r.NonceManager.PlusNonce(address)
	return txHash,nil
}

func (r *ETHRPCRequester) GetNonce(address string)(uint64,error){
	methodName := "eth_getTransactionCount"
	nonce :=""

	err := r.client.client.Call(&nonce,methodName,address,"pending")
	if err!= nil {
		return 0,fmt.Errorf("发送交易失败！ %s",err.Error())
	}
	n,_ :=new(big.Int).SetString(nonce[2:],16)
	return n.Uint64(),nil
}


func (r *ETHRPCRequester) SendETHTrasaction(fromStr, toStr, valueStr string,gasLimit,gasPrice uint64)(string, error) {
	if !common.IsHexAddress(fromStr)|| !common.IsHexAddress(toStr){
		return "", errors.New("invalid addresss")
	}

	to :=common.HexToxAddress(toStr)
	gasPrice_ := new(big.Int).SetUint64(gasPrice)

	realV :=tool.GetRealDecimalValue(valueStr,18)
	nonce :=r.NonceManager.GetNonce(fromStr)
	if realV == ""{
		return "", errors.New("invalid value")
	}
	amount,_ :=new(big.Int).SetString(realV,10)
	if nonce ==nil{
		n,err :=r.GetNonce(fromStr)
		if err!=nil {
			return "",fmt.Errorf("获取 nonce失败 %s", err.Error())
		}
		nonce =new(big.Int).SetUint64(n)
		r.NonceManager.SetNonce(fromStr,nonce)
	}
	data :=[]byte("")
	transaction :=types.NewTransaction(
		nonce.Uint64(),//unit64
		to,//Hex
		amount,//new(big.Int)
		gasLimit,//uint64
		gasPrice_,//uint64
		data,//string
		)
	return r.SendTransaction(fromStr,transaction)
}


