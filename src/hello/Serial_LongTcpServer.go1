package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/tarm/serial"
	"net"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type attributes struct {
	cmdContent string
	serialPort string
	baudRate   string
	dataBits   string
	stopBits   string
	parity     string
	times      int
}

func GetValueByCmdStringTimes(CmdContent string, serialPort string, baudRate string, dataBits string, stopBits string, parity string, times int) (string, []byte, error) {
	defer func() {
		// recover内置函数，可以捕获到异常
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
		}
	}()
	if times > 1 {
		for j := 0; j < times-1; j++ {
			value, by, err := GetValueByCmdString(CmdContent, serialPort, baudRate, dataBits, stopBits, parity, 0)
			if err == nil {
				return value, by, err
			}
		}
	}
	return GetValueByCmdString(CmdContent, serialPort, baudRate, dataBits, stopBits, parity, 0)
}

/*
串口连接，并且执行获取值
*/
func GetValueByCmdString(CmdContent string, serialPort string, baudRate string, dataBits string, stopBits string, parity string, i int) (string, []byte, error) {
	fmt.Println("进入执行命令")
	baudInt, err := strconv.Atoi(baudRate)
	parityBit := serial.ParityNone
	if parity == "odd" || parity == "o" {
		parityBit = serial.ParityOdd
	} else if parity == "even" || parity == "e" {
		parityBit = serial.ParityEven
	} else if parity == "mark" || parity == "m" {
		parityBit = serial.ParityMark
	} else if parity == "space" || parity == "s" {
		parityBit = serial.ParitySpace
	}
	stopBit := serial.Stop1
	if stopBits == "2" {
		stopBit = serial.Stop2
	} else if stopBits == "15" || stopBits == "1.5" {
		stopBit = serial.Stop1Half
	}
	if err != nil {
		return "Baud错误", nil, err
	}

	config := &serial.Config{
		Name:        serialPort,
		Baud:        baudInt,
		Parity:      parityBit,
		StopBits:    stopBit,
		ReadTimeout: 3 * time.Second,
	}
	now := time.Now()
	fmt.Println(now)
	fmt.Println("打开串口" + serialPort)
	s, err := serial.OpenPort(config)
	defer s.Close()
	if err != nil {
		fmt.Println(err)
		fmt.Println("串口被占用，沉睡5秒中")
		time.Sleep(5 * time.Second)
		sysType := runtime.GOOS
		if sysType == "linux" {
			// LINUX系统
			exec.Command("fuser -k " + serialPort)
		}
		if sysType == "windows" {
			// windows系统
			fmt.Println("Windows system")
		}
		s, err = serial.OpenPort(config)
		if err != nil {
			return "", nil, err
		}

	}
	fmt.Println("连接成功" + serialPort)
	if CmdContent != "" {
		fmt.Println("有命令，发送后获取值")
		//字符串的十六进制转成十进制
		serialStrArray := strings.Fields(CmdContent)
		long1 := len(serialStrArray)
		bufTemp := make([]byte, long1)
		var i int
		for i = 0; i < long1; i++ {
			temp, _ := strconv.ParseInt("0x"+serialStrArray[i], 0, 16)
			bufTemp[i] = byte(temp)
		}
		s.Write(bufTemp)
		time.Sleep(1 * time.Second)
		buf := make([]byte, 500)
		n, err := s.Read(buf)
		//把buf[:n]转成字符串
		str := ConvertByteToString(buf[:n])
		fmt.Println("转化完成的字符串：" + str)
		return str, buf[:n], err
	} else {
		fmt.Println("无命令，自动获取值 延时2秒")
		time.Sleep(1 * time.Second)
		buf := make([]byte, 500)
		n, err := s.Read(buf)
		fmt.Println(buf[:n])
		//把buf[:n]转成字符串
		str := ConvertByteToString(buf[:n])
		fmt.Println("转化完成的字符串：" + str)
		return str, buf[:n], err
	}
}

func ConvertByteToString(b []byte) string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	for j := range s {
		s[j] = Tran10StringTo16(string(s[j]))
	}
	return strings.Join(s, " ")
}

func Tran10StringTo16(content string) string {
	//使用空白字符进行分割
	strAarry := strings.Fields(content)
	var resultString string
	var i int
	for i = 0; i < len(strAarry); i++ {
		//10进制字符串转成16进制
		temp, _ := strconv.ParseInt(strAarry[i], 0, 16)
		tmpStr := toHex(int(temp))
		if len(tmpStr) == 1 {
			tmpStr = "0" + tmpStr
		}
		if resultString == "" {
			resultString += tmpStr
		} else {
			resultString += " " + tmpStr
		}
	}
	return resultString
}

//十进制转换为16进制
func toHex(ten int) string {
	m := 0
	hex := make([]int, 0)
	for {
		m = ten % 16
		ten = ten / 16
		if ten == 0 {
			hex = append(hex, m)
			break
		}
		hex = append(hex, m)
	}
	var hexStr []string
	for i := len(hex) - 1; i >= 0; i-- {
		if hex[i] >= 10 {
			hexStr = append(hexStr, fmt.Sprintf("%c", 'A'+hex[i]-10))
		} else {
			hexStr = append(hexStr, fmt.Sprintf("%d", hex[i]))
		}
	}
	return strings.Join(hexStr, "")
}

// TCP Server端测试
// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [12800]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err: ", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到Client端发来的数据：", recvStr)

		dataStr := recvStr
		var dataMap map[string]string
		err = json.Unmarshal([]byte(dataStr), &dataMap)
		if err != nil {
			fmt.Printf("Json串转化为Map失败,异常:%s\n", err)
			return
		}
		fmt.Printf("Json串(本质是string)转化为Map成功:%v", dataMap)

		//json转成对象
		attributes := attributes{}
		//attributes.cmdContent = "AA BB CC 01 02 00 03 DD EE FF"
		//attributes.serialPort = "COM121"
		//attributes.baudRate = "9600"
		//attributes.dataBits = "8"
		//attributes.stopBits = ""
		//attributes.parity = "1"
		//attributes.times = 2

		attributes.cmdContent = dataMap["cmdContent"]
		attributes.serialPort = dataMap["serialPort"]
		attributes.baudRate = dataMap["baudRate"]
		attributes.dataBits = dataMap["dataBits"]
		attributes.stopBits = dataMap["stopBits"]
		attributes.parity = dataMap["parity"]
		attributes.times = HF_Atoi(dataMap["times"])

		//conn.Write([]byte(recvStr)) // 发送数据
		value, _, err := GetValueByCmdStringTimes(
			attributes.cmdContent,
			attributes.serialPort,
			attributes.baudRate,
			attributes.dataBits,
			attributes.stopBits,
			attributes.parity,
			attributes.times)
		if err == nil {
			fmt.Println(value)
			conn.Write([]byte(value)) // 发送数据
		} else {
			fmt.Println("error")
		}
	}
}

//! 字符串转数字
func HF_Atoi(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
func main() {
	defer func() {
		// recover内置函数，可以捕获到异常
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
		}
	}()
	listen, err := net.Listen("tcp", ":19999")
	if err != nil {
		fmt.Println("Listen() failed, err: ", err)
		return
	}
	for {
		conn, err := listen.Accept() // 监听客户端的连接请求
		if err != nil {
			fmt.Println("Accept() failed, err: ", err)
			continue
		}
		go process(conn) // 启动一个goroutine来处理客户端的连接请求
	}
}
