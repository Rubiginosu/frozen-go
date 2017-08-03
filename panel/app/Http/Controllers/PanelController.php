<?php
namespace App\Http\Controllers;

/*
 * PanelController V0.1 Alpha
 * Author:XueluoPoi Date:2017.7.29
 * 此控制器针对进入面板后端逻辑
 */

use Log;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Cache;
use Illuminate\Support\Facades\Storage;
use App\Http\Controllers\Controller;
use Illuminate\Contracts\Encryption\DecryptException;
require_once __DIR__ . '/../../../public/FrozenGo.php';

class PanelController extends Controller
{
    public function index(Request $request)
    {
        return view('server');
    }

    public function readini()
    {
        $serverINI = DB::table('panel_config')->where('name', 'manyServers')->value('value');
        return $serverINI;
    }

    private function getSock($ip, $port, $code)
    {
        $SDK = new FrozenGo();
        if ($ip == 0 || $port == 0 || $code == 0) {
            $SDK->ip = DB::table('panel_config')->where('name', 'daemon_ip')->value('value');
            $SDK->port = DB::table('panel_config')->where('name', 'daemon_port')->value('value');
            $SDK->verifyCode = DB::table('panel_config')->where('name', 'daemon_verifyCode')->value('value');
        } else {
            $SDK->ip = $ip;
            $SDK->port = $port;
            $SDK->verifyCode = $code;
        }
        $data = $SDK->getServerList();
        if ($data[0] == -1) return false;
        else return $SDK;
    }

    public function portal(Request $request)
    {
        $sock = $this->getSock(0, 0, 0);//获取socket对象
        if ($sock != false) {
            switch ($request->input('action')) {
                case 'start': {
                    if ($this->chkUsers($request, $request->session()->get('login_token'), encrypt("startServer"), $request->input('serverid'), $request->input('playserverid'))) {
                        $this->start($sock, $request->input('playserverid'));
                        return true;
                    } else return false;
                };
                    break;
                case 'stop': {
                    if ($this->chkUsers($request, $request->session()->get('login_token'), encrypt("stopServer"), $request->input('serverid'), $request->input('playserverid'))) {
                        $this->stop($sock, $request->input('playserverid'));
                    } else return false;
                };
                    break;
                case 'restart': {
                    if ($this->chkUsers($request, $request->session()->get('login_token'), encrypt("restartServer"), $request->input('serverid'), $request->input('playserverid'))) {
                        $this->restart($sock, $request->input('playserverid'));
                    } else return false;
                };
                    break;
                case 'create': {
                    if ($this->chkUsers($request, $request->session()->get('login_token'), encrypt("createServer"), $request->input('serverid'), $request->input('playserverid'))) {
                        $status = $this->createServer($sock, $request->input('playserverid'));
                        if ($status!=true) return $status;
                    } else return false;
                };
                    break;
                case 'delete': {
                    if ($this->chkUsers($request, $request->session()->get('login_token'), encrypt("deleteServer"), $request->input('serverid'), $request->input('playserverid'))) {
                        $status=$this->deleteServer($sock, $request->input('playserverid'));
                        if ($status!=true) return $status;
                    } else return false;
                };
                default:
                    return false;
                    break;
            }
        } else {
            return false;
        }
    }

    private function start($sock, $serid)
    {
        $sock->startServer($serid);
    }

    private function stop($sock, $serid)
    {
        $sock->stopServer($serid);
    }

    private function restart($sock, $serid)
    {
        $sock->startServer($serid);
        $sock->stopServer($serid);
    }

    private function createServer($sock, $serid)
    {
        $data = $sock->createServer($serid);
        if ($data == 'OK') {
            return true;
        } else {
            return $data;
        }
    }

    private function deleteServer($sock,$serid){
        $data = $sock->deleteServer($serid);
        if ($data == 'OK') {
            return true;
        } else {
            return $data;
        }
    }

    public function try_bind(Request $request)
    {
        $ip = $request->input('ip');
        $port = $request->input('port');
        $code = $request->input('code');
        $sock = $this->getSock($ip, $port, $code);
        if (!$sock) return true;
        else return false;
    }

    public function checkUser(Request $request)
    {
        /**
         * 此功能充当Panel用户组的核心结构
         * 返回true代表有权限进行本次操作，false代表无权限
         * @param opeareID：服务器ID
         * 使用规则：当操作涉及到验证用户所属服务器时务必先调用relations_server(),当用户权限仅有三种基本权限时可调用relations_users的level 2
         */
        $status = $this->chkUsers($request, $request->session()->get('login_token'), $request->input('actionSign'), $request->input('serverID'), $request->input('play_serverID'));
        return $status;
    }

    private function chkUsers($request, $token, $actionSign, $serverID, $playserverID)
    {
        try {
            $sign = decrypt($actionSign);
        } catch (DecryptException $e) {
            return false;
        }
        $userData = $this->getUser($request);
        if ($token === $userData->token) {
            //这先验证当前用户是否有权限操作当前服务器，再验证
            if ($userData->permission == "superadmin") return true;
            if ($this->relations_users($userData, $sign, $userData->permission, 1)) {
                if ($this->before_chkUsers($userData) == 2 && $this->relations_server($serverID, $playserverID, $this->getUser($request), 1)) {//不等于00代表本操作是需要验证serverid的
                    $permitID = str_random(32);
                    $sessionLog = $permitID . '.' . $sign;
                    $request->session()->put('permitID', $sessionLog);
                    DB::table('panel_actions')->where('name', $sign)->update(['lastest_user' => $userData->username, 'permitID' => $permitID]);
                    Log::info("准许用户：" . $userData->username . "进行：" . $sign . "操作！");
                    return true;
                } else if ($this->before_chkUsers($userData) == 1 && $this->relations_server($serverID, $playserverID, $this->getUser($request), 2)) {
                    $permitID = str_random(32);
                    $sessionLog = $permitID . '.' . $sign;
                    $request->session()->put('permitID', $sessionLog);
                    DB::table('panel_actions')->where('name', $sign)->update(['lastest_user' => $userData->username, 'permitID' => $permitID]);
                    Log::info("准许用户：" . $userData->username . "进行：" . $sign . "操作！");
                    return true;
                } else {
                    Log::info("进行用户权限检查时发生错误：用户（" . $userData->username . "）没有进行操作：" . $sign . "的权限！");
                    return false;
                }
            } else {
                Log::info("进行用户权限检查时发生错误：用户（" . $userData->username . "）没有进行操作：" . $sign . "的权限！");
                return false;
            }
        } else {
            Log::info("进行用户权限检查时发生错误：用户登录记录不正确，疑似非法登录！");
            return false;
        }
    }

    private function before_chkUsers($userData)
    {
        if ($userData->permission == 'admin') {
            return "1";
        } else {
            return "2";
        }
    }

    private function gate_One()
    {
        //TODO:暂时不启用本功能，将用于用户开启安全模式之后的措施
    }

    private function relations_users($userData, $action, $obj, $level)
    {
        $flag = false;
        if ($level == 1) {//多服务器模式
            //查找action操作是否在当前用户可操作范围内
            $actiondata = DB::table('panel_relations')->where('basic_action', $action)->get();
            $alldata = DB::table('panel_relations')->where('group', 1)->get();
            $arr = array();//存入所有有权限操作action操作的权限名
            foreach ($actiondata as $va) {
                $arr = array_add($va->name);
            }
            foreach ($alldata as $value) {
                foreach ($arr as $values) {
                    if ($value->permission_bind == $values) $arr = array_add($value->name);
                }
            }
            foreach ($arr as $arrva) {
                if ($userData->permission == $arrva) {
                    $flag = true;
                    break;
                }
            }
        } else {//单服务器模式
            //判断当前用户权限是否在obj权限之上（例如当前用户权限是superadmin，obj规定为admin，即通过
            if ($obj == "buyer") {
                if ($userData->permission == "buyer" || $userData->permission == "admin" || $userData->permission == "superadmin") {
                    $flag = true;
                }
            } else if ($obj == "admin") {
                if ($userData->permission == "admin" || $userData->permission == "superadmin") {
                    $flag = true;
                }
            } else if ($obj == "superadmin") {
                if ($userData->permission == "superadmin") {
                    $flag = true;
                }
            } else {
                $flag = false;
            }
        }
        return $flag;
    }

    private function relations_server($obj1, $obj2, $user, $level)
    {
        //证明这个obj2游戏服务器属于obj1主服务器，并且user也属于这个服务器(前提是进行了relations_users操作并且true）
        $flag = false;
        $serverdata = DB::table('panel_servers_relations')->where('serverid', $obj1)->get();
        if ($level == 1) {//1代表版主级别，为多服务器做准备
            foreach ($serverdata as $va) {
                if ($va->play_serverid == $obj2 && $user->group_server == $obj1 && $user->group == $obj2) {
                    $flag = true;
                    break;
                }
            }
        } else if ($level == 2) {
            if ($obj1 == $user->group_server) $flag = true;
        } else {
            //单服务器模式
            if ($obj1 == $user->group) $flag = true;
        }
        if ($flag) return true;
        else return false;
    }

    private function getUser(Request $request)
    {
        $userid = $request->session()->get('userid');
        $data = DB::table('panel_users')->where('id', $userid)->first();
        return $data;
    }
}