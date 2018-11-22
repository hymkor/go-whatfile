@setlocal
@set "PROMPT=$G "
@call :"%1"
@endlocal
@exit /b

:""
:"all"
    set GOARCH=386
    go fmt || exit /b
    pushd "%~dp0cmd\wfile"
        go fmt
        go build -o "%~dp0wfile.exe" -ldflags="-s -w" || (popd & exit /b)
    popd
    pushd "%~dp0cmd\guifile"
        go fmt
        go build -o "%~dp0guifile.exe" -ldflags="-s -w"
    popd
    exit /b
