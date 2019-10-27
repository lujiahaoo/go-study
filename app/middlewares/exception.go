package middlewares

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/lujiahaoo/go-study/app/utils"
	"github.com/lujiahaoo/go-study/config"
)

func Exception(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("异常捕获发送邮件启动时间:%i", time.Now().Unix())
				DebugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					DebugStack += v + "<br>"
				}

				subject := "【重要错误】 来自项目——" + utils.ENV("appname").(string)
				body := strings.ReplaceAll(MailTemplate, "{ErrorMsg}", fmt.Sprintf("%s", err))
				body = strings.ReplaceAll(body, "{RequestTime}", time.Now().Format("2006-01-02 15:04:05"))
				body = strings.ReplaceAll(body, "{RequestURL}", r.Method+"  "+r.Host+r.RequestURI)
				body = strings.ReplaceAll(body, "{RequestUA}", r.UserAgent())
				body = strings.ReplaceAll(body, "{DebugStack}", DebugStack)

				receiver := config.EmailReceiver

				//查看gomail源码可以看到它发送邮件时传入参数时可以一次性传入多个邮件接收者，不过在send()函数里还是range一个一个发送，这样发送太慢了
				//于是改成了利用goroutine异步发送
				//测试群发了21人，利用goroutine的话不到一秒就发送完毕，比不用快了6秒多
				for _, value := range receiver {
					go utils.SendMail(value, subject, body)
				}
				fmt.Printf("异常捕获发送邮件结束时间:%s", time.Now().Unix())
				utils.ResponseWithJson(
					w,
					utils.Response{Code: http.StatusBadRequest, Msg: "系统故障,请联系管理员"},
					http.StatusBadRequest,
				)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
