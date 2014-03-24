package main

/*
#cgo CFLAGS: -I/usr/include/cups
#cgo LDFLAGS: -lcups
#include </usr/include/cups/cups.h>
#include <stdlib.h>
*/
import "C"
import "fmt"
import "unsafe"

func main() {
	// fmt.Println(int(C.random()))

	fmt.Println(C.GoString(C.cupsUser()))

	fmt.Println(C.GoString(C.cupsGetDefault()))

	var options C.cups_option_t

	options.name = C.CString("PageSize")
	options.value = C.CString("Postcard")

	fmt.Println(options)
	fmt.Println(C.cupsPrintFile(C.CString("EPSON_Epson_Stylus_Photo_T50"), C.CString("/Users/Raymond/Downloads/hz.png"), C.CString("Test Print"), C.int(0), unsafe.Pointer(options)))
}
