package errs

type IMErrorCode int32

var IMErrorCode_SUCCESS IMErrorCode = 0

// api errorcode
var (
	IMErrorCode_API_DEFAULT            IMErrorCode = 10000
	IMErrorCode_API_APPKEY_REQUIRED    IMErrorCode = 10001
	IMErrorCode_API_NONCE_REQUIRED     IMErrorCode = 10002
	IMErrorCode_API_TIMESTAMP_REQUIRED IMErrorCode = 10003
	IMErrorCode_API_SIGNATURE_REQUIRED IMErrorCode = 10004
	IMErrorCode_API_APP_NOT_EXISTED    IMErrorCode = 10005
	IMErrorCode_API_SIGNATURE_FAIL     IMErrorCode = 10006
	IMErrorCode_API_REQ_BODY_ILLEGAL   IMErrorCode = 10007
	IMErrorCode_API_INTERNAL_TIMEOUT   IMErrorCode = 10008
	IMErrorCode_API_INTERNAL_RESP_FAIL IMErrorCode = 10009
	IMErrorCode_API_PARAM_REQUIRED     IMErrorCode = 10010
	IMErrorCode_API_PARAM_ILLEGAL      IMErrorCode = 10011
)

func GetApiErrorByCode(code IMErrorCode) *ApiErrorMsg {
	return newApiErrorMsg(200, code, "")
}

type ApiErrorMsg struct {
	HttpCode int         `json:"-"`
	Code     IMErrorCode `json:"code"`
	Msg      string      `json:"msg"`
}

func newApiErrorMsg(httpCode int, code IMErrorCode, msg string) *ApiErrorMsg {
	return &ApiErrorMsg{
		HttpCode: httpCode,
		Code:     code,
		Msg:      msg,
	}
}

/*
type ErrorCode int32
const (
	ERR_SUCCESS ErrorCode = 0
	//api
	ERR_API_DEFAULT ErrorCode = 10000
	//conn
	ERR_CONN_DEFAULT         ErrorCode = 11000
	ERR_CONN_APPKEY_REQUIRED ErrorCode = 11001
	ERR_CONN_TOKEN_REQUIRED  ErrorCode = 11002
	ERR_CONN_APP_NOT_EXISTED ErrorCode = 11003
	ERR_CONN_TOKEN_ILLEGAL   ErrorCode = 11004
	ERR_CONN_TOKEN_AUTHFAIL  ErrorCode = 11005
	ERR_CONN_TOKEN_EXPIRED   ErrorCode = 11006
)
*/
