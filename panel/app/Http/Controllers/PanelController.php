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
        $sock = $this->getSock(0, 0, 0);//获取socket对象
        if ($sock != false) {
            switch ($request->input('action')) {
                case 'start': {
                    if ($this->chkUsers($request, $request->session()->get('userid'), $request->session()->get('login_token'), encrypt("startServer"), $request->input('id'))) {
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

    private function chkUsers($request, $uid, $token, $actionSign, $opeareID)
    {
        //TODO:未完持续........
        try {
            $sign = decrypt($actionSign);
        } catch (DecryptException $e) {
            return false;
        }
        $userData = DB::table('panel_users')->where('id', $uid)->first();
        if ($token === $userData->token) {
            //这里细分if是为了Log更有意义
            if (DB::table('panel_actions')->where('name', $sign)->value('permission') == $userData->permission || $userData->permission == "superadmin" || ($userData->permission == "admin" && ($userData->group_server == $opeareID) || DB::table('panel_servers_relations')->where('')) || ($userData->permission == "buyer" && $userData->group == $opeareID)) {
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
            Log::info("进行用户权限检查时发生错误：用户登录记录不正确，疑似非法登录！");
            return false;
        }
    }

    private function gate_One()
    {//审查门，用于检查各项操作是否为合法操作
        //TODO:暂时不启用本功能，将用于用户开启安全模式之后的措施
    }
}