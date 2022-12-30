package playground

import (
	"testing"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func TestXxx(t *testing.T) {
	err := errors.New("123")
	log.SetFormatter(&log.TextFormatter{
		DisableQuote: true, // 如果不为true,无法换行
	})
	log.Errorf("%+v\n", err)
}

//mqtt事件
type MQTTEvent struct {
}

// 设备事件
type DeviceEvent struct {
}

type IotDevice interface {
	// mqtt处理器
	MQTTHandler(event MQTTEvent)
	// 设备处理器
	DeviceHandler(event DeviceEvent)
	// 执行器（自由业务处理）
	Processor()
}

type Sw7 struct {
}

func (s Sw7) MQTTHandler(event MQTTEvent) {
}
func (s Sw7) DeviceHandler(event DeviceEvent) {
}
func (s Sw7) Processor() {
}

func start(d IotDevice) {
	for {
		mqtt_event := mqtt_recv()
		device_event := device_recv()
		d.Processor()
		d.MQTTHandler(mqtt_event)
		d.DeviceHandler(device_event)
	}
}

func mqtt_recv() MQTTEvent {
	return MQTTEvent{}
}

func device_recv() DeviceEvent {
	return DeviceEvent{}
}
