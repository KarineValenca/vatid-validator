package soap

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckValidVATIsValid(t *testing.T) {
	// using check vat test service to get a valid response
	c := Client{
		Client:             http.DefaultClient,
		CheckVatServiceURL: "http://ec.europa.eu/taxation_customs/vies/services/checkVatTestService",
	}

	countryCode := "DE"
	vatNumber := "100"

	isValid, err := c.CheckValidVAT(countryCode, vatNumber)

	assert.Equal(t, true, isValid)
	assert.Nil(t, err)
}

func TestCheckValidVATIsInvalid(t *testing.T) {
	c := Client{
		Client:             http.DefaultClient,
		CheckVatServiceURL: "https://ec.europa.eu/taxation_customs/vies/services/checkVatService",
	}

	countryCode := "DE"
	vatNumber := "123456789"

	isValid, err := c.CheckValidVAT(countryCode, vatNumber)

	assert.Equal(t, false, isValid)
	assert.Nil(t, err)
}

func TestCheckValidVATIsInvalidNotGerman(t *testing.T) {
	c := Client{
		Client:             http.DefaultClient,
		CheckVatServiceURL: "https://ec.europa.eu/taxation_customs/vies/services/checkVatService",
	}

	countryCode := "UK"
	vatNumber := "123456789"

	isValid, err := c.CheckValidVAT(countryCode, vatNumber)

	assert.Equal(t, false, isValid)
	assert.Nil(t, err)
}
