package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type ShareRecordKey struct {
	Key   string
	Index int
}

func generateKey(APIstub shim.ChaincodeStubInterface, key string) []byte {

	lastkeyAsBytes, err := APIstub.GetState(key)
	if err != nil {
		fmt.Println("Failed to get state")
	}

	lastkey := ShareRecordKey{}
	json.Unmarshal(lastkeyAsBytes, &lastkey)

	if len(lastkey.Key) == 0 || lastkey.Key == "" {
		lastkey.Key = "SR"
		lastkey.Index = 0
	} else {
		lastkey.Index++
	}

	returnValueBytes, _ := json.Marshal(lastkey)

	return returnValueBytes
}

func (s *SmartContract) startShare(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments.")
	}

	shareKey := ShareRecordKey{}
	shareKeyAsBytes := generateKey(APIstub, "lastKey")
	json.Unmarshal(shareKeyAsBytes, &shareKey)

	now := time.Now().Format("2006-01-02 15:04:05")

	share := ShareRecord{ID: args[0], Timestamp: now, Target: args[1], Type: "start", Location: args[2]}
	shareAsBytes, _ := json.Marshal(share)
	keyStr := shareKey.Key + strconv.Itoa(shareKey.Index)

	APIstub.PutState(keyStr, shareAsBytes)
	APIstub.PutState("lastKey", shareKeyAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) endShare(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments.")
	}

	shareKey := ShareRecordKey{}
	shareKeyAsBytes := generateKey(APIstub, "lastKey")
	json.Unmarshal(shareKeyAsBytes, &shareKey)

	now := time.Now().Format("2006-01-02 15:04:05")

	share := ShareRecord{ID: args[0], Timestamp: now, Target: args[1], Type: "end", Location: args[2]}
	shareAsBytes, _ := json.Marshal(share)
	keyStr := shareKey.Key + strconv.Itoa(shareKey.Index)

	APIstub.PutState(keyStr, shareAsBytes)
	APIstub.PutState("lastKey", shareKeyAsBytes)

	return shim.Success(nil)
}
