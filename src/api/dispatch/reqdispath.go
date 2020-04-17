package dispatch

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Request 请求数据
type Request struct {
	Msisdn          int    `json:"msisdn" form:"msisdn" binding:"required"`
	Amountrequested string `json:"amountrequested" form:"amountrequested"`
	TransactionID   int    `json:"transactionId" form:"transactionId" binding:"required"`
	RequestType     string `json:"requestType" form:"requestType"`
	ChannelType     string `json:"channelType" form:"channelType" binding:"required"`
	TransactionType string `json:"transactionType" form:"transactionType" binding:"required"`
	PendingLoan     string `json:"pendingLoan" form:"pendingLoan"`
	RecordCount     string `json:"recordCount" form:"recordCount"`
}

// Resp 响应数据
type Resp struct {
	Responsecode   int    `json:"responsecode"`
	Responsestring string `json:"responsestring"`
}

// Reqdispath moreCredit Service Provider
func Reqdispath(c *gin.Context) {
	var (
		r Request
		q Resp
	)

	// 接收数据，做数据校验
	err := c.ShouldBind(&r)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	// TODO
	// 继续写业务逻辑

	// 响应数据
	q.Responsecode = 0
	q.Responsestring = "Request received successfully"

	// 输出结果
	c.JSON(http.StatusOK, r)
	return
}
