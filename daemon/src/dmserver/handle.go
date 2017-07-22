package dmserver

import (
	"net"
	"encoding/json"
	"io/ioutil"
	"io"
	"auth"
	"conf"
)

var serverSaved []ServerLocal
var servers []ServerRun
func connErrorToExit(errorInfo string, c net.Conn) {
	res, _ := json.Marshal(Response{-1, errorInfo})
	c.Write(res)
	c.Close()
}

// 保存服务器信息
func saveServerInfo(config conf.Config) error {
	b, err := json.Marshal(serverSaved)
	if err != nil {
		return err
	}
	ioutil.WriteFile(config.ServerManagerConfig.Servers, b, 0664)
	return nil
}

// 处理本地命令

func handleConnection(c net.Conn,config conf.Config,pairs []auth.ValidationKeyPairTime) {
	buf := make([]byte, config.DaemonServerConfig.DefaultBufLength)
	length, _ := c.Read(buf)
	request := InterfaceRequest{}
	err := json.Unmarshal(buf[:length], &request)
	if err != nil {
		connErrorToExit(err.Error(),c)
	}
	if request.Req.Method == "GetInput" {
		if ioCheck(pairs,request,c) {
			for{
				io.Copy(servers[request.Req.OperateID].Stdin,c)
			}
		}

	} else if request.Req.Method == "GetOutput" {
		if ioCheck(pairs,request,c) {
			for {
				io.Copy(c,servers[request.Req.OperateID].Stdout)
			}
		}
	} else if auth.Auth([]byte(request.Auth),[]byte(config.DaemonServerConfig.VerifyCode)){
		res,_ := json.Marshal(handleRequest(request.Req,pairs,config))
		c.Write(res)
	} else {
		connErrorToExit("Auth failed.",c)
	}

}
// 命令处理器
func handleRequest(request Request, pairs []auth.ValidationKeyPairTime,config conf.Config) Response {
	switch request.Method {

	case "List":
		return outputListOfServers()
	case "Create":
		serverSaved = append(serverSaved, ServerLocal{len(serverSaved), request.Message, "", 0})
		serverSaved[len(serverSaved)-1].EnvPrepare()
		// 序列化b来储存。
		b, err := json.MarshalIndent(serverSaved, "", "\t")

		// 新创建的服务器写入data文件
		err2 := ioutil.WriteFile(config.ServerManagerConfig.Servers, b, 0666)
		if err2 != nil {
			return Response{
				-1,
				err2.Error(),
			}
		}
		if err != nil {
			return Response{
				-1,
				err.Error(),
			}
		}
		return Response{
			0,
			"OK",
		}
	case "Start":
		// 运行这个服务器
		if request.OperateID > len(serverSaved)-1 {
			return Response{
				-1, "Invalid server id",
			}
		}
		err := serverSaved[request.OperateID].Start()
		if err == nil {
			return Response{
				0, "OK",
			}
		} else {
			return Response{-1, err.Error()}
		}

	case "Stop":
		if request.OperateID > len(servers)-1 {
			return Response{0, "Invalid serverid"}
		}
		servers[request.OperateID].Close()

		return Response{
			0, "OK",
		}

	case "SetExecutable":
		serverSaved[request.OperateID].Executable = request.Message
		return Response{
			0, "OK",
		}

	case "GetPairs":
		if request.OperateID > len(serverSaved)-1 {
			return Response{
				-1, "Invalid server id",
			}
		}
		for i := 0; i < len(pairs); i++ {
			if pairs[i].ValidationKeyPair.ID == request.OperateID {
				responseData, _ := json.Marshal(pairs[i].ValidationKeyPair)
				return Response{
					0, string(responseData),
				}
			}
		}
		// 未找到已经存在的ValidationKey
		// 为请求者生成ValidationKey
		responseData, _ := json.Marshal(auth.ValidationKeyGenerate(request.OperateID))
		return Response{0, string(responseData)}
	}
	return Response{
		-1, "Unexpected err",
	}
}
