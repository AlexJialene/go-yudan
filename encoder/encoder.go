package encoder

import "strings"

type Encoder struct {
	Str string
}

func GetEncoder() *Encoder {
	e := &Encoder{
		Str: "",
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
	return encoder.Str+"\0"
}
