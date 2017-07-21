<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class PanelConfig extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        //面板设置
        Schema::create('panel_config', function (Blueprint $table) {
            $table->increments('id')->unique();
            $table->text('name');
            $table->text('value');
            $table->text('permission')->default('all');//默认所有人都可以访问
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
