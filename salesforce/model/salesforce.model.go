package salesforcemodel

import "github.com/nimajalali/go-force/sobjects"

type PBAListingSObject struct {
	sobjects.BaseSObject

	Id                     string  `force:"Id"`
	RegionID               string  `force:"pba__Area_pb__c"`
	CommercialResidential  string  `force:"Development_Type__c"`
	StreetNumber           int     `force:"House_Number__c"`
	StreetName             string  `force:"pba__Street_pb__c"`
	ApartmentNumber        int     `force:"Lot_Number__c"`
	PostalCode             string  `force:"pba__PostalCode_pb__c"`
	CityID                 string  `force:"pba__City_pb__c"`
	Latitude               string  `force:"pba__Latitude_pb__c"`
	Longitude              string  `force:"pba__Longitude_pb__c"`
	CurrentListingPrice    float64 `force:"pba__ListingPrice_pb__c"`
	CurrentListingCurrency string  `force:"Common_Charge_Currency__c"`
	RentalCommissionMonths float64 `force:"MonthlyRent__c"`
	SoldPrice              float64 `force:"pba__SoldPrice__c"`
	SoldDate               string  `force:"pba__SoldDate__c"`
	TransactionType        string  `force:"pba__ListingType__c"`
	PropertyType           string  `force:"pba__PropertyType__c"`
	PropertyCategory       string  `force:"Zone_Category__c"`
	ListingStatus          string  `force:"pba__Status__c"`
	TotalArea              float64 `force:"pba__TotalArea_pb__c"`
	LotSizeFirstPart       float64 `force:"pba__LotSize_pb__c"`
	LotSizeSecondPart      string  `force:"Lot_Size_Unit__c"`
	ParkingSpaces          float64 `force:"pba__NumberOfParkingSpaces__c"`
	TotalNumOfRooms        float64 `force:"Rooms__c"`
	NumberOfBathrooms      float64 `force:"pba__FullBathrooms_pb__c"`
	NumberOfBedrooms       float64 `force:"pba__Bedrooms_pb__c"`
	NumberOfToiletRooms    float64 `force:"Toilettes__c"`
	NumberOfFloors         float64 `force:"pba__NumberOfFloors__c"`
	YearBuild              float64 `force:"pba__YearBuilt_pb__c"`
	DescriptionText        string  `force:"pba__Description_pb__c"`
	RatingValue            string  `force:"Energy_Efficiency_Rating_by_Expert__c"`
	HeatingCost            int     `force:"Heating_Charges__c"`
	AlternateURL           string  `force:"Listing_URL__c"`
	//PriceType              int    `force:""` //TODO !!!!!! It is not CurrencyIsoCode
	//FloorLevel             int    `force:"Floor_Picklist__c"`//TODO !!!!!! Check
}

func (t *PBAListingSObject) ApiName() string {
	return "pba__Listing__c"
}

type PBAListingSObjectQueryResponse struct {
	sobjects.BaseQuery

	Records []*PBAListingSObject `force:"records"`
}

type PropertyMediaSObject struct {
	sobjects.BaseSObject

	SequenceNumber float64 `force:"pba__SortOnExpose__c"` // TODO to PropertyMedia
	FileName       string  `force:"pba__Filename__c"`     // TODO to PropertyMedia
	//DescriptiveName string  `force:"pba__Description__c"`  // TODO to PropertyMedia
	URL string `force:"pba__URL__c"` // TODO to PropertyMedia
}

type PropertyMediaSObjectQueryResponse struct {
	sobjects.BaseQuery

	Records []*PropertyMediaSObject `force:"records"`
}
