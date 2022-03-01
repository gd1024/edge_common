package mqtt

import "sync"

type MqttMsgEntity struct {
	TraceId  string                 `json:"trace_id"`
	MsgId    int64                  `json:"msg_id"`
	DeviceId string                 `json:"device_id"`
	Version  string                 `json:"version"`
	Source   MqttSource             `json:"source"`
	Mold     MqttCmdMold            `json:"mold"`
	Cmd      MqttCmd                `json:"cmd"`
	Content  map[string]interface{} `json:"content"`
}

var MqttMsgPool = sync.Pool{
	New: func() interface{} {
		return MqttMsgEntity{
			TraceId:  "",
			MsgId:    0,
			DeviceId: "",
			Version:  "",
			Source:   0,
			Mold:     0,
			Cmd:      0,
			Content:  nil,
		}
	},
}

//消息来源
type MqttSource int32

const (
	MqttSourceCollect   MqttSource = 100000 //采集器-100000
	MqttSourceEdge      MqttSource = 200000 //边缘计算器-200000
	MqttSourceEdgeAdmin MqttSource = 300000 //边缘计算管理-300000
	MqttSourceIot       MqttSource = 400000 //云端-400000
)

//指令类型
type MqttCmdMold int32

const (
	DeviceControl MqttCmdMold = 1000 //设备控制-10000
	DataHandle    MqttCmdMold = 2000 //数据处理-20000
)

//消息指令
type MqttCmd int32

const (
	CollectDeviceRegister MqttCmd = 101001 //采集器-设备-注册（上报）
	CollectDeviceDel      MqttCmd = 101003 //采集器-设备-销毁（上报）
	CollectDataReport     MqttCmd = 102001 //采集器-数据-上报（上报）

	EdgeDeviceRegister MqttCmd = 201001 //边缘计算器-设备-注册（上报）
	EdgeDeviceUpgrade  MqttCmd = 201002 //边缘计算器-设备-升级（下发）
	EdgeDataReport     MqttCmd = 202001 //边缘计算器-数据-上报（上报）

	EdgeAdminDeviceRegister MqttCmd = 301001 //边缘计算管理器-设备-注册（上报）
	EdgeAdminDeviceUpgrade  MqttCmd = 301002 //边缘计算管理器-设备-升级（下发）
	EdgeAdminDataReport     MqttCmd = 302001 //边缘计算管理器-数据-上报（上报）

	IotDeviceRegister MqttCmd = 401001 //云端-设备-注册（下发）
	IotDeviceUpgrade  MqttCmd = 401002 //云端-设备-升级（下发）
	IotData           MqttCmd = 402001 //云端-数据-下发
)
