package fetcher

import (
	"bufio"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"math/rand"
	"time"
)

 var my_headers = []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.153 Safari/537.36", "Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36",
"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:30.0) Gecko/20100101 Firefox/30.0", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.75.14 (KHTML, like Gecko) Version/7.0.3 Safari/537.75.14",
"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)"}

 var rateLimiter = time.Tick(20*time.Millisecond)
func Fetch(url string)([]byte,error){
	<- rateLimiter
	//request,err := http.NewRequest("POST",url,nil)
	//if err!=nil {
	//	beego.Info("生成request失败",err)
	//	return nil,err
	//}
	head,_ := Random(my_headers,len(my_headers))
	//request.Header.Add("User-Agent",head)
	//request.Header.Add("Content-Type","application/json")
	//request.Header.Add("Referer",url)
	//client := http.DefaultClient
	//resp, err := client.Do(request)
	//defer resp.Body.Close()
	//if resp.StatusCode != http.StatusOK {
	//	beego.Info(resp.StatusCode)
	//	return nil,errors.New("Wrong status code:"+string(resp.StatusCode))
	//}
	//reader := bufio.NewReader(resp.Body)
	//e := determineEncoding(reader)
	//newReader := transform.NewReader(resp.Body,e.NewDecoder())
	//return ioutil.ReadAll(newReader)
	req := httplib.Post(url)
	req.Header("Content-Type","application/json")
	req.Header("user-agent",head)
	req.Header("Referer",url)
	rs,err := req.Bytes()
	if err!=nil {
		beego.Info(err)
	}
	return rs,err
}

/**
转换为UTF8编码格式
 */
func determineEncoding(reader *bufio.Reader) encoding.Encoding{
	bytes,err := reader.Peek(2048)
	if err != nil{
		beego.Info("Fetcher error :",err)
		return unicode.UTF8
	}
	e,_,_ := charset.DetermineEncoding(bytes,"")
	return e
}

func Random(strings []string, length int) (string, error) {
	if len(strings) <= 0 {
		return "", errors.New("the length of the parameter strings should not be less than 0")
	}

	if length <= 0 || len(strings) <= length {
		return "", errors.New("the size of the parameter length illegal")
	}

	for i := len(strings) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		strings[i], strings[num] = strings[num], strings[i]
	}

	str := ""
	for i := 0; i < length; i++ {
		str += strings[i]
	}
	return str, nil
}
