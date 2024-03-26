package main

import (
	"fmt"
	"crypto/rand"
	"time"
	"math/big"
	"strconv"
)
var TABLE string = ".0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz";
var mp map[string]int = make(map[string]int);
var node string = "."
var maxBigInt = big.NewInt(1002342342334234213)

func sortedBase64Encode(inp string) (string) {
	var padding = len(inp) % 3;
	var length = len(inp) - padding
	var i int = 0
	var output string = ""
	for ; i<length; i++ {
		a := uint32(inp[i])
		x := a << 16
		i++
		a = uint32(inp[i])
		y := a << 8
		i++
		a = uint32(inp[i])
		z := a << 0

		buffer := x+y+z
		output += string(TABLE[buffer >> 18 & 0x3F]) + string(TABLE[buffer >> 12 & 0x3F]) + string(TABLE[buffer >> 6 & 0x3F]) + string(TABLE[buffer >> 0 & 0x3F])
	}

	if padding == 1 {
		buffer := uint32(inp[i])
		output += string(TABLE[buffer >> 2 & 0x3F]) + string(TABLE[buffer << 4 & 0x3F]) + "==";
	} else if padding == 2 {
		a := uint32(inp[i]) << 8;
		i++;
		b := uint32(inp[i]) << 0;
		buffer := a+b
		output += string(TABLE[buffer >> 10 & 0x3F]) + string(TABLE[buffer >> 4 & 0x3F]) + string(TABLE[buffer << 2 & 0x3F]) + "="
	}
	return output
}

func sorteBase64Decode(inp string) (string) {
	var i int = 0;
	var length = len(inp)
	output := ""
	cnt := 0

	for ;i<length; i++ {
		a := mp[string(inp[i])] << (((cnt+1)%4) * 2) & 0xFF
		b := mp[string(inp[i+1])] >> ((6 - ((cnt+1)%4)*2) & 0xFF)

		buffer := a+b

		output += string(buffer)
		if (cnt+1)%3 == 0 {
			i++;
			cnt = 0;
		} else {
			cnt++;
		}
	}

	return output

}


func GenerateId() string {
	randomNumber, _ := rand.Int(rand.Reader, maxBigInt)
	var id = string(strconv.FormatInt(time.Now().UnixMilli(), 10)) + node + randomNumber.String()
	return sortedBase64Encode(id)
}

func PerfTest(length int) {
	start := time.Now().UnixMilli()
	for i:=0; i<length; i++ {
		GenerateId()
	}
	end := time.Now().UnixMilli()

	fmt.Println(float64(end - start)/1000.0, "seconds for", length, "ids")

}

func main() {
	for i:=0; i<len(TABLE); i++ {
		mp[string(TABLE[i])] = i
	}

	length := 1000000
	PerfTest(length)
}
