setlocal
goto :"%1"
endlocal
exit /b

:""
    set GOARCH=386
    go build -ldflags="-s -w"
    exit /b

:"upgrade"
    for /F %%I in ('where wfile.exe') do copy /-Y wfile.exe "%%I"
    exit /b
