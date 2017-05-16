/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strings"
	"encoding/json"
)

// NumberPortabilityChaincode is a Smart Contract between CSPs for porting In or Porting Out Customers and settling the billing across the CSPs


type NumberPortabilityChaincode struct {
}

type EligibilityConfirm struct {

		Number string
		ServiceProviderOld string
		ServiceProviderNew string
		CustomerName string 
	    SSNNumber string
	    PortabilityIndicator string
		status string
}



type UserAcceptance struct {

		Number string
		ServiceProviderOld string
		PlanOld string 
	    ServiceValidityOld string
	    TalktimeBalanceOld string
		SMSbalanceOld string
		DataBalanceOld string
		ServiceProviderNew string
		PlanNew string 
	    ServiceValidityNew string
	    TalktimeBalanceNew string
		SMSbalanceNew string
		DataBalanceNew string
		CustomerAcceptance string
		status string
		
}

type Reserve struct {
	TollFreeno string;
	status string;
	
}


type UsageDetailsFromDonorandAcceptorCSP struct {

		Number string
		ServiceProviderOld string
		PlanOld string 
	    ServiceValidityOld string
	    TalktimeBalanceOld string
		SMSbalanceOld string
		DataBalanceOld string
		ServiceProviderNew string
		PlanNew string 
	    ServiceValidityNew string
	    TalktimeBalanceNew string
		SMSbalanceNew string
		DataBalanceNew string
		status string
		
}


type UsageDetailsFromCSP struct {

		Number string
		ServiceProviderOld string
		ServiceProviderNew string
		Plan string 
	    ServiceValidity string
	    TalktimeBalance string
		SMSbalance string
		DataBalance string
		status string
}


// Init method will be called during deployment.

func (t *NumberPortabilityChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Init Chaincode...")
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	fmt.Println("Init Chaincode...done")

	return nil, nil
}




// CGTA invoke function


// UserAcceptance Invoke function

func (t *NumberPortabilityChaincode) UserAcceptance(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

    fmt.Println("UserAcceptance Information invoke Begins...")

	if len(args) != 14 {
		return nil, errors.New("Incorrect number of arguments. Expecting 14")
	}

	// Check the User Acceptance paramater, if true then update world state with new status
	
	var status1 string
	key := args[0]+args[1]+args[7]
	Acceptance := args[13]
	if(Acceptance == "true"){
	 status1 = "CustomerAccepted"
	} else {
	  status1 = "CustomerRejected"
	}
	
	UserAcceptanceObj := UserAcceptance{Number: args[0], ServiceProviderOld: args[1], PlanOld: args[2], ServiceValidityOld: args[3], TalktimeBalanceOld: args[4], SMSbalanceOld: args[5], DataBalanceOld: args[6], ServiceProviderNew: args[7], PlanNew: args[8], ServiceValidityNew: args[9], TalktimeBalanceNew: args[10], SMSbalanceNew: args[11], DataBalanceNew: args[12], CustomerAcceptance: args[13], status: status1}
    fmt.Println("UserAcceptance Details Structure ",UserAcceptanceObj)
	err := stub.PutState(key,[]byte(fmt.Sprintf("%s",UserAcceptanceObj)))
	if err != nil {
		return nil, err
	}
	
	fmt.Println("UserAcceptance Information invoke ends...")
	return nil, nil
}
//Reserve Invoke function
func (t *NumberPortabilityChaincode) Reserve(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

    fmt.Println("Reserve Information invoke Begins...")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}

	// Check the Reseve paramater, if true then update world state with new status
	
	var status1 string
	key := args[0]+args[1]
	Acceptance := args[1]
	if(Acceptance == "true"){
	 status1 = "RequestInitiated"
	} 
	
	ReserveObj := Reserve{TollFreeno: args[0], status: status1}
    fmt.Println("Reserve Details Structure ",ReserveObj)
	value, e := json.Marshal(ReserveObj)
	if e != nil {
		return nil, e
	}
	err := stub.PutState(key,[]byte(fmt.Sprintf("%s",value)))
	if err != nil {
		return nil, err
	}
	
	fmt.Println("Reserve Information invoke ends...")
	return nil, nil
}


// FinalPortInfo Invoke function


func (t *NumberPortabilityChaincode) RegulatorQuery(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var key, jsonResp string
    var err error

    if len(args) != 3 {
        return nil, errors.New("Incorrect number of arguments. Expecting 3 arguments")
    }

    key = args[0]+args[1]+args[2]
    valAsbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    } else if len(valAsbytes) == 0{
	    jsonResp = "{\"Error\":\"Failed to get Query for " + key + "\"}"
        return nil, errors.New(jsonResp)
	}

	fmt.Println("Query NumberPortability Chaincode... end") 
    return valAsbytes, nil 

}


func (t *NumberPortabilityChaincode) RegulatorQuery1(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var key, jsonResp string
    var err error

    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting 1 arguments")
    }

    key = args[0]+args[1]
    valAsbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    } else if len(valAsbytes) == 0{
	    jsonResp = "{\"Error\":\"Failed to get Query for " + key + "\"}"
        return nil, errors.New(jsonResp)
	}

	fmt.Println("Query NumberPortability Chaincode... end") 
    return valAsbytes, nil 

}




// Invoke Function

func (t *NumberPortabilityChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
      
	 fmt.Println("Invoke NumberPortability Chaincode... start") 

	
	// Handle different functions UserAcceptance
	if function == "EligibilityConfirm" {
		return t.EligibilityConfirm (stub, args)
	} else if function == "UsageDetailsFromDonorCSP" {
		return t.UsageDetailsFromDonorCSP(stub, args)
	}else if function == "EntitlementFromRecipientCSP" {
		return t.EntitlementFromRecipientCSP(stub, args)
	}else if function == "UserAcceptance" {
		return t.UserAcceptance(stub, args)
	}else if function == "ConfirmationOfMNPRequest" {
		return t.ConfirmationOfMNPRequest(stub, args)
	} else{
	    return nil, errors.New("Invalid function name. Expecting 'EligibilityConfirm' or 'UsageDetailsFromDonorCSP' or 'EntitlementFromRecipientCSP' but found '" + function + "'")
	}
	
	
	fmt.Println("Invoke Numberportability Chaincode... end") 
	
	return nil,nil;
}




// Query to get CSP Service Details

func (t *NumberPortabilityChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Query NumberPortability Chaincode... start") 

	
	if function == "EntitlementFromRecipientCSPQuery" {
		return t.EntitlementFromRecipientCSPQuery(stub, args)
	} 
	
	if function == "RegulatorQuery" {
		return t.RegulatorQuery(stub, args)
	} 
	
	// else We can query WorldState to fetch value
	
	var key, jsonResp string
    var err error

    if len(args) < 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
    }
	fmt.Println(len(args))
	if len(args) == 3 {
	   key = args[0]+args[1]+args[2]
	} else if len(args) == 2 {
	   key = args[0]+args[1]
	} else {
	   key = args[0]
	}

    
    valAsbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    } else if len(valAsbytes) == 0{
	    jsonResp = "{\"Error\":\"Failed to get Query for " + key + "\"}"
        return nil, errors.New(jsonResp)
	}

	fmt.Println("Query NumberPoratbility Chaincode... end") 
    return valAsbytes, nil 
  
	
}



func main() {
	err := shim.Start(new(NumberPortabilityChaincode))
	if err != nil {
		fmt.Println("Error starting NumberPortabilityChaincode: %s", err)
	}
}
