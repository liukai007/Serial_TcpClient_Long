package main

import (
	"flag"
	"fmt"
	"github.com/tarm/serial"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

var ipPort string

//端口值
var serialPortVal string

//比特率
var baudVal string

//停止位
var stopBitsVal string

//奇偶校验
var parityVal string

//毫秒数
var noMilliseconds string

func init() {
	var n int
	n=len(os.Args)
	println("There are several parameters:" + strconv.Itoa(n))
	if n == 1 {
		//第一个参数是parameters1:C:\Users\85077\AppData\Local\Temp\___18838go_build_awesomeProject1_src_hello.exe
		println("parameters1:" + os.Args[0])
		println("enter 1")
		flag.StringVar(&serialPortVal, "serialPortVal", "COM1", "serialPortVal")
		flag.StringVar(&ipPort, "ipPort", ":19998", "ipPort")
		flag.StringVar(&baudVal, "baudVal", "9600", "baudVal")
		flag.StringVar(&parityVal, "parityVal", "N", "parity")
		flag.StringVar(&stopBitsVal, "stopBitsVal", "1", "StopBits")
		flag.StringVar(&noMilliseconds, "noMilliseconds", "3000", "noMilliseconds")
	} else if n == 2 {
		println("enter 2")
		flag.StringVar(&serialPortVal, "serialPortVal", os.Args[1], "serialPortVal")
		flag.StringVar(&ipPort, "ipPort", ":19998", "ipPort")
		flag.StringVar(&baudVal, "baudVal", "9600", "baudVal")
		flag.StringVar(&parityVal, "parityVal", "N", "parity")
		flag.StringVar(&stopBitsVal, "stopBitsVal", "1", "StopBits")
		flag.StringVar(&noMilliseconds, "noMilliseconds", "3000", "noMilliseconds")
	} else if n == 3 {
		println("enter 3")
		flag.StringVar(&serialPortVal, "serialPortVal", os.Args[1], "serialPortVal")
		flag.StringVar(&ipPort, "ipPort", os.Args[2], "ipPort")
		flag.StringVar(&baudVal, "baudVal", "9600", "baudVal")
		flag.StringVar(&parityVal, "parityVal", "N", "parity")
		flag.StringVar(&stopBitsVal, "stopBitsVal", "1", "StopBits")
		flag.StringVar(&noMilliseconds, "noMilliseconds", "3000", "noMilliseconds")
	} else if n == 4 {
		println("enter 4")
		flag.StringVar(&serialPortVal, "serialPortVal", os.Args[1], "serialPortVal")
		flag.StringVar(&ipPort, "ipPort", os.Args[2], "ipPort")
		flag.StringVar(&baudVal, "baudVal", os.Args[3], "baudVal")
		flag.StringVar(&parityVal, "parityVal", "N", "parity")
		flag.StringVar(&stopBitsVal, "stopBitsVal", "1", "StopBits")
		flag.StringVar(&noMilliseconds, "noMilliseconds", "3000", "noMilliseconds")
	} else if n == 5 {
		println("enter 5")
		flag.StringVar(&serialPortVal, "serialPortVal", os.Args[1], "serialPortVal")
		flag.StringVar(&ipPort, "ipPort", os.Args[2], "ipPort")
		flag.StringVar(&baudVal, "baudVal", os.Args[3], "baudVal")
		flag.StringVar(&parityVal, "parityVal", os.Args[4], "parity")
		flag.StringVar(&stopBitsVal, "stopBitsVal", "1", "StopBits")
		flag.StringVar(&noMilliseconds, "noMilliseconds", "3000", "noMilliseconds")
	} else if n == 6 {
		println("enter 6")
		flag.StringVar(&serialPortVal, "serialPortVal", os.Args[1], "serialPortVal")
		flag.StringVar(&ipPort, "ipPort", os.Args[2], "ipPort")
		flag.StringVar(&baudVal, "baudVal", os.Args[3], "baudVal")
		flag.StringVar(&parityVal, "parityVal", os.Args[4], "parity")
		flag.StringVar(&stopBitsVal, "stopBitsVal", os.Args[5], "StopBits")
		flag.StringVar(&noMilliseconds, "noMilliseconds", "3000", "noMilliseconds")
	} else if n >= 7 {
		println("enter 7")
		flag.StringVar(&serialPortVal, "serialPortVal", os.Args[1], "serialPortVal")
		flag.StringVar(&ipPort, "ipPort", os.Args[2], "ipPort")
		flag.StringVar(&baudVal, "baudVal", os.Args[3], "baudVal")
		flag.StringVar(&parityVal, "parityVal", os.Args[4], "parity")
		flag.StringVar(&stopBitsVal, "stopBitsVal", os.Args[5], "StopBits")
		flag.StringVar(&noMilliseconds, "noMilliseconds", os.Args[6], "noMilliseconds")
	}
}

func main() {
	//println(serialPortVal)
	//println(ipPort)
	//println(baudVal)
	//println(parityVal)
	//println(stopBitsVal)
	//println(noMilliseconds)
	baudInt, _ := strconv.Atoi(baudVal)
	parityBit := serial.ParityNone
	if parityVal == "odd" || parityVal == "o" {
		parityBit = serial.ParityOdd
	} else if parityVal == "even" || parityVal == "e" {
		parityBit = serial.ParityEven
	} else if parityVal == "mark" || parityVal == "m" {
		parityBit = serial.ParityMark
	} else if parityVal == "space" || parityVal == "s" {
		parityBit = serial.ParitySpace
	}
	stopBit := serial.Stop1
	if stopBitsVal == "2" {
		stopBit = serial.Stop2
	} else if stopBitsVal == "15" || stopBitsVal == "1.5" {
		stopBit = serial.Stop1Half
	}
	defer func() {
		// recover内置函数，可以捕获到异常
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
		}
	}()
	listen, err := net.Listen("tcp", ipPort)
	if err != nil {
		fmt.Println("Listen() failed, err: ", err)
		return
	}
	var tcpConn net.Conn
	go func() {
		for {
			tcpConn, err = listen.Accept() // 监听客户端的连接请求
			if err != nil {
				fmt.Println("Accept() failed, err: ", err)
				continue
			}
			numMilli, _ := strconv.Atoi(noMilliseconds)
			err = SerialBase(tcpConn, serialPortVal, baudInt, parityBit, stopBit, numMilli)
			if err != nil {
				continue
			}
		}
	}()

	for true {
		time.Sleep(time.Second)
	}

}

func SerialBase(tcpConn net.Conn, serialPort string, baudVal int, parityVal serial.Parity, stopBitsVal serial.StopBits, noMillisecondsV int) error {
	defer func() {
		// recover内置函数，可以捕获到异常
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
		}
	}()
	//设置串口编号
	//ser := &serial.Config{Name: "COM1", Baud: 9600}
	ser := &serial.Config{
		Name:        serialPort,
		Baud:        baudVal,
		Parity:      parityVal,
		StopBits:    stopBitsVal,
		ReadTimeout: 3 * time.Second,
	}

	//打开串口
	var errTcp error
	conn, err := serial.OpenPort(ser)
	for err != nil {
		log.Fatal(err)
		fmt.Println("串口被占用，沉睡5秒中")
		time.Sleep(5 * time.Second)
		sysType := runtime.GOOS
		if sysType == "linux" {
			exec.Command("fuser -k " + serialPort)
		} else {
			fmt.Println("Windows system")
		}
		conn, err = serial.OpenPort(ser)
	}

	//启动一个协程循环发送
	go func() {
		for {
			var n int
			buf := make([]byte, 1024)
			n, errTcp = tcpConn.Read(buf)
			if errTcp != nil {
				tcpConn.Close()
				conn.Close()
				break
			}
			revData := buf[:n]
			_, err := conn.Write(revData)
			if err != nil {
				log.Println(err)
				conn, err = serial.OpenPort(ser)
				continue
			}
			log.Printf("Tx:%X \n", revData)
			time.Sleep(time.Second)
		}
	}()

	//保持数据持续接收
	for {
		if errTcp != nil {
			conn.Close()
			return errTcp
		}
		buf := make([]byte, 1024)
		lens, err := conn.Read(buf)
		time.Sleep(time.Duration(noMillisecondsV) * time.Millisecond)
		if err != nil {
			log.Println(err)
			conn, err = serial.OpenPort(ser)
			continue
		}
		revData := buf[:lens]
		if len(revData) > 0 {
			log.Printf("Rx:%X \n", revData)
			_, errTcp = tcpConn.Write(revData)
			if errTcp != nil {
				//conn.Close()
				tcpConn.Close()
				return errTcp
			}
		}
	}
}
