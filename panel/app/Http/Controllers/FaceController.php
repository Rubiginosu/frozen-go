<?php
namespace App\Http\Controllers;

use Log;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Cache;
use Illuminate\Support\Facades\Storage;
use App\Http\Controllers\Controller;
use Illuminate\Contracts\Encryption\DecryptException;

class FaceController extends Controller{
    public function index(Request $request){
        //执行服务端注册登记，并注册授权
        if($this->checkServe($request)){
            //本页面等同于portal不做详细处理，直接转交panel/index，后做中间件处理
            return redirect()->action('PanelController@index');
        }else{
            $this->autoloader();//初始化
            $causes=$this->getCloseCauses();//获取失败原因
            return view('FirstUse',['causes'=>$causes]);
        }
    }
    function checkServe($request){
        //对远程服务器验证当前服务端是否合法，以及验证远程服务器信号
        if(DB::table('panel_config')->get()==null) return false;//如果是第一次使用就直接false
        $appid=DB::table('panel_config')->where('name','APPID')->first();
        if(!empty($appid)){
            $url="http://panel.dev/core/verify";//验证服务器，这里是测试用
            $key=$this->encrypt_self($appid,$request->getClientIp(),str_random(10),date("YmdHis"));
            $url_data=file_get_contents($url.'/'.$key);
        }else $status=false;
        return $status;
    }
    function encrypt_self($appid,$ip,$random,$date){
        //sign加密
    }
}
?>