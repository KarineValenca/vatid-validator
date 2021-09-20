package soap

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/KarineValenca/vatid-validator/pkg/validator/dto"
)

func SoapEncode(contents interface{}) ([]byte, error) {
	data, err := xml.Marshal(contents)
	if err != nil {
		return nil, err
	}
	data = append([]byte("\n"), data...)
	env := dto.SoapEnvelope{
		Val1: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: dto.SoapBody{
			Contents: data,
		},
	}
	return xml.Marshal(&env)
}

func SoapDecode(data []byte, contents interface{}) error {
	env := dto.SoapEnvelopeResponse{Body: dto.SoapBodyResponse{}}
	err := xml.Unmarshal(data, &env)

	if err != nil {
		return err
	}

	return xml.Unmarshal(env.Body.Contents, contents)
}

type Client struct {
	*http.Client
	CheckVatServiceURL string
}

func (c *Client) CheckValidVAT(countryCode string, vatNumber string) (bool, error) {
	//if not 'DE' as ContryCode, return invalid vat
	if countryCode != "DE" {
		return false, nil
	}

	request := &dto.CheckValidVATResquest{
		CountryCode: countryCode,
		VatNumber:   vatNumber,
	}

	contents, _ := SoapEncode(&request)

	req, err := http.NewRequest("POST", c.CheckVatServiceURL, bytes.NewReader(contents))
	if err != nil {
		log.Fatal("Error creating new request ", err)
		return false, err
	}

	req.Header.Set("Content-type", "text/xml")

	// dispatch the request
	res, err := c.Do(req)
	if err != nil {
		log.Fatal("Error on dispatching request ", err)
		return false, err
	}
	defer res.Body.Close()

	var result dto.CheckValidVATResponse

	bytesBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading body ", err)
		return false, err
	}

	err = SoapDecode([]byte(bytesBody), &result)
	if err != nil {
		log.Fatal("Error decoding the response ", err)
		return false, err
	}

	return result.Valid, nil
}
