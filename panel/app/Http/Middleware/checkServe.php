<?php

namespace App\Http\Middleware;

use Closure;
use Illuminate\Support\Facades\DB;

class checkServe
{
    /**
     * 用于查询panel是否合法
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \Closure  $next
     * @return mixed
     */
    public function handle($request, Closure $next)
    {
        $appid=DB::table('panel_config')->where('name','APPID')->value('value');
        $adminname=DB::table('panel_config')->where('name','adminname')->value('value');
        if(empty($appid)||empty($adminname)){
            return redirect()->action('FaceController@index');
        }
        return $next($request);
    }
}
