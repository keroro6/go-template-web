package internalapi

type CodeType int

var (
	CodeServerErr     CodeType = -4 //服务错误
	CodeDBErr         CodeType = -3 //系统数据库错误
	CodeUserNotExists CodeType = -2 //用户不存在
	CodeParamErr      CodeType = -1 //请求参数错误
	CodeOk            CodeType = 0  // 正常
)

var CodeMap = map[CodeType]string{
	CodeServerErr:     "server err",
	CodeDBErr:         "internet err",
	CodeUserNotExists: "user not exist",
	CodeParamErr:      "param err",
	CodeOk:            "success",
}

const (
	DaySeconds  = 86400
	HourSeconds = 3600
)
