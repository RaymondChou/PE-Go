package main

/*
#cgo CFLAGS: -I/usr/include/cups
#cgo LDFLAGS: -lcups -D_IPP_PRIVATE_STRUCTURES
#include <cups.h>
*/
import "C"

import "fmt"
import (
	"flag"
	"github.com/golang/glog"
	"time"
)

import _ "unsafe"

func main() {
	flag.Parse()
	timer := time.NewTicker(2 * time.Second)
	timer_two := time.NewTicker(5 * time.Second)

	fmt.Println(C.GoString(C.cupsUser()))

	fmt.Println(C.GoString(C.cupsGetDefault()))

	var options *C.cups_option_t

	num_options := C.int(0)

	num_options = C.cupsAddOption(C.CString("PageSize"), C.CString("Postcard"), C.int(num_options), &options)
	num_options = C.cupsAddOption(C.CString("MediaSize"), C.CString("A4"), C.int(num_options), &options)
	num_options = C.cupsAddOption(C.CString("MediaSize1"), C.CString("A5"), C.int(num_options), &options)

	dest := C.CString("Test")
	file := C.CString("/Users/ZhouYT/Downloads/follow.png")
	title := C.CString("Test Print")

	job_id, err := C.cupsPrintFile(dest, file, title, num_options, options)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(job_id)

	defer C.cupsFreeOptions(num_options, options)

	fmt.Println(C.GoString(options.name))

	for {
		select {
		case <-timer.C:

			go func() {
				fmt.Println("test2")
				glog.Info("test")
			}()

		case <-timer_two.C:

			go func() {
				fmt.Println("test5")
				glog.Info("test5")
			}()
		}
	}

	defer glog.Flush()
}
