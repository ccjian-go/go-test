package test

/**
	一. 执行单元测试
		go test ./test -run greet
 */
import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

//func Greet(writer *bytes.Buffer, name string) {
//	fmt.Fprintf(writer, "Hello, %s", name)
//}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer,"Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func main() {
	//Greet(os.Stdout, "Elodie")
	Greet(os.Stdout, "Elodie")
}