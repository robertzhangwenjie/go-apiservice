package grpcutil

import (
	"fmt"
	"go_python/internal/pkg/logging"
	"google.golang.org/grpc"
	"sync"
)

func NewClientConnPool(host string, port int) *sync.Pool {
	grpcClientConnPool := &sync.Pool{
		New: func() interface{} {
			userClientConn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure())
			if err != nil {
				logging.Fatalf("connect server %s:%d failed with err:\n%v", host, port, err)
			}
			return userClientConn
		},
	}
	return grpcClientConnPool
}
