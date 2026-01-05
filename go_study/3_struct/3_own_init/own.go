package main

// 定义一个自定义类型 Code，用于表示状态码，并为其定义常量和方法。
// 通过为 Code 类型定义 GetMsg 方法，可以根据状态码获取对应的字符串信息，提升代码的可读性和维护性。
type Code int

const (
	SuccessCode    Code = 0    // 成功
	ServiceErrCode Code = 1001 // 服务错误
	NetworkErrCode Code = 1002 // 网络错误
)

// GetMsg 方法用于将 Code 类型转换为对应的字符串信息
func (c Code) GetMsg() string {
	switch c {
	case SuccessCode:
		return "成功"
	case ServiceErrCode:
		return "服务错误"
	case NetworkErrCode:
		return "网络错误"
	default:
		return "未知错误"
	}
}

// getCodeMsg 根据状态码获取对应的信息
// func getCodeMsg(code Code) (msg string) {
// 	switch code {
// 	case SuccessCode:
// 		return "成功"
// 	case ServiceErrCode:
// 		return "服务错误"
// 	case NetworkErrCode:
// 		return "网络错误"
// 	default:
// 		return "未知错误"
// 	}
// }

//	func webServer(name string) (code Code, msg string) {
//		if name == "1" {
//			return ServiceErrCode, getCodeMsg(ServiceErrCode)
//		}
//		if name == "2" {
//			return NetworkErrCode, getCodeMsg(NetworkErrCode)
//		}
//		return SuccessCode, getCodeMsg(SuccessCode)
//	}

func webServer(name string) (code Code, msg string) {
	if name == "1" {
		return ServiceErrCode, ServiceErrCode.GetMsg()
	}
	if name == "2" {
		return NetworkErrCode, NetworkErrCode.GetMsg()
	}
	return SuccessCode, SuccessCode.GetMsg()
}

func main() {
	code, msg := webServer("2")
	println("code:", code, "msg:", msg)
}
