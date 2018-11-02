setlocal
set GOARCH=386
go build -ldflags="-H windowsgui"
endlocal
