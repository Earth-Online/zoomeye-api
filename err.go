package zoomeye

// api error json
type ApiError struct {
	// 	error id raised e.g. "rate limit"
	Error   string `json:"error"`
	// 	end user message of error raised . e.g. "Your account reached the API limit. Please wait a few minutes before making new requests"
	Message string `json:"message"`
	// references url with more information about the error . e.g. "https://developer.zoomeye.org/documentations#rate-limits"
	Url     string `json:"url"`
}
