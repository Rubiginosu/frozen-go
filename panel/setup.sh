#!/bin/sh
echo "Welcome to Frozen-GO panel setup bash!"
echo "Author:XueluoPoi Date:2017-7-18 Version:V1.0 Beta"
echo "This setup bash will help you set some config correct,so please make sure your input is correct!"
echo -n "Setting files permission......"
chmod -R 777 storage
chmod -R 777 bootstrap/cache
chown -R apache:apache ../panel
if [ "$?" != "0" ];
   then
   echo "Please install Apache 2.0!"
   else
   echo "Success!"
echo "Installing some important via composer......"
composer install
if [ "$?" != "0" ];
   then
   echo "Please install Composer!"
   else
   echo "Success!"
   echo "Wait a minute....."
   cp .env.example .env
   php artisan key:generate
   echo "Finish!Now you can open the config/database.php to set information about your database!"
   fi
fi
exit 1;
