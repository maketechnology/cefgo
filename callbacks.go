package main

// #include <stdlib.h>
// #include <stdio.h>
// #include "include/capi/cef_browser_process_handler_capi.h"
// #include "include/capi/cef_app_capi.h"
import "C"
import "fmt"

//export Callback
func Callback() {
	fmt.Println("in Callback")
}
