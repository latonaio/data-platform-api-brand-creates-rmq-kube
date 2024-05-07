package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-brand-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-brand-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-brand-creates-rmq-kube/sub_func_complementer"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToHeaderCreates(subfuncSDC *sub_func_complementer.SDC) (*Header, error) {
	data := subfuncSDC.Message.Header

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToPartnerCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Partner, error) {
	partners := make([]Partner, 0)

	for _, data := range *subfuncSDC.Message.Partner {
		partner, err := TypeConverter[*Partner](data)
		if err != nil {
			return nil, err
		}

		partners = append(partners, *partner)
	}

	return &partners, nil
}

func ConvertToHeaderUpdates(headerData dpfm_api_input_reader.Header) (*Header, error) {
	data := headerData

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToPartnerUpdates(partnerUpdates *[]dpfm_api_processing_formatter.PartnerUpdates) (*[]Partner, error) {
	partners := make([]Partner, 0)

	for _, data := range *partnerUpdates {
		partner, err := TypeConverter[*Partner](data)
		if err != nil {
			return nil, err
		}

		partners = append(partners, *partner)
	}

	return &partners, nil
}

func ConvertToHeader(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	subfuncSDC.Message.Header = &sub_func_complementer.Header{
		Brand:                          *input.Header.Brand,
		BrandType:                      input.Header.BrandType,
		BrandOwner:                     input.Header.BrandOwner,
		BrandOwnerBusinessPartnerRole:  input.Header.BrandOwnerBusinessPartnerRole,
		PersonResponsible:              input.Header.PersonResponsible,
		ValidityStartDate:              input.Header.ValidityStartDate,
		ValidityStartTime:              input.Header.ValidityStartTime,
		ValidityEndDate:                input.Header.ValidityEndDate,
		ValidityEndTime:                input.Header.ValidityEndTime,
		Description:                    input.Header.Description,
		LongText:                       input.Header.LongText,
		Introduction:                   input.Header.Introduction,
		Project:                        input.Header.Project,
		WBSElement:                     input.Header.WBSElement,
		Tag1:                           input.Header.Tag1,
		Tag2:                           input.Header.Tag2,
		Tag3:                           input.Header.Tag3,
		Tag4:                           input.Header.Tag4,
		CreationDate:                   input.Header.CreationDate,
		CreationTime:                   input.Header.CreationTime,
		LastChangeDate:                 input.Header.LastChangeDate,
		LastChangeTime:                 input.Header.LastChangeTime,
		IsReleased:                     input.Header.IsReleased,
		IsMarkedForDeletion:            input.Header.IsMarkedForDeletion,
	}

	return subfuncSDC
}

func ConvertToPartner(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var partners []sub_func_complementer.Partner

	partners = append(
		partners,
		sub_func_complementer.Partner{
			Brand:          		 *input.Header.Brand,
			PartnerFunction:         input.Header.Partner[0].PartnerFunction,
			BusinessPartner:         input.Header.Partner[0].BusinessPartner,
			BusinessPartnerFullName: input.Header.Partner[0].BusinessPartnerFullName,
			BusinessPartnerName:     input.Header.Partner[0].BusinessPartnerName,
			Organization:            input.Header.Partner[0].Organization,
			Country:                 input.Header.Partner[0].Country,
			Language:                input.Header.Partner[0].Language,
			Currency:                input.Header.Partner[0].Currency,
			ExternalDocumentID:      input.Header.Partner[0].ExternalDocumentID,
			AddressID:               input.Header.Partner[0].AddressID,
			EmailAddress:            input.Header.Partner[0].EmailAddress,
		},
	)

	subfuncSDC.Message.Partner = &partners

	return subfuncSDC
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
