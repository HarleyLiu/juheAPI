package juheAPI

import (
	"errors"
)

var (
	ErrorCodes map[int64]error
	//都是中国人，中文吧
	ParamErrorWrongMethod = errors.New("参数错误:请求方法")
	ParamErrorNoURL       = errors.New("参数错误:无请求URL")
	ParamErrorNoKey       = errors.New("参数错误:无请求key")
	ParamErrorNoRealName  = errors.New("参数错误:无姓名")
	ParamErrorWrongIdCard = errors.New("参数错误:身份证格式错误")
	ParamErrorNoUserId    = errors.New("参数错误:无用户id")

	ErrorNoData    = errors.New("无返回数据")
	ErrorUnknow    = errors.New("未知错误")
	ErrorCheckFail = errors.New("校验失败")

	ErrorMoreRequest = errors.New("频繁查询")
	ErrorNoRequestID = errors.New("未设置用户id")
)

// 服务级错误码参照(error_code)：
//   	210301 	库中无此身份证记录
//   	210302 	第三方服务器异常
//   	210303 	服务器维护
//   	210304 	参数异常
//   	210305 	网络错误，请重试
//   	210306 	参数错误，具体参照reason

// 系统级错误码参照：
//   	10001 	错误的请求KEY 	101
//   	10002 	该KEY无请求权限 	102
//   	10003 	KEY过期 	103
//   	10004 	错误的OPENID 	104
//   	10005 	应用未审核超时，请提交认证 	105
//   	10007 	未知的请求源 	107
//   	10008 	被禁止的IP 	108
//   	10009 	被禁止的KEY 	109
//   	10011 	当前IP请求超过限制 	111
//   	10012 	请求超过次数限制 	112
//   	10013 	测试KEY超过请求限制 	113
//   	10014 	系统内部异常(调用充值类业务时，请务必联系客服或通过订单查询接口检测订单，避免造成损失) 	114
//   	10020 	接口维护 	120
//   	10021 	接口停用 	121
// 错误码格式说明（示例：200201）：
// 2 	002 	01
// 服务级错误（1为系统级错误） 	服务模块代码(即数据ID) 	具体错误代码

func init() {
	ErrorCodes = make(map[int64]error)
	ErrorCodes[210301] = errors.New("库中无此身份证记录")
	ErrorCodes[210302] = errors.New("第三方服务器异常")
	ErrorCodes[210303] = errors.New("服务器维护")
	ErrorCodes[210304] = errors.New("参数异常")
	ErrorCodes[210305] = errors.New("网络错误请重试")
	ErrorCodes[210306] = errors.New("参数错误具体参照reason")
	ErrorCodes[10001] = errors.New("错误的请求KEY")
	ErrorCodes[10002] = errors.New("该KEY无请求权限")
	ErrorCodes[10003] = errors.New("KEY过期")
	ErrorCodes[10004] = errors.New("错误的OPENID")
	ErrorCodes[10005] = errors.New("应用未审核超时请提交认证")
	ErrorCodes[10007] = errors.New("未知的请求源")
	ErrorCodes[10008] = errors.New("被禁止的IP")
	ErrorCodes[10009] = errors.New("被禁止的KEY")
	ErrorCodes[10011] = errors.New("当前IP请求超过限制")
	ErrorCodes[10012] = errors.New("请求超过次数限制")
	ErrorCodes[10013] = errors.New("测试KEY超过请求限制")
	ErrorCodes[10014] = errors.New("系统内部异常(调用充值类业务时，请务必联系客服或通过订单查询接口检测订单，避免造成损失)")
	ErrorCodes[10020] = errors.New("接口维护")
	ErrorCodes[10021] = errors.New("接口停用")
}

func GetError(num int64) error {
	if err, ok := ErrorCodes[num]; ok {
		return err
	}
	return ErrorUnknow
}
