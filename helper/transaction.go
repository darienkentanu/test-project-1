package helper

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func get4Digit(input int) int {
	inputString := strconv.Itoa(input)
	outputString := fmt.Sprintf("%v%v%v%v",
		string(inputString[0]),
		string(inputString[1]),
		string(inputString[2]),
		string(inputString[3]),
	)
	output, err := strconv.Atoi(outputString)
	if err != nil {
		log.Println(err)
	}
	return output
}

func CreateTransactionNumber() int {
	idt := fmt.Sprint(time.Now().Year(),
		int(time.Now().Month()),
		get4Digit(time.Now().Nanosecond()),
	)
	idtrim := strings.Replace(idt, " ", "", -1)
	idtrans, _ := strconv.Atoi(idtrim)
	return idtrans
}
