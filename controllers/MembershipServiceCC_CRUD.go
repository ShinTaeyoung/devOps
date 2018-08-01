package controllers

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	utils "fabric/membershipService/utils"
	//controllers "github.com/chaincode/setlTxnChainCode/controllers"
	models "fabric/membershipService/models"
	"encoding/json"
	"fmt"
	"strconv"
)

type ValidC struct {
	STY utils.QueryState
}

func (t *ValidC) SetMembership(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	/*
	if err := utils.CheckNull(args, 1); err != nil {
		return shim.Success(utils.GenerateResponseError(400, err.Error()))
	}
	*/
	fmt.Println("start SetMembership")
	fmt.Println("args[0] : " + args[0])
	var request map[string]interface{}
	err := json.Unmarshal([]byte(args[0]), &request)

	if err != nil {
		return shim.Success(utils.GenerateResponseError(400, err.Error()))
	}

	certId := fmt.Sprint(request["certId"])
	userId := fmt.Sprint(request["userId"])
	userName := fmt.Sprint(request["userName"])
	ageStr := fmt.Sprint(request["age"])
	gender := fmt.Sprint(request["gender"])
	membershipPointStr := fmt.Sprint(request["membershipPoint"])
	age, err := strconv.ParseInt(ageStr, 10, 64)
	membershipPoint, err := strconv.ParseInt(membershipPointStr, 10, 64)


	fmt.Println("certId : " + certId)

/*
	certId := args[0]
	setlYn := args[1]
	assetType := args[2]
	amount, err := strconv.ParseInt(args[3], 10, 64)
	pblsCorp := args[4]
	from := args[5]
	to := args[6]
	pblsAmount, err := strconv.ParseInt(args[7], 10, 64)
    docType := "docTypeSetl"
*/

	if err != nil {
		return shim.Success(nil)
		}

	newMembershipService := models.NewMembershipService(certId, userId, userName, age, gender, membershipPoint)
	utils.Save(stub, certId, newMembershipService)
	
	jsonBytes, err:= json.Marshal(newMembershipService)
    if err !=nil {
		return shim.Error(err.Error())
	}
	
//return shim.Success(jsonBytes)
return shim.Success(utils.GenerateResponseBytes(200, "", jsonBytes, stub.GetTxID()))

}

func (t *ValidC) GetMembership(stub shim.ChaincodeStubInterface, args []string) pb.Response {
 
	var request map[string]interface{}
	err := json.Unmarshal([]byte(args[0]), &request)

	if err != nil {
		return shim.Success(utils.GenerateResponseError(400, err.Error()))
	}

    certId := fmt.Sprint(request["certId"])

 /*	
  certId := args[0]
  */
  membershipService, err := utils.Load(stub, certId)

  if err != nil {
	return shim.Error(err.Error())
}
return shim.Success(utils.GenerateResponseBytes(200, "", membershipService, stub.GetTxID()))

}


func (t *ValidC) UpdateMembership(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var request map[string]interface{}
	err := json.Unmarshal([]byte(args[0]), &request)

	if err != nil {
		return shim.Success(utils.GenerateResponseError(400, err.Error()))
	}

	certId := fmt.Sprint(request["certId"])
	userName := fmt.Sprint(request["userName"])

	/*
	certId := args[0]
	setYn := args[1]
	*/
	membershipServ, err := utils.Load(stub, certId)

	membershipService := models.MembershipService{}
	json.Unmarshal(membershipServ,&membershipService)

	membershipService.UserName = userName

	utils.Save(stub, certId, membershipService)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}