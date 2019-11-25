package rexmodel

const (
	True  = "True"
	False = "False"
)

type JobControl struct {
	Xmlns                     string `xml:"xmlns:xsi,attr"`
	NoNamespaceSchemaLocation string `xml:"xsi:noNamespaceSchemaLocation,attr"`
	IntegratorID              int    `xml:"IntegratorID"`
	//OfficeFileName            string `xml:"OfficeFile,omitempty"`
	//AgentFileName             string `xml:"AgentFile,omitempty"`
	PropertyFileName string `xml:"PropertyFile,omitempty"`
}

type UploadXML struct {
	//Xmlns                     string          `xml:"xmlns:xsi,attr"`
	//NoNamespaceSchemaLocation string          `xml:"xsi:noNamespaceSchemaLocation,attr"`
	Version        Version         `xml:"Version"`
	PropertiesData *PropertiesData `xml:"Properties,omitempty"`
	//ImageRemovals             *struct{}       `xml:"ImageRemovals,omitempty"`
	PropertiesToDisable string `xml:"PropertiesToDisable,omitempty"`
}

type Version struct {
	Version string `xml:"Version,attr"`
}

type innerXML struct {
	S string `xml:",innerxml"`
}

type PropertiesData struct {
	Properties []Property `xml:"Property"`
}

type PropertyDescription struct {
	LanguageID      LanguageID `xml:"LanguageID"`
	DescriptionText string     `xml:"DescriptionText"` //40000
}

type LanguageID struct {
	LanguageID int `xml:"LanguageID,attr"`
}

type PropertyDescriptions struct {
	Descriptions []PropertyDescription `xml:"PropertyDescription"`
}

type Image struct {
	SequenceNumber int    `xml:"SequenceNumber"`
	FileName       string `xml:"FileName"`
	//ImageQualityType int `xml:"ImageQualityType,omitempty"`//TODO NEED CHECK VERSION 3.3 XSD
	//DescriptiveName int       `xml:"DescriptiveName,omitempty"`
	IsDefault IsDefault `xml:"IsDefault,omitempty"`
	//SequenceNumber int `xml:"SequenceNumber,omitempty"`//TODO NEED CHECK VERSION 3.3 XSD
}

type IsDefault struct {
	IsDefault string `xml:"IsDefault,attr"`
}

type Images struct {
	DefaultImageSequenceNumber int     `xml:"DefaultImageSequenceNumber"`
	Images                     []Image `xml:"Image"`
}

type RegionID struct {
	RegionID int `xml:"RegionID,attr"`
}

type CommercialResidential struct {
	CommercialResidential int `xml:"CommercialResidential,attr"`
}

type FloorLevel struct {
	FloorLevel int `xml:"FloorLevel,attr"`
}

type CurrentListingCurrency struct {
	CurrentListingCurrency string `xml:"CurrentListingCurrency,attr"`
}

type PriceType struct {
	PriceType int `xml:"PriceType,attr"`
}

type ShowAddressOnWeb struct {
	ShowAddressOnWeb string `xml:"ShowAddressOnWeb,attr"`
}

type TransactionType struct {
	TransactionType int `xml:"TransactionType,attr"`
}

type PropertyType struct {
	PropertyType int `xml:"PropertyType,attr"`
}

type PropertyCategory struct {
	PropertyCategory int `xml:"PropertyCategory,attr"`
}

type ListingStatus struct {
	ListingStatus int `xml:"ListingStatus,attr"`
}

type ContractType struct {
	ContractType int `xml:"ContractType,attr"`
}

type PropertyStatus struct {
	PropertyStatus int `xml:"PropertyStatus,attr"`
}

type Property struct {
	IntegratorID               int                   `xml:"IntegratorID"`               //TODO !!! Config
	IntegratorPropertyID       string                `xml:"IntegratorPropertyID"`       //TODO !!! 50
	IntegratorOfficeID         string                `xml:"IntegratorOfficeID"`         //TODO !!! Config 50
	IntegratorSalesAssociateID string                `xml:"IntegratorSalesAssociateID"` //TODO !!! Config 100
	RegionID                   RegionID              `xml:"RegionID"`                   //TODO !!!
	CommercialResidential      CommercialResidential `xml:"CommercialResidential"`      //TODO !!!
	StreetNumber               string                `xml:"StreetNumber,omitempty"`     //  20
	StreetName                 string                `xml:"StreetName,omitempty"`       //  50
	ApartmentNumber            string                `xml:"ApartmentNumber,omitempty"`  //  15
	AddressLine2               string                `xml:"AddressLine2,omitempty"`     // ?? 50
	PostalCode                 string                `xml:"PostalCode,omitempty"`       // 15
	CityID                     int                   `xml:"CityID"`                     //TODO !!!
	//FloorLevel                 *FloorLevel             `xml:"FloorLevel,omitempty"`             //TODO Enum FloorLevel
	ShowAddressOnWeb       ShowAddressOnWeb       `xml:"ShowAddressOnWeb"`
	CurrentListingPrice    float64                `xml:"CurrentListingPrice"`    //TODO !!!
	CurrentListingCurrency CurrentListingCurrency `xml:"CurrentListingCurrency"` //TODO !!! Enum CurrencyCode
	PriceType              *PriceType             `xml:"PriceType,omitempty"`    //TODO Enum PriceType
	SoldPrice              float64                `xml:"SoldPrice,omitempty"`
	SoldDate               string                 `xml:"SoldDate"`                   //TODO ,omitempty
	TransactionType        TransactionType        `xml:"TransactionType"`            //TODO !!! Enum TransactionType
	ContractType           ContractType           `xml:"ContractType"`               //??TODO !!! Enum ContractType
	PropertyStatus         PropertyStatus         `xml:"PropertyStatus,omitempty"`   //??TODO !!! Enum PropertyStatus
	PropertyType           PropertyType           `xml:"PropertyType"`               //TODO !!! Enum PropertyType / CommPropertyType
	PropertyCategory       *PropertyCategory      `xml:"PropertyCategory,omitempty"` //TODO Enum PropertyCategory
	ListingStatus          ListingStatus          `xml:"ListingStatus"`              //TODO !!! Enum ListingStatus
	TotalArea              float64                `xml:"TotalArea,omitempty"`
	LotSize                string                 `xml:"LotSize,omitempty"` // 50
	LotSizeM2              float64                `xml:"LotSizeM2,omitempty"`
	TotalNumOfRooms        float64                `xml:"TotalNumOfRooms,omitempty"`
	NumberOfBathrooms      int                    `xml:"NumberOfBathrooms,omitempty"`
	NumberOfBedrooms       int                    `xml:"NumberOfBedrooms,omitempty"`
	NumberOfToiletRooms    int                    `xml:"NumberOfToiletRooms,omitempty"`
	NumberOfFloors         int                    `xml:"NumberOfFloors,omitempty"`
	YearBuild              string                 `xml:"YearBuild,omitempty"`
	AlternateURL           string                 `xml:"AlternateURL,omitempty"` //  200
	PropertyDescriptions   *PropertyDescriptions  `xml:"PropertyDescriptions,omitempty"`
	Images                 *Images                `xml:"Images,omitempty"`
	Latitude               float64                `xml:"Latitude,omitempty"`
	Longitude              float64                `xml:"Longitude,omitempty"`
	//Features                 *struct{}             `xml:"Features,omitempty"` //??
	//EnergyRating             *struct{}             `xml:"EnergyRating,omitempty"` //TODO NEED CHECK VERSION 3.3 XSD
	//RentalCommissionMonths     string                 `xml:"RentalCommissionMonths,omitempty"` // 64
	//ParkingSpaces              float64                `xml:"ParkingSpaces,omitempty"`
	//HeatingCost                int                    `xml:"HeatingCost,omitempty"`

	//LocalZoneID                  int       `xml:"LocalZoneID,omitempty"`                  // ??
	//LocalZone                    string    `xml:"LocalZone,omitempty"`                    // ?? 50
	//District                     string    `xml:"District,omitempty"`                     // ?? 50
	//RentalPriceGranularity       int       `xml:"RentalPriceGranularity ,omitempty"`      // ?? TODO Enum PaymentPeriod
	//SoldPriceCurrency            int       `xml:"SoldPriceCurrency,omitempty"`            //??TODO Enum CurrencyCode
	//PropertyStatus               int       `xml:"PropertyStatus,omitempty"`               //??TODO Enum PropertyStatus
	//MarketStatus                 int       `xml:"MarketStatus,omitempty"`                 //??TODO Enum MarketStatus
	//LivingArea                   float64   `xml:"LivingArea,omitempty"`                   // ??
	//CubicVolume                  float64   `xml:"CubicVolume,omitempty"`                    // ??
	//BuiltArea                    float64   `xml:"BuiltArea,omitempty"`                    // ??
	//NumberOfApartmentsInBuilding int       `xml:"NumberOfApartmentsInBuilding,omitempty"` // ??
	//PossessionDate               string    `xml:"PossessionDate,omitempty"`               //?? 50
	//AvailabilityDate             Date      `xml:"AvailabilityDate,omitempty"`             //??
	//OrigListingDate              Date      `xml:"OrigListingDate,omitempty"`              //??
	//ExpiryDate                   Date      `xml:"ExpiryDate,omitempty"`                   //??
	//VirtualTourURL               string    `xml:"VirtualTourURL,omitempty"`               //?? 200
	//Notes                        string    `xml:"Notes,omitempty"`                        //?? 500
	//CoopCommision                int       `xml:"CoopCommision,omitempty"`                //??
	//CoopProperty                 bool      `xml:"CoopProperty,omitempty"`                 //??
	//ComTotalPct                  int       `xml:"ComTotalPct,omitempty"`                  //??
	//ComTotalFix                  int       `xml:"ComTotalFix,omitempty"`                  //??
	//ComBuyAgentPctTot            int       `xml:"ComBuyAgentPctTot,omitempty"`            //??
	//ComBuyAgentPct               int       `xml:"ComBuyAgentPct,omitempty"`               //??
	//ComBuyAgentFix               int       `xml:"ComBuyAgentFix,omitempty"`               //??
	//ComReferralPct               int       `xml:"ComReferralPct,omitempty"`               //??
	//ComReferralFixed             int       `xml:"CoopCommision,omitempty"`                //??
	//ComRefSource                 string    `xml:"ComRefSource,omitempty"`                 //?? 50
	//Commercial                   *struct{} `xml:"Commercial,omitempty"`                   //??
	//TransactionCost              float64   `xml:"TransactionCost,omitempty"`              //??
	//ProcessImage                 bool      `xml:"ProcessImage,omitempty"`                 //??
	//Utilities                    int       `xml:"Utilities,omitempty"`                    //??
	//UtilityPeriod                int       `xml:"UtilityPeriod,omitempty"`                //??TODO Enum PaymentPeriod
	//HeatingCostPeriod            int       `xml:"HeatingCostPeriod,omitempty"`            //??TODO Enum PaymentPeriod
	//MaintenanceFee               int       `xml:"MaintenanceFee,omitempty"`               //??
	//MaintenanceFeePeriod         int       `xml:"MaintenanceFeePeriod,omitempty"`         //??TODO Enum PaymentPeriod
	//ParkingCost                  int       `xml:"ParkingCost,omitempty"`                  //??
	//ParkingCostPeriod            int       `xml:"ParkingCostPeriod,omitempty"`            //??TODO Enum PaymentPeriod
}
