package utils

import (
	"time"
	"math/rand"
    "net/http"
    "net/url"
    "io/ioutil"
    "fmt"
    "crypto/md5"
    "encoding/hex"
)

// 生成32位MD5
func MD5(text string) string {
    ctx := md5.New()
    ctx.Write([]byte(text))
    return hex.EncodeToString(ctx.Sum(nil))
}

//获取当前时间
func GetTime() string {
    const shortForm = "20060102150405"
    t := time.Now()
    temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
    str := temp.Format(shortForm)
    return str
}
// 随机生成置顶位数的大写字母和数字的组合
func  GetRandomString(l int) string {
    //str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    str := "0123456789"
    bytes := []byte(str)
    result := []byte{}
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    for i := 0; i < l; i++ {
        result = append(result, bytes[r.Intn(len(bytes))])
    }
    return GetTime() + string(result)
}

// get 网络请求
func Get(apiURL string,params url.Values)(rs[]byte ,err error){
    var Url *url.URL
    Url,err=url.Parse(apiURL)
    if err!=nil{
        fmt.Printf("解析url错误:\r\n%v",err)
        return nil,err
    }
    //如果参数中有中文参数,这个方法会进行URLEncode
    Url.RawQuery=params.Encode()
    resp,err:=http.Get(Url.String())
    if err!=nil{
        fmt.Println("err:",err)
        return nil,err
    }
    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}
 
// post 网络请求 ,params 是url.Values类型
func Post(apiURL string, params url.Values)(rs[]byte,err error){
    resp,err:=http.PostForm(apiURL, params)
    if err!=nil{
        return nil ,err
    }
    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}
