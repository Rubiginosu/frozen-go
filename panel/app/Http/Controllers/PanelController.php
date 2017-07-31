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

    private function getSock($ip,$port,$code)
    {
        $SDK=new FrozenGo();
        if($ip==0||$port==0||$code==0){
            $SDK->ip=DB::table('panel_config')->where('name','daemon_ip')->value('value');
            $SDK->port=DB::table('panel_config')->where('name','daemon_port')->value('value');
            $SDK->verifyCode=DB::table('panel_config')->where('name','daemon_verifyCode')->value('value');
        }else{
            $SDK->ip=$ip;
            $SDK->port=$port;
            $SDK->verifyCode=$code;
        }
        $data=$SDK->getServerList();
        if($data[0]==-1) return false;
        else return $SDK;
    }

    public function portal(Request $request)
    {
        $sock = $this->getSock(0,0,0);//获取socket对象
        if($sock!=false){
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
                case 'create': {
                    $status=$this->createServer($sock, $request->input('id'));
                    if(!$status) return $status;
                };
                    break;
                default:
                    return false;
                    break;
            }
        }else{
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
    private function createServer($sock, $serid){
        $data=$this->createServer($serid);
        if($data=='OK'){
            return true;
        }else{
            return $data;
        }
    }

    public function try_bind(Request $request){
        $ip=$request->input('ip');
        $port=$request->input('port');
        $code=$request->input('code');
        $sock=$this->getSock($ip,$port,$code);
        if(!$sock) return true;
        else return false;
    }
}