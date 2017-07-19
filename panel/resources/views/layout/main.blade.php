<?php
/*本文件为页面框架，不可单独使用
 *采用semantic-ui 2.2.7为主要UI
 */
?>
<!DOCTYPE html>
<html lang="en">
    <meta charset="utf_8">
    <meta name="csrf-token" content="{{ csrf_token() }}">
    <title>Xpoi - @yield('title')</title>
    <script src="//cdn.bootcss.com/jquery/3.1.1/jquery.js"></script>
    <link rel="stylesheet" type="text/css" href="//cdn.bootcss.com/semantic-ui/2.2.10/semantic.css">
    <script src="//cdn.bootcss.com/semantic-ui/2.2.10/semantic.min.js"></script>

    <body>
        <div class="ui container">
            @yield('content')
        </div>
    </body>
</html>