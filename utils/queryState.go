package utils

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
	//"fabric-dev/chaincode/controllableCoinDemo/constants"
	"errors"
	"strconv"
	"fmt"
)

type QueryState struct {
	Stub shim.ChaincodeStubInterface
}

// func NewQueryState(stub shim.ChaincodeStubInterface) *QueryState {
// 	q := new(QueryState)
// 	q.Stub = stub
// 	return q
// }

func (q *QueryState) QueryStateUsingPage(stub shim.ChaincodeStubInterface, queryString string, selectCount int, selectPage int) ([]byte, error) {


	fmt.Println(queryString)


	queryResultIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		//return nil, err
		fmt.Println(err.Error())
		return nil, errors.New("getQueryResultError!!!")
	}

	var results []map[string]interface{}
	for queryResultIterator.HasNext() {
		queryResponse, err := queryResultIterator.Next()
		if err != nil {
			//return nil, err
			return nil, errors.New("queryResultIterator")
		}
		var result map[string]interface{}
		if queryResponse.Value != nil {
			json.Unmarshal(queryResponse.Value, &result)
			results = append(results, result)
		}
	}
	if len(results) == 0 {
		return nil, errors.New("ERROR_EMPTY_STATE")
	}
	page := selectPage -1
	start := selectCount * page
	end := start + selectCount

	if len(results) < start {
		return nil, errors.New("ERROR_DATA_LENGTH" + " :" + strconv.Itoa(len(results)))
	}

	if len(results) < end {
		end = len(results)
	}
	resultsSlice := results[start:end]
	response := map[string]interface{}{
		"totalCount":  strconv.Itoa(len(results)),
		"pageCount":   strconv.Itoa(selectPage),
		"selectCount": strconv.Itoa(selectCount),
		"values":       ArrayToMap(resultsSlice),
	}

	return json.Marshal(response)
}