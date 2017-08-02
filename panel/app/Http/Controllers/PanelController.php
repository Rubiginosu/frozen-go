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
        //操作这里需要至少buyer的权限（同play_server_id)
        $sock = $this->getSock(0, 0, 0);//获取socket对象
        if ($sock != false) {
            switch ($request->input('action')) {
                case 'start': {
                    if ($this->chkUsers($request, $request->session()->get('userid'), $request->session()->get('login_token'), encrypt("startServer"), $request->input('serverid'))) {
                        $this->start($sock, $request->input('id'));
                        return true;
                    } else return false;
                };
                    break;
                case 'stop': {
                    if ($this->chkUsers($request, $request->session()->get('userid'), $request->session()->get('login_token'), encrypt("startServer"), $request->input('id'))) {
                        $this->stop($sock, $request->input('id'));
                    } else return false;
                };
                    break;
                case 'restart': {
                    if ($this->chkUsers($request, $request->session()->get('userid'), $request->session()->get('login_token'), encrypt("startServer"), $request->input('id'))) {
                        $this->restart($sock, $request->input('id'));
                    } else return false;
                };
                    break;
                case 'create': {
                    if ($this->chkUsers($request, $request->session()->get('userid'), $request->session()->get('login_token'), encrypt("startServer"), $request->input('id'))) {
                        $status = $this->createServer($sock, $request->input('id'));
                        if (!$status) return $status;
                    } else return false;
                };
                    break;
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
        $data = $this->createServer($serid);
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
         * @params opeareID：服务器ID
         */
        $status = $this->chkUsers($request, $request->session()->get('userid'), $request->session()->get('login_token'), $request->input('actionSign'), $request->input('serverID'));
        return $status;
    }

    private function chkUsers($request, $uid, $token, $actionSign, $serverID, $playserverID)
    {
        //TODO:未完持续........
        try {
            $sign = decrypt($actionSign);
        } catch (DecryptException $e) {
            return false;
        }
        $userData = $this->getUser($request);
        if ($token === $userData->token) {
            //这先验证当前用户是否有权限操作当前服务器，再验证
            if ($this->relations_users($sign, $userData->permission)) {
                if ($serverID != "00" && $playserverID != "00" && $this->relations_server($serverID, $playserverID, $this->getUser($request), 1)) {//不等于00代表本操作是需要验证serverid的
                    $permitID = str_random(32);
                    $sessionLog = $permitID . '.' . $sign;
                    $request->session()->put('permitID', $sessionLog);
                    DB::table('panel_actions')->where('name', $sign)->update(['lastest_user' => $userData->username, 'permitID' => $permitID]);
                    Log::info("准许用户：" . $userData->username . "进行：" . $sign . "操作！");
                    return true;
                } else if ($serverID != "00" && $playserverID == "00" && $this->relations_server($serverID, $playserverID, $this->getUser($request), 2)) {
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

    private function gate_One()
    {//审查门，用于检查各项操作是否为合法操作
        //TODO:暂时不启用本功能，将用于用户开启安全模式之后的措施
    }

    private function relations_users($action, $obj)
    {
        //绘制一个恶心的关系图,证明obj是不是action操作的儿子.首先提取所有有权限对action进行操作的权限名
        $actiondata = DB::table('panel_relations')->where('basic_action', $action)->get();
        $alldata = DB::table('panel_relations')->where('group', 1)->get();
        $arr = array();
        foreach ($actiondata as $va) {
            $arr = array_add($va->name);
        }
        foreach ($alldata as $value) {
            foreach ($arr as $values) {
                if ($value->permission_bind == $values) $arr = array_add($value->name);
            }
        }
        $perdataObj2 = DB::table('panel_relations')->where([['group', '1'], ['name', $obj]])->first();
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
        } else {
            if ($obj1 == $user->group_server) $flag = true;
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