package trys

import "fmt"
import bytes "bytes"
import "os"

func Byte() {
	//buffer :=new(bytes.Buffer)
	s := fmt.Sprintf("%s", "abc")
	fmt.Println(s)

}

func ByteFprinf() {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "%s\n", "qwe")
	buffer.WriteTo(os.Stdout)

}
