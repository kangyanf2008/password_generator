package main
import "C"
import (
	"encryption"
	"flag"
	"fmt"
	"os"
)

func main() {
	var sn string
	var snFile string
	flag.StringVar(&sn, "sn", "", "")
	flag.StringVar(&snFile, "snFile", "", "")
	flag.Parse()

	if snFile != "" {
		err := encryption.BatchCreateSnPWBySnFile(snFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		os.Exit(0)
	}

	if sn ==  "" {
		fmt.Println("sn is empty")
		os.Exit(-1)
	} else {
		//生成20位密码字符串
		fmt.Println(encryption.SecretAESBase64Md532Len20(sn))
		os.Exit(0)
	}

}
//export SecretAESBase64Md532
/*func SecretAESBase64Md532(sn string) string {
	currentPath, _ :=os.Getwd()
	filePath := filepath.Join(currentPath, utils.FormatTime(time.Now())+".txt")
	f, _ := os.Create(filePath) //创建文件
	defer f.Close()
	w := bufio.NewWriter(f) //创建新的 Writer 对象
	w.WriteString(sn)
	w.Flush()

	return encryption.SecretAESBase64Md532Len20(sn)
}
*/
//export SecretAESBase64Md532
func SecretAESBase64Md532(src *C.char) *C.char {
	sn := C.GoString(src)
	return C.CString(encryption.SecretAESBase64Md532Len20(sn))
}