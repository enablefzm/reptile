package proxy

type ProxyDb struct {
	ID                    int    `json:"id"`
	IP                    string `json:"ip"`
	Port                  string `json:"port"`
	SchemeType            int    `json:"scheme_type"`
	AssessTimes           int    `json:"assess_times"`
	SuccessTimes          int    `json:"success_times"`
	AvgResponseTime       int    `json:"avg_response_time"`
	ContinuousFailedTimes int    `json:"continuous_failed_times"`
	Score                 int    `json:"score"`
	InsertTime            int    `json:"insert_time"`
	UpdateTime            int    `json:"update_time"`
}

// 获取一个随机的代理IP
func GetRandomProxy() *ProxyDb {

}
