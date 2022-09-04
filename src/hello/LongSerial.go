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
	"strings"
	"time"
)

var ipPort string

//公有函数首字母得大写，私有函数首字母得小写
//端口值
var serialPortVal string

//比特率
var baudVal string

//数据位
var dataBits string

//停止位
var stopBitsVal string

//奇偶校验
var parityVal string

//毫秒数
var noMilliseconds string

//全局变量
var tcpConnMap map[net.Conn]struct{}

func init() {
	var n int
	n = len(os.Args)
	println("There are several parameters:" + strconv.Itoa(n))
	if n == 1 {
		//第一个参数是parameters1:C:\Users\85077\AppData\Local\Temp\___18838go_build_awesomeProject1_src_hello.exe
		println("parameters1:" + os.Args[0])
		println("enter 1")
		flag.StringVar(&serialPortVal, "serialPortVal", "COM1", "serialPortVal")
		flag.StringVar(&ipPort, "ipPort", ":19998", "ipPort")
		flag.StringVar(&baudVal, "baudVal", "9600", "baudVal")
		flag.StringVar(&parityVal, "parityVal", "N", "parity")
		flag.StringVar(&dataBits, "dataBits", "8", "dataBits")
		flag.StringVar(&stopBitsVal, "stopBitsVal", "1", "StopBits")
		flag.StringVar(&noMilliseconds, "noMilliseconds", "3000", "noMilliseconds")
	} else if n == 2 {
		println("enter 2")
		flag.StringVar(&serialPortVal, "serialPortVal", os.Args[1], "serialPortVal")
		flag.StringVar(&ipPort, "ipPort", ":19998", "ipPort")
		flag.StringVar(&baudVal, "baudVal", "9600", "baudVal")
		flag.StringVar(&parityVal, "parityVal", "N", "parity")
		flag.StringVar(&dataBits, "dataBits", "8", "dataBits")
		flag.StringVar(&stopBitsVal, "stopBitsVal", "1", "StopBits")
		flag.StringVar(&noMilliseconds, "noMilliseconds", "3000", "noMilliseconds")
	} else if n == 3 {
		println("enter 3")
		flag.StringVar(&serialPortVal, "serialPortVal", os.Args[1], "serialPortVal")
		flag.StringVar(&ipPort, "ipPort", os.Args[2], "ipPort")
		flag.StringVar(&baudVal, "baudVal", "9600", "baudVal")
		flag.StringVar(&parityVal, "parityVal", "N", "parity")
		flag.StringVar(&dataBits, "dataBits", "8", "dataBits")
		flag.StringVar(&stopBitsVal, "stopBitsVal", "1", "StopBits")
		flag.StringVar(&noMilliseconds, "noMilliseconds", "3000", "noMilliseconds")
	} else if n == 4 {
		println("enter 4")
		flag.StringVar(&serialPortVal, "serialPortVal", os.Args[1], "serialPortVal")
		flag.StringVar(&ipPort, "ipPort", os.Args[2], "ipPort")
		flag.StringVar(&baudVal, "baudVal", os.Args[3], "baudVal")
		flag.StringVar(&parityVal, "parityVal", "N", "parity")
		flag.StringVar(&dataBits, "dataBits", "8", "dataBits")
		flag.StringVar(&stopBitsVal, "stopBitsVal", "1", "StopBits")
		flag.StringVar(&noMilliseconds, "noMilliseconds", "3000", "noMilliseconds")
	} else if n == 5 {
		println("enter 5")
		flag.StringVar(&serialPortVal, "serialPortVal", os.Args[1], "serialPortVal")
		flag.StringVar(&ipPort, "ipPort", os.Args[2], "ipPort")
		flag.StringVar(&baudVal, "baudVal", os.Args[3], "baudVal")
		flag.StringVar(&parityVal, "parityVal", os.Args[4], "parity")
		flag.StringVar(&dataBits, "dataBits", "8", "dataBits")
		flag.StringVar(&stopBitsVal, "stopBitsVal", "1", "StopBits")
		flag.StringVar(&noMilliseconds, "noMilliseconds", "3000", "noMilliseconds")
	} else if n == 6 {
		println("enter 6")
		flag.StringVar(&serialPortVal, "serialPortVal", os.Args[1], "serialPortVal")
		flag.StringVar(&ipPort, "ipPort", os.Args[2], "ipPort")
		flag.StringVar(&baudVal, "baudVal", os.Args[3], "baudVal")
		flag.StringVar(&parityVal, "parityVal", os.Args[4], "parity")
		flag.StringVar(&dataBits, "dataBits", os.Args[5], "dataBits")
		flag.StringVar(&stopBitsVal, "stopBitsVal", "1", "StopBits")
		flag.StringVar(&noMilliseconds, "noMilliseconds", "3000", "noMilliseconds")
	} else if n == 7 {
		println("enter 7")
		flag.StringVar(&serialPortVal, "serialPortVal", os.Args[1], "serialPortVal")
		flag.StringVar(&ipPort, "ipPort", os.Args[2], "ipPort")
		flag.StringVar(&baudVal, "baudVal", os.Args[3], "baudVal")
		flag.StringVar(&parityVal, "parityVal", os.Args[4], "parity")
		flag.StringVar(&dataBits, "dataBits", os.Args[5], "dataBits")
		flag.StringVar(&stopBitsVal, "stopBitsVal", os.Args[6], "StopBits")
		flag.StringVar(&noMilliseconds, "noMilliseconds", "3000", "noMilliseconds")
	} else if n >= 8 {
		println("enter 8")
		flag.StringVar(&serialPortVal, "serialPortVal", os.Args[1], "serialPortVal")
		flag.StringVar(&ipPort, "ipPort", os.Args[2], "ipPort")
		flag.StringVar(&baudVal, "baudVal", os.Args[3], "baudVal")
		flag.StringVar(&parityVal, "parityVal", os.Args[4], "parity")
		flag.StringVar(&dataBits, "dataBits", os.Args[5], "dataBits")
		flag.StringVar(&stopBitsVal, "stopBitsVal", os.Args[6], "StopBits")
		flag.StringVar(&noMilliseconds, "noMilliseconds", os.Args[7], "noMilliseconds")
	}
}

func main() {
	println(serialPortVal)
	println(ipPort)
	println(baudVal)
	println(parityVal)
	println(stopBitsVal)
	println(noMilliseconds)
	baudInt, _ := strconv.Atoi(baudVal)
	//校验码
	parityBit := serial.ParityNone
	parityVal = strings.ToLower(parityVal)
	if parityVal == "odd" || parityVal == "o" {
		parityBit = serial.ParityOdd
	} else if parityVal == "even" || parityVal == "e" {
		parityBit = serial.ParityEven
	} else if parityVal == "mark" || parityVal == "m" {
		parityBit = serial.ParityMark
	} else if parityVal == "space" || parityVal == "s" {
		parityBit = serial.ParitySpace
	}
	//数据位 //字符串转成数字
	var dataBits1 uint8
	dataBits1 = 8
	if dataBits == "7" {
		dataBits1 = 7
	} else if dataBits == "5" {
		dataBits1 = 5
	} else if dataBits == "6" {
		dataBits1 = 6
	}
	//停止位
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
	tcpConnMap = map[net.Conn]struct{}{}
	numMilli, _ := strconv.Atoi(noMilliseconds)

	go SerialBase(serialPortVal, baudInt, parityBit, dataBits1, stopBit, numMilli)

	var tcpConn net.Conn
	go func() {
		for {
			tcpConn, err = listen.Accept() // 监听客户端的连接请求
			if err != nil {
				fmt.Println("Accept() failed, err: ", err)
				continue
			} else {
				tcpConnMap[tcpConn] = struct{}{}
			}

			//err = SerialBase(tcpConn, serialPortVal, baudInt, parityBit, stopBit, numMilli)
			//if err != nil {
			//	continue
			//}
		}
	}()

	for true {
		time.Sleep(time.Second)
	}

}

func SerialBase(serialPort string, baudVal int, parityVal serial.Parity, dataBits1 uint8, stopBitsVal serial.StopBits, noMillisecondsV int) error {
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
		Size:        dataBits1,
		Parity:      parityVal,
		StopBits:    stopBitsVal,
		ReadTimeout: 3 * time.Second,
	}

	//打开串口
	var errTcp error
	serialConn, err := serial.OpenPort(ser)
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
		serialConn, err = serial.OpenPort(ser)
	}

	//启动一个协程循环发送
	go func() {
		defer func() {
			// recover内置函数，可以捕获到异常
			err := recover()
			if err != nil {
				fmt.Println("err:", err)
			}
		}()
		for {
			var n int
			buf := make([]byte, 1024)
			if tcpConnMap != nil {
				for tcpConn, _ := range tcpConnMap {
					n, errTcp = tcpConn.Read(buf)
					if errTcp != nil {
						tcpConn.Close()
						delete(tcpConnMap, tcpConn)
						//serialConn.Close()
						continue
					}
					revData := buf[:n]
					_, err := serialConn.Write(revData)
					if err != nil {
						log.Println(err)
						serialConn, err = serial.OpenPort(ser)
						continue
					}
					log.Printf("Tx:%X \n", revData)
					time.Sleep(500 * time.Millisecond)
				}
			}
		}
	}()

	//保持数据持续接收
	for {
		//if errTcp != nil {
		//	serialConn.Close()
		//	return errTcp
		//}
		buf := make([]byte, 1024)
		lens, err := serialConn.Read(buf)
		time.Sleep(time.Duration(noMillisecondsV) * time.Millisecond)
		if err != nil {
			log.Println(err)
			serialConn, err = serial.OpenPort(ser)
			continue
		}
		revData := buf[:lens]
		if len(revData) > 0 {
			log.Printf("Rx:%X \n", revData)
			if tcpConnMap != nil {
				for tcpConn, _ := range tcpConnMap {
					_, errTcp = tcpConn.Write(revData)
					if errTcp != nil {
						//serialConn.Close()
						tcpConn.Close()
						delete(tcpConnMap, tcpConn)
						//return errTcp
					}
				}
			}
		}
	}
}
