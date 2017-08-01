<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class PanelActions extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('panel_actions', function (Blueprint $table) {
            $table->increments('id')->unique();
            $table->text('name');
            $table->text('chnName');//中文名的操作
            $table->text('permission')->default('standard');//用户所需权限（superadmin,admin,buyer)
            $table->text('lastest_user')->nullable();//最后调用本操作的用户
            $table->text('permitID')->nullable();//用户操作合法的证书ID
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
