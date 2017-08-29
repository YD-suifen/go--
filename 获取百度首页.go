import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func baiduDate() string {
	resp, err := http.Get("http://www.baidu.com")  //这里利用get拿到百度页面的html，然后返回
	if err != nil{              //这里判断get获取是否成功
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body) //这里使用ReadAll读取获取到的html内容，这里ReadAll返回的是一个数组
	if err != nil{
		fmt.Println(err)
	}
	return string(body)  //因为ReadAll返回的是一个数组，所以我们需要把它强制转换为string类型，
}

//这里开启一个简单的http应用，内容将刚刚获取的百度页面放上
func aaa(w http.ResponseWriter, r *http.Request)  {  //
	fmt.Fprintln(w, baiduDate())
}
func main()  {
	http.HandleFunc("/", aaa)
	http.ListenAndServe("127.0.0.1:8888",nil)

}