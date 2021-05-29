package error

import (
	"go_python/pkg/codeconv"
	"google.golang.org/grpc/status"
	"net/http"
)

type APIError struct {
	Status int    `json:"status"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
}

func (e APIError) Error() string {
	return e.Msg
}

/*
contract error info of grpcutil and convert to http
*/
func HttpInfoFromGrpc(err error) (int, string) {
	if err == nil {
		return http.StatusOK, ""

	}
	// extract grpcutil error info
	e := status.Convert(err)
	grpcErrorCode := e.Code()
	errMessage := e.Message()
	httpErrorCode := codeconv.HTTPStatusFromCode(grpcErrorCode)
	return httpErrorCode, errMessage
}
