package Models

type Response struct {
	Response           any
	ErrorCode          int32
	ThrottleSeconds    int32
	ErrorStatus        string
	Message            string
	MessageData        map[string]string
	DetailedErrorTrace string
}
