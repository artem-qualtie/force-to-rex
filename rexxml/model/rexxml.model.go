package model

type JobControl struct {
	Xmlns                     string `xml:"xmlns:xsi,attr"`
	NoNamespaceSchemaLocation string `xml:"xsi:noNamespaceSchemaLocation,attr"`
	IntegratorID              int    `xml:"IntegratorID"`
	OfficeFileName            string `xml:"OfficeFile,omitempty"`
	AgentFileName             string `xml:"AgentFile,omitempty"`
	PropertyFileName          string `xml:"PropertyFile,omitempty"`
}

type UploadXML struct {
	Xmlns                     string          `xml:"xmlns:xsi,attr"`
	NoNamespaceSchemaLocation string          `xml:"xsi:noNamespaceSchemaLocation,attr"`
	Version                   Version         `xml:"Version"`
	PropertiesData            *PropertiesData `xml:"Properties,omitempty"`
	ImageRemovals             *struct{}       `xml:"ImageRemovals,omitempty"`
	PropertiesToDisable       string          `xml:"PropertiesToDisable,omitempty"`
}

type Version struct {
	Version string `xml:"Version,attr"`
}

type Date struct {
}

type PropertiesData struct {
	Properties []Property `xml:"Property"`
}

type Property struct {
	IntegratorID                 int       `xml:"IntegratorID"`               //TODO !!!
	IntegratorPropertyID         string    `xml:"IntegratorPropertyID"`       //TODO !!!
	SystemPropertyIntegratorID   string    `xml:"SystemPropertyIntegratorID"` //TODO !!!
	IntegratorOfficeID           string    `xml:"IntegratorOfficeID"`         //TODO !!!
	IntegratorSalesAssociateID   string    `xml:"IntegratorSalesAssociateID"` //TODO !!!
	RepresentingAgentID          string    `xml:"RepresentingAgentID,omitempty"`
	RegionID                     int       `xml:"RegionID"` //TODO !!!
	Disabled                     bool      `xml:"Disabled"`
	CommercialResidential        int       `xml:"CommercialResidential"` //TODO !!!
	StreetNumber                 string    `xml:"StreetNumber,omitempty"`
	StreetName                   string    `xml:"StreetName,omitempty"`
	ApartmentNumber              string    `xml:"ApartmentNumber,omitempty"`
	AddressLine2                 string    `xml:"AddressLine2,omitempty"`
	PostalCode                   string    `xml:"PostalCode,omitempty"`
	CityID                       int       `xml:"CityID"` //TODO !!!
	LocalZoneID                  int       `xml:"LocalZoneID,omitempty"`
	LocalZone                    string    `xml:"LocalZone,omitempty"`
	District                     string    `xml:"District,omitempty"`
	Latitude                     float64   `xml:"Latitude,omitempty"`
	Longitude                    float64   `xml:"Longitude,omitempty"`
	FloorLevel                   int       `xml:"FloorLevel,omitempty"`              //TODO Enum FloorLevel
	CurrentListingPrice          int       `xml:"CurrentListingPrice"`               //TODO !!!
	CurrentListingCurrency       int       `xml:"CurrentListingCurrency"`            //TODO !!! Enum CurrencyCode
	PriceType                    int       `xml:"PriceType,omitempty"`               //TODO Enum PriceType
	RentalPriceGranularity       int       `xml:"RentalPriceGranularity ,omitempty"` //TODO Enum PaymentPeriod
	RentalCommissionMonths       string    `xml:"RentalCommissionMonths,omitempty"`
	HidePricePublic              bool      `xml:"HidePricePublic"`
	ShowAddressOnWeb             bool      `xml:"ShowAddressOnWeb"`
	SoldPrice                    int       `xml:"LocalZoneID,omitempty"`
	SoldPriceCurrency            int       `xml:"SoldPriceCurrency,omitempty"` //TODO Enum CurrencyCode
	SoldDate                     Date      `xml:"SoldDate,omitempty"`
	TransactionType              int       `xml:"TransactionType"`            //TODO !!! Enum TransactionType
	ContractType                 int       `xml:"ContractType"`               //TODO !!! Enum ContractType
	PropertyStatus               int       `xml:"PropertyStatus,omitempty"`   //TODO Enum PropertyStatus
	PropertyType                 int       `xml:"PropertyType"`               //TODO !!! Enum PropertyType / CommPropertyType
	PropertyCategory             int       `xml:"PropertyCategory,omitempty"` //TODO Enum PropertyCategory
	ListingStatus                int       `xml:"ListingStatus"`              //TODO !!! Enum ListingStatus
	MarketStatus                 int       `xml:"MarketStatus,omitempty"`     //TODO Enum MarketStatus
	TotalArea                    float64   `xml:"TotalArea,omitempty"`
	LivingArea                   float64   `xml:"LivingArea,omitempty"`
	CubicVolume                  float64   `xml:"TotalArea,omitempty"`
	LotSize                      string    `xml:"LotSize,omitempty"`
	LotSizeM2                    float64   `xml:"LotSizeM2,omitempty"`
	BuiltArea                    float64   `xml:"BuiltArea,omitempty"`
	ParkingSpaces                float64   `xml:"ParkingSpaces,omitempty"`
	TotalNumOfRooms              int       `xml:"TotalNumOfRooms,omitempty"`
	NumberOfBathrooms            int       `xml:"NumberOfBathrooms,omitempty"`
	NumberOfBedrooms             int       `xml:"NumberOfBedrooms,omitempty"`
	NumberOfToiletRooms          int       `xml:"NumberOfToiletRooms,omitempty"`
	NumberOfFloors               int       `xml:"NumberOfFloors,omitempty"`
	NumberOfApartmentsInBuilding int       `xml:"NumberOfApartmentsInBuilding,omitempty"`
	Rooms                        *struct{} `xml:"Rooms,omitempty"`
	YearBuild                    int       `xml:"YearBuild,omitempty"`
	PossessionDate               string    `xml:"PossessionDate,omitempty"`
	AvailabilityDate             Date      `xml:"AvailabilityDate,omitempty"`
	OrigListingDate              Date      `xml:"OrigListingDate,omitempty"`
	ExpiryDate                   Date      `xml:"ExpiryDate,omitempty"`
	AlternateURL                 string    `xml:"AlternateURL,omitempty"`
	VirtualTourURL               string    `xml:"VirtualTourURL,omitempty"`
	PropertyDescriptions         *struct{} `xml:"PropertyDescriptions,omitempty"`
	Notes                        string    `xml:"Notes,omitempty"`
	Features                     *struct{} `xml:"Features,omitempty"`
	CoopCommision                int       `xml:"CoopCommision,omitempty"`
	CoopProperty                 bool      `xml:"CoopProperty,omitempty"`
	ComTotalPct                  int       `xml:"ComTotalPct,omitempty"`
	ComTotalFix                  int       `xml:"ComTotalFix,omitempty"`
	ComBuyAgentPctTot            int       `xml:"ComBuyAgentPctTot,omitempty"`
	ComBuyAgentPct               int       `xml:"ComBuyAgentPct,omitempty"`
	ComBuyAgentFix               int       `xml:"ComBuyAgentFix,omitempty"`
	ComReferralPct               int       `xml:"ComReferralPct,omitempty"`
	ComReferralFixed             int       `xml:"CoopCommision,omitempty"`
	ComRefSource                 string    `xml:"ComRefSource,omitempty"`
	Commercial                   *struct{} `xml:"Commercial,omitempty"`
	TransactionCost              float64   `xml:"TransactionCost,omitempty"`
	ProcessImage                 bool      `xml:"ProcessImage,omitempty"`
	Images                       *struct{} `xml:"Images,omitempty"`
	EnergyRating                 *struct{} `xml:"EnergyRating,omitempty"`
	Utilities                    int       `xml:"Utilities,omitempty"`
	UtilityPeriod                int       `xml:"UtilityPeriod,omitempty"` //TODO Enum PaymentPeriod
	HeatingCost                  int       `xml:"HeatingCost,omitempty"`
	HeatingCostPeriod            int       `xml:"HeatingCostPeriod,omitempty"` //TODO Enum PaymentPeriod
	MaintenanceFee               int       `xml:"MaintenanceFee,omitempty"`
	MaintenanceFeePeriod         int       `xml:"MaintenanceFeePeriod,omitempty"` //TODO Enum PaymentPeriod
	ParkingCost                  int       `xml:"ParkingCost,omitempty"`
	ParkingCostPeriod            int       `xml:"ParkingCostPeriod,omitempty"` //TODO Enum PaymentPeriod
}
