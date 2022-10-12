package handlers

import (
	"context"
	"fmt"
	"net/http"
	pirpb "simplepir/grpc_gen/pir"
	"simplepir/internal/pirclient/rpc"
	"simplepir/third_party/forked/pir"
	"strconv"

	"github.com/gin-gonic/gin"
)

var hint pir.Msg

func Refresh(c *gin.Context) {
	fmt.Print("11111111111111111111111")
	resp, err := rpc.Refresh(context.Background(), &pirpb.PirRefreshRequest{})
	for k, v := range resp.Data {
		hint.Data[k].Cols = v.Cols
		hint.Data[k].Rows = v.Rows
		for o, p := range v.Data {
			hint.Data[k].Data[o] = pir.Elem(p)
		}

	}
	fmt.Print(hint)
	fmt.Print("22222222222222222222222")
	error_type := "刷新失败"
	if err != nil {
		fmt.Print(err)
		error_type = ""

	}
	c.HTML(http.StatusOK, "pir.html", gin.H{
		"error_type": error_type,
	})
}
func QueryPirBoundary(c *gin.Context) {
	fmt.Print("aaa")
	error_type := ""
	c.HTML(http.StatusOK, "pir.html", gin.H{
		"error_type": error_type,
	})
}

func QueryPir(c *gin.Context) {
	var QueryVar PirQueryParam
	s := c.PostForm("pirname")
	a, _ := strconv.Atoi(s)
	QueryVar.PhoneNumber = uint64(a)
	_, err := rpc.QueryPir(context.Background(), &pirpb.PirQueryRequest{
		PhoneNumber: QueryVar.PhoneNumber,
	})
	error_type := "存在"
	if err != nil {
		error_type = "不存在"
	}

	c.HTML(http.StatusOK, "pir.html", gin.H{
		"error_type": error_type,
	})
}
