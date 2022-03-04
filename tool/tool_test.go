package tool

import (
	json2 "encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"testing"
)

func Test_UnlockETHWallet(t *testing.T){
	address :="0xdcd2af33d7f00b0e4c84b15943b7d38672e73465"
	keysDir :="../keystores"
	err1:=UnlockETHWallet(keysDir,address,"789")
	if err1 !=nil{
		fmt.Println("第一次解锁错误：",err1.Error())
	}else{
		fmt.Println("第一次解锁成功！")
	}
	err2 := UnlockETHWallet(keysDir,address,"123456a")
	if err2 !=nil{
		fmt.Println("第二次解锁错误：",err1.Error())
	}else{
		fmt.Println("第二次解锁成功！")
	}

	tx := types.NewTransaction(
		123,
		common.Address{},
		new(big.Int).SetInt64(10),
		1000,
		new(big.Int).SetInt64(20),
		[]byte("交易"))
	signTx,err := SignETHTransaction(address,tx)
	if err!= nil{
		fmt.Println("签名失败！",err.Error())
		return
	}
	data,_ :=json2.Marshal(signTx)
	fmt.Println("签名成功\n",string(data))
}

