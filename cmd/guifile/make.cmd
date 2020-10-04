@echo off
setlocal
goto :"%1"
endlocal

:""
    set GOARCH=386
    if not exist rsrc.syso rsrc -manifest test.manifest -o rsrc.syso
    go build -ldflags="-s -w -H windowsgui"
    exit /b

:"upgrade"
    for /F %%I in ('where guifile.exe') do copy /-Y guifile.exe "%%I"
    exit /b
