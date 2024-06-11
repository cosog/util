// encode.go
package util_encode

// import (
// 	"encoding/binary"
// 	"math"
// 	"reflect"
// 	"runtime"
// 	"strings"
// 	"unsafe"

// 	"github.com/cosog/util"

// 	"github.com/cosog/util/crc"
// )

// func Encode_ReadHoldingRegisters(protocolType string, storeDataType string, slave byte, addr int, quantity int) ([]int, []string) {

// 	var crc util_crc.Crc
// 	var mbap []byte
// 	var code byte
// 	var addr_pdu int
// 	var addr_mbap int
// 	var quantity_pdu int
// 	var addr_res []int
// 	var adu_res []string

// 	code = 0x03

// 	switch strings.ToLower(storeDataType) { //asc bit byte int16 uint16 float32 bcd
// 	case "asc":
// 	case "bit":
// 		quantity = int(math.Ceil(float64(quantity) / 16))
// 	case "byte":
// 		quantity = int(math.Ceil(float64(quantity) / 2))
// 	case "int16":
// 	case "uint16":
// 	case "float32":
// 		quantity = 2 * quantity
// 	case "bcd":
// 		quantity = int(math.Ceil(float64(quantity) / 4))
// 	}

// 	sendQuantity := 100

// 	var times int
// 	if quantity%sendQuantity == 0 {
// 		times = quantity / sendQuantity
// 	} else {
// 		times = quantity/sendQuantity + 1 //当小于拆分发送寄存器长度时，按实际长度发送
// 	}

// 	for i := 0; i < times; i++ {

// 		addr_mbap = addr + i*sendQuantity
// 		// addr_pdu = addr - 1 - 40000 + i*sendQuantity
// 		if addr > 400000 {
// 			addr_pdu = addr - 1 - 400000 + i*sendQuantity
// 		} else {
// 			addr_pdu = addr - 1 - 40000 + i*sendQuantity
// 		}
// 		if i == times-1 { //当小于拆分发送寄存器长度时，按实际长度发送
// 			quantity_pdu = quantity - sendQuantity*i
// 		} else {
// 			quantity_pdu = sendQuantity
// 		}

// 		addr_pdu_b := make([]byte, 2)
// 		addr_mbap_b := make([]byte, 2)
// 		quantity_b := make([]byte, 2)

// 		binary.BigEndian.PutUint16(addr_mbap_b, uint16(int16(addr_mbap)))
// 		binary.BigEndian.PutUint16(addr_pdu_b, uint16(int16(addr_pdu)))
// 		binary.BigEndian.PutUint16(quantity_b, uint16(int16(quantity_pdu)))

// 		b := make([]byte, 0)

// 		switch strings.ToLower(protocolType) {

// 		case "modbus-rtu": //被动响应

// 			b = append(b, slave)
// 			b = append(b, code)
// 			b = append(b, addr_pdu_b...)
// 			b = append(b, quantity_b...)

// 			crc.Reset()
// 			crc.PushBytes(b)
// 			b = append(b, crc.Low, crc.High)

// 		case "modbus-tcp": //被动响应
// 			fallthrough
// 		default:

// 			mbap = make([]byte, 0)
// 			mbap = append(mbap, addr_mbap_b...)                                                                    //事务标识符 2个字节
// 			mbap = append(mbap, 0x00, 0x00)                                                                        //协议标识符 2个字节
// 			size := byte(unsafe.Sizeof(slave)) + byte(unsafe.Sizeof(code)) + byte(len(addr_pdu_b)+len(quantity_b)) //
// 			mbap = append(mbap, 0x00, size)                                                                        //数据帧长度 2个字节
// 			mbap = append(mbap, slave)

// 			b = append(b, mbap...)

// 			b = append(b, code)
// 			b = append(b, addr_pdu_b...)
// 			b = append(b, quantity_b...)

// 		}
// 		addr_res = append(addr_res, addr_mbap)
// 		adu_res = append(adu_res, string(b))
// 		runtime.Gosched()
// 	}
// 	return addr_res, adu_res
// }
// func Encode_WriteHoldingRegister(protocolType string, storeDataType string, ifDataType string, slave byte, addr int, ratio float64, value interface{}) string {

// 	var crc util_crc.Crc
// 	var start, end int
// 	var mbap []byte
// 	var code byte
// 	var length byte
// 	var quantity int
// 	var addr_pdu int
// 	var addr_mbap int

// 	addr_mbap = addr
// 	// addr_pdu = addr - 1 - 40000
// 	if addr > 400000 {
// 		addr_pdu = addr - 1 - 400000
// 	} else {
// 		addr_pdu = addr - 1 - 40000
// 	}

// 	addr_mbap_b := make([]byte, 2)
// 	addr_pdu_b := make([]byte, 2)

// 	binary.BigEndian.PutUint16(addr_mbap_b, uint16(int16(addr_mbap)))
// 	binary.BigEndian.PutUint16(addr_pdu_b, uint16(int16(addr_pdu)))

// 	quantity_b := make([]byte, 2)
// 	value_b := make([]byte, 0)
// 	var elements []interface{}
// 	var element interface{}
// 	var tempBool []bool
// 	var tempInt []int
// 	var tempFloat32 []float32
// 	var tempFloat64 []float64
// 	var tempString []string
// 	var ok bool

// 	switch reflect.TypeOf(value).String() {
// 	case "[]interface {}":
// 		elements, ok = value.([]interface{})
// 		if ok {
// 			for _, v := range elements {

// 				switch strings.ToLower(ifDataType) { //bool int float32 float64 string
// 				case "bool":
// 					switch v.(type) {
// 					case bool:
// 						tempBool = append(tempBool, v.(bool))
// 					}
// 				case "int":
// 					switch v.(type) {
// 					case int:
// 						tempInt = append(tempInt, v.(int))
// 					case float32:
// 						tempInt = append(tempInt, int(v.(float32)))
// 					case float64:
// 						tempInt = append(tempInt, int(v.(float64)))
// 					}

// 				case "float32":
// 					switch v.(type) {
// 					case int:
// 						tempFloat32 = append(tempFloat32, float32(v.(int)))
// 					case float32:
// 						tempFloat32 = append(tempFloat32, v.(float32))
// 					case float64:
// 						tempFloat32 = append(tempFloat32, float32(v.(float64)))
// 					}
// 				case "float64":
// 					switch v.(type) {
// 					case int:
// 						tempFloat64 = append(tempFloat64, float64(v.(int)))
// 					case float32:
// 						tempFloat64 = append(tempFloat64, float64(v.(float32)))
// 					case float64:
// 						tempFloat64 = append(tempFloat64, float64(v.(float64)))
// 					}
// 				case "string":
// 					tempString = append(tempString, v.(string))
// 				}
// 				runtime.Gosched()
// 			}
// 			switch strings.ToLower(ifDataType) { //bool int float32 float64 string

// 			case "bool":
// 				element = tempBool
// 			case "int":
// 				element = tempInt
// 			case "float32":
// 				element = tempFloat32
// 			case "float64":
// 				element = tempFloat64
// 			case "string":
// 				element = tempString
// 			}
// 		}

// 	default:
// 		element = value
// 	}
// 	switch strings.ToLower(ifDataType) {

// 	case "bool":
// 		info, ok := element.([]bool)
// 		if ok {
// 			cnt := len(info)
// 			quantity = int(math.Ceil(float64(len(info)) / 16))
// 			binary.BigEndian.PutUint16(quantity_b, uint16(int16(quantity)))
// 			length = byte(2 * quantity)

// 			switch quantity {
// 			case 0:
// 			case 1:
// 				code = 0x06
// 			default:
// 				code = 0x10
// 			}

// 			value_b = make([]byte, 2*quantity)

// 			switch storeDataType {
// 			case "bit":

// 				for i := 0; i < cnt; i++ {

// 					if info[i] == true {
// 						switch i % 8 {
// 						case 0:
// 							value_b[i/8] |= 0b00000001
// 						case 1:
// 							value_b[i/8] |= 0b00000010
// 						case 2:
// 							value_b[i/8] |= 0b00000100
// 						case 3:
// 							value_b[i/8] |= 0b00001000
// 						case 4:
// 							value_b[i/8] |= 0b00010000
// 						case 5:
// 							value_b[i/8] |= 0b00100000
// 						case 6:
// 							value_b[i/8] |= 0b01000000
// 						case 7:
// 							value_b[i/8] |= 0b10000000
// 						}
// 					}
// 					runtime.Gosched()
// 				}

// 				// case "uint16":

// 				// 	for i := 0; i < cnt; i++ {
// 				// 		start = i * 2
// 				// 		end = (i + 1) * 2
// 				// 		if math.Abs(ratio) < ad_global.CalculationError {
// 				// 			ratio = 1.0
// 				// 		}
// 				// 		binary.BigEndian.PutUint16(value_b[start:end], uint16(int16(int(float32(info[i])/float32(ratio)))))
// 				// 		runtime.Gosched()
// 				// 	}
// 				// case "int16":

// 				// 	for i := 0; i < cnt; i++ {
// 				// 		start = i * 2
// 				// 		end = (i + 1) * 2
// 				// 		if math.Abs(ratio) < ad_global.CalculationError {
// 				// 			ratio = 1.0
// 				// 		}
// 				// 		binary.BigEndian.PutUint16(value_b[start:end], uint16(int16(int(float32(info[i])/float32(ratio)))))
// 				// 		runtime.Gosched()
// 				// 	}

// 				// default:
// 			}
// 		}

// 	case "int":
// 		info, ok := element.([]int)
// 		if ok {
// 			cnt := len(info)
// 			quantity = len(info)
// 			binary.BigEndian.PutUint16(quantity_b, uint16(int16(quantity)))
// 			length = byte(2 * quantity)
// 			switch cnt {
// 			case 0:
// 			case 1:
// 				code = 0x06
// 			default:
// 				code = 0x10
// 			}

// 			value_b = make([]byte, 2*quantity)

// 			switch storeDataType {
// 			case "uint16":

// 				for i := 0; i < cnt; i++ {
// 					start = i * 2
// 					end = (i + 1) * 2
// 					if math.Abs(ratio) < util.CalculationError {
// 						ratio = 1.0
// 					}
// 					binary.BigEndian.PutUint16(value_b[start:end], uint16(int16(int(float32(info[i])/float32(ratio)))))
// 					runtime.Gosched()
// 				}
// 			case "int16":

// 				for i := 0; i < cnt; i++ {
// 					start = i * 2
// 					end = (i + 1) * 2
// 					if math.Abs(ratio) < util.CalculationError {
// 						ratio = 1.0
// 					}
// 					binary.BigEndian.PutUint16(value_b[start:end], uint16(int16(int(float32(info[i])/float32(ratio)))))
// 					runtime.Gosched()
// 				}

// 			default:
// 			}
// 		}

// 	case "float32":
// 		info, ok := element.([]float32)
// 		if ok {

// 			switch storeDataType {
// 			case "uint16":
// 				cnt := len(info)
// 				quantity = len(info)
// 				binary.BigEndian.PutUint16(quantity_b, uint16(int16(quantity)))
// 				length = byte(2 * quantity)

// 				switch cnt {
// 				case 0:
// 				case 1:
// 					code = 0x06
// 				default:
// 					code = 0x10
// 				}

// 				value_b = make([]byte, 2*quantity)
// 				for i := 0; i < cnt; i++ {
// 					start = i * 2
// 					end = (i + 1) * 2
// 					if math.Abs(ratio) < util.CalculationError {
// 						ratio = 1.0
// 					}
// 					binary.BigEndian.PutUint16(value_b[start:end], uint16(int16(int(float32(info[i])/float32(ratio)))))
// 					runtime.Gosched()
// 				}
// 			case "int16":
// 				cnt := len(info)
// 				quantity = len(info)
// 				binary.BigEndian.PutUint16(quantity_b, uint16(int16(quantity)))
// 				length = byte(2 * quantity)
// 				switch cnt {
// 				case 0:
// 				case 1:
// 					code = 0x06
// 				default:
// 					code = 0x10
// 				}
// 				value_b = make([]byte, 2*quantity)
// 				for i := 0; i < cnt; i++ {
// 					start = i * 2
// 					end = (i + 1) * 2
// 					if math.Abs(ratio) < util.CalculationError {
// 						ratio = 1.0
// 					}
// 					binary.BigEndian.PutUint16(value_b[start:end], uint16(int16(int(float32(info[i])/float32(ratio)))))
// 					runtime.Gosched()
// 				}
// 			case "float32":
// 				cnt := len(info)
// 				quantity = 2 * len(info)
// 				binary.BigEndian.PutUint16(quantity_b, uint16(int16(quantity)))
// 				length = byte(2 * quantity)
// 				code = 0x10
// 				value_b = make([]byte, 2*quantity)
// 				for i := 0; i < cnt; i++ {
// 					start = i * 4
// 					end = (i + 1) * 4
// 					if math.Abs(ratio) < util.CalculationError {
// 						ratio = 1.0
// 					}
// 					bits := math.Float32bits(float32(info[i]) / float32(ratio))
// 					binary.BigEndian.PutUint32(value_b[start:end], bits)
// 					runtime.Gosched()
// 				}

// 			default:
// 			}
// 		}
// 	case "float64":
// 		info, ok := element.([]float64)
// 		if ok {
// 			switch storeDataType {
// 			case "uint16":
// 				cnt := len(info)
// 				quantity = len(info)
// 				binary.BigEndian.PutUint16(quantity_b, uint16(int16(quantity)))
// 				length = byte(2 * quantity)
// 				switch cnt {
// 				case 0:
// 				case 1:
// 					code = 0x06
// 				default:
// 					code = 0x10
// 				}
// 				value_b = make([]byte, 2*quantity)
// 				for i := 0; i < cnt; i++ {
// 					start = i * 2
// 					end = (i + 1) * 2
// 					if math.Abs(ratio) < util.CalculationError {
// 						ratio = 1.0
// 					}
// 					binary.BigEndian.PutUint16(value_b[start:end], uint16(int16(int(float32(info[i])/float32(ratio)))))
// 					runtime.Gosched()
// 				}
// 			case "int16":
// 				cnt := len(info)
// 				quantity = len(info)
// 				binary.BigEndian.PutUint16(quantity_b, uint16(int16(quantity)))
// 				length = byte(2 * quantity)
// 				switch cnt {
// 				case 0:
// 				case 1:
// 					code = 0x06
// 				default:
// 					code = 0x10
// 				}

// 				value_b = make([]byte, 2*quantity)
// 				for i := 0; i < cnt; i++ {
// 					start = i * 2
// 					end = (i + 1) * 2
// 					if math.Abs(ratio) < util.CalculationError {
// 						ratio = 1.0
// 					}
// 					binary.BigEndian.PutUint16(value_b[start:end], uint16(int16(int(float32(info[i])/float32(ratio)))))
// 					runtime.Gosched()
// 				}
// 			case "float32":

// 				cnt := len(info)
// 				quantity = 2 * len(info)
// 				binary.BigEndian.PutUint16(quantity_b, uint16(int16(quantity)))
// 				length = byte(2 * quantity)
// 				code = 0x10
// 				value_b = make([]byte, 2*quantity)
// 				for i := 0; i < cnt; i++ {
// 					start = i * 4
// 					end = (i + 1) * 4
// 					if math.Abs(ratio) < util.CalculationError {
// 						ratio = 1.0
// 					}
// 					bits := math.Float32bits(float32(info[i]) / float32(ratio))
// 					binary.BigEndian.PutUint32(value_b[start:end], bits)
// 					runtime.Gosched()
// 				}
// 			default:
// 			}
// 		}
// 	case "string":
// 		info, ok := element.([]string)
// 		if ok {
// 			switch storeDataType {
// 			case "bcd":
// 				cnt := len(info[0])
// 				if cnt%2 == 0 {
// 					if cnt%4 == 0 {
// 						quantity = len(info[0]) / 4
// 					} else {
// 						quantity = len(info[0])/4 + 1
// 					}
// 					binary.BigEndian.PutUint16(quantity_b, uint16(int16(quantity)))
// 					length = byte(2 * quantity)
// 					code = 0x10
// 					value_b = make([]byte, 2*quantity)
// 					a := []byte(info[0])

// 					for i := 0; i < cnt; i = i + 2 {
// 						value_b[i/2] = a[i]<<4 | a[i+1]
// 						runtime.Gosched()
// 					}
// 				}

// 			default:
// 			}
// 		}

// 	}

// 	b := make([]byte, 0)
// 	switch strings.ToLower(protocolType) {
// 	case "modbus-rtu":

// 		switch code {
// 		case 0x06:
// 			b = append(b, slave, code)
// 			b = append(b, addr_pdu_b...)
// 			b = append(b, value_b...)
// 		case 0x10:
// 			b = append(b, slave, code)
// 			b = append(b, addr_pdu_b...)
// 			b = append(b, quantity_b...)
// 			b = append(b, length)
// 			b = append(b, value_b...)
// 		}

// 		crc.Reset()
// 		crc.PushBytes(b)
// 		b = append(b, crc.Low, crc.High)
// 	case "modbus-tcp":
// 		fallthrough
// 	default:

// 		switch code {
// 		case 0x06:
// 			mbap = append(mbap, addr_mbap_b...) //事务标识符 2个字节
// 			mbap = append(mbap, 0x00, 0x00)     //协议标识符 2个字节

// 			size := byte(unsafe.Sizeof(slave)) + byte(unsafe.Sizeof(code)) + byte(len(addr_pdu_b)) + byte(len(value_b)) //
// 			mbap = append(mbap, 0x00, size)                                                                             //数据帧长度 2个字节
// 			mbap = append(mbap, slave)

// 			b = append(b, mbap...)

// 			b = append(b, code)
// 			b = append(b, addr_pdu_b...)
// 			b = append(b, value_b...)
// 		case 0x10:
// 			mbap = append(mbap, addr_mbap_b...) //事务标识符 2个字节
// 			mbap = append(mbap, 0x00, 0x00)     //协议标识符 2个字节

// 			size := byte(unsafe.Sizeof(slave)) + byte(unsafe.Sizeof(code)) + byte(len(addr_pdu_b)+len(quantity_b)) + byte(unsafe.Sizeof(length)) + byte(len(value_b)) //
// 			mbap = append(mbap, 0x00, size)                                                                                                                           //数据帧长度 2个字节
// 			mbap = append(mbap, slave)

// 			b = append(b, mbap...)

// 			b = append(b, code)
// 			b = append(b, addr_pdu_b...)
// 			b = append(b, quantity_b...)
// 			b = append(b, length)
// 			b = append(b, value_b...)
// 		}

// 	}

// 	return string(b)

// }
// func Encode_ReadCoils(protocolType string, storeDataType string, slave byte, addr int, quantity int) ([]int, []string) {

// 	var crc util_crc.Crc
// 	var mbap []byte
// 	var code byte
// 	var addr_mbap int
// 	var addr_pdu int
// 	var quantity_pdu int
// 	var addr_res []int
// 	var adu_res []string

// 	code = 0x01

// 	sendQuantity := 1000
// 	var times int
// 	if quantity%sendQuantity == 0 {
// 		times = quantity / sendQuantity
// 	} else {
// 		times = quantity/sendQuantity + 1 //当小于拆分发送寄存器长度时，按实际长度发送
// 	}
// 	// times := quantity/sendQuantity + 1 //当小于拆分发送数量时，按实际数量发送

// 	for i := 0; i < times; i++ {

// 		addr_mbap = addr + i*sendQuantity
// 		addr_pdu = addr - 1 + i*sendQuantity

// 		if i == times-1 { //当小于拆分发送数量时，按实际数量发送
// 			quantity_pdu = quantity - sendQuantity*i
// 		} else {
// 			quantity_pdu = sendQuantity
// 		}

// 		addr_mbap_b := make([]byte, 2)
// 		addr_pdu_b := make([]byte, 2)
// 		quantity_b := make([]byte, 2)

// 		binary.BigEndian.PutUint16(addr_mbap_b, uint16(int16(addr_mbap)))
// 		binary.BigEndian.PutUint16(addr_pdu_b, uint16(int16(addr_pdu)))
// 		binary.BigEndian.PutUint16(quantity_b, uint16(int16(quantity_pdu)))

// 		b := make([]byte, 0)
// 		switch strings.ToLower(protocolType) {

// 		case "modbus-rtu": //Modbus RTU协议,被动响应

// 			b = append(b, slave)
// 			b = append(b, code)
// 			b = append(b, addr_pdu_b...)
// 			b = append(b, quantity_b...)

// 			crc.Reset()
// 			crc.PushBytes(b)
// 			b = append(b, crc.Low, crc.High)

// 		case "modbus-tcp": //Modbus TCP协议,被动响应
// 			fallthrough
// 		default:

// 			mbap = make([]byte, 0)
// 			mbap = append(mbap, addr_mbap_b...)                                                                    //事务标识符 2个字节
// 			mbap = append(mbap, 0x00, 0x00)                                                                        //协议标识符 2个字节
// 			size := byte(unsafe.Sizeof(slave)) + byte(unsafe.Sizeof(code)) + byte(len(addr_pdu_b)+len(quantity_b)) //
// 			mbap = append(mbap, 0x00, size)                                                                        //数据帧长度 2个字节
// 			mbap = append(mbap, slave)

// 			b = append(b, mbap...)
// 			b = append(b, code)
// 			b = append(b, addr_pdu_b...)
// 			b = append(b, quantity_b...)

// 		}
// 		addr_res = append(addr_res, addr_mbap) //响应的地址
// 		adu_res = append(adu_res, string(b))
// 		runtime.Gosched()
// 	}

// 	return addr_res, adu_res
// }
// func Encode_WriteCoils(protocolType string, storeDataType string, ifDataType string, slave byte, addr int, value interface{}) string {

// 	var crc util_crc.Crc
// 	// var start, end int
// 	var mbap []byte
// 	var addr_mbap int
// 	var addr_pdu int
// 	var code byte
// 	var length byte
// 	var quantity int

// 	// var addr_pdu int
// 	// var quantity_pdu int
// 	// var addr_res []int
// 	// var adu_res []string

// 	// sendQuantity := 1000
// 	// times := quantity/sendQuantity + 1 //当小于拆分发送数量时，按实际数量发送

// 	// for i := 0; i < times; i++ {

// 	// 	addr_pdu = addr - 1 + i*sendQuantity
// 	// 	binary.BigEndian.PutUint16(addr_pdu_b, uint16(int16(addr_pdu)))
// 	// 	if i == times-1 { //当小于拆分发送数量时，按实际数量发送
// 	// 		quantity_pdu = quantity - sendQuantity*i
// 	// 	} else {
// 	// 		quantity_pdu = sendQuantity
// 	// 	}

// 	// 	addr_pdu_b := make([]byte, 2)
// 	// 	quantity_b := make([]byte, 2)
// 	// 	binary.BigEndian.PutUint16(quantity_b, uint16(int16(quantity_pdu)))

// 	// }
// 	addr_mbap = addr
// 	addr_pdu = addr - 1

// 	addr_mbap_b := make([]byte, 2)
// 	addr_pdu_b := make([]byte, 2)
// 	binary.BigEndian.PutUint16(addr_mbap_b, uint16(int16(addr_mbap)))
// 	binary.BigEndian.PutUint16(addr_pdu_b, uint16(int16(addr_pdu)))

// 	quantity_b := make([]byte, 2)
// 	value_b := make([]byte, 0)
// 	b := make([]byte, 0)
// 	var elements []interface{}
// 	var element interface{}
// 	var tempBool []bool
// 	var ok bool
// 	switch reflect.TypeOf(value).String() { //18200界面json请求时生成的格式为[]interface,代码直接请求时生成的为interface
// 	case "[]interface {}":
// 		elements, ok = value.([]interface{})
// 		if ok {
// 			for _, v := range elements {
// 				t, ok := v.(bool)
// 				if ok {
// 					tempBool = append(tempBool, t)
// 				}
// 				runtime.Gosched()
// 			}
// 			element = tempBool
// 		}
// 	default:
// 		element = value
// 	}
// 	switch strings.ToLower(ifDataType) {
// 	case "bool":

// 		info, ok := element.([]bool)
// 		if ok {
// 			quantity = len(info)
// 			switch quantity {
// 			case 0:
// 			case 1:
// 				binary.BigEndian.PutUint16(quantity_b, uint16(int16(quantity)))

// 				if info[0] == true { //对应ON
// 					value_b = append(value_b, 0xFF, 0x00)
// 				} else { //对应OFF
// 					value_b = append(value_b, 0x00, 0x00)
// 				}

// 			default:
// 				binary.BigEndian.PutUint16(quantity_b, uint16(int16(quantity)))
// 				if quantity%8 == 0 {
// 					length = byte(quantity / 8)
// 					value_b = make([]byte, quantity/8)
// 				} else {
// 					length = byte(quantity/8 + 1)
// 					value_b = make([]byte, quantity/8+1)
// 				}
// 				for i := 0; i < quantity; i++ {

// 					if info[i] == true {
// 						switch i % 8 {
// 						case 0:
// 							value_b[i/8] |= 0b00000001
// 						case 1:
// 							value_b[i/8] |= 0b00000010
// 						case 2:
// 							value_b[i/8] |= 0b00000100
// 						case 3:
// 							value_b[i/8] |= 0b00001000
// 						case 4:
// 							value_b[i/8] |= 0b00010000
// 						case 5:
// 							value_b[i/8] |= 0b00100000
// 						case 6:
// 							value_b[i/8] |= 0b01000000
// 						case 7:
// 							value_b[i/8] |= 0b10000000
// 						}
// 					}
// 					runtime.Gosched()
// 				}
// 			}
// 		}
// 	}

// 	switch strings.ToLower(protocolType) {
// 	case "modbus-rtu":

// 		switch quantity { //此处用数量，不用len(value),不同于写保持寄存器
// 		case 0:
// 		case 1: //写单个线圈
// 			code = 0x05

// 			b = append(b, slave, code)
// 			b = append(b, addr_pdu_b...)
// 			b = append(b, value_b...)

// 		default: //写多个寄存器
// 			code = 0x0f

// 			b = append(b, slave)
// 			b = append(b, code)
// 			b = append(b, addr_pdu_b...)
// 			b = append(b, quantity_b...)
// 			b = append(b, length)
// 			b = append(b, value_b...)
// 		}
// 		crc.Reset()
// 		crc.PushBytes(b)
// 		b = append(b, crc.Low, crc.High)
// 	case "modbus-tcp":
// 		fallthrough
// 	default:

// 		mbap = append(mbap, addr_mbap_b...) //事务标识符 2个字节
// 		mbap = append(mbap, 0x00, 0x00)     //协议标识符 2个字节
// 		switch quantity {                   //此处用数量，不用len(value),不同于写保持寄存器

// 		case 0:
// 		case 1: //写单个线圈
// 			code = 0x05

// 			size := byte(unsafe.Sizeof(slave)) + byte(unsafe.Sizeof(code)) + byte(len(addr_pdu_b)+len(value_b)) //
// 			mbap = append(mbap, 0x00, size)                                                                     //数据帧长度 2个字节
// 			mbap = append(mbap, slave)

// 			b = append(b, mbap...)

// 			b = append(b, code)
// 			b = append(b, addr_pdu_b...)
// 			b = append(b, value_b...)

// 		default: //写多个线圈
// 			code = 0x0f

// 			size := byte(unsafe.Sizeof(slave)) + byte(unsafe.Sizeof(code)) + byte(len(addr_pdu_b)+len(quantity_b)) + byte(unsafe.Sizeof(length)) + byte(len(value_b)) //
// 			mbap = append(mbap, 0x00, size)                                                                                                                           //数据帧长度 2个字节
// 			mbap = append(mbap, slave)

// 			b = append(b, mbap...)

// 			b = append(b, code)
// 			b = append(b, addr_pdu_b...)
// 			b = append(b, quantity_b...)
// 			b = append(b, length)
// 			b = append(b, value_b...)
// 		}

// 	}

// 	return string(b)

// }
// func Encode_ReadInputRegisters(protocolType string, storeDataType string, slave byte, addr int, quantity int) ([]int, []string) {

// 	var crc util_crc.Crc
// 	var mbap []byte
// 	var code byte
// 	var addr_mbap int
// 	var addr_pdu int
// 	var quantity_pdu int
// 	var addr_res []int
// 	var adu_res []string

// 	code = 0x04

// 	switch strings.ToLower(storeDataType) { //asc bit byte int16 uint16 float32 bcd
// 	case "asc":
// 	case "bit":
// 		quantity = int(math.Ceil(float64(quantity) / 16))
// 	case "byte":
// 		quantity = int(math.Ceil(float64(quantity) / 2))
// 	case "int16":
// 	case "uint16":
// 	case "float32":
// 		quantity = 2 * quantity
// 	case "bcd":
// 		quantity = int(math.Ceil(float64(quantity) / 4))
// 	}

// 	sendQuantity := 100
// 	var times int
// 	if quantity%sendQuantity == 0 {
// 		times = quantity / sendQuantity
// 	} else {
// 		times = quantity/sendQuantity + 1 //当小于拆分发送寄存器长度时，按实际长度发送
// 	}
// 	// times := quantity/sendQuantity + 1 //当小于拆分发送寄存器长度时，按实际长度发送

// 	for i := 0; i < times; i++ {

// 		addr_mbap = addr + i*sendQuantity
// 		if addr > 300000 {
// 			addr_pdu = addr - 1 - 300000 + i*sendQuantity
// 		} else {
// 			addr_pdu = addr - 1 - 30000 + i*sendQuantity
// 		}

// 		if i == times-1 { //当小于拆分发送寄存器长度时，按实际长度发送
// 			quantity_pdu = quantity - sendQuantity*i
// 		} else {
// 			quantity_pdu = sendQuantity
// 		}

// 		addr_mbap_b := make([]byte, 2)
// 		addr_pdu_b := make([]byte, 2)
// 		quantity_b := make([]byte, 2)

// 		binary.BigEndian.PutUint16(addr_mbap_b, uint16(int16(addr_mbap)))
// 		binary.BigEndian.PutUint16(addr_pdu_b, uint16(int16(addr_pdu)))
// 		binary.BigEndian.PutUint16(quantity_b, uint16(int16(quantity_pdu)))

// 		b := make([]byte, 0)

// 		switch strings.ToLower(protocolType) {

// 		case "modbus-rtu": //被动响应

// 			b = append(b, slave)
// 			b = append(b, code)
// 			b = append(b, addr_pdu_b...)
// 			b = append(b, quantity_b...)

// 			crc.Reset()
// 			crc.PushBytes(b)
// 			b = append(b, crc.Low, crc.High)

// 		case "modbus-tcp": //被动响应
// 			fallthrough
// 		default:

// 			mbap = make([]byte, 0)
// 			mbap = append(mbap, addr_mbap_b...)                                                                    //事务标识符 2个字节
// 			mbap = append(mbap, 0x00, 0x00)                                                                        //协议标识符 2个字节
// 			size := byte(unsafe.Sizeof(slave)) + byte(unsafe.Sizeof(code)) + byte(len(addr_pdu_b)+len(quantity_b)) //
// 			mbap = append(mbap, 0x00, size)                                                                        //数据帧长度 2个字节
// 			mbap = append(mbap, slave)

// 			b = append(b, mbap...)

// 			b = append(b, code)
// 			b = append(b, addr_pdu_b...)
// 			b = append(b, quantity_b...)

// 		}
// 		addr_res = append(addr_res, addr_mbap)
// 		adu_res = append(adu_res, string(b))
// 		runtime.Gosched()
// 	}

// 	return addr_res, adu_res
// }
// func Encode_ReadDiscreteInputs(protocolType string, storeDataType string, slave byte, addr int, quantity int) ([]int, []string) {

// 	var crc util_crc.Crc
// 	var mbap []byte
// 	var code byte
// 	var addr_mbap int
// 	var addr_pdu int
// 	var quantity_pdu int
// 	var addr_res []int
// 	var adu_res []string

// 	code = 0x02

// 	sendQuantity := 1000
// 	var times int
// 	if quantity%sendQuantity == 0 {
// 		times = quantity / sendQuantity
// 	} else {
// 		times = quantity/sendQuantity + 1 //当小于拆分发送寄存器长度时，按实际长度发送
// 	}
// 	// times := quantity/sendQuantity + 1 //当小于拆分发送数量时，按实际数量发送

// 	for i := 0; i < times; i++ {

// 		addr_mbap = addr + i*sendQuantity
// 		if addr > 100000 {
// 			addr_pdu = addr - 1 - 100000 + i*sendQuantity
// 		} else {
// 			addr_pdu = addr - 1 - 10000 + i*sendQuantity
// 		}

// 		if i == times-1 { //当小于拆分发送数量时，按实际数量发送
// 			quantity_pdu = quantity - sendQuantity*i
// 		} else {
// 			quantity_pdu = sendQuantity
// 		}

// 		addr_mbap_b := make([]byte, 2)
// 		addr_pdu_b := make([]byte, 2)
// 		quantity_b := make([]byte, 2)

// 		binary.BigEndian.PutUint16(addr_mbap_b, uint16(int16(addr_mbap)))
// 		binary.BigEndian.PutUint16(addr_pdu_b, uint16(int16(addr_pdu)))
// 		binary.BigEndian.PutUint16(quantity_b, uint16(int16(quantity_pdu)))

// 		b := make([]byte, 0)

// 		switch strings.ToLower(protocolType) {

// 		case "modbus-rtu": //Modbus RTU协议,被动响应

// 			b = append(b, slave)
// 			b = append(b, code)
// 			b = append(b, addr_pdu_b...)
// 			b = append(b, quantity_b...)

// 			crc.Reset()
// 			crc.PushBytes(b)
// 			b = append(b, crc.Low, crc.High)

// 		case "modbus-tcp": //Modbus TCP协议,被动响应
// 			fallthrough
// 		default:

// 			mbap = make([]byte, 0)
// 			mbap = append(mbap, addr_mbap_b...)                                                                    //事务标识符 2个字节
// 			mbap = append(mbap, 0x00, 0x00)                                                                        //协议标识符 2个字节
// 			size := byte(unsafe.Sizeof(slave)) + byte(unsafe.Sizeof(code)) + byte(len(addr_pdu_b)+len(quantity_b)) //
// 			mbap = append(mbap, 0x00, size)                                                                        //数据帧长度 2个字节
// 			mbap = append(mbap, slave)

// 			b = append(b, mbap...)
// 			b = append(b, code)
// 			b = append(b, addr_pdu_b...)
// 			b = append(b, quantity_b...)

// 		}
// 		addr_res = append(addr_res, addr_mbap)
// 		adu_res = append(adu_res, string(b))
// 		runtime.Gosched()
// 	}

// 	return addr_res, adu_res
// }

// Bit access:

// Read Discrete Inputs
// Read Coils
// Write Single Coil
// Write Multiple Coils
// 16-bit acess:

// Read Input Registers
// Read Multiple Holding Registers
// Write Single Holding Register
// Write Multiple Holding Registers

// http://blog.sina.com.cn/s/blog_72d911930102wrlq.html
// Modbus协议定义的寄存器地址是5位十进制地址，即：
// 线圈（DO）地址：00001~09999
// 触点（DI）地址：10001~19999
// 输入寄存器（AI）地址：30001~39999
// 输出寄存器（AO）地址：40001~49999
// 由于上述各类地址是唯一对应的，因此有些资料就以其第一个数字区分各类地址，即：0x代表线圈（DO）类地址，1x代表触点（DI）类地址、 3x代表输入寄存器（AI）类地址、4x代表输出寄存器（AO）类地址。
// 功能码01（读线圈状态）对应的地址是线圈的地址(即位地址)，若要求其所在的寄存器地址，计算如下：
// 线圈的位地址/16=整数商（即寄存器地址）+余数（即位偏移地址）
