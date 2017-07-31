@extends('layout.her')

@section('title','系统设置')

@section('content')
    <script>
function kick(user,serverID){
 $.ajax({
                url: "{{ url('/panel/portal/kick') }}",
                type: "post",
                data: {
                   ServerID : serverID,
                   UserName : user,
                   type : 'json',
                   Csrf_token : csrf_token 
               },
                dataType: 'json',
    success:function(data) { 
    		if(data['success'])
    		{
 toastr.success('踢出玩家'+user+'成功');
}
    		else
    		{
   toastr.error(data['ErrorMsg']);
    		}
}
})}
function kick(user,serverID){
 $.ajax({
                url: "{{ url('/panel/portal/kick') }}",
                type: "post",
                data: {
                   ServerID : serverID,
                   UserName : user,
                   type : 'json',
                   Csrf_token : csrf_token 
               },
                dataType: 'json',
    success:function(data) { 
    		if(data['success'])
    		{
 toastr.success('踢出玩家'+user+'成功');
}
    		else
    		{
   toastr.error(data['ErrorMsg']);
    		}
}
})}
function srs(button,serverID){
 $.ajax({
                url: '{{ url("/panel/portal/'+button+'") }}',
                type: "post",
                data: {
                   ServerID : serverID,
                   type : 'json',
                   Csrf_token : csrf_token 
               },
                dataType: 'json',
    success:function(data) { 
    		if(data['success'])
    		{
 toastr.success('执行操作'+button+'成功');
}
    		else
    		{
   toastr.error(data['ErrorMsg']);
    		}
}
})}
function getserinfo(){
 $.ajax({
                url: "{{ url('/panel/portal/getinfo') }}",
                type: "post",
                data: {
                   serid : 37,
                    type : 'json',
Csrf_token : 'ACOW-DSAC-1SAW-1DC3F-3AZC-ZZQ1S'                },
                dataType: 'json',
    success:function(data) { 
    		if(data['success'])
    		{
var neicun = data['maxmemory']*(data['memory']/100);
$("#CPU").attr("style","width: "+data['CPU']);
$("#CPU").html(data['CPU']);
$("#neicun").attr("style","width: "+data['memory']+'%');
$("#neicun").html(neicun+'M/'+data['maxmemory']+"M");
$("#pl1").html(data['player']+'/'+data['maxplayer']+'负载：'+data['load']);
$("#pl2").html(data['OnlineTime']);
$("#players").html('');//重置为空
$("#text1111").html('');//重置为空
$.each(data['Players'],function(i,v){
 $("#players").append('<tr><td>'+v+'</td><td><button type="button" class="am-btn am-btn-primary tpl-btn-bg-color-success " onclick="kick("'+v+data['ServerID']+'")">Kick</button></td></tr>');
})
switch(data['status'])
{
case 'start':
  $("#status").html('<i class="am-icon-check"></i>');
  break;
case 'restart':
  $("#status").html('<i class="am-icon-refresh am-icon-spin"></i>');
  break;
case 'shutdown':
  $("#status").html('<i class="am-icon-times"></i>');
  break;
default:
$("#status").html('<i class="am-icon-info"></i>');
}

			}
    		else
    		{
   toastr.error(data['data']);
    		}
     },  
    error : function() {  
                toastr.error('连接服务器出现未知错误 请联系亦欢');
    } 
            });
}
setInterval(getserinfo, 1000); 
</script>


    <header class="am-topbar am-topbar-inverse admin-header">
        <div class="am-topbar-brand">
            <a href="javascript:;" class="tpl-logo">
                <img src="assets/img/logo.png" alt="">
            </a>
        </div>
        <div class="am-icon-list tpl-header-nav-hover-ico am-fl am-margin-right">

        </div>

        <button class="am-topbar-btn am-topbar-toggle am-btn am-btn-sm am-btn-success am-show-sm-only" data-am-collapse="{target: '#topbar-collapse'}"><span class="am-sr-only">导航切换</span> <span class="am-icon-bars"></span></button>

        <div class="am-collapse am-topbar-collapse" id="topbar-collapse">

            <ul class="am-nav am-nav-pills am-topbar-nav am-topbar-right admin-header-list tpl-header-list">
                <li class="am-dropdown" data-am-dropdown data-am-dropdown-toggle>
                    <a class="am-dropdown-toggle tpl-header-list-link" href="javascript:;">
                        <span class="am-icon-bell-o"></span> 提醒 <span class="am-badge tpl-badge-success am-round">5</span></span>
                    </a>
                    <ul class="am-dropdown-content tpl-dropdown-content">
                        <li class="tpl-dropdown-content-external">
                            <h3>你有 <span class="tpl-color-success">5</span> 条提醒</h3><a href="###">全部</a></li>
                        <li class="tpl-dropdown-list-bdbc"><a href="#" class="tpl-dropdown-list-fl"><span class="am-icon-btn am-icon-plus tpl-dropdown-ico-btn-size tpl-badge-success"></span> 【预览模块】移动端 查看时 手机、电脑框隐藏。</a>
                            <span class="tpl-dropdown-list-fr">3小时前</span>
                        </li>
                        <li class="tpl-dropdown-list-bdbc"><a href="#" class="tpl-dropdown-list-fl"><span class="am-icon-btn am-icon-check tpl-dropdown-ico-btn-size tpl-badge-danger"></span> 移动端，导航条下边距处理</a>
                            <span class="tpl-dropdown-list-fr">15分钟前</span>
                        </li>
                        <li class="tpl-dropdown-list-bdbc"><a href="#" class="tpl-dropdown-list-fl"><span class="am-icon-btn am-icon-bell-o tpl-dropdown-ico-btn-size tpl-badge-warning"></span> 追加统计代码</a>
                            <span class="tpl-dropdown-list-fr">2天前</span>
                        </li>
                    </ul>
                </li>
                <li class="am-dropdown" data-am-dropdown data-am-dropdown-toggle>
                    <a class="am-dropdown-toggle tpl-header-list-link" href="javascript:;">
                        <span class="am-icon-comment-o"></span> 消息 <span class="am-badge tpl-badge-danger am-round">9</span></span>
                    </a>
                    <ul class="am-dropdown-content tpl-dropdown-content">
                        <li class="tpl-dropdown-content-external">
                            <h3>你有 <span class="tpl-color-danger">9</span> 条新消息</h3><a href="###">全部</a></li>
                        <li>
                            <a href="#" class="tpl-dropdown-content-message">
                                <span class="tpl-dropdown-content-photo">
                      <img src="assets/img/user02.png" alt=""> </span>
                                <span class="tpl-dropdown-content-subject">
                      <span class="tpl-dropdown-content-from"> 禁言小张 </span>
                                <span class="tpl-dropdown-content-time">10分钟前 </span>
                                </span>
                                <span class="tpl-dropdown-content-font"> Amaze UI 的诞生，依托于 GitHub 及其他技术社区上一些优秀的资源；Amaze UI 的成长，则离不开用户的支持。 </span>
                            </a>
                            <a href="#" class="tpl-dropdown-content-message">
                                <span class="tpl-dropdown-content-photo">
                      <img src="assets/img/user03.png" alt=""> </span>
                                <span class="tpl-dropdown-content-subject">
                      <span class="tpl-dropdown-content-from"> Steam </span>
                                <span class="tpl-dropdown-content-time">18分钟前</span>
                                </span>
                                <span class="tpl-dropdown-content-font"> 为了能最准确的传达所描述的问题， 建议你在反馈时附上演示，方便我们理解。 </span>
                            </a>
                        </li>

                    </ul>
                </li>




                <li class="am-dropdown" data-am-dropdown data-am-dropdown-toggle>
                    <a class="am-dropdown-toggle tpl-header-list-link" href="javascript:;">
                        <span class="tpl-header-list-user-nick">Yihuan</span><span class="tpl-header-list-user-ico"> <img src="assets/img/user01.png"></span>
                    </a>
                    <ul class="am-dropdown-content">
                        <li><a href="#"><span class="am-icon-bell-o"></span> 资料</a></li>
                        <li><a href="#"><span class="am-icon-cog"></span> 设置</a></li>
                        <li><a href="#"><span class="am-icon-power-off"></span> 退出</a></li>
                    </ul>
                </li>

            </ul>
        </div>
    </header>



    <div class="tpl-page-container tpl-page-header-fixed">
        <div class="tpl-left-nav tpl-left-nav-hover">
		        <div class="tpl-left-nav-title">
                <ol class="am-breadcrumb">
                <li><a href="#" class="am-icon-home">服务器列表</a></li>
                <li class="am-active">服务器</li>
            </ol>
            </div>
            <div class="tpl-left-nav-list">
                <ul class="tpl-left-nav-menu">
                    <li class="tpl-left-nav-item">
                        <a href="index.html" class="nav-link">
                            <i class="am-icon-commenting"></i>
                            <span>在线聊天</span>
                        </a>
                    </li>
                    <li class="tpl-left-nav-item">
                        <a href="index.html" class="nav-link">
                            <i class="am-icon-code"></i>
                            <span>控制台</span>
                        </a>
                    </li>
	<li class="tpl-left-nav-item">
                        <a href="index.html" class="nav-link">
                            <i class="am-icon-file"></i>
                            <span>上传文件</span>
                        </a>
                    </li>
                    <li class="tpl-left-nav-item">
                        <a href="index.html" class="nav-link">
                            <i class="am-icon-home"></i>
                            <span>定时任务</span>
                        </a>
                    </li>
                    <li class="tpl-left-nav-item">
                        <a href="index.html" class="nav-link">
                            <i class="am-icon-shopping-cart"></i>
                            <span>插件商店</span>
                        </a>
                    </li>
                    <li class="tpl-left-nav-item">
                        <a href="index.html" class="nav-link">
                            <i class="am-icon-times"></i>
                            <span>删除服务器</span>
                        </a>
                    </li>

                </ul>
            </div>
        </div>




        <div class="tpl-content-wrapper">

           
           <div class="tpl-portlet-components">
                <div class="portlet-title">
                    <div class="caption font-green bold">
                        <span class="am-icon-code"></span> Yihuan的Minecraft服务器 (ID:37)
                    </div>
                    <div class="tpl-portlet-input tpl-fz-ml">
                        <div class="portlet-input input-small input-inline">
                            <div class="input-icon ">
                             <ul class="am-nav am-nav-pills">

</ul> </div>
                        </div>
                    </div>


                </div>


                <div class="tpl-block">

                    <div class="am-g">
                        <div class="tpl-form-body tpl-form-line">
                            <form class="am-form tpl-form-line-form">
							<div class="am-form-group">
<label for="user-name" class="am-u-sm-3 am-form-label" id="status"><i class="am-icon-refresh am-icon-spin"></i></label>
<div class="am-u-sm-9 am-btn-group">
  <button type="button" class="am-btn am-btn-success am-round am-btn-xs" onclick="srs('open')">开启</button>
  <button type="button" class="am-btn am-btn-secondary am-round am-btn-xs" onclick="srs('restart')">重启</button>
  <button type="button" class="am-btn am-btn-danger am-round am-btn-xs" onclick="srs('stop')">关闭</button>
  </div>
</div>
                                <div class="am-form-group">
                                    <label for="user-name" class="am-u-sm-3 am-form-label">在线玩家 </label>
                                    <div class="am-u-sm-9">
                                        <label for="player" class="am-form-label" id="pl1"><i class="am-icon-refresh am-icon-spin"></i></label> 
                                    </div>
                                </div>
                                <div class="am-form-group">
                                    <label for="user-name" class="am-u-sm-3 am-form-label">服务器已持续运行</label>
                                    <div class="am-u-sm-9">
                                        <label for="player" class="am-form-label" id="pl2"><i class="am-icon-refresh am-icon-spin"></i></label> 
                                    </div>
                                </div>
                                <div class="am-form-group">
                                    <label for="user-name" class="am-u-sm-3 am-form-label">服务器名称</label>
                                    <div class="am-u-sm-9">
                                        <input type="text" class="tpl-form-input" id="user-name" placeholder="请输入服务器名称">
                                        <small>请填写服务器名称10字以下。</small>
                                    </div>
                                </div>

                                <div class="am-form-group">
                                    <label for="user-email" class="am-u-sm-3 am-form-label">到期时间</label>
                                    <div class="am-u-sm-9">
                                        <input type="text" class="am-form-field tpl-form-no-bg" placeholder="到期时间" data-am-datepicker="" readonly/>
                                        <small>填写后，到指定时间就可以自动停服啦~~~</small>
                                    </div>
                                </div>

                                <div class="am-form-group">
                                    <label for="user-phone" class="am-u-sm-3 am-form-label">腐竹</label>
                                    <div class="am-u-sm-9">
                                        <select data-am-selected="{searchBox: 1}">
  <option value="a">Yihuan</option>
  <option value="b">Axo12</option>
  <option value="o">雪落</option>
</select>
                                    </div>
                                </div>
																                                <div class="am-form-group">
                                    <label for="user-phone" class="am-u-sm-3 am-form-label">设定核心</label>
                                    <div class="am-u-sm-9">
                                        <select data-am-selected="{searchBox: 1}">
  <option value="a">[MCPE]1.10.1</option>
  <option value="b">[MCPC]1.7.1</option>
  <option value="o">自定义JAR</option>
</select>
                                    </div>
                                </div>

                                <div class="am-form-group">
                                    <label class="am-u-sm-3 am-form-label">服务器IP</label>
                                    <div class="am-u-sm-9">
                                        <input type="text" placeholder="输入服务器IP">
                                    </div>
                                </div>


                                <div class="am-form-group">
                                    <label for="user-weibo" class="am-u-sm-3 am-form-label">设定人数 </label>
                                    <div class="am-u-sm-9">
                                        <input type="text" id="user-weibo" placeholder="请输入整数">
                                        <div>

                                        </div>
                                    </div>
                                </div>
								                               <div class="am-form-group">
                                    <label for="user-weibo" class="am-u-sm-3 am-form-label">内存设定 </label>
                                    <div class="am-u-sm-9">
                                        <input type="text" id="user-weibo" placeholder="请输入整数">
                                        <div>

                                        </div>
                                    </div>
                                </div>



                                <div class="am-form-group">
                                    <div class="am-u-sm-9 am-u-sm-push-3">
                                        <button type="button" class="am-btn am-btn-primary tpl-btn-bg-color-success ">保存设定</button>
                                    </div>
                                </div>
								  <div class="am-form-group">
								  
                                    <div class="am-u-sm-5 ">CPU使用量：
                                        <div class="am-progress am-progress-striped  am-active ">
  <div class="am-progress-bar am-progress-bar-success"  id="CPU" style="width: 100%"><i class="am-icon-refresh am-icon-spin"></i></div>
</div>
                                    </div>
									                                    <div class="am-u-sm-5 ">内存使用量：
                                        <div class="am-progress am-progress-striped am-active ">
  <div class="am-progress-bar am-progress-bar-secondary"  id="neicun" style="width: 100%"><i class="am-icon-refresh am-icon-spin"></i></div>
</div>
                                    </div>
                                </div>
								
                            </form>

<label for="user-name" class="am-form-label">在线玩家  <span class="tpl-form-line-small-title">Player  <a href="#">实时聊天</a></span></label>
<br>
<a id="text1111"><i class="am-icon-refresh am-icon-spin"></i>加载是一件很重要的事情~~~~~</a>
<table class="am-table">
<tbody id="players">
    </tbody>
</table>
                        </div>
                    </div>
                </div>


            </div>


        </div>

    </div>

@endsection