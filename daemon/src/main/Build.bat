::::::::::::::::::::::::::::::
:: 内部自用编译发布脚本
:: Powered by Axoford12
::::::::::::::::::::::::::::::
@echo off
mkdir ~TEMP
color 3b
:: Build win 32
set GOOS=windows
set GOARCH=386
go build frozen.go
move .\frozen.exe .\~TEMP\
rename .\~TEMP\frozen.exe frozen_windows_386.exe
:: Build win 64
set GOARCH=amd64
go build frozen.go
move .\frozen.exe .\~TEMP\
rename .\~TEMP\frozen.exe frozen_windows_amd64.exe
:: Build linux 64
set GOOS=linux
go build frozen.go
move .\frozen .\~TEMP\
rename .\~TEMP\frozen frozen_linux_amd64

:: Build linux 32
set GOARCH=386
go build frozen.go
move .\frozen .\~TEMP\
rename .\~TEMP\frozen frozen_linux_386

7z a frozengo.7z .\~TEMP\*

rmdir /s /q .\~TEMP
