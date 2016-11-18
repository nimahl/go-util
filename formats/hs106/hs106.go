package hs106

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

// Mapping for the HS106 file format
type (
	HS106 struct {
		Type1 HS106Type1
		Type2 HS106Type2
		Type3 []*HS106Type3
		Type4 []*HS106Type4
		Type5 []*HS106Type5
		Type6 []*HS106Type6
	}
	HS106Type1 struct {
		OrderID                string
		CustomerID             string
		OrderDateTime          string
		OrderStatus            string
		OrderTotalValue        string
		OrderCurrency          string
		DeliveryMethod         string
		EarliestCollectionDate string
		LatestDeliveryDate     string
		DeliveryCharge         string
		OrderSource            string
		Locale                 string
		DestinationType        string
		CustomerCollectionDate string
		DeliveryStoreNumber    string
		CollectPlusDetails     string
	}
	HS106Type2 struct {
		OrderID             string
		PaymentProvider     string
		PaymentReference    string
		AuthorisationCode   string
		AuthorisedAmount    string
		PaymentType         string
		MerchantAccount     string
		PaypalAddress       string
		CardLast4           string
		Subclient           string
		ConcatenatedAddress string
	}
	HS106Type3 struct {
		OrderID              string
		OrderLineSequenceID  string
		SkuID                string
		SkuUnitPrice         string
		OrderQuantity        string
		OrderLineDescription string
	}
	HS106Type4 struct {
		OrderID         string
		ContactType     string
		Title           string
		FirstName       string
		Surname         string
		EmailAddress    string
		TelephoneNumber string
		ServiceNumber   string
	}
	HS106Type5 struct {
		OrderId                     string
		AddressType                 string
		AddressLine1                string
		AddressLine2                string
		AddressLine3                string
		Town                        string
		County                      string
		PostCode                    string
		CountryCode                 string
		SpecialDeliveryInstructions string
	}
	HS106Type6 struct {
		OrderID               string
		OrderLineSequenceID   string
		DiscountID            string
		DiscountName          string
		PromoCode             string
		PromoCodeDefinitionID string
		DiscountAmount        string
	}
)

// FileParser will parse a raw file into events
type FileParser interface {
	Parser
}

type (
	// Parser will parse a raw byte stream of a file and parse records from it
	Parser interface {
		Parse([]byte) (records []byte)
	}
	// ParserFunc is a function adaptor to Parser
	ParserFunc func([]byte) []byte
)

type hs106Parser struct {
	delimiter       string
	groupIdentifier string
}

func NewHS106(data []byte) (*hs106Parser, error) {
	h := new(hs106Parser)
	h.delimiter = "^^"
	h.groupIdentifier = "1"
	h.record(data)
	return h, nil
}

func (h *hs106Parser) record(data []byte) []byte {
	scanner := bufio.NewScanner(bytes.NewBuffer(data))
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), h.delimiter)
		attributeID := row[0]
		fmt.Println(attributeID)
		if attributeID == h.groupIdentifier {
			t1 := newHS106Type1(row)
			fmt.Println(t1)
		}
	}
	return nil
}

func newHS106Type1(row []string) *HS106Type1 {
	return &HS106Type1{
		OrderID:                row[0],
		CustomerID:             row[1],
		OrderDateTime:          row[2],
		OrderStatus:            row[3],
		OrderTotalValue:        row[4],
		OrderCurrency:          row[5],
		DeliveryMethod:         row[6],
		EarliestCollectionDate: row[7],
		LatestDeliveryDate:     row[8],
		DeliveryCharge:         row[9],
		OrderSource:            row[10],
		Locale:                 row[11],
		DestinationType:        row[12],
		CustomerCollectionDate: row[13],
		DeliveryStoreNumber:    row[14],
		CollectPlusDetails:     row[15],
	}
}
func newHS106Type2(row []string) *HS106Type2 {
	return &HS106Type2{
		OrderID:             row[0],
		PaymentProvider:     row[1],
		PaymentReference:    row[2],
		AuthorisationCode:   row[3],
		AuthorisedAmount:    row[4],
		PaymentType:         row[5],
		MerchantAccount:     row[6],
		PaypalAddress:       row[7],
		CardLast4:           row[8],
		Subclient:           row[9],
		ConcatenatedAddress: row[10],
	}
}

func newHS106Type3(row []string) *HS106Type3 {
	return &HS106Type3{
		OrderID:              row[0],
		OrderLineSequenceID:  row[1],
		SkuID:                row[2],
		SkuUnitPrice:         row[3],
		OrderQuantity:        row[4],
		OrderLineDescription: row[5],
	}
}

func newHS106Type4(row []string) *HS106Type4 {
	return &HS106Type4{
		OrderID:         row[0],
		ContactType:     row[1],
		Title:           row[2],
		FirstName:       row[3],
		Surname:         row[4],
		EmailAddress:    row[5],
		TelephoneNumber: row[6],
		ServiceNumber:   row[7],
	}
}
func newHS106Type5(row []string) *HS106Type5 {
	return &HS106Type5{
		OrderId:                     row[0],
		AddressType:                 row[1],
		AddressLine1:                row[2],
		AddressLine2:                row[3],
		AddressLine3:                row[4],
		Town:                        row[5],
		County:                      row[6],
		PostCode:                    row[7],
		CountryCode:                 row[8],
		SpecialDeliveryInstructions: row[9],
	}
}
func newHS106Type6(row []string) *HS106Type6 {
	return &HS106Type6{
		OrderID:               row[0],
		OrderLineSequenceID:   row[1],
		DiscountID:            row[2],
		DiscountName:          row[3],
		PromoCode:             row[4],
		PromoCodeDefinitionID: row[5],
		DiscountAmount:        row[6],
	}
}
