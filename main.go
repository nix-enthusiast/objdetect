package main

/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -L./lib -lunildd
#include "unildd.h"
#include <stdint.h>
#include <string.h>
*/
import "C"

import (
	"fmt"
	"os"
	"unsafe"
)

func toGoString(cStr *C.char) string {
	if cStr == nil {
		return "Undefined"
	} else {
		return C.GoString(cStr)
	}
}

func printObjects(fileNames []string, isMultiple bool) {
	for _, fileName := range fileNames {
		CFileName := C.CString(fileName)

		fileContent, err := os.ReadFile(fileName)

		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)

			if isMultiple {
				fmt.Println()
				fmt.Println()
				continue
			} else {
				os.Exit(-25)
			}
		}

		buffer := (*C.uint8_t)(unsafe.Pointer(&fileContent[0]))

		size := C.size_t(len(fileContent))

		readObjects := C.read_obj(CFileName, buffer, size, false)
		objectsLength := int(readObjects.length)
		objectArray := unsafe.Slice(readObjects.vec, objectsLength)

		for i := 0; i < objectsLength; i++ {
			var objectResult C.ULDDObjResult = objectArray[i]

			object := objectResult.obj
			err := objectResult.error

			errorCode := int64(err.code)

			if errorCode != 0 {
				_, _ = fmt.Fprintln(os.Stderr, C.GoString(err.explanation))

				if isMultiple {
					fmt.Println()
					fmt.Println()
					continue
				} else {
					if errorCode < 0 {
						os.Exit(int(err.code))
					} else {
						os.Exit(-25)
					}
				}
			}

			fmt.Println("File name:           " + C.GoString(object.file_name))

			memberName := ""

			memberNamesLength := int(object.member_name.length)
			memberNames := unsafe.Slice(object.member_name.vec, memberNamesLength)
			for j := 0; j < memberNamesLength; j++ {
				member := memberNames[j]
				memberName += C.GoString(member)

				if j+1 != int(object.member_name.length) {
					memberName += " -> "
				}
			}

			fmt.Println("Member of:           " + memberName)

			fmt.Println("Executable format:   " + toGoString(object.executable_format))

			is64 := ""

			if object.is_64 {
				is64 = "64-bit"
			} else {
				is64 = "32-bit"
			}

			fmt.Println("Word size:           " + is64)

			fmt.Println("OS:                  " + toGoString(object.os_type))

			fmt.Println("File type:           " + toGoString(object.file_type))

			isStripped := ""

			if object.is_stripped {
				isStripped = "Yes"
			} else {
				isStripped = "No"
			}

			fmt.Println("Is stripped:         " + isStripped)

			fmt.Println("CPU type:            " + toGoString(object.cpu_type))

			fmt.Println("CPU subtype:         " + toGoString(object.cpu_subtype))

			fmt.Println("Linker:              " + toGoString(object.interpreter))

			fmt.Println("Libraries:")
			librariesLength := int(object.libraries.length)
			libraries := unsafe.Slice(object.libraries.vec, librariesLength)
			for _, library := range libraries {
				fmt.Println("  -" + toGoString(library))
			}

			if i+1 != int(readObjects.length) {
				fmt.Println()
				fmt.Println()
			}
		}

		C.free_obj(readObjects, false)
		C.free(unsafe.Pointer(CFileName))
	}
}

func main() {
	argv := os.Args[1:]
	argc := len(argv)

	helpMenu := ` OBJDetect - A Tool To Get Information About The Executable/Library Files
--------------------------------------------------------------------------

 DESCRIPTION
-------------
  cn [-h | /H | --help    | help]            : prints this message
  cn [-i | /I | --input   | ipt]  <file(s)>  : parses files to get information`

	if argc == 0 {
		fmt.Println(helpMenu)
		return
	}

	isMultiple := argc > 2
	flag := argv[0]

	switch flag {
	case "-h":
		fmt.Println(helpMenu)
		return

	case "-i":
		if argc < 2 {
			_, _ = fmt.Fprintf(os.Stderr, "Enter a file. Type 'objdetect -h' to get help")
			os.Exit(1)
		}

		printObjects(argv[1:], isMultiple)
		return

	default:
		_, _ = fmt.Fprintf(os.Stderr, "Invalid flag. Type 'objdetect -h' to get help")
		os.Exit(1)
	}
}
