// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "cryptopages": Application User Types
//
// Command:
// $ goagen
// --design=github.com/tutley/cryptopages/design
// --out=$(GOPATH)/src/github.com/tutley/cryptopages
// --regen=true
// --version=v1.3.1

package app

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// userPayload user type.
type userPayload struct {
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
		// Accepts Some Other Coin
		Other *bool `form:"other,omitempty" json:"other,omitempty" xml:"other,omitempty"`
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
	JobCategory    *string `form:"jobCategory,omitempty" json:"jobCategory,omitempty" xml:"jobCategory,omitempty"`
	JobDescription *string `form:"jobDescription,omitempty" json:"jobDescription,omitempty" xml:"jobDescription,omitempty"`
	// the user's location geographically
	Location *struct {
		// Should the location be shown to others
		MakePublic *bool `form:"makePublic,omitempty" json:"makePublic,omitempty" xml:"makePublic,omitempty"`
		// the location
		Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
	} `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// The user's full name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Name of the other coin a user accepts
	OtherCoin *string `form:"otherCoin,omitempty" json:"otherCoin,omitempty" xml:"otherCoin,omitempty"`
	// A password (only exposed to user)
	Password *string  `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	Skills   []string `form:"skills,omitempty" json:"skills,omitempty" xml:"skills,omitempty"`
	// The unique username for this person
	Username *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
}

// Validate validates the userPayload type instance.
func (ut *userPayload) Validate() (err error) {
	if ut.JobCategory != nil {
		if !(*ut.JobCategory == "hardware" || *ut.JobCategory == "software" || *ut.JobCategory == "writing" || *ut.JobCategory == "legal" || *ut.JobCategory == "labor" || *ut.JobCategory == "automotive" || *ut.JobCategory == "services" || *ut.JobCategory == "others") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`request.jobCategory`, *ut.JobCategory, []interface{}{"hardware", "software", "writing", "legal", "labor", "automotive", "services", "others"}))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 8, true))
		}
	}
	if ut.Username != nil {
		if utf8.RuneCountInString(*ut.Username) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.username`, *ut.Username, utf8.RuneCountInString(*ut.Username), 2, true))
		}
	}
	return
}

// Publicize creates UserPayload from userPayload
func (ut *userPayload) Publicize() *UserPayload {
	var pub UserPayload
	if ut.Available != nil {
		pub.Available = ut.Available
	}
	if ut.Coins != nil {
		pub.Coins = &struct {
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
			// Accepts Some Other Coin
			Other *bool `form:"other,omitempty" json:"other,omitempty" xml:"other,omitempty"`
			// Accepts Lumen
			Xlm *bool `form:"xlm,omitempty" json:"xlm,omitempty" xml:"xlm,omitempty"`
			// Accepts Ripple
			Xrp *bool `form:"xrp,omitempty" json:"xrp,omitempty" xml:"xrp,omitempty"`
		}{}
		if ut.Coins.Bcc != nil {
			pub.Coins.Bcc = ut.Coins.Bcc
		}
		if ut.Coins.Btc != nil {
			pub.Coins.Btc = ut.Coins.Btc
		}
		if ut.Coins.Eth != nil {
			pub.Coins.Eth = ut.Coins.Eth
		}
		if ut.Coins.Ltc != nil {
			pub.Coins.Ltc = ut.Coins.Ltc
		}
		if ut.Coins.Neo != nil {
			pub.Coins.Neo = ut.Coins.Neo
		}
		if ut.Coins.Other != nil {
			pub.Coins.Other = ut.Coins.Other
		}
		if ut.Coins.Xlm != nil {
			pub.Coins.Xlm = ut.Coins.Xlm
		}
		if ut.Coins.Xrp != nil {
			pub.Coins.Xrp = ut.Coins.Xrp
		}
	}
	if ut.Email != nil {
		pub.Email = &struct {
			// Should the email address be shown to others
			MakePublic *bool `form:"makePublic,omitempty" json:"makePublic,omitempty" xml:"makePublic,omitempty"`
			// the email address
			Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
		}{}
		if ut.Email.MakePublic != nil {
			pub.Email.MakePublic = ut.Email.MakePublic
		}
		if ut.Email.Value != nil {
			pub.Email.Value = ut.Email.Value
		}
	}
	if ut.JobCategory != nil {
		pub.JobCategory = ut.JobCategory
	}
	if ut.JobDescription != nil {
		pub.JobDescription = ut.JobDescription
	}
	if ut.Location != nil {
		pub.Location = &struct {
			// Should the location be shown to others
			MakePublic *bool `form:"makePublic,omitempty" json:"makePublic,omitempty" xml:"makePublic,omitempty"`
			// the location
			Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
		}{}
		if ut.Location.MakePublic != nil {
			pub.Location.MakePublic = ut.Location.MakePublic
		}
		if ut.Location.Value != nil {
			pub.Location.Value = ut.Location.Value
		}
	}
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	if ut.OtherCoin != nil {
		pub.OtherCoin = ut.OtherCoin
	}
	if ut.Password != nil {
		pub.Password = ut.Password
	}
	if ut.Skills != nil {
		pub.Skills = ut.Skills
	}
	if ut.Username != nil {
		pub.Username = ut.Username
	}
	return &pub
}

// UserPayload user type.
type UserPayload struct {
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
		// Accepts Some Other Coin
		Other *bool `form:"other,omitempty" json:"other,omitempty" xml:"other,omitempty"`
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
	JobCategory    *string `form:"jobCategory,omitempty" json:"jobCategory,omitempty" xml:"jobCategory,omitempty"`
	JobDescription *string `form:"jobDescription,omitempty" json:"jobDescription,omitempty" xml:"jobDescription,omitempty"`
	// the user's location geographically
	Location *struct {
		// Should the location be shown to others
		MakePublic *bool `form:"makePublic,omitempty" json:"makePublic,omitempty" xml:"makePublic,omitempty"`
		// the location
		Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
	} `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// The user's full name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Name of the other coin a user accepts
	OtherCoin *string `form:"otherCoin,omitempty" json:"otherCoin,omitempty" xml:"otherCoin,omitempty"`
	// A password (only exposed to user)
	Password *string  `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	Skills   []string `form:"skills,omitempty" json:"skills,omitempty" xml:"skills,omitempty"`
	// The unique username for this person
	Username *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
}

// Validate validates the UserPayload type instance.
func (ut *UserPayload) Validate() (err error) {
	if ut.JobCategory != nil {
		if !(*ut.JobCategory == "hardware" || *ut.JobCategory == "software" || *ut.JobCategory == "writing" || *ut.JobCategory == "legal" || *ut.JobCategory == "labor" || *ut.JobCategory == "automotive" || *ut.JobCategory == "services" || *ut.JobCategory == "others") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`type.jobCategory`, *ut.JobCategory, []interface{}{"hardware", "software", "writing", "legal", "labor", "automotive", "services", "others"}))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`type.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 8, true))
		}
	}
	if ut.Username != nil {
		if utf8.RuneCountInString(*ut.Username) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`type.username`, *ut.Username, utf8.RuneCountInString(*ut.Username), 2, true))
		}
	}
	return
}
