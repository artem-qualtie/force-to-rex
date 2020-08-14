package rex

import (
	"encoding/xml"
	"fmt"
	"force-to-rex/conf"
	rexmodel "force-to-rex/rex/model"
	"force-to-rex/salesforce"
	salesforcemodel "force-to-rex/salesforce/model"
	"github.com/nimajalali/go-force/force"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var IntegratorID, _ = conf.Config.Int("IntegratorID")
var IntegratorOfficeID = conf.Config.Str("IntegratorOfficeID")
var IntegratorSalesAssociateID = conf.Config.Str("IntegratorSalesAssociateID")
var FileStart = `<UploadXML xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="https://www.gryphtech.com/REX/XSD/Property_v3_2.xsd"><Version Version="3.2"></Version><Properties>`
var FileEnd = `</Properties></UploadXML>`

func GenerateFiles(fileDateId string) {
	IntegratorID, _ = conf.Config.Int("integratorID")
	IntegratorOfficeID = conf.Config.Str("integratorOfficeID")
	IntegratorSalesAssociateID = conf.Config.Str("integratorSalesAssociateID")

	GeneratePropertiesFile(fileDateId)
	GenerateJobFile(fileDateId)
}

func GeneratePropertiesFile(fileDateId string) {
	os.RemoveAll("temp")
	err := os.Mkdir("temp", 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	os.Mkdir("temp/images", 0777)

	imageForceApi, err := force.Create(
		conf.Config.Str("apiVersion"),
		conf.Config.Str("clientID"),
		conf.Config.Str("clientSecret"),
		conf.Config.Str("username"),
		conf.Config.Str("password"),
		conf.Config.Str("securityToken"),
		conf.Config.Str("environment"),
	)

	file, err := os.OpenFile("temp/P_"+fileDateId+".xml", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.WriteString(FileStart)

	forceApi, resultSObjects, err := salesforce.GetTestData()
	if err != nil {
		fmt.Println(err)
	}
	for {
		for _, data := range resultSObjects.Records {
			rex := ForceToRex(data)
			rex.Images = GetPropertyImages(imageForceApi, rex.IntegratorPropertyID)
			b, err := xml.Marshal(rex)
			if err != nil {
				fmt.Println("error: ", err)
			}
			file.Write(b)
		}
		if resultSObjects.Done || resultSObjects.NextRecordsUri == "" {
			break
		}
		err = forceApi.QueryNext(resultSObjects.NextRecordsUri, resultSObjects)
	}

	file.WriteString(FileEnd)
}

func GenerateJobFile(fileDateId string) {
	file, err := os.OpenFile("temp/JOB_"+fileDateId+".xml", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	job := rexmodel.JobControl{
		Xmlns:                     "http://www.w3.org/2001/XMLSchema-instance",
		NoNamespaceSchemaLocation: "http://www.gryphtech.com/REX/XSD/JobControl_v3_2.xsd",
		IntegratorID:              IntegratorID,
		PropertyFileName:          "P_" + fileDateId + ".xml",
	}
	b, err := xml.Marshal(job)
	if err != nil {
		fmt.Println("error: ", err)
	}
	file.Write(b)
}

func ValidateString(unvalidate string, max int, notEmpty bool) string {
	var answer string
	runs := []rune(unvalidate)
	if notEmpty && len(runs) < 1 {
		answer = " "
	} else if len(runs) > max {
		answer = string(runs[:max])
	} else {
		answer = unvalidate
	}
	return answer
}

func ValidateNumber(unvalidate int, min int, max int) int {
	if unvalidate < min {
		unvalidate = min
	} else if unvalidate > max {
		unvalidate = max
	}
	return unvalidate
}

func TranslateFloat64(float string) float64 {
	f, _ := strconv.ParseFloat(float, 64)
	return f
}

func TranslateRegionID(region string) int {
	return 1 //TODO translate mapping: pba__Area_pb__c => RegionID
}

var ResidentialPropertyType = map[string]int{
	"Single Family":      11,
	"House":              12,
	"Apartment":          4,
	"Townhouse":          39,
	"Villa":              42,
	"Studio":             53,
	"Duplex":             13,
	"Condo":              4,
	"Penthouse":          45,
	"Multi Family":       35,
	"Land":               21,
	"Loft":               95,
	"Mobile Home":        25,
	"Lot/Land":           28,
	"Farm":               15,
	"Boat Dock":          18,
	"Other":              93,
	"Parking":            86,
	"Coop":               103,
	"Rental":             102,
	"Residential":        102,
	"Residential Lease":  22,
	"Residential Rental": 102,
	"Waterfront":         68,
}
var CommercialPropertyType = map[string]int{
	"Manufactured":                  4,
	"Business Opportunity":          9,
	"Commercial/Industrial":         4,
	"Commercial/Industrial - Lease": 4,
	"Hotel":                         12,
	"Commercial":                    1,
	"Commercial Lease":              1,
	"Commercial Sale":               1,
	"Business/Commercial":           1,
	"Timeshare":                     1,
	"Office":                        7,
}

//translate mapping: Development_Type__c => CommercialResidential
func TranslateCommercialResidential(propertyType string) int {
	if _, OK := CommercialPropertyType[propertyType]; OK {
		return 2
	}
	return 1
}

//translate mapping: pba__PropertyType__c => PropertyType
func TranslatePropertyType(propertyType string) int {
	if answer, OK := ResidentialPropertyType[propertyType]; OK {
		return answer
	} else if answer, OK = CommercialPropertyType[propertyType]; OK {
		return answer
	}
	return 1
}

func TranslateCityID(cityID string) int {
	return 1 //TODO translate mapping:
}

func TranslateCurrentListingCurrency(currentListingCurrency string) string {
	return currentListingCurrency
}

func TranslateDate(soldDate string) string {
	currentTime, err := time.Parse("2006-01-02 15:04:05", soldDate)
	if err != nil {
		fmt.Println(err)
	}
	return currentTime.Format("2006-01-02")
}

var TransactionType = map[string]int{
	"Sale": 2,
	"Rent": 1,
}

//translate mapping: pba__ListingType__c => TransactionType
func TranslateTransactionType(transactionType string) int {
	return TransactionType[transactionType]
}

var PropertyCategory = map[string]int{
	"Residential area":     17,
	"Resort area":          6,
	"Non residantial area": 3,
	"Other":                10,
}

//mapping: Zone_Category__c => PropertyCategory
func TranslatePropertyCategory(propertyCategory string) int {
	return PropertyCategory[propertyCategory]
}

var ListingStatus = map[string]int{
	"New":         1,
	"Active":      1,
	"Offer":       12,
	"Cancelled":   5,
	"Sold_Rented": 2,
	"Expired":     4,
	"Closing":     8,
}

//translate mapping: pba__Status__c => ListingStatus
func TranslateListingStatus(listingStatus string) int {
	return ListingStatus[listingStatus]
}

func DownloadImage(url, filename string) error {
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	//open a file for writing
	file, err := os.Create("temp/images/" + filename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func ForceImagesToRexImages(records []*salesforcemodel.PropertyMediaSObject) []rexmodel.Image {
	answer := make([]rexmodel.Image, 0)
	for _, record := range records {
		var isDefault rexmodel.IsDefault
		if record.SequenceNumber == 0 {
			isDefault.IsDefault = "True"
		} else {
			isDefault.IsDefault = "False"
		}
		image := rexmodel.Image{
			SequenceNumber: int(record.SequenceNumber),
			FileName:       record.FileName,
			//DescriptiveName: record.DescriptiveName,
			IsDefault: isDefault,
		}
		err := DownloadImage(record.URL, record.FileName)
		if err != nil {
			continue
		}
		answer = append(answer, image)
	}
	return answer
}

func GetPropertyImages(imageForceApi *force.ForceApi, propertyID string) *rexmodel.Images {
	answer := rexmodel.Images{
		DefaultImageSequenceNumber: 0,
		Images:                     make([]rexmodel.Image, 0),
	}
	pbaImages, err := salesforce.GetPropertyImages(imageForceApi, propertyID)
	if err != nil {
		fmt.Println(err)
	}

	for {
		images := ForceImagesToRexImages(pbaImages.Records)
		answer.Images = append(answer.Images, images...)
		if pbaImages.Done || pbaImages.NextRecordsUri == "" {
			break
		}
		err = imageForceApi.QueryNext(pbaImages.NextRecordsUri, pbaImages)
	}
	if len(answer.Images) < 1 {
		return nil
	}
	return &answer
}

func ForceToRex(forceDoc *salesforcemodel.PBAListingSObject) *rexmodel.Property {
	return &rexmodel.Property{
		IntegratorID:               IntegratorID,
		IntegratorPropertyID:       ValidateString(forceDoc.Id, 50, true),
		IntegratorOfficeID:         IntegratorOfficeID,
		IntegratorSalesAssociateID: IntegratorSalesAssociateID,
		RegionID: rexmodel.RegionID{
			RegionID: TranslateRegionID(forceDoc.RegionID),
		},
		CommercialResidential: rexmodel.CommercialResidential{
			CommercialResidential: ValidateNumber(TranslateCommercialResidential(forceDoc.PropertyType), 1, 2),
		},
		StreetNumber:    ValidateString(fmt.Sprint(forceDoc.StreetNumber), 20, true),
		StreetName:      ValidateString(forceDoc.StreetName, 50, true),
		ApartmentNumber: ValidateString(fmt.Sprint(forceDoc.ApartmentNumber), 15, true),
		AddressLine2:    ValidateString(forceDoc.CityID, 50, false),
		PostalCode:      ValidateString(forceDoc.PostalCode, 15, false),
		CityID:          TranslateCityID(forceDoc.CityID),
		//FloorLevel:                 &rexmodel.FloorLevel{FloorLevel:1},// TODO Check
		ShowAddressOnWeb:    rexmodel.ShowAddressOnWeb{ShowAddressOnWeb: "True"},
		CurrentListingPrice: forceDoc.CurrentListingPrice,
		CurrentListingCurrency: rexmodel.CurrentListingCurrency{
			CurrentListingCurrency: TranslateCurrentListingCurrency(forceDoc.CurrentListingCurrency),
		},
		//PriceType:                  &rexmodel.PriceType{PriceType:},// TODO Check
		SoldPrice: forceDoc.SoldPrice,
		SoldDate:  TranslateDate(forceDoc.SoldDate),
		TransactionType: rexmodel.TransactionType{
			TransactionType: ValidateNumber(TranslateTransactionType(forceDoc.TransactionType), 1, 3),
		},
		ContractType: rexmodel.ContractType{
			ContractType: 1, //TODO find
		},
		PropertyStatus: rexmodel.PropertyStatus{
			PropertyStatus: 1, //TODO find
		},
		PropertyType: rexmodel.PropertyType{
			PropertyType: ValidateNumber(TranslatePropertyType(forceDoc.PropertyType), 1, 103),
		},
		PropertyCategory: &rexmodel.PropertyCategory{
			PropertyCategory: ValidateNumber(TranslatePropertyCategory(forceDoc.PropertyCategory), 1, 24),
		},
		ListingStatus: rexmodel.ListingStatus{
			ListingStatus: ValidateNumber(TranslateListingStatus(forceDoc.ListingStatus), 1, 12),
		},
		TotalArea:           forceDoc.TotalArea,
		LotSize:             ValidateString(fmt.Sprint(forceDoc.LotSizeFirstPart, forceDoc.LotSizeSecondPart), 50, false),
		LotSizeM2:           forceDoc.LotSizeFirstPart,
		TotalNumOfRooms:     forceDoc.TotalNumOfRooms,
		NumberOfBathrooms:   int(forceDoc.NumberOfBathrooms),
		NumberOfBedrooms:    int(forceDoc.NumberOfBedrooms),
		NumberOfToiletRooms: int(forceDoc.NumberOfToiletRooms),
		NumberOfFloors:      int(forceDoc.NumberOfFloors),
		YearBuild:           ValidateString(fmt.Sprint(forceDoc.YearBuild), 50, false),
		//AlternateURL:               TranslateAlternateURL(forceDoc.AlternateURL),// TODO Check
		PropertyDescriptions: &rexmodel.PropertyDescriptions{
			Descriptions: []rexmodel.PropertyDescription{
				{
					LanguageID:      rexmodel.LanguageID{LanguageID: 1},
					DescriptionText: ValidateString(forceDoc.DescriptionText, 40000, false),
				},
			},
		},
		Images:    nil,
		Latitude:  TranslateFloat64(forceDoc.Latitude),
		Longitude: TranslateFloat64(forceDoc.Longitude),
	}
}
