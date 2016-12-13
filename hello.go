package main

// #cgo CFLAGS: -I./../../../../../../Downloads/cef_binary_3.2883.1539.gd7f087e_linux64/
// #cgo LDFLAGS: -lcef -L/home/gzunino/Downloads/cef_binary_3.2883.1539.gd7f087e_linux64/Release/
// #include <stdlib.h>
// #include "include/capi/cef_app_capi.h"
//
// typedef struct _cef_browser_process_handler_t* (*get_browser_process_handler)(struct _cef_app_t* self);
import "C"
import (
	"fmt"
	"os"
)

var _Argv []*C.char = make([]*C.char, len(os.Args))

// export
func GetBrowserProcessHandler(self *C.struct__cef_app_t) *C.struct__cef_browser_process_handler_t {
	return nil
}

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

	// Initialize cef_settings_t structure.
	var cefSettings *C.struct__cef_settings_t
	cefSettings = (*C.struct__cef_settings_t)(
		C.calloc(1, C.sizeof_struct__cef_settings_t))
	cefSettings.size = C.sizeof_struct__cef_settings_t

	createSettings(cefSettings)

	var _AppHandler *C.cef_app_t // requires reference counting
	_AppHandler = (*C.cef_app_t)(
		C.calloc(1, C.sizeof_cef_app_t))

	f := C.get_browser_process_handler(GetBrowserProcessHandler)
	//f := getBrowserProcessHandler
	//fp := unsafe.Pointer(&f)
	//_AppHandler.get_browser_process_handler = fp

	fmt.Println("cef_initialize")
	C.cef_initialize(_MainArgs, cefSettings, _AppHandler, nil)

	fmt.Println("Calling runMessageLoop")
	// Run the Cef message loop. This will block until CefQuitMessageLoop() is
	// called.
	C.cef_run_message_loop()

	fmt.Println("Calling shutdown")
	// Shut down Cef.
	C.cef_shutdown()

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

func createSettings(cefSettings *C.struct__cef_settings_t) {
	//C.cef_string_from_utf8(cachePath, C.strlen(cachePath), &cefSettings.cache_path)

	//cefSettings.log_severity = (C.cef_log_severity_t)(C.int(settings.LogSeverity))

	//var logFile *C.char = C.CString("")
	//defer C.free(unsafe.Pointer(logFile))
	//C.cef_string_from_utf8(logFile, C.strlen(logFile), &cefSettings.log_file)

	//var resourcesDirPath *C.char = C.CString("")
	//defer C.free(unsafe.Pointer(resourcesDirPath))
	//C.cef_string_from_utf8(resourcesDirPath, C.strlen(resourcesDirPath), &cefSettings.resources_dir_path)

	//var localesDirPath *C.char = C.CString(settings.LocalesDirPath)
	//defer C.free(unsafe.Pointer(localesDirPath))
	//C.cef_string_from_utf8(localesDirPath, C.strlen(localesDirPath),
	//	&cefSettings.locales_dir_path)

	cefSettings.no_sandbox = C.int(1)
}
