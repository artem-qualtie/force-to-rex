package salesforce

import (
	"fmt"
	"force-to-rex/conf"
	"force-to-rex/salesforce/model"
	"github.com/nimajalali/go-force/force"
)

func GetForceConnection() (*force.ForceApi, error) {
	forceApi, err := force.Create(
		conf.Config.Str("apiVersion"),
		conf.Config.Str("clientID"),
		conf.Config.Str("clientSecret"),
		conf.Config.Str("username"),
		conf.Config.Str("password"),
		conf.Config.Str("securityToken"),
		conf.Config.Str("environment"),
	)

	if err != nil {
		fmt.Println("Create Error: ", err)
		return nil, err
	}
	return forceApi, nil
}

func GetTestData() (*force.ForceApi, *salesforcemodel.PBAListingSObjectQueryResponse, error) {
	forceApi, err := force.Create(
		conf.Config.Str("apiVersion"),
		conf.Config.Str("clientID"),
		conf.Config.Str("clientSecret"),
		conf.Config.Str("username"),
		conf.Config.Str("password"),
		conf.Config.Str("securityToken"),
		conf.Config.Str("environment"),
	)

	if err != nil {
		fmt.Println("Create Error: ", err)
		return nil, nil, err
	}
	//pba__SoldDate__c
	//err = forceApi.Query("SELECT Id, pba__Area_pb__c, Development_Type__c, House_Number__c, pba__Street_pb__c, Lot_Number__c, pba__PostalCode_pb__c, pba__City_pb__c, pba__Latitude_pb__c, pba__Longitude_pb__c, pba__ListingPrice_pb__c, Common_Charge_Currency__c, MonthlyRent__c, pba__SoldPrice__c, pba__SoldDate__c, pba__ListingType__c, pba__PropertyType__c, Zone_Category__c, pba__Status__c, pba__TotalArea_pb__c, pba__LotSize_pb__c, Lot_Size_Unit__c, pba__NumberOfParkingSpaces__c, Rooms__c, pba__FullBathrooms_pb__c, pba__Bedrooms_pb__c, Toilettes__c, pba__NumberOfFloors__c, pba__YearBuilt_pb__c, pba__Description_pb__c, Energy_Efficiency_Rating_by_Expert__c, Heating_Charges__c FROM pba__Listing__c LIMIT 10", pbaListingSObjects)
	//    err = forceApi.Query("SELECT pba__Description_pb__c, Energy_Efficiency_Rating_by_Expert__c, Heating_Charges__c FROM pba__Listing__c LIMIT 10", pbaListingSObjects)
	pbaListingSObjects := &salesforcemodel.PBAListingSObjectQueryResponse{}
	err = forceApi.QueryAll("SELECT Id, pba__Area_pb__c, Development_Type__c, House_Number__c, pba__Street_pb__c, Lot_Number__c, pba__PostalCode_pb__c, pba__City_pb__c, pba__Latitude_pb__c, pba__Longitude_pb__c, pba__ListingPrice_pb__c, Common_Charge_Currency__c, MonthlyRent__c, pba__SoldPrice__c, pba__SoldDate__c, pba__ListingType__c, pba__PropertyType__c, Zone_Category__c, pba__Status__c, pba__TotalArea_pb__c, pba__LotSize_pb__c, Lot_Size_Unit__c, pba__NumberOfParkingSpaces__c, Rooms__c, pba__FullBathrooms_pb__c, pba__Bedrooms_pb__c, Toilettes__c, pba__NumberOfFloors__c, pba__YearBuilt_pb__c, pba__Description_pb__c, Energy_Efficiency_Rating_by_Expert__c, Heating_Charges__c, LastModifiedDate FROM pba__Listing__c", pbaListingSObjects)
	if err != nil {
		fmt.Println("Select Error: ", err)
		return forceApi, nil, err
	}

	////fmt.Println("%#v", pbaListingSObjects)
	//for _, so := range pbaListingSObjects.Records {
	//	//fmt.Println("%v", so)
	//	b, err := json.Marshal(so)
	//	if err != nil {
	//		fmt.Println("error: ", err)
	//	}
	//	fmt.Println(string(b))
	//}
	return forceApi, pbaListingSObjects, nil
}

func GetPropertyImages(forceApi *force.ForceApi, propertyID string) (*salesforcemodel.PropertyMediaSObjectQueryResponse, error) {
	propertyMediaSObjects := &salesforcemodel.PropertyMediaSObjectQueryResponse{}
	err := forceApi.Query(`SELECT Id, pba__SortOnExpose__c, pba__Filename__c, pba__Description__c, pba__URL__c FROM pba__PropertyMedia__c WHERE pba__URL__c='`+propertyID+`'`, propertyMediaSObjects)
	if err != nil {
		fmt.Println("Select Error: ", err)
		return nil, err
	}

	return propertyMediaSObjects, nil
}
