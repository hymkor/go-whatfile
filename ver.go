package wfile

import (
	"errors"
	"syscall"
	"unsafe"
)

var versionDll = syscall.NewLazyDLL("version")
var procGetFileVersionInfoSize = versionDll.NewProc("GetFileVersionInfoSizeW")
var procGetFileVersionInfo = versionDll.NewProc("GetFileVersionInfoW")
var procVerQueryValue = versionDll.NewProc("VerQueryValueW")

type vsFixedFileInfo struct {
	Signature        uint32
	StrucVersion     uint32
	FileVersionMS    uint32
	FileVersionLS    uint32
	ProductVersionMS uint32
	ProductVersionLS uint32
	FileFlagsmask    uint32
	FileFlags        uint32
	FileOs           uint32
	FileType         uint32
	FileSubtype      uint32
	FileDateMS       uint32
	FileDateLS       uint32
}

func lower16bit(n uint32) uint {
	return uint(n & 0xFFFF)
}

func upper16bit(n uint32) uint {
	return uint(n>>16) & 0xFFFF
}

func getVersionInfo(fname string) ([]uint, error) {
	_fname, err := syscall.UTF16PtrFromString(fname)
	if err != nil {
		return nil, err
	}
	size, _, err := procGetFileVersionInfoSize.Call(
		uintptr(unsafe.Pointer(_fname)),
		0)
	if size == 0 {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("GetFileVersionInfoSize failed.")
	}
	buffer := make([]byte, size)
	rc, _, err := procGetFileVersionInfo.Call(
		uintptr(unsafe.Pointer(_fname)),
		0,
		size,
		uintptr(unsafe.Pointer(&buffer[0])))

	if rc == 0 {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("GetFileVersioninfo failed.")
	}

	subBlock, err := syscall.UTF16PtrFromString(`\`)
	if err != nil {
		return nil, err
	}
	var f *vsFixedFileInfo
	var queryLen uintptr

	procVerQueryValue.Call(
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(unsafe.Pointer(subBlock)),
		uintptr(unsafe.Pointer(&f)),
		uintptr(unsafe.Pointer(&queryLen)))

	return []uint{
		upper16bit(f.FileVersionMS),
		lower16bit(f.FileVersionMS),
		upper16bit(f.FileVersionLS),
		lower16bit(f.FileVersionLS),
	}, nil
}
