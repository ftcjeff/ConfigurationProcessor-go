package logger

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

func Enter() string {
	method := MethodInfo()
	log.Printf("%s - [ENTER]\n", method)
	return method
}

func Trace(method string) {
	log.Printf("%s - [EXIT]\n", method)
}

func Log(s string) {
	log.Printf("%s - %s\n", MethodInfo(), s)
}

func LogMap(m map[string]interface{}) {
	log.Printf("%s - %v\n", MethodInfo(), m)
}

func LogFatal(s string) {
	log.Fatalf("%s - %s\n", MethodInfo(), s)
}

func MethodInfo() string {
	var method string

	funcAddr, file, line, _ := runtime.Caller(2)

	funcPtr := runtime.FuncForPC(funcAddr)
	funcFullName := funcPtr.Name()
	tokens := strings.Split(funcFullName, ".")
	funcName := tokens[1]

	tokens = strings.Split(file, "/")
	baseName := tokens[len(tokens)-1]

	method = fmt.Sprintf("%s|%s|%d", baseName, funcName, line)

	return method
}
