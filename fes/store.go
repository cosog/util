// storemodbusa11.go
package util_fes

import (
	"encoding/binary"
	"math"
	"runtime"
	"strconv"

	"github.com/cosog/util/crc"
)

func StoreModbusA11(fes *FESStruct) []string {
	var fesA11 FESStruct
	if fes.CNT > 200 || len(fes.S) > 200 {
		fesA11 = FESInterpolationPoint(fes, 200)
	} else {
		fesA11 = *fes
	}

	b := make([]byte, 0)
	c := make([]byte, 0)
	w := make([]string, 0)
	var crc util_crc.Crc
	var b_41001 []byte

	cnt := uint16(fesA11.CNT)
	b_40984 := make([]byte, 2)
	binary.BigEndian.PutUint16(b_40984, cnt)

	s := fesA11.AcqTime
	y1, _ := strconv.Atoi(string(s[0]))
	y2, _ := strconv.Atoi(string(s[1]))
	y3, _ := strconv.Atoi(string(s[2]))
	y4, _ := strconv.Atoi(string(s[3]))

	month1, _ := strconv.Atoi(string(s[5]))
	month2, _ := strconv.Atoi(string(s[6]))

	d1, _ := strconv.Atoi(string(s[8]))
	d2, _ := strconv.Atoi(string(s[9]))

	h1, _ := strconv.Atoi(string(s[11]))
	h2, _ := strconv.Atoi(string(s[12]))

	min1, _ := strconv.Atoi(string(s[14]))
	min2, _ := strconv.Atoi(string(s[15]))

	s1, _ := strconv.Atoi(string(s[17]))
	s2, _ := strconv.Atoi(string(s[18]))

	bcd_40985_high := byte(y1<<4 | y2)        //年
	bcd_40985_low := byte(y3<<4 | y4)         //年
	bcd_40986_high := byte(0x00)              //-
	bcd_40986_low := byte(month1<<4 | month2) //月
	bcd_40987_high := byte(0x00)              //-
	bcd_40987_low := byte(d1<<4 | d2)         //日
	bcd_40988_high := byte(0x00)              //空格
	bcd_40988_low := byte(h1<<4 | h2)         //时
	bcd_40989_high := byte(0x00)              //:
	bcd_40989_low := byte(min1<<4 | min2)     //分
	bcd_40990_high := byte(0x00)              //:
	bcd_40990_low := byte(s1<<4 | s2)         //秒

	b_40991 := make([]byte, 4)

	bits := math.Float32bits(float32(fesA11.SPM))
	binary.BigEndian.PutUint32(b_40991, bits)

	b_40993 := make([]byte, 4)
	bits = math.Float32bits(float32(fesA11.Stroke))

	binary.BigEndian.PutUint32(b_40993, bits)

	b = append(b, 0x01, 0x10, 0x03, 0xD7, 0x00, 0x0B, 0x16)
	b = append(b, b_40984...)
	b = append(b, bcd_40985_high, bcd_40985_low, bcd_40986_high, bcd_40986_low, bcd_40987_high, bcd_40987_low, bcd_40988_high, bcd_40988_low, bcd_40989_high, bcd_40989_low, bcd_40990_high, bcd_40990_low)
	b = append(b, b_40991...)
	b = append(b, b_40993...)

	c = append(c, 0x01, 0x10, 0x03, 0xD7, 0x00, 0x0B, 0x16)
	c = append(c, b_40984...)
	c = append(c, bcd_40985_high, bcd_40985_low, bcd_40986_high, bcd_40986_low, bcd_40987_high, bcd_40987_low, bcd_40988_high, bcd_40988_low, bcd_40989_high, bcd_40989_low, bcd_40990_high, bcd_40990_low)
	c = append(c, b_40991...)
	c = append(c, b_40993...)

	crc.Reset()
	crc.PushBytes(c)
	b = append(b, crc.Low, crc.High)

	// box_global.RTUClientChan.Client.Cmd <- string(b)
	w = append(w, string(b))

	////////////////////////////////////////////////////////////////////////////////////////////////
	// high-高字节 low-低字节
	// 大端：高字节对应内存低地址
	if cnt <= 250 {
		totalByteLength := 2000
		sendRegisterLength := 100
		times := totalByteLength / sendRegisterLength / 2

		b_41001 = make([]byte, totalByteLength)
		for k, v := range fesA11.S {
			high := 2 * k
			low := 2*k + 1
			b2 := make([]byte, 2)
			binary.BigEndian.PutUint16(b2, uint16(int16(int32(float32(v*100)))))
			b_41001[high] = b2[0]
			b_41001[low] = b2[1]
		}
		for k, v := range fesA11.F {
			high := 2*k + 500
			low := 2*k + 1 + 500
			b2 := make([]byte, 2)
			binary.BigEndian.PutUint16(b2, uint16(int16(int32(float32(v*100)))))
			b_41001[high] = b2[0]
			b_41001[low] = b2[1]
		}

		for k, v := range fesA11.I {
			high := 2*k + 1000
			low := 2*k + 1 + 1000
			b2 := make([]byte, 2)
			binary.BigEndian.PutUint16(b2, uint16(int16(int32(float32(v*100)))))
			b_41001[high] = b2[0]
			b_41001[low] = b2[1]
		}
		for k, v := range fesA11.Watt {
			high := 2*k + 1500
			low := 2*k + 1 + 1500
			b2 := make([]byte, 2)
			binary.BigEndian.PutUint16(b2, uint16(int16(int32(float32(v*100)))))
			b_41001[high] = b2[0]
			b_41001[low] = b2[1]
		}

		for i := 0; i < times; i++ {

			num := 1001 + i*sendRegisterLength - 1
			start := 2 * i * sendRegisterLength
			end := 2 * (i + 1) * sendRegisterLength
			b = make([]byte, 0)
			c = make([]byte, 0)
			addr := make([]byte, 2)

			binary.BigEndian.PutUint16(addr, uint16(num))
			b = append(b, 0x01, 0x10)
			b = append(b, addr...)
			b = append(b, 0x00, 0x64)
			b = append(b, 0xC8)
			b = append(b, b_41001[start:end]...)

			c = append(c, 0x01, 0x10)
			c = append(c, addr...)
			c = append(c, 0x00, 0x64)
			c = append(c, 0xC8)
			c = append(b, b_41001[start:end]...)

			crc.Reset()
			crc.PushBytes(c)
			b = append(b, crc.Low, crc.High)

			// box_global.RTUClientChan.Client.Cmd <- string(b)
			w = append(w, string(b))

			runtime.Gosched()
		}
	} else if cnt <= 500 {

		totalByteLength := 4000
		sendRegisterLength := 100
		times := totalByteLength / sendRegisterLength / 2

		b_41001 = make([]byte, totalByteLength)
		for k, v := range fesA11.S {
			high := 2 * k
			low := 2*k + 1
			b2 := make([]byte, 2)
			binary.BigEndian.PutUint16(b2, uint16(int16(int32(float32(v*100)))))
			b_41001[high] = b2[0]
			b_41001[low] = b2[1]
		}
		for k, v := range fesA11.F {
			high := 2*k + 1000
			low := 2*k + 1 + 1000
			b2 := make([]byte, 2)
			binary.BigEndian.PutUint16(b2, uint16(int16(int32(float32(v*100)))))
			b_41001[high] = b2[0]
			b_41001[low] = b2[1]
		}

		for k, v := range fesA11.I {
			high := 2*k + 2000
			low := 2*k + 1 + 2000
			b2 := make([]byte, 2)
			binary.BigEndian.PutUint16(b2, uint16(int16(int32(float32(v*100)))))
			b_41001[high] = b2[0]
			b_41001[low] = b2[1]
		}
		for k, v := range fesA11.Watt {
			high := 2*k + 3000
			low := 2*k + 1 + 3000
			b2 := make([]byte, 2)
			binary.BigEndian.PutUint16(b2, uint16(int16(int32(float32(v*100)))))
			b_41001[high] = b2[0]
			b_41001[low] = b2[1]
		}

		for i := 0; i < times; i++ {

			num := 1001 + i*sendRegisterLength - 1
			start := 2 * i * sendRegisterLength
			end := 2 * (i + 1) * sendRegisterLength
			b = make([]byte, 0)
			c = make([]byte, 0)
			addr := make([]byte, 2)

			binary.BigEndian.PutUint16(addr, uint16(num))
			b = append(b, 0x01, 0x10)
			b = append(b, addr...)
			b = append(b, 0x00, 0x64)
			b = append(b, 0xC8)
			b = append(b, b_41001[start:end]...)

			c = append(c, 0x01, 0x10)
			c = append(c, addr...)
			c = append(c, 0x00, 0x64)
			c = append(c, 0xC8)
			c = append(b, b_41001[start:end]...)

			var crc util_crc.Crc
			crc.Reset()
			crc.PushBytes(c)
			b = append(b, crc.Low, crc.High)

			// box_global.RTUClientChan.Client.Cmd <- string(b)
			w = append(w, string(b))

			runtime.Gosched()
		}
	}
	return w
}
