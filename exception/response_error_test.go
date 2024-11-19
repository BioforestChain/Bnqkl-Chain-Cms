package exception

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewResponseError(t *testing.T) {
	errorResp := NewExceptionWithParam(USER_IS_NOT_EXISTED, map[string]string{
		"loginName": "123",
	})
	jsonData, err := json.Marshal(errorResp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(jsonData))
}
