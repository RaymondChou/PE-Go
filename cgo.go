package main

/*
#cgo CFLAGS: -I/usr/include/cups
#cgo LDFLAGS: -lcups -D_IPP_PRIVATE_STRUCTURES
#include <cups.h>
*/
import "C"

import "fmt"

import _ "unsafe"

func main() {

	fmt.Println(C.GoString(C.cupsUser()))

	fmt.Println(C.GoString(C.cupsGetDefault()))

	var options *C.cups_option_t

	num_options := C.int(0)

	num_options = C.cupsAddOption(C.CString("PageSize"), C.CString("Postcard"), C.int(num_options), &options)

	num_options = C.cupsAddOption(C.CString("Media"), C.CString("Photo"), C.int(num_options), &options)

	dest := C.CString("EPSON_Epson_Stylus_Photo_T50")
	file := C.CString("/Users/Raymond/Downloads/hz.png")
	title := C.CString("Test Print")

	job_id, err := C.cupsPrintFile(dest, file, title, num_options, options)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(job_id)

	defer C.cupsFreeOptions(num_options, options)

}
