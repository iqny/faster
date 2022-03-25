package taobaologisticsaddresssearch

type LogisticsAddressSearchRequest struct {
	apiParas map[string]string
}
func (t *LogisticsAddressSearchRequest) GetApiMethodName() string {
	return "taobao.logistics.address.search"
}

func (t *LogisticsAddressSearchRequest) GetApiParas() map[string]string {
	return t.apiParas
}

func (t *LogisticsAddressSearchRequest) SetRdef(rdef string) {
	t.apiParas["rdef"] = rdef
}
func New() *LogisticsAddressSearchRequest {
	return &LogisticsAddressSearchRequest{apiParas: map[string]string{}}
}