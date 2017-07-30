<?php
/**
 * Created by PhpStorm.
 * User: axoford12
 * Date: 17-7-29
 * Time: 下午9:38
 */
// 这是用于连接Daemon的SDK..
class Request{
    public $Method;
    public $OperateID;
    public $Message;
}
class InterfaceRequest{
    public $Auth;
    public $Req;
}

/*
 * 公共返回对象字段：
 * Status : 包含运行的状态
 * Message : 当状态为-1时包含错误，其余包含返回信息(可能是string也可能是另一个对象)
 * 在下面的函数注释中，我会对Message进行说明.
 */

class FrozenGo
{
    /**
     * @var $ip
     * 要连接的ip地址
     * @var $port
     * 端口
     * @var $verifyCode
     * 在Daemon中配置的验证Code.
     */
    private $ip;
    private $port;
    private $verifyCode;

    /**
     * FrozenGo constructor.
     * @param $ip
     * @param $port
     * @param $verifyCode
     */
    public function __construct($ip,$port,$verifyCode)
    {
        $this->ip = $ip;
        $this->port = $port;
        $this->verifyCode = $verifyCode;
    }

    /**
     * @return mixed|string
     * 返回服务器列表，类型是对象数组
     * 对象保存了如下字段：
     * ID: 服务器的ID
     * Name: 服务器名
     * Executable : 可执行文件配置
     * Status: 运行状态
     * UserUid: 运行时期制定用户的Uid
     */
    public function getServerList(){
        $servers = $this->SockResult("List");
        $servers->Message = json_decode($servers->Message);
        return $servers;
    }

    /**
     * @param $name
     * 服务器名。
     * @return mixed|string
     * 成功时返回"OK"
     * 失败返回错误信息
     */
    public function createServer($name){
        return $this->SockResult("Create",0,$name);
    }

    /**
     * @param $id
     * 服务器id (只能是ID)否则daemon无法解析
     * @return mixed|string
     * 返回同上
     */
    public function deleteServer($id) {
        return $this->SockResult("Delete",$id);
    }

    /**
     * @param $id
     * 这个密钥对应的服务器id
     * @return mixed|string
     * 对象，包含两个字段，
     * ValidationKeyPair : 对象，包含两个字段：
     *         ID:这个Key对应的ID
     *         Key: 密钥，20位字符
     * GeneratedTIme: 生成的时间。格式大致如下：
     * 2017-07-29T22:35:15.184376223+08:00
     */
    public function getValidationKeyPairs($id){
        $result = $this->SockResult("GetPairs",$id);
        $result->Message = json_decode($result->Message);
        return $result;
    }

    /**
     * @param $id
     * 要设置的id
     * @param $exec
     * 可执行文件名
     * @return mixed|string
     * 返回同Create.
     */
    public function setExecutable($id,$exec){
        return $this->SockResult("SetExecutable",$id,$exec);
    }
    public function execInstall($url,$id){
        return $this->SockResult("ExecInstall",$id,$url);

    }
    // TODO 尽快修好！
    // 下面两个正在调试
    // !!! With Bug...
    public function startServer($id){
        return $this->SockResult("Start",$id);
    }
    public function stopServer($id){
        return $this->SockResult("Stop",$id);
    }


    private function SockResult($method,$operateId = 0,$message = ""){
        $socket = socket_create(AF_INET, SOCK_STREAM, SOL_TCP);
        $conn = socket_connect($socket,$this->ip,$this->port);
        if($conn < 0){
            return  "5" . socket_strerror($conn);
        }
        $Req = new Request();
        $Req->Method = $method;
        $Req->OperateID = $operateId;
        $Req->Message = $message;
        $InReq = new InterfaceRequest();
        $InReq->Auth = $this->verifyCode;
        $InReq->Req = $Req;
        $sending = json_encode($InReq);
        socket_write($socket,$sending,strlen($sending));
        $result = "";
        while($resultBuf = socket_read($socket,1024)){
            $result .= $resultBuf;
        }
        socket_close($socket);
        return json_decode($result);
    }
}
