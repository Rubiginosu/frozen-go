<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class PanelServersRelations extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('panel_servers_relations', function (Blueprint $table) {
            $table->increments('id')->unique();
            $table->integer('group');//关系所属类别（1为新增记录，2为继承记录）
            $table->text('serverid');//主服务器id
            $table->text('play_serverid');//游戏服务器id
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        //
    }
}
