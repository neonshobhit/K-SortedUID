package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

var TABLE string = ".0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"
var mp map[string]int = make(map[string]int)
var node string = "."
var maxBigInt = new(big.Int).SetUint64(18446744073709551615)
var lastTime int64
var lastRandom uint64

func SortedBase64Encode(inp string) string {
	var padding = len(inp) % 3
	var length = len(inp) - padding
	var i int = 0
	var output string = ""
	for ; i < length; i++ {
		a := uint32(inp[i])
		x := a << 16
		i++
		a = uint32(inp[i])
		y := a << 8
		i++
		a = uint32(inp[i])
		z := a << 0

		buffer := x + y + z
		output += string(TABLE[buffer>>18&0x3F]) + string(TABLE[buffer>>12&0x3F]) + string(TABLE[buffer>>6&0x3F]) + string(TABLE[buffer>>0&0x3F])
	}

	if padding == 1 {
		buffer := uint32(inp[i])
		output += string(TABLE[buffer>>2&0x3F]) + string(TABLE[buffer<<4&0x3F]) + "=="
	} else if padding == 2 {
		a := uint32(inp[i]) << 8
		i++
		b := uint32(inp[i]) << 0
		buffer := a + b
		output += string(TABLE[buffer>>10&0x3F]) + string(TABLE[buffer>>4&0x3F]) + string(TABLE[buffer<<2&0x3F]) + "="
	}
	return output
}

func SorteBase64Decode(inp string) string {
	var i int = 0
	var length = len(inp)
	output := ""
	cnt := 0

	for ; i < length; i++ {
		a := mp[string(inp[i])] << (((cnt + 1) % 4) * 2) & 0xFF
		b := mp[string(inp[i+1])] >> ((6 - ((cnt+1)%4)*2) & 0xFF)

		buffer := a + b

		output += string(buffer)
		if (cnt+1)%3 == 0 {
			i++
			cnt = 0
		} else {
			cnt++
		}
	}

	return output

}

func GenerateId() string {
	var now int64 = time.Now().UnixMilli()
	var delta uint64 = 0
	randomNumber, _ := rand.Int(rand.Reader, maxBigInt)
	randomNumberWithDelta := randomNumber.Uint64()
	if lastTime == now && lastRandom == randomNumberWithDelta {
		delta = 1
		randomNumberWithDelta = randomNumberWithDelta + delta
	}
	lastTime = now
	lastRandom = randomNumberWithDelta
	// var id = string(strconv.FormatInt(time.Now().UnixMilli(), 10)) + node + randomNumber.String()
	// return SortedBase64Encode(id)
	kSortedId := make([]byte, 19)
	for i := 18; i >= 8; i-- {
		kSortedId[i] = TABLE[randomNumberWithDelta%64]
		randomNumberWithDelta = randomNumberWithDelta >> 6
	}
	// fmt.Println(kSortedId)
	for i := 7; i >= 0; i-- {
		kSortedId[i] = TABLE[now%64]
		now = now >> 6
	}
	// fmt.Println(kSortedId)
	return string(kSortedId)
}

func PerfTest(length int, shallPrint bool) {
	var collisions map[string]int = make(map[string]int)
	collisionsCnt := 0
	start := time.Now().UnixMilli()
	strs := []string{}
	for i := 0; i < length; i++ {
		ksortedid := GenerateId()
		if shallPrint {
			fmt.Println(ksortedid)
		}
		strs = append(strs, ksortedid)
		_, ok := collisions[ksortedid]
		if ok {
			collisionsCnt++
			collisions[ksortedid]++
		} else {
			collisions[ksortedid] = 1
		}
		_ = ksortedid
	}
	end := time.Now().UnixMilli()
	fmt.Println("Collisions:", collisionsCnt)
	fmt.Println(float64(end-start)/1000.0, "seconds for", length, "ids")

}

func main() {
	for i := 0; i < len(TABLE); i++ {
		mp[string(TABLE[i])] = i
	}
	// length := 10000000
	length := 10
	PerfTest(length, true)
}
