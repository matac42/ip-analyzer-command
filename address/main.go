package address

import (
	"fmt"
	"strconv"
	"strings"
)

func AddrString2AddrArray(addr string) []int64 {
	// "192.168.58.1" -> [192, 168, 58, 1]

	stringArray := strings.Split(addr, ".")
	intArray := []int64{}

	for _, s := range stringArray {
		i, _ := strconv.Atoi(s)
		intArray = append(intArray, int64(i))
	}

	return intArray
}

func AddrArray2String(addrArray []int64) string {
	// [192, 168, 58, 1] -> "192.168.58.1"

	return fmt.Sprintf(
		"%s.%s.%s.%s",
		strconv.FormatInt(addrArray[0], 10),
		strconv.FormatInt(addrArray[1], 10),
		strconv.FormatInt(addrArray[2], 10),
		strconv.FormatInt(addrArray[3], 10),
	)
}

func Prefix2AddrArray(prefix int) []int64 {
	// 24 -> [255, 255, 255, 0]

	bit := strings.Repeat("1", prefix) + strings.Repeat("0", 32-prefix)

	return bit2AddrArray(bit)
}

func CalcBroadCastAddr(networkAddr []int64, maskAddr []int64) []int64 {
	// [192, 168, 58, 0], [255, 255, 255, 0] -> [192, 168, 58, 255]

	broadCastAddr := []int64{}

	for i := 0; i < 4; i++ {
		broadCastAddr = append(broadCastAddr, networkAddr[i]^(maskAddr[i]^255))
	}

	return broadCastAddr
}

func bit2AddrArray(bit string) []int64 {
	// "11111111111111111111111100000000" -> [255, 255, 255, 0]

	addrArray := []int64{}
	for i := 0; i < len(bit); i += 8 {
		res, _ := strconv.ParseInt(bit[i:(i+8)], 2, 64)
		addrArray = append(addrArray, res)
	}

	return addrArray
}
