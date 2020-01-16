package main

import ("net"
"fmt"
"bufio")

func main() {

      // connect to this socket

   conn, _ := net.Dial("tcp", "69.166.48.102:8081")

      // continuosly listen and print anytime we receive a message

   for {
         // listen for reply, read until an x is found

      message, _ := bufio.NewReader(conn).ReadString('x')

      fmt.Print(message + "\n")
   }
}
