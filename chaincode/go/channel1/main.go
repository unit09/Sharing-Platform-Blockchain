package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct{}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) pb.Response {
	function, args := APIstub.GetFunctionAndParameters()

	MSPid, err := cid.GetMSPID(APIstub)
	if err != nil {
		return shim.Error(err.Error())
	}

	if function == "getAllShareRecord" {

		return s.getAllShareRecord(APIstub)
	} else if function == "getShareStart" {

		return s.getShareStart(APIstub)
	} else if function == "getShareEnd" {

		return s.getShareEnd(APIstub)
	} else if function == "startShare" {

		if MSPid == "Share1Org" {
			return s.startShare(APIstub, args)
		}

		return shim.Error("Access Denied : " + MSPid)
	} else if function == "endShare" {

		if MSPid == "Share1Org" {
			return s.endShare(APIstub, args)
		}

		return shim.Error("Access Denied : " + MSPid)
	} else if function == "setPlace" {

		if MSPid == "AgencyOrg" {
			return s.setPlace(APIstub, args)
		}

		return shim.Error("Access Denied : " + MSPid)
	} else if function == "getShareRecordByLocation" {

		return s.getShareRecordByLocation(APIstub, args)
	} else if function == "countAllShareRecordByLocation" {

		return s.countAllShareRecordByLocation(APIstub)
	} else if function == "getAllPlace" {

		return s.getAllPlace(APIstub)
	} else if function == "getUserShareRecord" {

		if MSPid == "Share1Org" || MSPid == "MonitorOrg" {
			return s.getUserShareRecord(APIstub, args)
		}

		return shim.Error("Access Denied : " + MSPid)
	}

	fmt.Println("Please check your function : " + function)
	return shim.Error("Unknown function")
}

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
