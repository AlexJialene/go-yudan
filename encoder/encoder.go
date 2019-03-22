package encoder

import (
	"strings"
	"bytes"
	"encoding/binary"
)

type Encoder struct {
	Str string
}

func GetEncoder() *Encoder {
	e := &Encoder{
	}
	return e
}

func AddItem(encoder *Encoder, key, value string) int {
	key = strings.Replace(key, "/", "@S", -1)
	key = strings.Replace(key, "@", "@A", -1)
	encoder.Str = encoder.Str + key
	encoder.Str = encoder.Str + "@="
	value = strings.Replace(value, "/", "@S", -1)
	value = strings.Replace(value, "@", "@A", -1)
	encoder.Str = encoder.Str + value
	encoder.Str = encoder.Str + "/"

	return 0
}

func Result(encoder *Encoder) string {
	return encoder.Str
}

func ByteResult(encoder *Encoder , messageClient int) []byte {
	strBytes := []byte(encoder.Str)

	data := make([]byte, len(strBytes)+1)
	copy(data, strBytes)

	//末尾补0说明字符串已结束
	data[len(strBytes)] = 0
	//组装
	var b byte = 0
	strLen := len(data)+8
	//dataLen := bLen + 12

	buff := new(bytes.Buffer)
	binary.Write(buff, binary.LittleEndian, int32Byte(strLen))
	binary.Write(buff, binary.LittleEndian, int32Byte(strLen))
	binary.Write(buff, binary.LittleEndian, int16Byte(messageClient))
	binary.Write(buff, binary.LittleEndian, b)
	binary.Write(buff, binary.LittleEndian, b)
	binary.Write(buff, binary.LittleEndian, data)

	return buff.Bytes()
}

func int32Byte(i int) []byte {
	x := int32(i)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x)
	return bytesBuffer.Bytes()
}

func int16Byte(i int) []byte {
	x := int16(i)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x)
	return bytesBuffer.Bytes()
}
