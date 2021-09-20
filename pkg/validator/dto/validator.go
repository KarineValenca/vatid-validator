package dto

type ValidateVATIDRequest struct {
	CountryCode string
	VatNumber   string
}

type ValidateVATIDResponse struct {
	IsValid bool
}

type SoapEnvelope struct {
	XMLName struct{} `xml:"soapenv:Envelope"`
	Val1    string   `xml:"xmlns:soapenv,attr"`
	Body    SoapBody
}

type SoapBody struct {
	XMLName  struct{} `xml:"soapenv:Body"`
	Contents []byte   `xml:",innerxml"`
}

type SoapEnvelopeResponse struct {
	XMLName struct{} `xml:"Envelope"`
	Body    SoapBodyResponse
}

type SoapBodyResponse struct {
	XMLName  struct{} `xml:"Body"`
	Contents []byte   `xml:",innerxml"`
}

type CheckValidVATResquest struct {
	XMLName     struct{} `xml:"urn:ec.europa.eu:taxud:vies:services:checkVat:types checkVat"`
	CountryCode string   `xml:"countryCode"`
	VatNumber   string   `xml:"vatNumber"`
}

type CheckValidVATResponse struct {
	XMLName   struct{} `xml:"urn:ec.europa.eu:taxud:vies:services:checkVat:types checkVatResponse"`
	Valid     bool     `xml:"valid"`
	VatNumber string   `xml:"vatNumber"`
}
