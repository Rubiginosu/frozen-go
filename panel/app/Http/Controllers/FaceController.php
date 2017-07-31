<?php
namespace App\Http\Controllers;

/*
 * FaceController V0.1 Alpha
 * Author:XueluoPoi Date:2017.7.22
 * 此控制器针对进入面板前的处理任务
 */

use Log;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Cache;
use Illuminate\Support\Facades\Storage;
use App\Http\Controllers\Controller;
use Illuminate\Contracts\Encryption\DecryptException;
require_once __DIR__ . '/../../../public/FrozenGo.php';

class FaceController extends Controller
{
    public function index(Request $request)
    {
        $request->session()->forget('fail_cause');
        if (DB::table('panel_config')->first() == null) {
            $this->autoloader_first($request);
            $request->session()->put('fail_cause', '0001');
            return view('first_regis', ['issetuser' => false]);
        } else {
            if ($this->checkServe($request)) {
                $this->autoloader($request);
                return redirect()->action('PanelController@index');
            } else {
                $causes = $this->getCloseCauses($request);//获取失败原因
                return view('FirstUse', ['causes' => $causes]);
            }
        }
    }

    function checkServe($request)
    {
        if (DB::table('panel_config')->where('name', 'checkServe')->value('value') != true) return true;
        $appid = DB::table('panel_config')->where('name', 'APPID')->value('value');
        if (!empty($appid) && DB::table('panel_config')->where('name', 'checkServe')->value('value') == 'true') {
            echo 'yes';
            $url = "http://panel.dev/core/verify";
            $random = str_random(10);
            $key = $this->encrypt_self($appid, $request->getClientIp(), $random, date("YmdHis"));
            $url_data = file_get_contents($url . '/' . $key . '/' . $random);
            if ($url_data != "success") {
                $request->session()->put('fail_cause', $url_data);
                $status = false;
            } else $status = true;
        } else $status = true;
        return $status;
    }

    function encrypt_self($appid, $ip, $random, $date)
    {
        $str = $appid . '.' . $ip . '.' . $date . '+' . md5(md5($random));
        return $str;
    }

    function getCloseCauses($request)
    {
        $cause = "啊哦，这个错误面板酱表示一脸懵逼";
        $causesNum = $request->session()->get('fail_cause');
        if (isset($causesNum)) {
            $numBook = explode('&', Storage::get('fail_Num.txt'));
            foreach ($numBook as $va) {
                $value = explode(':', $va);
                if ($value[0] == $causesNum) {
                    $cause = $value[1];
                    break;
                }
            }
        }
        Log::debug("面板启动过程中发生错误：" . $cause);
        return $cause;
    }

    public function register()
    {
        if (DB::table('panel_config')->where('name', 'adminname')->value('value') != null) $issetuser = true;
        else $issetuser = false;
        return view('first_regis', ['issetuser' => $issetuser]);
    }

    public function register_post(Request $request)
    {
        if ($this->getSock($request->input('ip'), $request->input('port'), $request->input('code')) == true){
            DB::table('panel_config')->insert(['name' => 'adminname', 'value' => $request->input('username'), 'permission' => 'system']);
            DB::table('panel_config')->insert(['name' => 'adminpass', 'value' => encrypt($request->input('password')), 'permission' => 'system']);
            DB::table('panel_config')->insert(['name' => 'daemon_ip', 'value' => $request->input('ip'), 'permission' => 'system']);
            DB::table('panel_config')->insert(['name' => 'daemon_port', 'value' => $request->input('port'), 'permission' => 'system']);
            DB::table('panel_config')->insert(['name' => 'daemon_verifyCode', 'value' => $request->input('code'), 'permission' => 'system']);
            $status=true;
            $message='null';
        }else{
            $status=false;
            $message="请您检查Daemon连接信息是否正确！";
        }
        return response()->json(['success'=>$status,'message'=>$message]);
    }

    function autoLoader_first($request)
    {
        DB::table('panel_config')->insert(['name' => 'APPID', 'value' => str_random(32) . $request->getClientIp() . date("Ymd"), 'permission' => 'system']);
        DB::table('panel_config')->insert(['name' => 'checkServe', 'value' => 'false', 'permission' => 'system']);
        return redirect()->action('FaceController@register');
    }

    function autoloader($request)
    {
        //TODO:有待后续添加.....
        //return true;
    }

    private function getSock($ip, $port, $code)
    {
        $SDK = new \FrozenGo($ip,$port,$code);
        $data = $SDK->getServerList();
        print_r($data);
        if($data->Status==-1) return false;
        else return true;
        /**if ($SDK->SockResult('List')=="5") return false;
        else{
            if($data==-1) return false;
            else return true;
        }*/
    }
}
?>
