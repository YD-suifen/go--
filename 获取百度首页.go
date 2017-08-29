import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func baiduDate() string {
	resp, err := http.Get("http://www.baidu.com")  //��������get�õ��ٶ�ҳ���html��Ȼ�󷵻�
	if err != nil{              //�����ж�get��ȡ�Ƿ�ɹ�
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body) //����ʹ��ReadAll��ȡ��ȡ����html���ݣ�����ReadAll���ص���һ������
	if err != nil{
		fmt.Println(err)
	}
	return string(body)  //��ΪReadAll���ص���һ�����飬����������Ҫ����ǿ��ת��Ϊstring���ͣ�
}

//���￪��һ���򵥵�httpӦ�ã����ݽ��ոջ�ȡ�İٶ�ҳ�����
func aaa(w http.ResponseWriter, r *http.Request)  {  //
	fmt.Fprintln(w, baiduDate())
}
func main()  {
	http.HandleFunc("/", aaa)
	http.ListenAndServe("127.0.0.1:8888",nil)

}