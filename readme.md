wfile - what file is it?
========================

- `wfile` is a command like `file` of UNIX.
    - `cd cmd/wfile` and `go build`
- `guifile.exe` is Windows GUI version of wfile (output to MessageBox instead of STDOUT)
    - `cd cmd\guifile` and `make.cmd`

Support these files.

Windows Executables
-------------------
Check magic number: `MZ`

    $ wfile.exe wfile.exe
    wfile.exe: Windows CUI, Executable Image, 64bit Header

Zip Archives
------------
Check magic number: `PK\003\004`

    $ wfile.exe hoge.zip test.zip
    hoge.zip: Zip Archive,utf8-flag-off
    test.zip: Zip Archive,utf8-flag-on

Text files
----------
Check BOM, `\0` position for UTF16, whether is utf8 valid data or not.

    $ wfile.exe utf8 mbcs text.go
    utf8: UTF8,CRLF text data
    mbcs: ANSI(MBCS),CRLF text data
    text.go: ANSI(SBCS),LF text data

Others
-------

Run `wfile` without parameters to show database.

    wfile.exe - what file is it ?

    Usage: wfile.exe {filenames}

    Signature database
    for *.arx:
      MZ ... Portable Executable
    for *.brx:
      MZ ... Portable Executable
    for *.class:
      \xCA\xFE\xBA\xBE ... Java class file
    for *.dll:
      MZ ... Portable Executable
    for *.dwg:
      AC1003 ... AutoCAD EX-II
      AC1006 ... AutoCAD GX-III
      AC1009 ... AutoCAD 12,12,GX-5
      AC1012 ... AutoCAD 13
      AC1014 ... AutoCAD 14
      AC1015 ... AutoCAD 2000,2000i,2002
      AC1018 ... AutoCAD 2004,2005,2006
      AC1021 ... AutoCAD 2007,2008,2009
      AC1024 ... AutoCAD 2010,2011,2012
      AC1027 ... AutoCAD 2013,2014,2015,2016,2017
      AC1032 ... AutoCAD 2018
    for *.exe:
      MZ ... Portable Executable
    for *.gz:
      \x1F\x8B ... gzip compressed
    for *.jpeg:
      \xFF\xD8\xFF\xDB ... JPEG Image
      \xFF\xD8\xFF\xE0\x00\x10JFIF\x00\x01 ... JPEG Image
      \xFF\xD8\xFF\xEE ... JPEG Image
      \xFF\xD8\xFF\xE1 ... JPEG Image
    for *.jpg:
      \xFF\xD8\xFF\xDB ... JPEG Image
      \xFF\xD8\xFF\xE0\x00\x10JFIF\x00\x01 ... JPEG Image
      \xFF\xD8\xFF\xEE ... JPEG Image
      \xFF\xD8\xFF\xE1 ... JPEG Image
    for *.lzh:
      -lh from 2 ... LHA Archive
    for *.pdf:
      %PDF- ... PDF
    for *.png:
      \x89PNG\r\n\x1A\n ... Portable Network Graphics
    for *.zip:
      PK\x03\x04 ... ZIP Archive
    for *:
      #! ... UNIX Executables
      \xEF\xBB\xBF#! ... Broken UNIX Executables(BOM)
      \x7FELF ... ELF - Executable and Linkable Format
