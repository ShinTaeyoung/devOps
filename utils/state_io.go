package utils

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
	//"fabric-dev/chaincode/controllableCoinDemo/constants"
	"errors"
	"fmt"
	"strconv"
)


func Save(stub shim.ChaincodeStubInterface, key string, o interface{}) error {

	if out, err := json.Marshal(o); err != nil {
		return err
	} else {

		return stub.PutState(key, out)
	}
}
func Load(stub shim.ChaincodeStubInterface, key string) ([]byte,  error) {
	if data, err := stub.GetState(key); len(data) == 0 || err != nil {
		if err != nil {
			return nil, err
		} else {
			return nil, errors.New("ERROR_EMPTY_STATE")
		}
	} else {
		return data, err
	}
}

func CheckNull(args []string, length int) (error) {
	fmt.Println("len(args)!!")
	var lengthArray string
	lengthArray = strconv.Itoa(len(args))
	fmt.Println("!!"+lengthArray)
	if len(args) != length {
		return fmt.Errorf("ERROR_INVALID_ARGUMENT_NUM_WHYWHYWHY")
	}
	for _, val := range args {
		if len(val) < 0 {
			return fmt.Errorf("ERROR_INVALID_FORMAT")
		}
	}

	return nil
}