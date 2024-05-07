package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-brand-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
			Brand:							data.Brand,
			BrandType:						data.BrandType,
			BrandOwner:						data.BrandOwner,
			BrandOwnerBusinessPartnerRole:	data.BrandOwnerBusinessPartnerRole,
			PersonResponsible:				data.PersonResponsible,
			ValidityStartDate:				data.ValidityStartDate,
			ValidityStartTime:				data.ValidityStartTime,
			ValidityEndDate:				data.ValidityEndDate,
			ValidityEndTime:				data.ValidityEndTime,
			Description:					data.Description,
			LongText:						data.LongText,
			Introduction:					data.Introduction,
			Project:						data.Project,
			WBSElement:						data.WBSElement,
			Tag1:							data.Tag1,
			Tag2:							data.Tag2,
			Tag3:							data.Tag3,
			Tag4:							data.Tag4,
			LastChangeDate:					data.LastChangeDate,
			LastChangeTime:					data.LastChangeTime,
	}
}

func ConvertToPartnerUpdates(header dpfm_api_input_reader.Header, partner dpfm_api_input_reader.Partner) *PartnerUpdates {
	dataHeader := header
	data := partner

	return &PartnerUpdates{
		Brand:                    dataHeader.Brand,
		PartnerFunction:         data.PartnerFunction,
		BusinessPartner:         data.BusinessPartner,
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Organization:            data.Organization,
		Country:                 data.Country,
		Language:                data.Language,
		Currency:                data.Currency,
		ExternalDocumentID:      data.ExternalDocumentID,
		AddressID:               data.AddressID,
		EmailAddress:            data.EmailAddress,
	}
}
