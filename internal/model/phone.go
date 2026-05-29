package model

type PhoneNumber struct {
	ExitCode    *string `json:"exitCode,omitempty"`
	CountryCode *uint8  `json:"countryCode,omitempty"`
	AreaCode    uint16  `json:"areaCode"`
	Prefix      uint16  `json:"prefix"`
	Subscriber  uint16  `json:"subscriber"`
}
