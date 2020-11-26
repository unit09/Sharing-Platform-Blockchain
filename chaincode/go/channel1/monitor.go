package main

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (s *SmartContract) getUserShareRecord(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
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

	var idx int
	idx = 0

	for resultsIter.HasNext() {
		if idx >= len(args) {
			break
		}

		query, err := resultsIter.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		item := ShareRecord{}
		json.Unmarshal(query.Value, &item)

		if item.ID == args[idx] {
			idx++

			if isWritten == true {
				buffer.WriteString(",")
			}

			/*buffer.WriteString("{\"Key\":")
			buffer.WriteString("\"")
			buffer.WriteString(query.Key)
			buffer.WriteString("\"")

			buffer.WriteString("{\"Record\":")*/

			buffer.WriteString(string(query.Value))
			//buffer.WriteString("}")
			isWritten = true
		}
	}
	buffer.WriteString("]")

	return shim.Success(buffer.Bytes())
}
