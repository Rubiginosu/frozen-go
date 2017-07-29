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

class PanelController extends Controller
{
    public function index(Request $request)
    {
        return true;
    }

    function getsock()
    {
        $address = '127.0.0.1';
        $port = '52023';
        $sock = socket_create(AF_INET, SOCK_STREAM, SOL_TCP);
        socket_bind($sock, $address, $port);
        socket_listen($sock, 5);
        return $sock;
    }

    public function portal(Request $request)
    {
        $sock = $this->getsock();//获取socket对象
        switch ($request->input('action')) {
            case 'start': {
                $this->start($sock, $request->input('id'));
            };
                break;
            case 'stop': {
                $this->stop($sock, $request->input('id'));
            };
                break;
            case 'restart': {
                $this->restart($sock, $request->input('id'));
            };
                break;
            default:
                return false;
                break;
        }
        socket_close($sock);
    }

    function start($sock, $serid)
    {
        $msgsock = socket_accept($sock);
        $arr1 = array('Method' => 'Start', 'OpeareID' => $serid, 'Message' => '');
        $arr2 = array('Auth' => 'Test', 'Req' => $arr1);
        $msg = json_encode($arr2);
        socket_write($msgsock, $msg, strlen($msg));
        socket_close($msgsock);
    }

    function stop($sock, $serid)
    {
        $msgsock = socket_accept($sock);
        $arr1 = array('Method' => 'Stop', 'OpeareID' => $serid, 'Message' => '');
        $arr2 = array('Auth' => 'Test', 'Req' => $arr1);
        $msg = json_encode($arr2);
        socket_write($msgsock, $msg, strlen($msg));
        socket_close($msgsock);
    }

    function restart($sock, $serid)
    {
        $msgsock = socket_accept($sock);
        $arr1 = array('Method' => 'Stop', 'OpeareID' => $serid, 'Message' => '');
        $arr2 = array('Auth' => 'Test', 'Req' => $arr1);
        $msg = json_encode($arr2);
        socket_write($msgsock, $msg, strlen($msg));
        $arr1 = array('Method' => 'Start', 'OpeareID' => $serid, 'Message' => '');
        $arr2 = array('Auth' => 'Test', 'Req' => $arr1);
        $msg = json_encode($arr2);
        socket_write($msgsock, $msg, strlen($msg));
        socket_close($msgsock);
    }
}