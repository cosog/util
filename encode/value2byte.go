// value2byte.go
package util_encode

import (
	"encoding/binary"
	"math"

	"strconv"
	"strings"
)

func Value2Byte(storeMode string, quantity int, StoreDataType string, IFDataType string, ratio float64, value interface{}) []byte {
	var b []byte
	var s []byte
	switch strings.ToLower(IFDataType) {
	case "string":
		v, ok := value.(string)
		if ok {
			switch strings.ToLower(storeMode) {
			case "little":
				switch strings.ToLower(StoreDataType) {
				case "bcd":
					var bcd_msb int
					var bcd_lsb int
					var valid bool = true
					b := make([]byte, 0)
					temp := make([]byte, int(math.Ceil(float64(len(v))/2.0)))
					for i := 0; i < len(v); i++ {
						if !(v[i] == '0' || v[i] == '1' || v[i] == '2' || v[i] == '3' || v[i] == '4' || v[i] == '5' || v[i] == '6' || v[i] == '7' || v[i] == '8' || v[i] == '9') {
							valid = false
							break
						}
					}
					if valid == true {
						for i := 0; i < len(v); i = i + 2 {
							if i+1 < len(v) {
								bcd_msb, _ = strconv.Atoi(string(v[i]))
								bcd_lsb, _ = strconv.Atoi(string(v[i+1]))
							} else {
								bcd_msb = 0b00000000
								bcd_lsb, _ = strconv.Atoi(string(v[i]))
							}
							temp[i/2] = (byte(bcd_msb) << 4 & 0b11110000) | (byte(bcd_lsb) & 0b00001111)
						}
						// b = append(b, temp...)
					}
					for j := len(temp) - 1; j >= 0; j-- {
						b = append(b, temp[j])
					}
				}
			case "big":
				fallthrough
			default:
				switch strings.ToLower(StoreDataType) {
				case "bcd":
					var bcd_msb int
					var bcd_lsb int
					var valid bool = true
					b := make([]byte, 0)
					temp := make([]byte, int(math.Ceil(float64(len(v))/2.0)))
					for i := 0; i < len(v); i++ {
						if !(v[i] == '0' || v[i] == '1' || v[i] == '2' || v[i] == '3' || v[i] == '4' || v[i] == '5' || v[i] == '6' || v[i] == '7' || v[i] == '8' || v[i] == '9') {
							valid = false
							break
						}
					}
					if valid == true {
						for i := 0; i < len(v); i = i + 2 {
							if i+1 < len(v) {
								bcd_msb, _ = strconv.Atoi(string(v[i]))
								bcd_lsb, _ = strconv.Atoi(string(v[i+1]))
							} else {
								bcd_msb = 0b00000000
								bcd_lsb, _ = strconv.Atoi(string(v[i]))
							}
							temp[i/2] = (byte(bcd_msb) << 4 & 0b11110000) | (byte(bcd_lsb) & 0b00001111)
						}
						b = append(b, temp...)
					}
				case "byte":
					b := make([]byte, 0)
					b = append(b, []byte(v)...)
				}
			}
		}
	case "float64":
		v, ok := value.(float64)
		if ok {
			switch strings.ToLower(storeMode) {
			case "little":
				switch strings.ToLower(StoreDataType) {
				case "uint16":
					b = make([]byte, 2)
					binary.LittleEndian.PutUint16(b, uint16(uint64(v/float64(ratio))))
				case "int16":
					b = make([]byte, 2)
					binary.LittleEndian.PutUint16(b, uint16(int16(int64(v/float64(ratio)))))
				case "float32":
					b = make([]byte, 4)
					bits := math.Float32bits(float32(v / ratio))
					binary.LittleEndian.PutUint32(b, bits)
				}
			case "big":
				fallthrough
			default:
				switch strings.ToLower(StoreDataType) {
				case "uint16":
					b = make([]byte, 2)
					binary.BigEndian.PutUint16(b, uint16(uint64(v/float64(ratio))))
				case "int16":
					b = make([]byte, 2)
					binary.BigEndian.PutUint16(b, uint16(int16(int64(v/float64(ratio)))))
				case "float32":
					b = make([]byte, 4)
					bits := math.Float32bits(float32(v / ratio))
					binary.BigEndian.PutUint32(b, bits)
				}
			}
		}
	case "float32":
		v, ok := value.(float32)
		if ok {
			switch strings.ToLower(storeMode) {
			case "little":
				switch strings.ToLower(StoreDataType) {
				case "uint16":
					b = make([]byte, 2)
					binary.LittleEndian.PutUint16(b, uint16(uint32(v/float32(ratio))))
				case "int16":
					b = make([]byte, 2)
					binary.LittleEndian.PutUint16(b, uint16(int16(int32(v/float32(ratio)))))
				case "float32":
					b = make([]byte, 4)
					bits := math.Float32bits(v / float32(ratio))
					binary.LittleEndian.PutUint32(b, bits)
				}
			case "big":
				fallthrough
			default:
				switch strings.ToLower(StoreDataType) {
				case "uint16":
					b = make([]byte, 2)
					binary.BigEndian.PutUint16(b, uint16(uint32(v/float32(ratio))))
				case "int16":
					b = make([]byte, 2)
					binary.BigEndian.PutUint16(b, uint16(int16(int32(v/float32(ratio)))))
				case "float32":
					b = make([]byte, 4)
					bits := math.Float32bits(v / float32(ratio))
					binary.BigEndian.PutUint32(b, bits)
				}
			}

		}
	case "int":
		v, ok := value.(int)
		if ok {

			switch strings.ToLower(storeMode) {
			case "little":
				switch strings.ToLower(StoreDataType) {
				case "uint16":
					b = make([]byte, 2)
					binary.LittleEndian.PutUint16(b, uint16(uint(float64(v)/ratio)))
				case "int16":
					b = make([]byte, 2)
					binary.LittleEndian.PutUint16(b, uint16(int16(int(float64(v)/ratio))))
				}
			case "big":
				fallthrough
			default:
				switch strings.ToLower(StoreDataType) {
				case "uint16":
					b = make([]byte, 2)
					binary.BigEndian.PutUint16(b, uint16(uint(float64(v)/ratio)))
				case "int16":
					b = make([]byte, 2)
					binary.BigEndian.PutUint16(b, uint16(int16(int(float64(v)/ratio))))
				}
			}
		}
	case "int64":
		v, ok := value.(int64)
		if ok {

			switch strings.ToLower(storeMode) {
			case "little":
				switch strings.ToLower(StoreDataType) {
				case "uint16":
					b = make([]byte, 2)
					binary.LittleEndian.PutUint16(b, uint16(uint64(float64(v)/ratio)))
				case "int16":
					b = make([]byte, 2)
					binary.LittleEndian.PutUint16(b, uint16(int16(int64(float64(v)/ratio))))
				}
			case "big":
				fallthrough
			default:
				switch strings.ToLower(StoreDataType) {
				case "uint16":
					b = make([]byte, 2)
					binary.BigEndian.PutUint16(b, uint16(uint64(float64(v)/ratio)))
				case "int16":
					b = make([]byte, 2)
					binary.BigEndian.PutUint16(b, uint16(int16(int64(float64(v)/ratio))))

				}
			}
		}
	case "int32":
		v, ok := value.(int32)
		if ok {

			switch strings.ToLower(storeMode) {
			case "little":
				switch strings.ToLower(StoreDataType) {
				case "uint16":
					b = make([]byte, 2)
					binary.LittleEndian.PutUint16(b, uint16(uint32(float32(v)/float32(ratio))))
				case "int16":
					b = make([]byte, 2)
					binary.LittleEndian.PutUint16(b, uint16(int16(int32(float32(v)/float32(ratio)))))
				}
			case "big":
				fallthrough
			default:
				switch strings.ToLower(StoreDataType) {
				case "uint16":
					b = make([]byte, 2)
					binary.BigEndian.PutUint16(b, uint16(uint32(float32(v)/float32(ratio))))
				case "int16":
					b = make([]byte, 2)
					binary.BigEndian.PutUint16(b, uint16(int16(int32(float32(v)/float32(ratio)))))
				}
			}
		}
	case "int16":
		v, ok := value.(int16)
		if ok {

			switch strings.ToLower(storeMode) {
			case "little":
				switch strings.ToLower(StoreDataType) {
				case "uint16":
					b = make([]byte, 2)
					binary.LittleEndian.PutUint16(b, uint16(uint64(float64(v)/ratio)))
				case "int16":
					b = make([]byte, 2)
					binary.LittleEndian.PutUint16(b, uint16(int16(int64(float64(v)/ratio))))
				}
			case "big":
				fallthrough
			default:
				switch strings.ToLower(StoreDataType) {
				case "uint16":
					b = make([]byte, 2)
					binary.BigEndian.PutUint16(b, uint16(uint64(float64(v)/ratio)))
				case "int16":
					b = make([]byte, 2)
					binary.BigEndian.PutUint16(b, uint16(int16(int64(float64(v)/ratio))))
				}
			}
		}
	case "[]float64":
		slice, ok := value.([]float64)
		if ok {
			for _, v := range slice {
				switch strings.ToLower(storeMode) {
				case "little":
					switch strings.ToLower(StoreDataType) {
					case "uint16":
						s = make([]byte, 2)
						binary.LittleEndian.PutUint16(s, uint16(uint64(v/float64(ratio))))
					case "int16":
						s = make([]byte, 2)
						binary.LittleEndian.PutUint16(s, uint16(int16(int64(v/float64(ratio)))))
					case "float32":
						s = make([]byte, 4)
						bits := math.Float32bits(float32(v) / float32(ratio))
						binary.LittleEndian.PutUint32(s, bits)
					}
				case "big":
					fallthrough
				default:
					switch strings.ToLower(StoreDataType) {
					case "uint16":
						s = make([]byte, 2)
						binary.BigEndian.PutUint16(s, uint16(uint64(v/float64(ratio))))
					case "int16":
						s = make([]byte, 2)
						binary.BigEndian.PutUint16(s, uint16(int16(int64(v/float64(ratio)))))
					case "float32":
						s = make([]byte, 4)
						bits := math.Float32bits(float32(v) / float32(ratio))
						binary.BigEndian.PutUint32(s, bits)
					}
				}
				b = append(b, s...)
			}

		}
	case "[]float32":
		slice, ok := value.([]float32)
		if ok {
			for _, v := range slice {
				switch strings.ToLower(storeMode) {
				case "little":
					switch strings.ToLower(StoreDataType) {
					case "uint16":
						s = make([]byte, 2)
						binary.LittleEndian.PutUint16(s, uint16(uint32(v/float32(ratio))))
					case "int16":
						s = make([]byte, 2)
						binary.LittleEndian.PutUint16(s, uint16(int16(int32(v/float32(ratio)))))
					case "float32":
						s = make([]byte, 4)
						bits := math.Float32bits(v / float32(ratio))
						binary.LittleEndian.PutUint32(s, bits)
					}
				case "big":
					fallthrough
				default:
					switch strings.ToLower(StoreDataType) {
					case "uint16":
						s = make([]byte, 2)
						binary.BigEndian.PutUint16(s, uint16(uint32(v/float32(ratio))))
					case "int16":
						s = make([]byte, 2)
						binary.BigEndian.PutUint16(s, uint16(int16(int32(v/float32(ratio)))))
					case "float32":
						s = make([]byte, 4)
						bits := math.Float32bits(v / float32(ratio))
						binary.BigEndian.PutUint32(s, bits)
					}
				}
				b = append(b, s...)
			}

		}
	case "[]int":
		slice, ok := value.([]int)
		if ok {
			for _, v := range slice {
				switch strings.ToLower(storeMode) {
				case "little":
					switch strings.ToLower(StoreDataType) {
					case "uint16":
						s = make([]byte, 2)
						binary.LittleEndian.PutUint16(s, uint16(uint(float64(v)/ratio)))
					case "int16":
						s = make([]byte, 2)
						binary.LittleEndian.PutUint16(s, uint16(int16(int(float64(v)/ratio))))
					}
				case "big":
					fallthrough
				default:
					switch strings.ToLower(StoreDataType) {
					case "uint16":
						s = make([]byte, 2)
						binary.BigEndian.PutUint16(s, uint16(uint(float64(v)/ratio)))
					case "int16":
						s = make([]byte, 2)
						binary.BigEndian.PutUint16(s, uint16(int16(int(float64(v)/ratio))))
					}
				}
				b = append(b, s...)
			}
		}
	case "[]int64":
		slice, ok := value.([]int64)
		if ok {
			for _, v := range slice {
				switch strings.ToLower(storeMode) {
				case "little":
					switch strings.ToLower(StoreDataType) {
					case "uint16":
						s = make([]byte, 2)
						binary.LittleEndian.PutUint16(s, uint16(uint64(float64(v)/ratio)))
					case "int16":
						s = make([]byte, 2)
						binary.LittleEndian.PutUint16(s, uint16(int16(int64(float64(v)/ratio))))
					}
				case "big":
					fallthrough
				default:
					switch strings.ToLower(StoreDataType) {
					case "uint16":
						s = make([]byte, 2)
						binary.BigEndian.PutUint16(s, uint16(uint64(float64(v)/ratio)))
					case "int16":
						s = make([]byte, 2)
						binary.BigEndian.PutUint16(s, uint16(int16(int64(float64(v)/ratio))))
					}
				}
				b = append(b, s...)
			}
		}
	case "[]int32":
		slice, ok := value.([]int32)
		if ok {
			for _, v := range slice {
				switch strings.ToLower(storeMode) {
				case "little":
					switch strings.ToLower(StoreDataType) {
					case "uint16":
						s = make([]byte, 2)
						binary.LittleEndian.PutUint16(s, uint16(uint32(float32(v)/float32(ratio))))
					case "int16":
						s = make([]byte, 2)
						binary.LittleEndian.PutUint16(s, uint16(int16(int32(float32(v)/float32(ratio)))))
					}
				case "big":
					fallthrough
				default:
					switch strings.ToLower(StoreDataType) {
					case "uint16":
						s = make([]byte, 2)
						binary.BigEndian.PutUint16(s, uint16(uint32(float32(v)/float32(ratio))))
					case "int16":
						s = make([]byte, 2)
						binary.BigEndian.PutUint16(s, uint16(int16(int32(float32(v)/float32(ratio)))))
					}
				}
				b = append(b, s...)
			}
		}
	case "[]int16":
		slice, ok := value.([]int16)
		if ok {
			for _, v := range slice {
				switch strings.ToLower(storeMode) {
				case "little":
					switch strings.ToLower(StoreDataType) {
					case "uint16":
						s = make([]byte, 2)
						binary.LittleEndian.PutUint16(s, uint16(uint64(float64(v)/ratio)))
					case "int16":
						s = make([]byte, 2)
						binary.LittleEndian.PutUint16(s, uint16(int16(int64(float64(v)/ratio))))
					}
				case "big":
					fallthrough
				default:
					switch strings.ToLower(StoreDataType) {
					case "uint16":
						s = make([]byte, 2)
						binary.BigEndian.PutUint16(s, uint16(uint64(float64(v)/ratio)))
					case "int16":
						s = make([]byte, 2)
						binary.BigEndian.PutUint16(s, uint16(int16(int64(float64(v)/ratio))))
					}
				}
			}
			b = append(b, s...)
		}
	}
	return b
}
