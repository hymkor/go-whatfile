package wfile

import (
	"errors"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
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

type versionInfo struct {
	buffer []byte
	size   uintptr
	fname  *uint16
}

func NewVersionInfo(fname string) (*versionInfo, error) {
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

	return &versionInfo{
		buffer: buffer,
		size:   size,
		fname:  _fname,
	}, nil
}

func (vi *versionInfo) Query(key string, f uintptr) (uintptr, error) {
	subBlock, err := syscall.UTF16PtrFromString(key)
	if err != nil {
		return 0, err
	}
	var queryLen uintptr
	procVerQueryValue.Call(
		uintptr(unsafe.Pointer(&vi.buffer[0])),
		uintptr(unsafe.Pointer(subBlock)),
		f,
		uintptr(unsafe.Pointer(&queryLen)))

	return queryLen, nil
}

func (vi *versionInfo) Number() (file []uint, product []uint, err error) {
	var f *vsFixedFileInfo

	_, err = vi.Query(`\`, uintptr(unsafe.Pointer(&f)))
	if err != nil {
		return nil, nil, err
	}
	return []uint{
			upper16bit(f.FileVersionMS),
			lower16bit(f.FileVersionMS),
			upper16bit(f.FileVersionLS),
			lower16bit(f.FileVersionLS),
		},
		[]uint{
			upper16bit(f.ProductVersionMS),
			lower16bit(f.ProductVersionMS),
			upper16bit(f.ProductVersionLS),
			lower16bit(f.ProductVersionLS),
		},
		nil
}

func (vi *versionInfo) Translation() (uint32, uint32) {
	var pLangCode *uint32
	vi.Query(`\VarFileInfo\Translation`, uintptr(unsafe.Pointer(&pLangCode)))
	return *pLangCode, *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(pLangCode)) + unsafe.Sizeof(*pLangCode)))
}

func utf16PtrToArray(p uintptr, size uintptr) []uint16 {
	buffer := make([]uint16, 0, size)
	for size > 0 {
		ch := *(*uint16)(unsafe.Pointer(p))
		if ch == 0 {
			break
		}
		buffer = append(buffer, ch)
		p += unsafe.Sizeof(uint16(0))
	}
	return buffer
}

func (vi *versionInfo) Item(key string) string {
	var pStrInfo *uint16
	length, _ := vi.Query(key, uintptr(unsafe.Pointer(&pStrInfo)))
	if length <= 0 {
		return ""
	}
	utf16array := utf16PtrToArray(uintptr(unsafe.Pointer(pStrInfo)), length)
	return windows.UTF16ToString(utf16array)
}

//func {
//	lang1,lang2 := vi.Translation()
//	key := fmt.Sprintf(`\StringFileInfo\%04X%04X\LegalTranslation`, lang1,lang2)
//}

func GetVersionInfo(fname string) (*Version, error) {
	vi, err := NewVersionInfo(fname)
	if err != nil {
		return nil, err
	}
	file, product, err := vi.Number()
	if err != nil {
		return nil, err
	}
	v := &Version{}
	copy(v.File[:], file)
	copy(v.Product[:], product)
	return v, nil
}
