package model

import (
	"fmt"
)

type PhoneNumber struct {
	ExitCode    *string `json:"exitCode,omitempty"`
	CountryCode *uint8  `json:"countryCode,omitempty"`
	AreaCode    uint16  `json:"areaCode"`
	Prefix      uint16  `json:"prefix"`
	Subscriber  uint16  `json:"subscriber"`
}

func (p *PhoneNumber) HrefValue() string {
	var exitCode, countryCode string

	if p.ExitCode != nil {
		exitCode = *p.ExitCode
	}
	if p.CountryCode != nil {
		countryCode = fmt.Sprintf("+%d", *p.CountryCode)
	}

	return fmt.Sprintf(
		"tel:%s%s%03d%03d%04d",
		exitCode,
		countryCode,
		p.AreaCode,
		p.Prefix,
		p.Subscriber,
	)
}

func (p *PhoneNumber) TextValue() string {
	// TODO: Support international phone numbers
	return fmt.Sprintf("(%03d) %03d-%04d", p.AreaCode, p.Prefix, p.Subscriber)
}
