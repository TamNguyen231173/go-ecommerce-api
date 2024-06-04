package response

const (
	ErrCodeSuccess      = 0
	ErrCodeParamInvalid = 2003
)

var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "param invalid",
}
