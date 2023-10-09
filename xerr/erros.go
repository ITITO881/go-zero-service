package xerr

import (
	"fmt"
)

/**
常用通用固定错误
*/

type CodeError struct {
	code uint32
	msg  string
}

// GetErrCode 返回给前端的错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.code
}

// GetErrMsg 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.msg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("code:%d，msg:%s", e.code, e.msg)
}
