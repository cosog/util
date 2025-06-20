// decode.go
package util_decode

import (
	"encoding/binary"

	"github.com/cosog/util/crc"
)

func Decode_Modbus_TCP(content string) (byte, int, interface{}, bool, bool, bool, bool) {

	// MBAP为报文头，长度为7字节
	// 事务处理标识 ：可以理解为报文的序列号，一般每次通信之后就要加1以区别不同的通信数据报文。
	// 协议标识符 ：00 00表示ModbusTCP协议。
	// 长度 ：表示接下来的数据长度，单位为字节。
	// 单元标识符 ：可以理解为设备地址。
	// PDU由功能码+数据组成。功能码为1字节，数据长度不定，由具体功能决定。

	// modbus的操作对象有四种：线圈、离散输入、输入寄存器、保持寄存器。
	// 线圈：PLC的输出位，开关量，在MODBUS中可读可写
	// 离散量：PLC的输入位，开关量，在MODBUS中只读
	// 输入寄存器：PLC中只能从模拟量输入端改变的寄存器，在MODBUS中只读
	// 保持寄存器：PLC中用于输出模拟量信号的寄存器，在MODBUS中可读可写
	// 根据对象的不同，modbus的功能码有：
	// 0x01：读线圈
	// 0x05：写单个线圈
	// 0x0F：写多个线圈
	// 0x02：读离散量输入
	// 0x04：读输入寄存器
	// 0x03：读保持寄存器
	// 0x06：写单个保持寄存器
	// 0x10：写多个保持寄存器

	// – slave ID：从站编号（事务标识符）
	// – function：功能码，0x01对应线圈操作，0x02对应离散量操作，0x03对应保持寄存器操作，0x04对应输入寄存器操作
	// – address：开始地址
	// – quantity：寄存器/线圈/离散量 的数量
	// modbus设备可分为主站(poll)和从站(slave)。主站只有一个，从站有多个，主站向各从站发送请求帧，从站给予响应。在使用TCP通信时，主站为client端，主动建立连接；从站为server端，等待连接。
	// 主站请求：功能码+数据
	// 从站正常响应：请求功能码+响应数据
	// 从站异常响应：异常功能码+异常码，其中异常功能码即将请求功能码的最高有效位置1，异常码指示差错类型

	var slave byte
	var addr int
	var value interface{}
	// var quantity int
	var size int
	var code byte
	var m_err bool = false
	var exception_code_err bool = false
	var crc_err bool = false
	var filter bool = false //过滤器  true-过滤  false-不过滤

	r := []byte(content)
	// log.Println("r:", r)

	if len(r) >= 8 {
		// addr = int(binary.BigEndian.Uint16(r[0:2])) //事务处理标识
		addr = int(binary.BigEndian.Uint16(r[0:2])) //事务处理标识
		// binary.BigEndian.Uint16(r[2:4])//协议标识符
		// size = int(binary.BigEndian.Uint16(r[4:6]))//长度

		slave = r[6] //单元标识符 ：可以理解为设备地址
		code = r[7]  //功能码

		// aduGroups, ok := m_slave_acqADUGroup[slave]
		// if ok {

		// if len(aduGroups) > 0 {

		// for _, aduGroup := range aduGroups {
		// for k, g_addr := range aduGroup.Addr {
		// if g_addr == addr {

		crc_err = false
		if code&0x80 == 0x80 {
			exception_code_err = true
			switch code {
			case 0x01 | 0x80:
			case 0x02 | 0x80:
			case 0x03 | 0x80:
			case 0x04 | 0x80:
			case 0x05 | 0x80:
			case 0x06 | 0x80:
			case 0x0f | 0x80:
			case 0x10 | 0x80:
			case 0x14 | 0x80:
			case 0x15 | 0x80:
			case 0x16 | 0x80:
			case 0x17 | 0x80:
			case 0x2B | 0x80:
			}
		} else {

			switch code {
			case 0x01: //读线圈	响应pdu:功能码1字节+字节数1字节(N*)+线圈状态(N个字节) N＝输出数量/8，如果余数不等于0，那么N = N+1
				if len(r) >= 9 {
					size = int(r[8])
					if len(r) >= 9+size {
						value = r[9 : 9+size]
					} else {
						crc_err = true
					}
				} else {
					crc_err = true
				}

			case 0x02: //读离散量输入	响应pdu:功能码1字节+字节数1字节(N*)+输入状态(N个字节) N＝输出数量/8，如果余数不等于0，那么N = N+1
				if len(r) >= 9 {
					size = int(r[8])
					if len(r) >= 9+size {
						value = r[9 : 9+size]
					} else {
						crc_err = true
					}
				} else {
					crc_err = true
				}
			case 0x03: //读保持寄存器	响应pdu:功能码1字节+字节数1字节(2*N)+寄存器值(N*2字节)
				if len(r) >= 9 {
					size = int(r[8])
					if len(r) >= 9+size {
						value = r[9 : 9+size]
					} else {
						crc_err = true
					}
				} else {
					crc_err = true
				}
			case 0x04: //读输入寄存器  响应pdu  功能码1字节+字节数1字节(2*N)+寄存器值(N*2字节)
				if len(r) >= 9 {
					size = int(r[8])
					if len(r) >= 9+size {
						value = r[9 : 9+size]
					} else {
						crc_err = true
					}
				} else {
					crc_err = true
				}
			case 0x05: //写单个线圈	响应pdu:功能码1字节+输出地址2字节+输出值2字节

				// if aduGroup.AcqMode[k] == "active" {
				// 	//预留modbus-tcp拓展主动上传模式
				// } else {
				filter = true

				if len(r) >= 12 {

				}
				// }
			case 0x06: //写单个保持寄存器	响应pdu:功能码1字节+寄存器地址2字节+寄存器值2字节
				// if aduGroup.AcqMode[k] == "active" {
				// 	//预留modbus-tcp拓展主动上传模式
				// } else {
				filter = true

				if len(r) >= 12 {
					// addr = append(addr, int(binary.BigEndian.Uint16(r[8:10])))
					// value = r[10:12]

				}
				// }

			case 0x0f: //写多个线圈	响应pdu:功能码1字节+起始地址2字节+输出数量2字节
				// if aduGroup.AcqMode[k] == "active" {
				// 	//预留modbus-tcp拓展主动上传模式
				// } else {
				filter = true

				// }

			case 0x10: //写多个保持寄存器	响应pdu:功能码1字节+起始地址2字节+寄存器数量2字节
				// if aduGroup.AcqMode[k] == "active" {
				// 	//预留modbus-tcp拓展主动上传模式
				// } else {
				filter = true

				if len(r) >= 12 {
					// addr = append(addr, int(binary.BigEndian.Uint16(r[8:10])))
					// quantity = int(binary.BigEndian.Uint16(r[10:12]))
				}
				// }

			// case 0x14:
			// case 0x15:
			// case 0x16:
			// case 0x17:
			// case 0x2B:
			default:
				crc_err = true
			}
		}
		// }
		// runtime.Gosched()
		// }
		// runtime.Gosched()
		// }
		// } else {
		// 	m_err = true
		// }

		// } else {
		// 	m_err = true
		// }

	} else {
		crc_err = true
	}
	return slave, addr, value, m_err, exception_code_err, crc_err, filter
}
func Decode_Modbus_RTU(addr int, content string) (byte, int, interface{}, bool, bool, bool, bool) {

	var crc util_crc.Crc
	var slave byte
	var value interface{}
	// var quantity int
	var size int
	var code byte
	var m_err bool = false
	var exception_code_err bool = false
	var crc_err bool = false
	var filter bool = false //过滤器  true-过滤  false-不过滤

	r := []byte(content)

	if len(r) >= 2 {

		slave = r[0]
		code = r[1]

		crc_err = false
		if code&0x80 == 0x80 {
			exception_code_err = true
			switch code {
			case 0x01 | 0x80:
			case 0x02 | 0x80:
			case 0x03 | 0x80:
			case 0x04 | 0x80:
			case 0x05 | 0x80:
			case 0x06 | 0x80:
			case 0x0f | 0x80:
			case 0x10 | 0x80:
			case 0x14 | 0x80:
			case 0x15 | 0x80:
			case 0x16 | 0x80:
			case 0x17 | 0x80:
			case 0x2B | 0x80:
			}
		} else {
			switch code {
			case 0x01: //读线圈	响应pdu:功能码1字节+字节数1字节(N*)+线圈状态(N个字节) N＝输出数量/8，如果余数不等于0，那么N = N+1
				if len(r) >= 3 {
					size = int(r[2])
					if len(r) >= 3+size {
						value = r[3 : 3+size]

					} else {
						crc_err = true
					}
				} else {
					crc_err = true
				}

			case 0x02: //读离散量输入	响应pdu:功能码1字节+字节数1字节(N*)+输入状态(N个字节) N＝输出数量/8，如果余数不等于0，那么N = N+1
				if len(r) >= 3 {
					size = int(r[2])
					if len(r) >= 3+size {
						value = r[3 : 3+size]
					} else {
						crc_err = true
					}
				} else {
					crc_err = true
				}

			case 0x03: //读保持寄存器	响应pdu  功能码1字节+字节数1字节(2*N)+寄存器值(N*2字节)
				if len(r) >= 3 {
					size = int(r[2])
					if len(r) >= 3+size {

						value = r[3 : 3+size]
						if len(r) >= 3+size+2 {
							crc.Reset()               //
							crc.PushBytes(r[:3+size]) //
							if crc.Low == r[3+size] && crc.High == r[3+size+1] {
								crc_err = false
							} else {
								crc_err = true
							}
						} else {
							crc_err = true
						}
					} else {
						crc_err = true
					}
				} else {
					crc_err = true
				}
			case 0x04: //读输入寄存器  响应pdu  功能码1字节+字节数1字节(2*N)+寄存器值(N*2字节)
				if len(r) >= 3 {
					size = int(r[2])
					if len(r) >= 3+size {
						value = r[3 : 3+size]
					} else {
						crc_err = true
					}
				} else {
					crc_err = true
				}

			case 0x05: //写单个线圈	响应pdu:功能码1字节+输出地址2字节+输出值2字节

				// if aduGroup.AcqMode[k] == "active" {
				// 	//预留modbus-tcp拓展主动上传模式
				// } else {
				filter = true

				if len(r) >= 6 {

				}
				// }

			case 0x06: //写单个保持寄存器	响应pdu:功能码1字节+寄存器地址2字节+寄存器值2字节
				// if aduGroup.AcqMode[k] == "active" {
				// 	//预留modbus-tcp拓展主动上传模式
				// } else {
				filter = true

				if len(r) >= 6 {
					// addr = append(addr, int(binary.BigEndian.Uint16(r[8:10])))
					// value = r[10:12]

				}
				// }
			case 0x0f: //写多个线圈	响应pdu:功能码1字节+起始地址2字节+输出数量2字节
				// if aduGroup.AcqMode[k] == "active" {
				// 	//预留modbus-tcp拓展主动上传模式
				// } else {
				filter = true

				// }

			case 0x10: //写多个保持寄存器	响应pdu  功能码1字节+起始地址2字节+寄存器数量2字节

				// if aduGroup.AcqMode[k] == "active" {
				// 	//预留modbus-tcp拓展主动上传模式
				// } else {
				filter = true

				if len(r) >= 6 {
					// addr = append(addr, int(binary.BigEndian.Uint16(r[8:10])))
					// quantity = int(binary.BigEndian.Uint16(r[10:12]))
				}
				// }
			// case 0x14:
			// case 0x15:
			// case 0x16:
			// case 0x17:
			// case 0x2B:
			default:
				crc_err = true
			}
		}

	} else {
		crc_err = true
	}

	return slave, addr, value, m_err, exception_code_err, crc_err, filter
}
