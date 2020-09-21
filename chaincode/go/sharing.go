package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct{}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) pb.Response {
	function, args := APIstub.GetFunctionAndParameters()

	if function == "checkUser" {
		return s.checkUser(APIstub, args)
	}
	/*else if function == "voting" {
		return s.voting(APIstub, args)
	}
	else if function == "getVote" {
		return s.getVote(APIstub, args)
	}
	else if function == "isEnable" {
		return s.isEnable(APIstub, args)
	}
	else if function == "setVote" {
		return s.setVote(APIstub, args)
	}
	else if function == "getResult" {
		return s.getResult(APIstub, args)
	}
	else if function == "getAllVotes" {
		return s.getAllVotes(APIstub)
	}*/

	fmt.Println("Please check your function : " + function)
	return shim.Error("Unknown function")
}

func (s *SmartContract) checkUser(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. [Id1] [Id2]")
	}

	DataAsBytes, err := APIstub.GetState(args[0] + args[1])
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if DataAsBytes == nil {
		return shim.Error("Not Found")
	}

	return shim.Success([]byte("i don't know"))
}

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
