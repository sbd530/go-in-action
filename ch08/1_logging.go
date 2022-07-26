package main

import "log"

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// Println 은 표준 로거에 메시지를 출력한다.
	log.Println("Message")
	// Fatalln 은 Println() 을 실행한 후 os.Exit(1)을 호출한다.
	log.Fatalln("Fatal err msg")
	// Panicln 은 Println() 을 호출한 후 panic() 을 호출한다.
	log.Panicln("Panic msg")
}
