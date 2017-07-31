<!doctype html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>FrozenGo - @yield('title')</title>
  <meta name="keywords" content="index">
  <meta name="csrf-token" content="{{ csrf_token() }}">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="renderer" content="webkit">
  <meta http-equiv="Cache-Control" content="no-siteapp" />
  <link rel="icon" type="image/png" href="assets/i/favicon.png">
  <link rel="stylesheet" href="//cdn.bootcss.com/amazeui/2.7.2/css/amazeui.min.css"/>
  <link rel="stylesheet" href="//static.mcpe.cc/css/admin.css">
  <link rel="stylesheet" href="//cdn.bootcss.com/toastr.js/latest/css/toastr.css">
<body data-type="generalComponents">
@yield('content')
</body>
<!--[if lt IE 9]>
<script src="http://libs.baidu.com/jquery/1.11.1/jquery.min.js"></script>
<script src="http://cdn.staticfile.org/modernizr/2.8.3/modernizr.js"></script>
<script src="//static.mcpe.cc/assets/js/amazeui.ie8polyfill.min.js"></script>
<![endif]-->

<!--[if (gte IE 9)|!(IE)]><!-->
<script src="//cdn.bootcss.com/jquery/3.2.1/jquery.js"></script>
<!--<![endif]-->
<script src="//cdn.bootcss.com/amazeui/2.7.2/js/amazeui.min.js"></script>
<script src="//static.mcpe.cc/js/app.js"></script>
<script src="//cdn.bootcss.com/toastr.js/latest/js/toastr.min.js" type="text/javascript"></script>
</html>