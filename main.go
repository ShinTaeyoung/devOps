package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	//"fabric-dev/chaincode/controllableCoinDemo/constants"
	//utils "github.com/chaincode/setlTxnChainCode/utils"
	//controllers "./controllers"
	controllers "fabric/membershipService/controllers"
	//models "github.com/chaincode/setlTxnChainCode/models"

	//"strconv"
)

type SimpleChaincode struct { 
}


func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	
	return shim.Success(nil)

}



func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	//test
	function, args := stub.GetFunctionAndParameters()
	vc := controllers.ValidC{}
	
	if function == "RegMembership" {
		return vc.SetMembership(stub, args)
	} else if function == "RetvMembership" {
		return vc.GetMembership(stub, args)
	} else if function == "ModifyMembership" {
		return vc.UpdateMembership(stub, args)
	}
	//return shim.Success(utils.GenerateResponseError(constants.CODE_NOT_FOUND, constants.ERROR_INVALID_FUNCTION_NAME))
	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf(err.Error())
	}
}
