package main

import (
	"it-chain/domain"
	"it-chain/db/leveldbhelper"
	"encoding/json"
	"fmt"
	"os"
)

type SampleSmartContract struct {
}

func (sc *SampleSmartContract) Init(args []string) error {
	fmt.Println("in Init func")
	tx := domain.Transaction{}
	err := json.Unmarshal([]byte(args[0]), &tx)
	if err != nil {
		return err
	}

	/* Init WorldStateDB
	---------------------*/
	path := "/go/src/worldstatedb"
	dbProvider := leveldbhelper.CreateNewDBProvider(path)
	defer func(){
		dbProvider.Close()
	}()

	wsDB := "worldStateDB"
	wsDBHandle := dbProvider.GetDBHandle(wsDB)
	//wsDBHandle.Put([]byte("test"), []byte("value1"), true)

	if tx.TxData.Method == "Query" {
		sc.Query(tx, wsDBHandle)
	} else if tx.TxData.Method == "Invoke" {
		sc.Invoke(tx, wsDBHandle)
	}

	return nil
}

func (sc *SampleSmartContract) Query(transaction domain.Transaction, wsDBHandle *leveldbhelper.DBHandle) {
	fmt.Println("func Query")
}

func (sc *SampleSmartContract) Invoke(transaction domain.Transaction, wsDBHandle *leveldbhelper.DBHandle) {
	wsDBHandle.Put([]byte("test"), []byte("success"), true)
	fmt.Println("func Invoke")
	test, _ := wsDBHandle.Get([]byte("test"))
	fmt.Println("test : " + string(test))
}

func main()  {
	fmt.Println("func main")
	args := os.Args[1:]
	ssc := new(SampleSmartContract)

	ssc.Init(args)
}