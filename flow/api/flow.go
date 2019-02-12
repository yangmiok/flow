package api
import (
	"github.com/gin-gonic/gin" 
    "fmt"
    "flow/utils"
    "net/url"
    "encoding/json"
)
const APPKEY = "0908edab1c3427d45de4162541449e57" //您申请的APPKEY
const OpenID = "JH5bfefbce885337b560fb841e6f185628" //OpenID在个人中心查询
// 全部流量套餐列表
func FlowList(ctx *gin.Context) {
    //请求地址
    juheURL :="http://v.juhe.cn/flow/list"
    //初始化参数
    param:=url.Values{}
    //配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
    param.Set("key",APPKEY) //应用APPKEY(应用详细页查询)
 
    //发送请求
    data,err:=utils.Post(juheURL,param)
    if err!=nil{
        fmt.Errorf("请求失败,错误信息:\r\n%v",err)
        ctx.JSON(404, gin.H{
            "code": "404",
            "message": err,
            })
    }else{
        var netReturn map[string]interface{}
        json.Unmarshal(data,&netReturn)

        ctx.JSON(200, gin.H{
            "error_code": netReturn["error_code"],
            "message": netReturn["reason"],
            "result":netReturn["result"],
        })
    }

}

//根据手机号码获取支持的流量套餐
func FlowTelcheck(ctx *gin.Context) {
    //获取请求参数
    phone := ctx.PostForm("phone")
    //请求地址
    juheURL :="http://v.juhe.cn/flow/telcheck"
    //初始化参数
    param:=url.Values{}
    //配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
    param.Set("phone",phone) //手机号码
    param.Set("key",APPKEY) //应用APPKEY(应用详细页查询)
 
    //发送请求
    data,err:=utils.Post(juheURL,param)
    if err!=nil{
        fmt.Errorf("请求失败,错误信息:\r\n%v",err)
        ctx.JSON(404, gin.H{
            "code": "404",
            "message": err,
            })
    }else{
        var netReturn map[string]interface{}
        json.Unmarshal(data,&netReturn)

        ctx.JSON(200, gin.H{
            "error_code": netReturn["error_code"],
            "message": netReturn["reason"],
            "result":netReturn["result"],
        })
    }

}

//提交流量充值
func FlowRecharge(ctx *gin.Context) {
    phone := ctx.PostForm("phone")
    pid := ctx.PostForm("pid")
    orderid := utils.GetRandomString(6)
        //请求地址
    juheURL :="http://op.juhe.cn/ofpay/mobile/onlineorder"
 
    //初始化参数
    param:=url.Values{}
    
    //校验值，md5(OpenID+key+phoneno+cardnum+orderid)
    sign := utils.MD5(OpenID+APPKEY+phoneno+pid+orderid)

    //配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
    param.Set("phone",phone) //手机号码
    param.Set("pid",pid) //流量套餐ID
    param.Set("orderid",orderid) //商家订单号，8-32位字母数字组合
    param.Set("key",APPKEY) //应用APPKEY(应用详细页查询)
    param.Set("sign",sign) //校验值，md5(OpenID+key+phoneno+cardnum+orderid)，OpenID在个人中心查询
 
 
    //发送请求
    data,err:=utils.Post(juheURL,param)
    if err!=nil{
        fmt.Errorf("请求失败,错误信息:\r\n%v",err)
        ctx.JSON(404, gin.H{
            "code": "404",
            "message": err,
            })
    }else{
        var netReturn map[string]interface{}
        json.Unmarshal(data,&netReturn)

        ctx.JSON(200, gin.H{
            "error_code": netReturn["error_code"],
            "message": netReturn["reason"],
            "result":netReturn["result"],
        })
    }

}

//订单状态查询
func FlowOrdersta(ctx *gin.Context) {


}

//状态回调配置
func Setpush(ctx *gin.Context) {


}
