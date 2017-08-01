<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class PanelUsers extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('panel_users', function (Blueprint $table) {
            $table->increments('id')->unique();
            $table->text('username');
            $table->text('password');
            $table->text('email')->nullable();
            $table->text('group_server')->default('unknown');//所属服务器uid（版主账户需要此项）
            $table->text('group')->default('unknown');//所属游戏服务器uid(仅买家账户需要此项（腐竹））
            $table->text('lastest_ip')->nullable();
            $table->boolean('black_list')->default(false);
            $table->boolean('double_check')->default(false);
            $table->text('token')->nullable();
            $table->text('permission')->default('standard');//用户权限（superadmin,admin,buyer)
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
