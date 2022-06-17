package main

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"log"
)

func main() {
	conn, err := dbus.SystemBus()
	fmt.Println(err)

	/*go func() {
		for {
			fmt.Println(111)
			err = conn.Emit("/com/deepin/sync/Helper", "org.freedesktop.DBus.Properties.PropertiesChanged",
				"com.deepin.sync.Helper", map[string]interface{}{"DeveloperMode": true}, []string{})
			time.Sleep(2*time.Second)
		}
	}()*/

	caller := conn.Object("com.deepin.sync.Helper", "/com/deepin/sync/Helper")
	caller.AddMatchSignal("org.freedesktop.DBus.Properties", "PropertiesChanged",
		dbus.WithMatchObjectPath(dbus.ObjectPath("/com/deepin/sync/Helper")))

	sc := make(chan *dbus.Signal, 20)
	conn.Signal(sc)

	go func() {
		for {
			select {
			case v := <-sc:
				log.Println("--- get signal from %s, name:[%s], path:[%s], body:{ %v } ---", v.Sender, v.Name, v.Path, v.Body)
			} //end select
		} //end for
	}()
	select {}
}
