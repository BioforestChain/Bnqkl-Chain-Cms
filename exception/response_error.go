package exception

import (
	"fmt"
	"regexp"
)

var msgReg *regexp.Regexp

func init() {
	msgReg = regexp.MustCompile(`\{([^\\]+?)\}`)
}

type ResponseError struct {
	Code    string `json:"code"`    // 错误码
	Message string `json:"message"` // 错误信息
}

func (responseError *ResponseError) Error() string {
	return fmt.Sprintf("Code: %s, Message: %s", responseError.Code, responseError.Message)
}

func NewExceptionWithoutParam(errorCode *ErrorCode) *ResponseError {
	return &ResponseError{
		Code:    errorCode.Code,
		Message: errorCode.Message,
	}
}

func NewExceptionWithParam(errorCode *ErrorCode, paramMap map[string]string) *ResponseError {
	message := msgReg.ReplaceAllStringFunc(errorCode.Message, func(key string) string {
		return paramMap[key[1:len(key)-1]]
	})
	return NewExceptionWithoutParam(NewErrorCode(errorCode.Code, message))
}
