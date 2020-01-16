package main

import ("net"
"fmt"
"time")

func main() {

      // print message to user

   fmt.Println("Launching server . . .")

      // listen on all interfaces

   ln, _ := net.Listen("tcp", ":8081")

      // accept connection on port

   conn, _ := ln.Accept()

      // get client ip

   clientIP := conn.RemoteAddr().String()

      // run loop forever (or until ctrl-c), loop by seconds:

   for range time.Tick(time.Second){

         // set date variable

      now := time.Now().String()

         // send the current time to the client

      conn.Write([]byte("Date: " + now + "\n"))
      conn.Write([]byte("Client IP: " + clientIP + "\n" + "x"))
   }
}
