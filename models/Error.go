package models

//{"error_no":210,"error_code":"ximalaya.common.app_validate_failed","error_desc":"client_id is invalid, or app is inactive."}
import (
	"encoding/json"
	"fmt"
)

type Error struct {
	ErrorNo   int    `json:"error_no"`
	ErrorCode string `json:"error_code"`
	ErrorDesc string `json:"error_desc"`
}

func (r Error) String() {

	fmt.Printf("[ error_no:%d, error_code:%s, error_desc:%s]", r.ErrorNo, r.ErrorCode, r.ErrorDesc)
}

func (r Error) JsonDecode(body []byte) Error {

    json.Unmarshal(body, &r)

	return r
}
