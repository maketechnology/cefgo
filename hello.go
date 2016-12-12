package main

// #cgo CFLAGS: -I./../../../../../../Downloads/cef_binary_3.2883.1539.gd7f087e_linux64/
// #cgo LDFLAGS: -lcef -L/home/gzunino/Downloads/cef_binary_3.2883.1539.gd7f087e_linux64/Release/
// #include <stdlib.h>
// #include "include/capi/cef_app_capi.h"
import "C"
import (
	"fmt"
	"os"
)

var _Argv []*C.char = make([]*C.char, len(os.Args))

func main() {
	fmt.Println("IN GO")
	var _MainArgs *C.struct__cef_main_args_t

	_MainArgs = (*C.struct__cef_main_args_t)(
		C.calloc(1, C.sizeof_struct__cef_main_args_t))
	fillMainArgs(_MainArgs)

	fmt.Println("GO cef_execute_process")
	var exitCode = C.cef_execute_process(_MainArgs, nil, nil)
	if exitCode >= 0 {
		fmt.Println("EXIT GO", exitCode)
		os.Exit(int(exitCode))
	}
	fmt.Println("DONE GO")
}

func fillMainArgs(mainArgs *C.struct__cef_main_args_t) {
	//Logger.Println("FillMainArgs, argc=", len(os.Args))
	fmt.Println("FillMainArgs, argc=", len(os.Args))
	fmt.Println("FillMainArgs, argv=", os.Args)
	for i, arg := range os.Args {
		_Argv[C.int(i)] = C.CString(arg)
	}
	mainArgs.argc = C.int(len(os.Args))
	mainArgs.argv = &_Argv[0]
}
