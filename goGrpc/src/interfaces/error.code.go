package interfaces

// Error Code
const (
	NOT_DEFINED_ERROR            string = "not defined error. check error stack trace"
	NOT_EXISTS_METADATA          string = "metadata is empty"
	INVALID_TOKEN                string = "token is invalid"
	EXPIRED_SESSION              string = "token has expired"
	UNDEFINED_ISSUER             string = "undefined issue request"
	FAILED_V2_LOCAL_A_GENERATOR  string = "PASETO V2 Local AccessToken Generate Error"
	FAILED_V2_LOCAL_R_GENERATOR  string = "PASETO V2 Local RefreshToken Generate Error"
	UNSUPPORTED_ISSUANCE_VERSION string = "This issuance request is no longer supported"
	TERMINATED_RPC_FUNCTION      string = "This feature is no longer supported"
)

// Error Stack Trace Message
const (
	UNSIGNED_ISSUANCE_STATUS_0 string = "error stack trace for \"crypt_issuance_status\" is 0"
	UNSIGNED_ISSUANCE_STATUS_3 string = "error stack trace for \"crypt_issuance_status\" is 3"
	NONE                       string = "none"
)
