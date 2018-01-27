// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "cryptopages": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/tutley/cryptopages/design
// --out=$(GOPATH)/src/github.com/tutley/cryptopages
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	"net/http"
	"unicode/utf8"
)

// The common media type to auth responses (default view)
//
// Identifier: application/vnd.cryptopages.success; view=default
type Success struct {
	// Always true
	OK bool `form:"ok" json:"ok" xml:"ok"`
}

// DecodeSuccess decodes the Success instance encoded in resp body.
func (c *Client) DecodeSuccess(resp *http.Response) (*Success, error) {
	var decoded Success
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A cryptopages user (default view)
//
// Identifier: application/vnd.cryptopages.user; view=default
type CryptopagesUser struct {
	// Is this user available to provide work
	Available *bool `form:"available,omitempty" json:"available,omitempty" xml:"available,omitempty"`
	// The coins this user will accept as payment
	Coins *struct {
		// Accepts Bitcoin Cash
		Bcc *bool `form:"bcc,omitempty" json:"bcc,omitempty" xml:"bcc,omitempty"`
		// Accepts Bitcoin
		Btc *bool `form:"btc,omitempty" json:"btc,omitempty" xml:"btc,omitempty"`
		// Accepts Ethereum
		Eth *bool `form:"eth,omitempty" json:"eth,omitempty" xml:"eth,omitempty"`
		// Accepts Litecoin
		Ltc *bool `form:"ltc,omitempty" json:"ltc,omitempty" xml:"ltc,omitempty"`
		// Accepts Neo
		Neo *bool `form:"neo,omitempty" json:"neo,omitempty" xml:"neo,omitempty"`
		// Accepts some other coin
		Other *struct {
			// Name of the other coin accepted
			Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
		} `form:"other,omitempty" json:"other,omitempty" xml:"other,omitempty"`
		// Accepts Lumen
		Xlm *bool `form:"xlm,omitempty" json:"xlm,omitempty" xml:"xlm,omitempty"`
		// Accepts Ripple
		Xrp *bool `form:"xrp,omitempty" json:"xrp,omitempty" xml:"xrp,omitempty"`
	} `form:"coins,omitempty" json:"coins,omitempty" xml:"coins,omitempty"`
	// The user's email address
	Email *struct {
		// Should the email address be shown to others
		MakePublic *bool `form:"makePublic,omitempty" json:"makePublic,omitempty" xml:"makePublic,omitempty"`
		// the email address
		Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
	} `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// API href for making requests on the user
	Href           string  `form:"href" json:"href" xml:"href"`
	JobCategory    string  `form:"jobCategory" json:"jobCategory" xml:"jobCategory"`
	JobDescription *string `form:"jobDescription,omitempty" json:"jobDescription,omitempty" xml:"jobDescription,omitempty"`
	// the user's location geographically
	Location *struct {
		// Should the location be shown to others
		MakePublic *bool `form:"makePublic,omitempty" json:"makePublic,omitempty" xml:"makePublic,omitempty"`
		// the location
		Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
	} `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// The user's full name
	Name string `form:"name" json:"name" xml:"name"`
	// A password (only exposed to user)
	Password *string  `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	Skills   []string `form:"skills,omitempty" json:"skills,omitempty" xml:"skills,omitempty"`
	// The unique username for this person
	Username string `form:"username" json:"username" xml:"username"`
}

// Validate validates the CryptopagesUser media type instance.
func (mt *CryptopagesUser) Validate() (err error) {
	if mt.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "username"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.JobCategory == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "jobCategory"))
	}
	if mt.Email != nil {
		if mt.Email.Value != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Email.Value); err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email.value`, *mt.Email.Value, goa.FormatEmail, err2))
			}
		}
		if mt.Email.Value != nil {
			if utf8.RuneCountInString(*mt.Email.Value) < 5 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email.value`, *mt.Email.Value, utf8.RuneCountInString(*mt.Email.Value), 5, true))
			}
		}
	}
	if !(mt.JobCategory == "hardware" || mt.JobCategory == "software" || mt.JobCategory == "writing" || mt.JobCategory == "legal" || mt.JobCategory == "labor" || mt.JobCategory == "automotive" || mt.JobCategory == "services" || mt.JobCategory == "others") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.jobCategory`, mt.JobCategory, []interface{}{"hardware", "software", "writing", "legal", "labor", "automotive", "services", "others"}))
	}
	if mt.Location != nil {
		if mt.Location.Value != nil {
			if utf8.RuneCountInString(*mt.Location.Value) < 2 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response.location.value`, *mt.Location.Value, utf8.RuneCountInString(*mt.Location.Value), 2, true))
			}
		}
	}
	if utf8.RuneCountInString(mt.Name) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 2, true))
	}
	if mt.Password != nil {
		if utf8.RuneCountInString(*mt.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *mt.Password, utf8.RuneCountInString(*mt.Password), 8, true))
		}
	}
	if utf8.RuneCountInString(mt.Username) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.username`, mt.Username, utf8.RuneCountInString(mt.Username), 2, true))
	}
	return
}

// A cryptopages user (tiny view)
//
// Identifier: application/vnd.cryptopages.user; view=tiny
type CryptopagesUserTiny struct {
	// API href for making requests on the user
	Href        string `form:"href" json:"href" xml:"href"`
	JobCategory string `form:"jobCategory" json:"jobCategory" xml:"jobCategory"`
	// The user's full name
	Name string `form:"name" json:"name" xml:"name"`
	// The unique username for this person
	Username string `form:"username" json:"username" xml:"username"`
}

// Validate validates the CryptopagesUserTiny media type instance.
func (mt *CryptopagesUserTiny) Validate() (err error) {
	if mt.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "username"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.JobCategory == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "jobCategory"))
	}
	if !(mt.JobCategory == "hardware" || mt.JobCategory == "software" || mt.JobCategory == "writing" || mt.JobCategory == "legal" || mt.JobCategory == "labor" || mt.JobCategory == "automotive" || mt.JobCategory == "services" || mt.JobCategory == "others") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.jobCategory`, mt.JobCategory, []interface{}{"hardware", "software", "writing", "legal", "labor", "automotive", "services", "others"}))
	}
	if utf8.RuneCountInString(mt.Name) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 2, true))
	}
	if utf8.RuneCountInString(mt.Username) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.username`, mt.Username, utf8.RuneCountInString(mt.Username), 2, true))
	}
	return
}

// DecodeCryptopagesUser decodes the CryptopagesUser instance encoded in resp body.
func (c *Client) DecodeCryptopagesUser(resp *http.Response) (*CryptopagesUser, error) {
	var decoded CryptopagesUser
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeCryptopagesUserTiny decodes the CryptopagesUserTiny instance encoded in resp body.
func (c *Client) DecodeCryptopagesUserTiny(resp *http.Response) (*CryptopagesUserTiny, error) {
	var decoded CryptopagesUserTiny
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// CryptopagesUserCollection is the media type for an array of CryptopagesUser (default view)
//
// Identifier: application/vnd.cryptopages.user; type=collection; view=default
type CryptopagesUserCollection []*CryptopagesUser

// Validate validates the CryptopagesUserCollection media type instance.
func (mt CryptopagesUserCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// CryptopagesUserCollection is the media type for an array of CryptopagesUser (tiny view)
//
// Identifier: application/vnd.cryptopages.user; type=collection; view=tiny
type CryptopagesUserTinyCollection []*CryptopagesUserTiny

// Validate validates the CryptopagesUserTinyCollection media type instance.
func (mt CryptopagesUserTinyCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeCryptopagesUserCollection decodes the CryptopagesUserCollection instance encoded in resp body.
func (c *Client) DecodeCryptopagesUserCollection(resp *http.Response) (CryptopagesUserCollection, error) {
	var decoded CryptopagesUserCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeCryptopagesUserTinyCollection decodes the CryptopagesUserTinyCollection instance encoded in resp body.
func (c *Client) DecodeCryptopagesUserTinyCollection(resp *http.Response) (CryptopagesUserTinyCollection, error) {
	var decoded CryptopagesUserTinyCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
