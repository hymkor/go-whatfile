wfile - what file is it?
========================

`wfile` is a command like `file` of UNIX.

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

Misc.
-----

Run `wfile` without parameters to show database.

    wfile.exe - what file is it ?

    Usage: wfile.exe {filenames}

    Signature database
    for *.class:
      [202 254 186 190] ... Java class file
    for *.dwg:
      [65 67 49 48 48 51] ... AutoCAD EX-II
      [65 67 49 48 48 54] ... AutoCAD GX-III
      [65 67 49 48 48 57] ... AutoCAD 12,12,GX-5
      [65 67 49 48 49 50] ... AutoCAD 13
      [65 67 49 48 49 52] ... AutoCAD 14
      [65 67 49 48 49 53] ... AutoCAD 2000,2000i,2002
      [65 67 49 48 49 56] ... AutoCAD 2004,2005,2006
      [65 67 49 48 50 49] ... AutoCAD 2007,2008,2009
      [65 67 49 48 50 52] ... AutoCAD 2010,2011,2012
      [65 67 49 48 50 55] ... AutoCAD 2013,2014,2015,2016,2017
      [65 67 49 48 51 50] ... AutoCAD 2018
    for *.exe:
      [77 90] ... Portable Executable
    for *.gz:
      [31 139] ... gzip compressed
    for *.jpeg:
      [255 216 255 219] ... JPEG Image
      [255 216 255 224 0 16 74 70 73 70 0 1] ... JPEG Image
      [255 216 255 238] ... JPEG Image
      [255 216 255 225] ... JPEG Image
    for *.jpg:
      [255 216 255 219] ... JPEG Image
      [255 216 255 224 0 16 74 70 73 70 0 1] ... JPEG Image
      [255 216 255 238] ... JPEG Image
      [255 216 255 225] ... JPEG Image
    for *.lzh:
      [45 108 104] from 2 ... LHA Archive
    for *.pdf:
      [37 80 68 70 45] ... PDF
    for *.png:
      [137 80 78 71 13 10 26 10] ... Portable Network Graphics
    for *.zip:
      [80 75 3 4] ... ZIP Archive
