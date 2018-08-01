package models

//import "fabric-dev/chaincode/controllableCoinDemo/constants"

type MembershipService struct {
	CertId    		   string     `json:"certId"`
	UserId    		   string     `json:"userId"`
	UserName  		   string     `json:"userName"`
	Age  	  		   int64      `json:"age,string"`
	Gender    		   string     `json:"gender"`
	MembershipPoint    int64      `json:"membershipPoint,string"`
}

func NewMembershipService(certId string, userId string, userName string, age int64, gender string, membershipPoint int64) MembershipService {
	membershipService := MembershipService{}
	membershipService.CertId = certId
	membershipService.UserId = userId
	membershipService.UserName = userName
	membershipService.Age = age
	membershipService.Gender = gender
	membershipService.MembershipPoint = membershipPoint

	return membershipService
}