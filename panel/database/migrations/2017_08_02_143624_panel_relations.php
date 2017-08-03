<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class PanelRelations extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('panel_relations', function (Blueprint $table) {
            //记录规则：所有新增记录均不填写自定义数据，一旦拥有自定义数据就一定是继承记录
            $table->increments('id')->unique();
            $table->integer('group');//关系所属类别（1为新增记录，2为继承记录）
            $table->text('permitID')->nullable();//关系记录授权证书ID（若无ID则视为无效记录）
            $table->text('name');//权限名（superadmin和admin以及buyer为基本权限不可覆盖也不可同名）
            $table->text('permission_bind')->nullable();//emmmmm在这里可以填写已有的权限名，填写后将实现对应权限名已有的所有权限
            $table->text('permission_grade')->nullable();//快捷通道，代表权限等级，只能填写(super,admin,buyer)，分别代表：最高权限，中等权限，最低权限
            //接下来是自定义数据
            $table->text('basic_action')->nullable();//填写系统内的权限名（通常是填写actions表内的操作名）
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
