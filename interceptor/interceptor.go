package interceptor

type httpRequest struct {
	RequestUrl string `json:"requestUrl"`
	Latency    string `json:"latency"`
}
