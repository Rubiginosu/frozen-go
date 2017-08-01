<?php

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/

Route::get('/', function () {
    return view('welcome');
});

Route::get('/index','FaceController@index');//将初始请求接入此处
Route::get('/firstuse','FaceController@register');
Route::post('/firstregis','FaceController@register_post');
Route::group(['prefix'=>'panel'],function() {
    Route::get('/index','PanelController@index')->middleware('checkServe');
    Route::post('/trybind','PanelController@try_bind');
});

?>
