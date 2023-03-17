package errors

type Error struct {
	ErrCode int32  `json:"errCode"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}
