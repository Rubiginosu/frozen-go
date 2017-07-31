@extends('layout.main')

@section('title','系统设置')

@section('content')

    <br /><br />
    <h2 class="ui dividing header">
        <i class="settings icon"></i>
        <div class="content">
            首次设置
            <div class="sub header">接下来将设置一些重要内容，这个过程不会很长</div>
        </div>
    </h2>
    <br />
    <div class="ui negative message">
        <i class="close icon"></i>
        <div class="header">
            重要提示！
        </div>
        <p>本页面可多次打开重复设置，请记录本页面域名：{{ url('/firstuse') }}
        </p>
    </div>
    <div class="errors"></div>
    <h3 class="ui top attached header">
        管理员设置
    </h3>
    <form class="ui settings form">
        {{ csrf_field() }}
    <div class="ui attached segment">
        @if(!$issetuser)
            <div class="field">
                <label>用户名</label>
                <input type="text" name="username" placeholder="设置管理员用户名">
            </div>
            <div class="field">
                <label>密码</label>
                <input type="password" name="password" placeholder="设置管理员密码">
            </div>
    </div>
        @else
            <p>管理员信息请登录面板后台后修改！</p>
        @endif
    <h3 class="ui attached header">
        Daemon信息设置
    </h3>
    <div class="ui attached segment">
        <div class="field">
            <label>Daemon连接IP</label>
            <input type="text" name="ip" placeholder="请输入安装了Daemon的服务器ip，本机请填127.0.0.1">
        </div>
        <div class="field">
            <label>Daemon连接端口</label>
            <input type="number" name="port" placeholder="请输入Daemon服务器的端口，默认为52023">
        </div>
        <div class="field">
            <label>Daemon鉴权Code</label>
            <input type="text" name="code" placeholder="请输入Daemon的唯一鉴权Code，一般在您安装daemon时可以得到">
        </div>
    </div>
    <h3 class="ui attached header">
        其他设置
    </h3>
    <div class="ui attached segment">
        <p>以下设置仅限专业人员改动，如您不清楚以下设置项有何用处请勿动！</p>
        <div class="ui segment">
            <div class="field">
                <div class="ui toggle checkbox">
                    <input type="checkbox" name="manyServer" tabindex="0" class="hidden">
                    <label>多服务器模式（不开启时默认单服务器模式）</label>
                </div>
            </div>
        </div>
        <div class="ui segment">
            <div class="field">
                <div class="ui toggle checkbox">
                    <input type="checkbox" name="safe" tabindex="0" class="hidden">
                    <label>安全模式（开启后将提升服务安全性，但部分功能可能受限）</label>
                </div>
            </div>
        </div>
    </div>
    <h3 class="ui bottom attached header">
        <button class="ui primary button" type="submit">保存设置</button>
    </h3>
    </form>
    <br />

    <script type="text/javascript">
        $(function(){
            $('.ui.checkbox').checkbox();
            $.fn.api.settings.api={
                'trybind':'{{ url('/firstregis') }}'
            };
            $('.settings.form').form({
                inline:'true',
                fields:{
                    username:{
                        identifier:'username',
                        rules:[{
                            type:'empty',
                            prompt:'用户名不得为空！'
                        }]
                    },
                    password:{
                        identifire:'password',
                        rules:[{
                            type:'empty',
                            prompt:'密码不得为空！'
                        }]
                    },
                    ip:{
                        identifire:'ip',
                        rules:[{
                            type:'empty',
                            prompt:'Daemon连接ip不得为空！'
                        }]
                    },
                    port:{
                        identifire:'port',
                        rules:[{
                            type:'empty',
                            prompt:'Daemon连接端口不得为空！'
                        }]
                    },
                    code:{
                        identifire:'code',
                        rules:[{
                            type:'empty',
                            prompt:'Daemon鉴权Code不得为空！'
                        }]
                    },
                }
            }).api({
                action: 'trybind',
                serializeForm: true,
                method:'POST',
                beforeXHR: function(xhr){
                    xhr.setRequestHeader('X-CSRF-TOKEN',$('meta[name="csrf-token"]').attr('content'));
                    return xhr;
                },
                onSuccess: function(response){
                    if(response.success==true){
                        window.location = '{{ url('/panel/index') }}';
                    }else{
                        $('.errors').html('<div class="ui red message">'+response.message+'</div>');
                    }
                },
                onFailure: function(response){
                    $('.errors').html('<div class="ui red message">'+response.message+'</div>');
                },
                onError: function(response){
                    $('.errors').html('<div class="ui red message">无法连接到Daemon!请检查Daemon信息是否正确！</div><br /><p>若提示underfined即为无效链接，请检查Daemon连接信息！</p>');
                }
            });
        });
    </script>
@endsection