package tracing_tools

type TracingMetadata struct {
	RequestID string `json:"requestID"`
	AppName   string `json:"appName"`
}
