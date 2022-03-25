package tmallexchangereceivegetrequest

type TmallExchangeReceiveGetRequest struct {
	apiParas map[string]string
}

func (t *TmallExchangeReceiveGetRequest) GetApiMethodName() string {
	return "tmall.exchange.receive.get"
}

func (t *TmallExchangeReceiveGetRequest) GetApiParas() map[string]string {
	return t.apiParas
}
func (t *TmallExchangeReceiveGetRequest) SetFields(fields string) {
	t.apiParas["fields"] = fields
}
func (t *TmallExchangeReceiveGetRequest) SetStartGmtModifiedTime(startGmtModifiedTime string) {
	t.apiParas["start_gmt_modified_time"] = startGmtModifiedTime
}
func (t *TmallExchangeReceiveGetRequest) SetEndGmtModifiedTime(endGmtModifiedTime string) {
	t.apiParas["end_gmt_modifed_time"] = endGmtModifiedTime
}
func (t *TmallExchangeReceiveGetRequest) SetPageSize(pageSize string) {
	t.apiParas["page_size"] = pageSize
}
func New() *TmallExchangeReceiveGetRequest {
	return &TmallExchangeReceiveGetRequest{apiParas: map[string]string{}}
}
func (t TmallExchangeReceiveGetRequest) SetPageNo(page string) {
	t.apiParas["page_no"] = page
}
