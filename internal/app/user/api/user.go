package api

import (
	"context"
	"github.com/gin-gonic/gin"
	proto "go_python/api/user"
	"go_python/internal/pkg/error"
	"go_python/internal/pkg/grpcutil"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
	"sync"
)

var userClientConnPool *sync.Pool

func init() {
	// 初始化一个grpc连接池
	var host = "127.0.0.1"
	var port = 50050
	userClientConnPool = grpcutil.NewClientConnPool(host, port)
}

func GetUserList(ctx *gin.Context) {
	pageSize, _ := strconv.ParseUint(ctx.DefaultQuery("pageSize", "10"), 10, 32)

	pageNum, _ := strconv.ParseUint(ctx.DefaultQuery("pageNum", "1"), 10, 32)

	// 从连接池中获取一个user服务的conn，并转换为*grpcutil.ClientConn类型
	userClientConn := userClientConnPool.Get().(*grpc.ClientConn)
	// 使用完毕后放回连接池
	defer userClientConnPool.Put(userClientConn)
	// 初始化一个client
	userClient := proto.NewUserClient(userClientConn)

	userList, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		PageNum:  uint32(pageNum),
		PageSize: uint32(pageSize),
	})

	if err != nil {
		errCode, errMessage := error.HttpInfoFromGrpc(err)
		ctx.JSON(errCode, gin.H{
			"data":    "",
			"message": errMessage,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    userList,
		"message": "get user list success",
	})

}

func GetUserByMobile(ctx *gin.Context) {

}
