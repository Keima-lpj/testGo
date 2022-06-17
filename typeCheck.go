package main

import "fmt"

type usb interface {
	start()
	stop()
}

type phone struct {
	Name string
}

type camera struct {
	Name string
}

func (p phone) start() {
	fmt.Println("phone start")
}

func (p phone) stop() {
	fmt.Println("phone stop")
}

func (p phone) call() {
	fmt.Println("phone call")
}

func (p camera) start() {
	fmt.Println("camera start")
}

func (p camera) stop() {
	fmt.Println("camera stop")
}

func (p camera) photo() {
	fmt.Println("camera take a photo")
}

/**
测试类型断言
*/
func main() {
	var usbArr = make([]usb, 3)
	usbArr[0] = phone{"apple"}
	usbArr[1] = phone{"xiaomi"}
	usbArr[2] = camera{"nikon"}

	for _, v := range usbArr {
		working(v)
	}

}

func working(usb usb) {
	usb.start()

	//这里类型断言，判断是phone还是camera
	if phone, ok := usb.(phone); ok {
		phone.call()
	}
	if camera, ok := usb.(camera); ok {
		camera.photo()
	}

	usb.stop()
}
