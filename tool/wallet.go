package tool

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var ETHUnlockMap map[string]accounts.Account

var UnlocKs *keystore.KeyStore

func UnlockETHWallet(keysDir string,address,password string) error{
	if UnlocKs==nil{
		UnlocKs =keystore.NewKeyStore(
			keysDir,
			keystore.StandardScryptN,
			keystore.StandardScryptP)
		if  UnlocKs ==nil{
			return errors.New("ks is nil")
		}
	}
	unlock := accounts.Account{Address: common.HexToAddress(address)}

	if err:= UnlocKs.Unlock(unlock,password);nil !=err{
		return errors.New("unlock err: " + err.Error())
	}
	if ETHUnlockMap==nil {
		ETHUnlockMap =map[string]accounts.Account{}
	}
	ETHUnlockMap[address] = unlock
	return  nil
}

func SignETHTransaction(addres string,transaction *types.Transaction)(*types.Transaction, error){
	if UnlocKs ==nil{
		return nil,errors.New("you need to init keystore first")
	}
	account :=ETHUnlockMap[addres]
	if !common.IsHexAddress(account.Address.String()) {
		return nil,errors.New("account need to unlock first!")
	}
	return  UnlocKs.SignTx(account,transaction,nil)
}

