package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/H5eye/go-pefile"
	"github.com/floyernick/fleep-go"
)

func peSubsystem(pe *pefile.PE) string {
	opt := pe.OptionalHeader
	if opt == nil {
		return "Subsystem Unknown: No Optional Header"
	}
	switch opt.Subsystem {
	default:
		return "Subsystem Unknown: Unknown Subsystem code"
	case pefile.IMAGE_SUBSYSTEM_UNKNOWN:
		return "Subsystem Unknown: Subsystem code for unknown"
	case pefile.IMAGE_SUBSYSTEM_NATIVE:
		return "Native"
	case pefile.IMAGE_SUBSYSTEM_WINDOWS_GUI:
		return "Windows GUI"
	case pefile.IMAGE_SUBSYSTEM_WINDOWS_CUI:
		return "Windows CUI"
	case pefile.IMAGE_SUBSYSTEM_OS2_CUI:
		return "OS2 CUI"
	case pefile.IMAGE_SUBSYSTEM_POSIX_CUI:
		return "POSIX CUI"
	case pefile.IMAGE_SUBSYSTEM_NATIVE_WINDOWS:
		return "Native Windows"
	case pefile.IMAGE_SUBSYSTEM_WINDOWS_CE_GUI:
		return "Windows CE GUI"
	case pefile.IMAGE_SUBSYSTEM_EFI_APPLICATION:
		return "EFI Application"
	case pefile.IMAGE_SUBSYSTEM_EFI_BOOT_SERVICE_DRIVER:
		return "EFI BOOT Service Driver"
	case pefile.IMAGE_SUBSYSTEM_EFI_RUNTIME_DRIVER:
		return "EFI Runtime Driver"
	case pefile.IMAGE_SUBSYSTEM_EFI_ROM:
		return "EFI ROM"
	case pefile.IMAGE_SUBSYSTEM_XBOX:
		return "XBOX:"
	case pefile.IMAGE_SUBSYSTEM_WINDOWS_BOOT_APPLICATION:
		return "Windows Boot Application"
	}
}

func imageCharacteristics(pe *pefile.PE) []string {
	result := []string{}
	ch := pe.FileHeader.Characteristics

	if (ch & pefile.IMAGE_FILE_EXECUTABLE_IMAGE) != 0 {
		result = append(result, "Executable Image")
	}
	if (ch & pefile.IMAGE_FILE_16BIT_MACHINE) != 0 {
		result = append(result, "16bit")
	}
	if (ch & pefile.IMAGE_FILE_32BIT_MACHINE) != 0 {
		result = append(result, "32bit")
	}
	if (ch & pefile.IMAGE_FILE_DLL) != 0 {
		result = append(result, "DLL")
	}
	return result
}

var extensions = map[string]func(fname string, bin []byte) string{
	"exe": func(fname string, bin []byte) string {
		pe, err := pefile.Parse(bin)
		if err != nil {
			return ""
		}
		tags := []string{peSubsystem(pe)}
		tags = append(tags, imageCharacteristics(pe)...)
		if pe.OptionalHeader64 != nil {
			tags = append(tags, "64bit Header")
		}
		return "<" + strings.Join(tags, "><") + ">"
	},
}

func eachFile(fname string) error {
	file, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}
	info, err := fleep.GetInfo(file)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", fname)
	functions := make([]func(string, []byte) string, 0, 10)
	for i := range info.Type {
		fmt.Printf("  %-13s: %-4s: %s\n", info.Type[i], info.Extension[i], info.Mime[i])
		if extension1, ok := extensions[info.Extension[i]]; ok {
			functions = append(functions, extension1)
		}
	}
	for _, f := range functions {
		if detail := f(fname, file); detail != "" {
			for _, line := range strings.Split(detail, "\n") {
				fmt.Printf("  %s\n", line)
			}
		}
	}
	return nil
}

func main() {
	for _, fname := range os.Args[1:] {
		if err := eachFile(fname); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", fname, err.Error())
		}
		fmt.Println()
	}
}
