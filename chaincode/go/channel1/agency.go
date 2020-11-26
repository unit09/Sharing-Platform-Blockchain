package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type PlaceRecord struct {
	ID        string `json:"id"`
	Location  string `json:"location"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

type PlaceRecordKey struct {
	Key   string
	Index int
}

func generatePlaceKey(APIstub shim.ChaincodeStubInterface, key string) []byte {

	lastkeyAsBytes, err := APIstub.GetState(key)
	if err != nil {
		fmt.Println("Failed to get state")
	}

	lastkey := PlaceRecordKey{}
	json.Unmarshal(lastkeyAsBytes, &lastkey)

	if len(lastkey.Key) == 0 || lastkey.Key == "" {
		lastkey.Key = "PR"
		lastkey.Index = 0
	} else {
		lastkey.Index++
	}

	returnValueBytes, _ := json.Marshal(lastkey)

	return returnValueBytes
}

func (s *SmartContract) setPlace(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments.")
	}

	placeKey := PlaceRecordKey{}
	placeKeyAsBytes := generatePlaceKey(APIstub, "lastPlaceKey")
	json.Unmarshal(placeKeyAsBytes, &placeKey)

	place := PlaceRecord{ID: args[0], Location: args[1], Longitude: args[2], Latitude: args[3]}
	placeAsBytes, _ := json.Marshal(place)
	keyStr := placeKey.Key + strconv.Itoa(placeKey.Index)

	APIstub.PutState(keyStr, placeAsBytes)
	APIstub.PutState("lastPlaceKey", placeKeyAsBytes)

	return shim.Success(nil)
}
