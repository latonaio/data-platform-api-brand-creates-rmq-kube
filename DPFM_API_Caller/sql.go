package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-brand-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-brand-creates-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_formatter "data-platform-api-brand-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-brand-creates-rmq-kube/sub_func_complementer"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *DPFMAPICaller) createSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *dpfm_api_output_formatter.Header
	var partner *[]dpfm_api_output_formatter.Partner
	for _, fn := range accepter {
		switch fn {
		case "Header":
			header = c.headerCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "Partner":
			partner = c.partnerCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:             header,
		Partner:            partner,
	}

	return data
}

func (c *DPFMAPICaller) updateSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *dpfm_api_output_formatter.Header
	var partner *[]dpfm_api_output_formatter.Partner
	for _, fn := range accepter {
		switch fn {
		case "Header":
			header = c.headerUpdateSql(mtx, input, output, errs, log)
		case "Partner":
			partner = c.partnerUpdateSql(mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:             header,
		Partner:            partner,
	}

	return data
}

func (c *DPFMAPICaller) headerCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	headerData := subfuncSDC.Message.Header
	res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "BrandHeader", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		*errs = append(*errs, err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "Header Data cannot insert"
		return nil
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) partnerCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Partner {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, partnerData := range *subfuncSDC.Message.Partner {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": partnerData, "function": "BrandPartner", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Partner Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPartnerCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) headerUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	header := input.Header
	headerData := dpfm_api_processing_formatter.ConvertToHeaderUpdates(header)

	sessionID := input.RuntimeSessionID
	if headerIsUpdate(headerData) {
		res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "BrandHeader", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Header Data cannot update"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderUpdates(header)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) partnerUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Partner {
	req := make([]dpfm_api_processing_formatter.PartnerUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, partner := range header.Partner {
		partnerData := *dpfm_api_processing_formatter.ConvertToPartnerUpdates(header, partner)

		if partnerIsUpdate(&partnerData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": partnerData, "function": "BrandPartner", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Partner Data cannot update"
				return nil
			}
		}
		req = append(req, partnerData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPartnerUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func headerIsUpdate(header *dpfm_api_processing_formatter.HeaderUpdates) bool {
	brand := header.Brand

	return !(brand == 0)
}

func partnerIsUpdate(partner *dpfm_api_processing_formatter.PartnerUpdates) bool {
	brand := partner.Brand
	partnerFunction := partner.PartnerFunction
	businessPartner := partner.BusinessPartner

	return !(brand == 0 || partnerFunction == "" || businessPartner == 0)
}
