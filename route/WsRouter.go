package route

import (
	"fmt"

	"github.com/bytedance/sonic"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/common/BaseModel/TokenModel"
	"main.go/tuuz"
	"main.go/tuuz/RET"
)

func MainWsRouter() {
	for c := range Net.WsServer_ReadChannel {
		fmt.Println(c.Conn.RemoteAddr().String(), string(c.Message), c.Status)
		nd, err := sonic.Get(c.Message, "route")
		if err != nil {
			continue
		}
		r, err := nd.String()
		if err != nil {
			continue
		}
		switch r {
		case "login":
			_uid, err := sonic.Get(c.Message, "uid")
			if err != nil {
				fmt.Println(tuuz.FUNCTION_ALL(), err)
				break
			}
			_token, err := sonic.Get(c.Message, "token")
			if err != nil {
				fmt.Println(tuuz.FUNCTION_ALL(), err)
				break
			}
			uid, err := _uid.String()
			if err != nil {
				fmt.Println(tuuz.FUNCTION_ALL(), err)
				break
			}
			token, err := _token.String()
			if err != nil {
				fmt.Println(tuuz.FUNCTION_ALL(), err)
				break
			}
			loginData := TokenModel.Api_find(uid, token)
			if len(loginData) > 0 {
				Net.WsServer_WriteChannel <- Net.WsData{Conn: c.Conn, Message: []byte(RET.Ws_succ("login", 0, nil, nil))}
				break
			} else {
				Net.WsServer_WriteChannel <- Net.WsData{Conn: c.Conn, Message: []byte(RET.Ws_fail("login", 401, nil, "登录失败，抖店助手暂未登录，你可以重新登录后再使用"))}
				break
			}

		default:
			Net.WsServer_WriteChannel <- c
			break
		}
	}
}
