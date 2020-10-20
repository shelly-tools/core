package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	jsoniter "github.com/json-iterator/go"
	"github.com/shelly-tools/coiot"
	"github.com/tidwall/gjson"
)

var tmpl = template.Must(template.ParseGlob("ui/views/*"))

type Payload struct {
	G [][]interface{} `json:"G"`
}

var netClient = &http.Client{
	Timeout: time.Second * 5,
}
var Data = map[string]map[int]interface{}{}
var Mutex = &sync.Mutex{}

func CoIoTHandler(l *net.UDPConn, a *net.UDPAddr, m *coiot.Message) *coiot.Message {

	//s := string(m.Payload)
	// non-coiot message? just skip..
	if m.OptionDevice() == nil {
		return nil
	}
	mp := make(map[int]interface{})
	pl := make(map[int]interface{})
	if m != nil {
		keys := gjson.GetBytes(m.Payload, "G.#.1")
		values := gjson.GetBytes(m.Payload, "G.#.2")
		for k, v := range values.Array() {
			mp[k] = v.Value()
		}
		for k, v := range keys.Array() {
			pl[int(v.Int())] = mp[k]
		}
	}
	Mutex.Lock()
	Data[m.DeviceID()+"#"+m.DeviceType()+"#"+a.String()] = pl
	Mutex.Unlock()

	return nil
}

func main() {

	r := mux.NewRouter()
	staticDir := "/ui/assets/"
	// Create the route
	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
	// Routes consist of a path and a handler function.
	r.HandleFunc("/api/v1/devices", DeviceIndex)
	r.HandleFunc("/", DashboardIndex)

	// Rooms
	r.HandleFunc("/ui/manage/rooms", RoomList)
	r.HandleFunc("/ui/manage/rooms/add", RoomAdd)
	r.HandleFunc("/ui/manage/rooms/edit/{id}", RoomEdit)
	r.HandleFunc("/ui/manage/rooms/delete/{id}", RoomDelete)
	// Buildings
	r.HandleFunc("/ui/manage/buildings", BuildingList)
	r.HandleFunc("/ui/manage/buildings/add", BuildingAdd)
	r.HandleFunc("/ui/manage/buildings/insert", BuildingInsert)
	r.HandleFunc("/ui/manage/buildings/edit/{id}", BuildingEdit)
	r.HandleFunc("/ui/manage/buildings/update", BuildingUpdate)
	r.HandleFunc("/ui/manage/buildings/delete/{id}", BuildingDelete)
	r.HandleFunc("/ui/manage/buildings/remove/{id}", BuildingRemove)

	// Devices
	r.HandleFunc("/ui/manage/devices/discovered", DiscoveredDevices)

	//r.HandleFunc("/api/v1/device/{id}", ApiDeviceHandler)

	go func() {
		fmt.Println("CoIoT Listener started")
		log.Fatal(http.ListenAndServe(":8000", r))
	}()

	mux := coiot.NewServeMux()
	mux.Handle("/cit/s", coiot.FuncHandler(CoIoTHandler))
	log.Fatal(coiot.ListenAndServe("udp", "224.0.1.187:5683", mux))
}
func DeviceIndex(w http.ResponseWriter, r *http.Request) {

	Mutex.Lock()
	//jsonData, err := json.Marshal(Data)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonData, _ := json.Marshal(&Data)
	Mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
func DashboardIndex(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "DashboardIndex", "")
}
