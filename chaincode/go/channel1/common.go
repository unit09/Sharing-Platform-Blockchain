package main

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type ShareRecord struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Target    string `json:"target"`
	Type      string `json:"type"`
	Location  string `json:"location"`
}

func (s *SmartContract) getAllShareRecord(APIstub shim.ChaincodeStubInterface) pb.Response {

	lastkeyAsBytes, err := APIstub.GetState("lastKey")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if lastkeyAsBytes == nil {
		return shim.Error("No Data")
	}

	lastkey := ShareRecordKey{}
	json.Unmarshal(lastkeyAsBytes, &lastkey)

	startKey := "SR0"
	endKey := "SR" + strconv.Itoa(lastkey.Index+1)

	resultsIter, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIter.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")
	isWritten := false

	for resultsIter.HasNext() {
		query, err := resultsIter.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		if isWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(query.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")

		buffer.WriteString(string(query.Value))
		buffer.WriteString("}")
		isWritten = true
	}
	buffer.WriteString("]\n")

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) getShareStart(APIstub shim.ChaincodeStubInterface) pb.Response {

	lastkeyAsBytes, err := APIstub.GetState("lastKey")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if lastkeyAsBytes == nil {
		return shim.Error("No Data")
	}

	lastkey := ShareRecordKey{}
	json.Unmarshal(lastkeyAsBytes, &lastkey)

	startKey := "SR0"
	endKey := "SR" + strconv.Itoa(lastkey.Index+1)

	resultsIter, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIter.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")
	isWritten := false

	for resultsIter.HasNext() {
		query, err := resultsIter.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		item := ShareRecord{}
		json.Unmarshal(query.Value, &item)

		if item.Type == "start" {
			if isWritten == true {
				buffer.WriteString(",")
			}

			buffer.WriteString("{\"Key\":")
			buffer.WriteString("\"")
			buffer.WriteString(query.Key)
			buffer.WriteString("\"")

			buffer.WriteString(", \"Record\":")

			buffer.WriteString(string(query.Value))
			buffer.WriteString("}")
			isWritten = true
		}
	}
	buffer.WriteString("]\n")

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) getShareEnd(APIstub shim.ChaincodeStubInterface) pb.Response {

	lastkeyAsBytes, err := APIstub.GetState("lastKey")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if lastkeyAsBytes == nil {
		return shim.Error("No Data")
	}

	lastkey := ShareRecordKey{}
	json.Unmarshal(lastkeyAsBytes, &lastkey)

	startKey := "SR0"
	endKey := "SR" + strconv.Itoa(lastkey.Index+1)

	resultsIter, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIter.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")
	isWritten := false

	for resultsIter.HasNext() {
		query, err := resultsIter.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		item := ShareRecord{}
		json.Unmarshal(query.Value, &item)

		if item.Type == "end" {
			if isWritten == true {
				buffer.WriteString(",")
			}

			buffer.WriteString("{\"Key\":")
			buffer.WriteString("\"")
			buffer.WriteString(query.Key)
			buffer.WriteString("\"")

			buffer.WriteString(", \"Record\":")

			buffer.WriteString(string(query.Value))
			buffer.WriteString("}")
			isWritten = true
		}
	}
	buffer.WriteString("]\n")

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) getShareRecordByLocation(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments.")
	}

	lastkeyAsBytes, err := APIstub.GetState("lastKey")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if lastkeyAsBytes == nil {
		return shim.Error("No Data")
	}

	lastkey := ShareRecordKey{}
	json.Unmarshal(lastkeyAsBytes, &lastkey)

	startKey := "SR0"
	endKey := "SR" + strconv.Itoa(lastkey.Index+1)

	resultsIter, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIter.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")
	isWritten := false

	for resultsIter.HasNext() {
		query, err := resultsIter.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		item := ShareRecord{}
		json.Unmarshal(query.Value, &item)

		if item.Location == args[0] {
			if isWritten == true {
				buffer.WriteString(",")
			}

			buffer.WriteString("{\"Key\":")
			buffer.WriteString("\"")
			buffer.WriteString(query.Key)
			buffer.WriteString("\"")

			buffer.WriteString(", \"Record\":")

			buffer.WriteString(string(query.Value))
			buffer.WriteString("}")
			isWritten = true
		}
	}
	buffer.WriteString("]\n")

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) countAllShareRecordByLocation(APIstub shim.ChaincodeStubInterface) pb.Response {

	lastkeyAsBytes, err := APIstub.GetState("lastKey")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if lastkeyAsBytes == nil {
		return shim.Error("No Data")
	}

	lastkey := ShareRecordKey{}
	json.Unmarshal(lastkeyAsBytes, &lastkey)

	startKey := "SR0"
	endKey := "SR" + strconv.Itoa(lastkey.Index+1)

	resultsIter, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIter.Close()

	locationMap := make(map[string]int)

	for resultsIter.HasNext() {
		query, err := resultsIter.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		item := ShareRecord{}
		json.Unmarshal(query.Value, &item)

		tempCount := locationMap[item.Location]
		if tempCount == 0 {
			locationMap[item.Location] = 1
		} else {
			locationMap[item.Location] = tempCount + 1
		}
	}

	var buffer bytes.Buffer
	buffer.WriteString("[")
	isWritten := false

	for key, val := range locationMap {
		if isWritten == true {
			buffer.WriteString(",")
		}

		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Count\":")

		buffer.WriteString(strconv.Itoa(val))
		buffer.WriteString("}")
	}

	buffer.WriteString("]\n")

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) getAllPlace(APIstub shim.ChaincodeStubInterface) pb.Response {

	lastkeyAsBytes, err := APIstub.GetState("lastPlaceKey")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if lastkeyAsBytes == nil {
		return shim.Error("No Data")
	}

	lastkey := PlaceRecordKey{}
	json.Unmarshal(lastkeyAsBytes, &lastkey)

	startKey := "PR0"
	endKey := "PR" + strconv.Itoa(lastkey.Index+1)

	resultsIter, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIter.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")
	isWritten := false

	for resultsIter.HasNext() {
		query, err := resultsIter.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		if isWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(query.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")

		buffer.WriteString(string(query.Value))
		buffer.WriteString("}")
		isWritten = true
	}
	buffer.WriteString("]\n")

	return shim.Success(buffer.Bytes())
}
