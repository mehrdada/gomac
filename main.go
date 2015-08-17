package main

import (
	"fmt"
	//	"net"
	//	"net/http"
	//	"io"
	"io"
	"log"
	"os"
	"strconv"
	//"syscall"
)

func main() {
	fmt.Printf("Starting go server (%s)...\n", os.Args[0])
	fd64, err := strconv.ParseInt(os.Args[0], 16, 64)
	if err != nil {
		log.Panic(err)
	}
	buf := make([]byte, 10240)
	/*syscall.Recvmsg(fd, buf, nil, syscall.recvm);
		syscall.ParseSocketControlMessage(buf)

		memset(&child_msg,   0, sizeof(child_msg));
	char cmsgbuf[CMSG_SPACE(sizeof(int))];
	child_msg.msg_control = cmsgbuf; // make place for the ancillary message to be received
	child_msg.msg_controllen = sizeof(cmsgbuf);

	printf("Waiting on recvmsg\n");
	rc = recvmsg(worker_sd, &child_msg, 0);
	struct cmsghdr *cmsg = CMSG_FIRSTHDR(&child_msg);
	if (cmsg == NULL || cmsg -> cmsg_type != SCM_RIGHTS) {
	     printf("The first control structure contains no file descriptor.\n");
	     exit(0);
	}
	memcpy(&pass_sd, CMSG_DATA(cmsg), sizeof(pass_sd));
	printf("Received descriptor = %d\n", pass_sd);
	*/
	file := os.NewFile(uintptr(fd64), "")
	for {
		count, err := io.ReadAtLeast(file, buf, 1)

		if err != nil {
			panic(err)
		}
		fmt.Printf("count:%d\n", count)
		fmt.Println(string(buf[:count]))
		if err == nil {
			file.Write(append([]byte("<html><style>body{background-color:white;color:green</style><a href=\"/\">Hello</a>, "), buf[:count]...))
			//			syscall.Sendmsg(fd, append([]byte("<html><style>body{background-color:white;color:black</style>Hello, "), buf[:count]...), nil, from, 0)
		}
	}
	fmt.Scanf("%s", &buf)
}
