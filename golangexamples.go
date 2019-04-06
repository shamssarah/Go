//Sarah Shams
//i150097
package golangexamples

import "github.com/ehteshamz/greetings"

func ConcatSlice(sliceToConcat []byte) string {

	var concatenated string
	for _, v := range sliceToConcat[:len(sliceToConcat)-1] {
		concatenated = concatenated + string(v) + "-"
	}
	return concatenated + string(sliceToConcat[len(sliceToConcat)-1])
}

func Encrypt(sliceToEncrypt *[]byte, ceaserCount int) {
	slice := *sliceToEncrypt
	for i, _ := range slice {
		slice[i] = byte(int(slice[i]) + ceaserCount)
	}
}

func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)

}
