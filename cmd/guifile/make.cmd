setlocal
set GOARCH=386
go build -ldflags="-s -w -H windowsgui"
endlocal
