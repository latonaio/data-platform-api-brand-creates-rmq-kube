package dpfm_api_processing_formatter

type HeaderUpdates struct {
	Brand							int		`json:"Brand"`
	BrandType						string	`json:"BrandType"`
	BrandOwner						int		`json:"BrandOwner"`
	BrandOwnerBusinessPartnerRole	string	`json:"BrandOwnerBusinessPartnerRole"`
	PersonResponsible				string	`json:"PersonResponsible"`
	ValidityStartDate				string	`json:"ValidityStartDate"`
	ValidityStartTime				string	`json:"ValidityStartTime"`
	ValidityEndDate					string	`json:"ValidityEndDate"`
	ValidityEndTime					string	`json:"ValidityEndTime"`
	Description						string	`json:"Description"`
	LongText						string	`json:"LongText"`
	Introduction					*string	`json:"Introduction"`
	Project							*int	`json:"Project"`
	WBSElement						*int	`json:"WBSElement"`
	Tag1							*string	`json:"Tag1"`
	Tag2							*string	`json:"Tag2"`
	Tag3							*string	`json:"Tag3"`
	Tag4							*string	`json:"Tag4"`
	LastChangeDate					string	`json:"LastChangeDate"`
	LastChangeTime					string	`json:"LastChangeTime"`
}

type PartnerUpdates struct {
	Brand                 	int     `json:"Brand"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
	EmailAddress            *string `json:"EmailAddress"`
}
