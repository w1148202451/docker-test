package prco
import (
	"bufio"
	"fmt"
	"net"
	"os"
)
func ReadHost() {
	fmt.Println("=== 配置githubDNS ===")

	file,err := os.Open("./github_domain.txt")
	if err == nil {
		r := bufio.NewReader(file)
		for tmp,_,err := r.ReadLine();err == nil;tmp,_,err = r.ReadLine() {
			host := string(tmp)
			addrs,err := net.LookupHost(host)
			if err == nil {
				for _,v := range addrs {
					fmt.Println(v,host)
				}
			}
		}
	}
}
