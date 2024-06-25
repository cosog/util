// byte2value.go
package util_decode

import (
	"encoding/binary"
	"math"
	"runtime"
	"strconv"
	"strings"
)

// func Byte2Value(storeMode string, quantity int, StoreDataType string, IFDataType string, ratio float64, b ...byte) interface{} {

// 	switch strings.ToLower(StoreDataType) {
// 	case "asc": //不是切片格式？？
// 		switch strings.ToLower(storeMode) {
// 		case "little":
// 			fallthrough
// 		case "big":
// 			fallthrough
// 		default:
// 			s := string(b)
// 			value := s
// 			return value
// 		}
// 	case "byte":
// 		if len(b) > 0 {

// 			switch strings.ToLower(IFDataType) {
// 			case "bool":
// 				value := make([]bool, 0)
// 				for i := 0; i < len(b); i++ {
// 					for j := 0; j < 8 && quantity > i*8+j; j++ {
// 						if (b[i]>>j)&0b00000001 == 0b00000001 {
// 							value = append(value, true)
// 						} else {
// 							value = append(value, false)
// 						}
// 					}
// 				}
// 				return value
// 			case "byte":
// 				value := make([]byte, 0)
// 				for i := 0; i < len(b); i++ {
// 					d := float32(int(uint(b[i])))
// 					d *= float32(ratio)
// 					value = append(value, byte(d))
// 				}
// 				return value
// 			case "int":
// 				value := make([]int, 0)
// 				for i := 0; i < len(b); i++ {
// 					d := float32(int(uint(b[i])))
// 					d *= float32(int(ratio))
// 					value = append(value, int(d))
// 				}
// 				return value
// 			case "float32":
// 				value := make([]float32, 0)
// 				for i := 0; i < len(b); i++ {
// 					d := float32(int(uint(b[i])))
// 					d *= float32(ratio)

// 					a := strconv.FormatFloat(float64(d), 'f', 2, 32)
// 					f, _ := strconv.ParseFloat(a, 32)

// 					value = append(value, float32(f))

// 				}
// 				// util.FormatOutput(value)
// 				return value
// 			case "float64":
// 				value := make([]float64, 0)
// 				for i := 0; i < len(b); i++ {
// 					d := float64(int(uint(b[i])))
// 					d *= float64(ratio)

// 					a := strconv.FormatFloat(d, 'f', 2, 64)
// 					f, _ := strconv.ParseFloat(a, 64)

// 					value = append(value, f)
// 				}
// 				// util.FormatOutput(value)
// 				return value
// 			}
// 		}
// 	case "uint16":
// 		switch strings.ToLower(storeMode) {
// 		case "little":
// 			if len(b) > 1 {

// 				switch strings.ToLower(IFDataType) {
// 				case "bool":
// 					value := make([]bool, 0)

// 					for i := 0; i < len(b); i = i + 2 {
// 						for j := 0; j < 8 && quantity > i*8+j; j++ {
// 							if (b[i]>>j)&0b00000001 == 0b00000001 {
// 								value = append(value, true)
// 							} else {
// 								value = append(value, false)
// 							}
// 						}
// 						for j := 0; j < 8 && quantity > (i+1)*8+j; j++ {
// 							if (b[i+1]>>j)&0b00000001 == 0b00000001 {
// 								value = append(value, true)
// 							} else {
// 								value = append(value, false)
// 							}
// 						}
// 					}

// 					return value
// 				case "int":
// 					value := make([]int, 0)
// 					for i := 0; i < len(b); i = i + 2 {
// 						d := int(binary.LittleEndian.Uint16(b[i : i+2]))
// 						f := float32(d) * float32(ratio)
// 						value = append(value, int(f))
// 					}
// 					return value
// 				case "float32":
// 					value := make([]float32, 0)
// 					for i := 0; i < len(b); i = i + 2 {
// 						d := int(binary.LittleEndian.Uint16(b[i : i+2]))
// 						f32 := float32(d) * float32(ratio)

// 						a := strconv.FormatFloat(float64(f32), 'f', 2, 32)
// 						f, _ := strconv.ParseFloat(a, 32)

// 						value = append(value, float32(f))
// 					}
// 					// util.FormatOutput(value)
// 					return value
// 				case "float64":
// 					value := make([]float64, 0)
// 					for i := 0; i < len(b); i = i + 2 {
// 						d := int(binary.LittleEndian.Uint16(b[i : i+2]))
// 						f64 := float64(d) * float64(ratio)

// 						a := strconv.FormatFloat(f64, 'f', 2, 64)
// 						f, _ := strconv.ParseFloat(a, 64)

// 						value = append(value, f)
// 					}
// 					// util.FormatOutput(value)
// 					return value
// 				}

// 			}

// 		case "big":
// 			fallthrough
// 		default:
// 			if len(b) > 1 {

// 				switch strings.ToLower(IFDataType) {
// 				case "bool":
// 					value := make([]bool, 0)

// 					for i := 0; i < len(b); i = i + 2 {
// 						for j := 0; j < 8 && quantity > i*8+j; j++ {
// 							if (b[i+1]>>j)&0b00000001 == 0b00000001 {
// 								value = append(value, true)
// 							} else {
// 								value = append(value, false)
// 							}
// 						}
// 						for j := 0; j < 8 && quantity > (i+1)*8+j; j++ {
// 							if (b[i]>>j)&0b00000001 == 0b00000001 {
// 								value = append(value, true)
// 							} else {
// 								value = append(value, false)
// 							}
// 						}
// 					}

// 					return value
// 				case "int":
// 					value := make([]int, 0)
// 					for i := 0; i < len(b); i = i + 2 {
// 						d := int(binary.BigEndian.Uint16(b[i : i+2]))
// 						f := float32(d) * float32(ratio)
// 						value = append(value, int(f))
// 					}
// 					return value
// 				case "float32":
// 					value := make([]float32, 0)
// 					for i := 0; i < len(b); i = i + 2 {
// 						d := int(binary.BigEndian.Uint16(b[i : i+2]))
// 						f32 := float32(d) * float32(ratio)

// 						a := strconv.FormatFloat(float64(f32), 'f', 2, 32)
// 						f, _ := strconv.ParseFloat(a, 32)
// 						value = append(value, float32(f))
// 					}
// 					// util.FormatOutput(value)
// 					return value
// 				case "float64":
// 					value := make([]float64, 0)
// 					for i := 0; i < len(b); i = i + 2 {
// 						d := int(binary.BigEndian.Uint16(b[i : i+2]))
// 						f64 := float64(d) * float64(ratio)

// 						a := strconv.FormatFloat(f64, 'f', 2, 64)
// 						f, _ := strconv.ParseFloat(a, 64)

// 						value = append(value, f)
// 					}
// 					// util.FormatOutput(value)
// 					return value
// 				}
// 			}
// 		}
// 	case "int16":
// 		switch strings.ToLower(storeMode) {
// 		case "little":
// 			if len(b) > 1 {

// 				switch strings.ToLower(IFDataType) {
// 				case "int":
// 					value := make([]int, 0)
// 					for i := 0; i < len(b); i = i + 2 {
// 						d := int(int16(binary.LittleEndian.Uint16(b[i : i+2])))
// 						f := float32(d) * float32(ratio)
// 						value = append(value, int(f))
// 					}
// 					return value
// 				case "float32":
// 					value := make([]float32, 0)
// 					for i := 0; i < len(b); i = i + 2 {
// 						d := int(int16(binary.LittleEndian.Uint16(b[i : i+2])))
// 						f32 := float32(d) * float32(ratio)

// 						a := strconv.FormatFloat(float64(f32), 'f', 2, 32)
// 						f, _ := strconv.ParseFloat(a, 32)

// 						value = append(value, float32(f))
// 					}
// 					// util.FormatOutput(value)
// 					return value
// 				case "float64":
// 					value := make([]float64, 0)
// 					for i := 0; i < len(b); i = i + 2 {
// 						d := int(int16(binary.LittleEndian.Uint16(b[i : i+2])))
// 						f64 := float64(d) * float64(ratio)

// 						a := strconv.FormatFloat(f64, 'f', 2, 64)
// 						f, _ := strconv.ParseFloat(a, 64)

// 						value = append(value, f)
// 					}
// 					// util.FormatOutput(value)
// 					return value
// 				}
// 			}
// 		case "big":
// 			fallthrough
// 		default:
// 			if len(b) > 1 {

// 				switch strings.ToLower(IFDataType) {
// 				case "int":
// 					value := make([]int, 0)
// 					for i := 0; i < len(b); i = i + 2 {
// 						d := int(int16(binary.BigEndian.Uint16(b[i : i+2])))
// 						f := float32(d) * float32(ratio)
// 						value = append(value, int(f))
// 					}
// 					return value
// 				case "float32":
// 					value := make([]float32, 0)
// 					for i := 0; i < len(b); i = i + 2 {
// 						d := int(int16(binary.BigEndian.Uint16(b[i : i+2])))
// 						f32 := float32(d) * float32(ratio)

// 						a := strconv.FormatFloat(float64(f32), 'f', 2, 32)
// 						f, _ := strconv.ParseFloat(a, 32)

// 						value = append(value, float32(f))
// 					}
// 					// util.FormatOutput(value)
// 					return value
// 				case "float64":
// 					value := make([]float64, 0)
// 					for i := 0; i < len(b); i = i + 2 {
// 						d := int(int16(binary.BigEndian.Uint16(b[i : i+2])))
// 						f64 := float64(d) * float64(ratio)

// 						a := strconv.FormatFloat(f64, 'f', 2, 64)
// 						f, _ := strconv.ParseFloat(a, 64)

// 						value = append(value, f)
// 					}
// 					// util.FormatOutput(value)
// 					return value
// 				}

// 			}
// 		}
// 	case "int64":
// 		switch strings.ToLower(storeMode) {
// 		case "little":
// 			if len(b) > 1 {

// 				switch strings.ToLower(IFDataType) {

// 				case "int64":
// 					value := make([]int64, 0)
// 					for i := 0; i < len(b); i = i + 8 {
// 						d := int64(binary.LittleEndian.Uint64(b[i : i+8]))
// 						f := float64(d) * float64(ratio)
// 						value = append(value, int64(f))
// 					}
// 					return value

// 				}
// 			}
// 		case "big":
// 			fallthrough
// 		default:
// 			if len(b) > 1 {

// 				switch strings.ToLower(IFDataType) {

// 				case "int64":
// 					value := make([]int64, 0)
// 					for i := 0; i < len(b); i = i + 8 {
// 						d := int64(binary.BigEndian.Uint64(b[i : i+8]))
// 						f := float64(d) * float64(ratio)
// 						value = append(value, int64(f))
// 					}
// 					return value

// 				}

// 			}
// 		}
// 	case "float32":
// 		switch strings.ToLower(storeMode) {
// 		case "little":
// 			if len(b) > 3 {

// 				switch strings.ToLower(IFDataType) {
// 				case "int":
// 					value := make([]int, 0)
// 					for i := 0; i < len(b); i = i + 4 {
// 						f := math.Float32frombits(binary.LittleEndian.Uint32(b[i : i+4]))
// 						d := f * float32(ratio)
// 						value = append(value, int(d))
// 					}
// 					return value
// 				case "float32":
// 					value := make([]float32, 0)
// 					for i := 0; i < len(b); i = i + 4 {
// 						f := math.Float32frombits(binary.LittleEndian.Uint32(b[i : i+4]))
// 						f32 := float32(f) * float32(ratio)

// 						a := strconv.FormatFloat(float64(f32), 'f', 2, 32)
// 						ff, _ := strconv.ParseFloat(a, 32)

// 						value = append(value, float32(ff))
// 					}
// 					// util.FormatOutput(value)
// 					return value
// 				case "float64":
// 					value := make([]float64, 0)
// 					for i := 0; i < len(b); i = i + 4 {
// 						f := math.Float32frombits(binary.LittleEndian.Uint32(b[i : i+4]))
// 						f64 := float64(f) * float64(ratio)

// 						a := strconv.FormatFloat(f64, 'f', 2, 64)
// 						ff, _ := strconv.ParseFloat(a, 64)

// 						value = append(value, ff)
// 					}
// 					// util.FormatOutput(value)
// 					return value
// 				}

// 			}
// 		case "big":
// 			fallthrough
// 		default:
// 			if len(b) > 3 {

// 				switch strings.ToLower(IFDataType) {
// 				case "int":
// 					value := make([]int, 0)
// 					for i := 0; i < len(b); i = i + 4 {
// 						f := math.Float32frombits(binary.BigEndian.Uint32(b[i : i+4]))
// 						d := f * float32(ratio)
// 						value = append(value, int(d))
// 					}
// 					return value
// 				case "float32":
// 					value := make([]float32, 0)
// 					for i := 0; i < len(b); i = i + 4 {
// 						f := math.Float32frombits(binary.BigEndian.Uint32(b[i : i+4]))
// 						f32 := float32(f) * float32(ratio)

// 						a := strconv.FormatFloat(float64(f32), 'f', 2, 32)
// 						ff, _ := strconv.ParseFloat(a, 32)

// 						value = append(value, float32(ff))
// 					}
// 					// util.FormatOutput(value)
// 					return value
// 				case "float64":
// 					value := make([]float64, 0)
// 					for i := 0; i < len(b); i = i + 4 {
// 						f := math.Float32frombits(binary.BigEndian.Uint32(b[i : i+4]))
// 						f64 := float64(f) * float64(ratio)

// 						a := strconv.FormatFloat(f64, 'f', 2, 64)
// 						ff, _ := strconv.ParseFloat(a, 64)

// 						value = append(value, ff)
// 					}
// 					// util.FormatOutput(value)
// 					return value
// 				}

// 			}
// 		}

// 	case "bcd":
// 		switch strings.ToLower(storeMode) {
// 		case "little":
// 			fallthrough
// 		case "big":
// 			fallthrough
// 		default:
// 			c := make([]string, 2*len(b))
// 			s := ""
// 			if len(b) > 0 {

// 				switch strings.ToLower(IFDataType) {
// 				case "int":
// 					d, err := strconv.Atoi(s)
// 					if err != nil {

// 					}
// 					f32 := float32(d) * float32(ratio)
// 					value := make([]int, 0)
// 					value = append(value, int(f32))
// 					return value
// 				case "float32":
// 					d, err := strconv.Atoi(s)
// 					if err != nil {

// 					}

// 					f32 := float32(d) * float32(ratio)

// 					value := make([]float32, 0)

// 					a := strconv.FormatFloat(float64(f32), 'f', 2, 32)
// 					f, _ := strconv.ParseFloat(a, 32)

// 					value = append(value, float32(f))
// 					// util.FormatOutput(value)
// 					return value

// 				case "float64":
// 					d, err := strconv.Atoi(s)
// 					if err != nil {

// 					}
// 					f64 := float64(d) * float64(ratio)

// 					value := make([]float64, 0)

// 					a := strconv.FormatFloat(f64, 'f', 2, 64)
// 					f, _ := strconv.ParseFloat(a, 64)

// 					value = append(value, f)
// 					// util.FormatOutput(value)
// 					return value
// 				case "string":
// 					for i := 0; i < len(b); i++ {
// 						c[i*2] = strconv.Itoa(int((b[i] & 0b11110000) >> 4))
// 						c[i*2+1] = strconv.Itoa(int(b[i] & 0b00001111))
// 						s += c[i*2]
// 						s += c[i*2+1]
// 					}
// 					value := make([]string, 0)
// 					value = append(value, s)
// 					return value
// 				}
// 			}

// 		}
// 	}

// 	var value interface{} //???
// 	return value

// }

func Byte2Value(storeMode string, quantity int, StoreDataType string, IFDataType string, prec int, ratio float64, b ...byte) interface{} {

	if prec <= 0 {
		prec = 2
	}
	switch strings.ToLower(StoreDataType) {
	case "asc": //不是切片格式？？
		switch strings.ToLower(storeMode) {
		case "little":
			fallthrough
		case "big":
			fallthrough
		default:
			s := string(b)
			value := s
			return value
		}
	case "bit":
		if len(b) > 0 {

			switch strings.ToLower(IFDataType) {
			case "bool":
				value := make([]bool, 0)
				for i := 0; i < len(b); i++ {
					for j := 0; j < 8 && quantity > i*8+j; j++ {
						if (b[i]>>j)&0b00000001 == 0b00000001 {
							value = append(value, true)
						} else {
							value = append(value, false)
						}
						runtime.Gosched()
					}
					runtime.Gosched()
				}
				return value
			}
		}
	case "byte":
		if len(b) > 0 {

			switch strings.ToLower(IFDataType) {
			case "bool":
				value := make([]bool, 0)
				for i := 0; i < len(b); i++ {
					for j := 0; j < 8 && quantity > i*8+j; j++ {
						if (b[i]>>j)&0b00000001 == 0b00000001 {
							value = append(value, true)
						} else {
							value = append(value, false)
						}
						runtime.Gosched()
					}
					runtime.Gosched()
				}
				return value
			case "byte":
				value := make([]byte, 0)
				for i := 0; i < len(b); i++ {
					d := float32(int(uint(b[i])))
					d *= float32(ratio)
					value = append(value, byte(d))
					runtime.Gosched()
				}
				return value
			case "int":
				value := make([]int, 0)
				for i := 0; i < len(b); i++ {
					d := float32(int(uint(b[i])))
					d *= float32(int(ratio))
					value = append(value, int(d))
					runtime.Gosched()
				}
				return value
			case "float32":
				value := make([]float32, 0)
				for i := 0; i < len(b); i++ {
					d := float32(int(uint(b[i])))
					d *= float32(ratio)

					a := strconv.FormatFloat(float64(d), 'f', prec, 32)
					f, _ := strconv.ParseFloat(a, 32)

					value = append(value, float32(f))
					runtime.Gosched()
				}
				// util.FormatOutput(value)
				return value
			case "float64":
				value := make([]float64, 0)
				for i := 0; i < len(b); i++ {
					d := float64(int(uint(b[i])))
					d *= float64(ratio)

					a := strconv.FormatFloat(d, 'f', prec, 64)
					f, _ := strconv.ParseFloat(a, 64)

					value = append(value, f)
					runtime.Gosched()
				}
				// util.FormatOutput(value)
				return value
			}
		}
	case "uint16":
		switch strings.ToLower(storeMode) {
		case "little":
			if len(b) > 1 {

				switch strings.ToLower(IFDataType) {
				case "bool":
					value := make([]bool, 0)

					for i := 0; i+2 <= len(b); i = i + 2 {
						for j := 0; j < 8 && quantity > i*8+j; j++ {
							if (b[i]>>j)&0b00000001 == 0b00000001 {
								value = append(value, true)
							} else {
								value = append(value, false)
							}
							runtime.Gosched()
						}
						for j := 0; j < 8 && quantity > (i+1)*8+j; j++ {
							if (b[i+1]>>j)&0b00000001 == 0b00000001 {
								value = append(value, true)
							} else {
								value = append(value, false)
							}
							runtime.Gosched()
						}
						runtime.Gosched()
					}

					return value
				case "int":
					value := make([]int, 0)
					for i := 0; i+2 <= len(b); i = i + 2 {
						d := int(binary.LittleEndian.Uint16(b[i : i+2]))
						f := float32(d) * float32(ratio)
						value = append(value, int(f))
						runtime.Gosched()
					}
					return value
				case "float32":
					value := make([]float32, 0)
					for i := 0; i+2 <= len(b); i = i + 2 {
						d := int(binary.LittleEndian.Uint16(b[i : i+2]))
						f32 := float32(d) * float32(ratio)

						a := strconv.FormatFloat(float64(f32), 'f', prec, 32)
						f, _ := strconv.ParseFloat(a, 32)

						value = append(value, float32(f))
						runtime.Gosched()
					}
					// util.FormatOutput(value)
					return value
				case "float64":
					value := make([]float64, 0)
					for i := 0; i+2 <= len(b); i = i + 2 {
						d := int(binary.LittleEndian.Uint16(b[i : i+2]))
						f64 := float64(d) * float64(ratio)

						a := strconv.FormatFloat(f64, 'f', prec, 64)
						f, _ := strconv.ParseFloat(a, 64)

						value = append(value, f)
						runtime.Gosched()
					}
					// util.FormatOutput(value)
					return value
				}

			}

		case "big":
			fallthrough
		default:
			if len(b) > 1 {

				switch strings.ToLower(IFDataType) {
				case "bool":
					value := make([]bool, 0)

					for i := 0; i+2 <= len(b); i = i + 2 {
						for j := 0; j < 8 && quantity > i*8+j; j++ {
							if (b[i+1]>>j)&0b00000001 == 0b00000001 {
								value = append(value, true)
							} else {
								value = append(value, false)
							}
							runtime.Gosched()
						}
						for j := 0; j < 8 && quantity > (i+1)*8+j; j++ {
							if (b[i]>>j)&0b00000001 == 0b00000001 {
								value = append(value, true)
							} else {
								value = append(value, false)
							}
							runtime.Gosched()
						}
						runtime.Gosched()
					}

					return value
				case "int":
					value := make([]int, 0)
					for i := 0; i+2 <= len(b); i = i + 2 {
						d := int(binary.BigEndian.Uint16(b[i : i+2]))
						f := float32(d) * float32(ratio)
						value = append(value, int(f))
						runtime.Gosched()
					}
					return value
				case "float32":
					value := make([]float32, 0)
					for i := 0; i+2 <= len(b); i = i + 2 {
						d := int(binary.BigEndian.Uint16(b[i : i+2]))
						f32 := float32(d) * float32(ratio)

						a := strconv.FormatFloat(float64(f32), 'f', prec, 32)
						f, _ := strconv.ParseFloat(a, 32)
						value = append(value, float32(f))
						runtime.Gosched()
					}
					// util.FormatOutput(value)
					return value
				case "float64":
					value := make([]float64, 0)
					for i := 0; i+2 <= len(b); i = i + 2 {
						d := int(binary.BigEndian.Uint16(b[i : i+2]))
						f64 := float64(d) * float64(ratio)

						a := strconv.FormatFloat(f64, 'f', prec, 64)
						f, _ := strconv.ParseFloat(a, 64)

						value = append(value, f)
						runtime.Gosched()
					}
					// util.FormatOutput(value)
					return value
				}
			}
		}
	case "int16":
		switch strings.ToLower(storeMode) {
		case "little":
			if len(b) > 1 {

				switch strings.ToLower(IFDataType) {
				case "int":
					value := make([]int, 0)
					for i := 0; i+2 <= len(b); i = i + 2 {
						d := int(int16(binary.LittleEndian.Uint16(b[i : i+2])))
						f := float32(d) * float32(ratio)
						value = append(value, int(f))
						runtime.Gosched()
					}
					return value
				case "float32":
					value := make([]float32, 0)
					for i := 0; i+2 <= len(b); i = i + 2 {
						d := int(int16(binary.LittleEndian.Uint16(b[i : i+2])))
						f32 := float32(d) * float32(ratio)

						a := strconv.FormatFloat(float64(f32), 'f', prec, 32)
						f, _ := strconv.ParseFloat(a, 32)

						value = append(value, float32(f))
						runtime.Gosched()
					}
					// util.FormatOutput(value)
					return value
				case "float64":
					value := make([]float64, 0)
					for i := 0; i+2 <= len(b); i = i + 2 {
						d := int(int16(binary.LittleEndian.Uint16(b[i : i+2])))
						f64 := float64(d) * float64(ratio)

						a := strconv.FormatFloat(f64, 'f', prec, 64)
						f, _ := strconv.ParseFloat(a, 64)

						value = append(value, f)
						runtime.Gosched()
					}
					// util.FormatOutput(value)
					return value
				}
			}
		case "big":
			fallthrough
		default:
			if len(b) > 1 {

				switch strings.ToLower(IFDataType) {
				case "int":
					value := make([]int, 0)
					for i := 0; i+2 <= len(b); i = i + 2 {
						d := int(int16(binary.BigEndian.Uint16(b[i : i+2])))
						f := float32(d) * float32(ratio)
						value = append(value, int(f))
						runtime.Gosched()
					}
					return value
				case "float32":
					value := make([]float32, 0)
					for i := 0; i+2 <= len(b); i = i + 2 {
						d := int(int16(binary.BigEndian.Uint16(b[i : i+2])))
						f32 := float32(d) * float32(ratio)

						a := strconv.FormatFloat(float64(f32), 'f', prec, 32)
						f, _ := strconv.ParseFloat(a, 32)

						value = append(value, float32(f))
						runtime.Gosched()
					}
					// util.FormatOutput(value)
					return value
				case "float64":
					value := make([]float64, 0)
					for i := 0; i+2 <= len(b); i = i + 2 {
						d := int(int16(binary.BigEndian.Uint16(b[i : i+2])))
						f64 := float64(d) * float64(ratio)

						a := strconv.FormatFloat(f64, 'f', prec, 64)
						f, _ := strconv.ParseFloat(a, 64)

						value = append(value, f)
						runtime.Gosched()
					}
					// util.FormatOutput(value)
					return value
				}
			}
		}

	case "float32":
		switch strings.ToLower(storeMode) {
		case "little":
			if len(b) > 3 {

				switch strings.ToLower(IFDataType) {
				case "int":
					value := make([]int, 0)
					for i := 0; i+4 <= len(b); i = i + 4 {
						f := math.Float32frombits(binary.LittleEndian.Uint32(b[i : i+4]))
						d := f * float32(ratio)
						value = append(value, int(d))
						runtime.Gosched()
					}
					return value
				case "float32":
					value := make([]float32, 0)
					for i := 0; i+4 <= len(b); i = i + 4 {
						f := math.Float32frombits(binary.LittleEndian.Uint32(b[i : i+4]))
						f32 := float32(f) * float32(ratio)

						a := strconv.FormatFloat(float64(f32), 'f', prec, 32)
						ff, _ := strconv.ParseFloat(a, 32)

						value = append(value, float32(ff))
						runtime.Gosched()
					}
					// util.FormatOutput(value)
					return value
				case "float64":
					value := make([]float64, 0)
					for i := 0; i+4 <= len(b); i = i + 4 {
						f := math.Float32frombits(binary.LittleEndian.Uint32(b[i : i+4]))
						f64 := float64(f) * float64(ratio)

						a := strconv.FormatFloat(f64, 'f', prec, 64)
						ff, _ := strconv.ParseFloat(a, 64)

						value = append(value, ff)
						runtime.Gosched()
					}
					// util.FormatOutput(value)
					return value
				}

			}
		case "big":
			fallthrough
		default:
			if len(b) > 3 {

				switch strings.ToLower(IFDataType) {
				case "int":
					value := make([]int, 0)
					for i := 0; i+4 <= len(b); i = i + 4 {
						f := math.Float32frombits(binary.BigEndian.Uint32(b[i : i+4]))
						d := f * float32(ratio)
						value = append(value, int(d))
					}
					return value
				case "float32":
					value := make([]float32, 0)
					for i := 0; i+4 <= len(b); i = i + 4 {
						f := math.Float32frombits(binary.BigEndian.Uint32(b[i : i+4]))
						f32 := float32(f) * float32(ratio)

						a := strconv.FormatFloat(float64(f32), 'f', prec, 32)
						ff, _ := strconv.ParseFloat(a, 32)

						value = append(value, float32(ff))
						runtime.Gosched()
					}
					// util.FormatOutput(value)
					return value
				case "float64":
					value := make([]float64, 0)
					for i := 0; i+4 <= len(b); i = i + 4 {
						f := math.Float32frombits(binary.BigEndian.Uint32(b[i : i+4]))
						f64 := float64(f) * float64(ratio)

						a := strconv.FormatFloat(f64, 'f', prec, 64)
						ff, _ := strconv.ParseFloat(a, 64)

						value = append(value, ff)
						runtime.Gosched()
					}
					// util.FormatOutput(value)
					return value
				}

			}
		}

	case "bcd":
		switch strings.ToLower(storeMode) {
		case "little":
			fallthrough
		case "big":
			fallthrough
		default:
			c := make([]string, 2*len(b))
			s := ""
			if len(b) > 0 {

				switch strings.ToLower(IFDataType) {
				case "int":
					d, err := strconv.Atoi(s)
					if err != nil {

					}
					f32 := float32(d) * float32(ratio)
					value := make([]int, 0)
					value = append(value, int(f32))
					return value
				case "float32":
					d, err := strconv.Atoi(s)
					if err != nil {

					}

					f32 := float32(d) * float32(ratio)

					value := make([]float32, 0)

					a := strconv.FormatFloat(float64(f32), 'f', prec, 32)
					f, _ := strconv.ParseFloat(a, 32)

					value = append(value, float32(f))
					// util.FormatOutput(value)
					return value

				case "float64":
					d, err := strconv.Atoi(s)
					if err != nil {

					}
					f64 := float64(d) * float64(ratio)

					value := make([]float64, 0)

					a := strconv.FormatFloat(f64, 'f', prec, 64)
					f, _ := strconv.ParseFloat(a, 64)

					value = append(value, f)
					// util.FormatOutput(value)
					return value
				case "string":
					for i := 0; i < len(b); i++ {
						c[i*2] = strconv.Itoa(int((b[i] & 0b11110000) >> 4))
						c[i*2+1] = strconv.Itoa(int(b[i] & 0b00001111))
						s += c[i*2]
						s += c[i*2+1]
						runtime.Gosched()
					}
					value := make([]string, 0)
					value = append(value, s)
					return value
				}
			}

		}
	}

	var value interface{} //???
	return value

}
