package handle

import (
	"GateWay/balancer"
	"GateWay/proxy"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

// GenHandlerFunc General entrance of handler function
func GenHandlerFunc(removePrefix bool) func(*gin.Context) {
	return func(c *gin.Context) {
		// 获取Path前缀
		suffix := c.Param("action")
		path := c.Request.URL.Path
		prefix := path[0 : len(path)-len(suffix)+1]

		// 匹配Pattern，并获取balancer，实现负载均衡
		r := c.Request
		lb := balancer.P2bMap[prefix]
		host, err := lb.Balance(proxy.GetIP(r))
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"code": 502,
				"msg":  err,
			})
		}
		lb.Inc(host)
		defer lb.Done(host)

		// 向远端服务器集群发起请求
		cli := &http.Client{}
		// 组装request

		if removePrefix {
			path = path[strings.IndexByte(path[1:], '/')+1:]
		}
		reqUrl := host + path
		// 请求方法，请求地址，请求体
		proxyReq, err := http.NewRequest(c.Request.Method, reqUrl, c.Request.Body.(io.Reader))
		if err != nil {
			fmt.Println("http.NewRequest(to target server addr):", err.Error())
			c.JSON(http.StatusNotFound, gin.H{
				"code": 502,
				"msg":  err,
			})
		}
		// 请求头
		for k, v := range c.Request.Header {
			proxyReq.Header.Set(k, v[0])
		}

		// sent proxy_req
		resp, err := cli.Do(proxyReq)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"code": 404,
				"msg":  err,
			})
		}
		defer resp.Body.Close()

		// 将 代理响应 复制给ResponseWriter
		io.Copy(c.Writer, resp.Body)
	}
}
