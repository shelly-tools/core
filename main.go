package main

import (
	"fmt"
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

type Relay struct {
	IsOn bool `json:"ison"`
}
type Input struct {
	IsOn bool `json:"ison"`
}

type Roller struct {
	Direction string `json:"direction"`
	Position  int    `json:"position"`
}
type ShellyDevice struct {
	Id          string   `json:"id" storm:"id,unique"`
	Ip          string   `json:"ip" storm:"index"`
	DeviceType  string   `json:"type" storm:"index"`
	Temperature float64  `json:"temperature"`
	Relay       []Relay  `json:"relays"`
	Input       []Input  `json:"inputs"`
	Roller      []Roller `json:"rollers"`
}

type Shelly struct {
	Id         string `storm:"id,unique"`
	DeviceType string `storm:"index"`
	Payload    string
	Ip         string    `storm:"unique"`
	CreatedAt  time.Time `storm:"index"`
}
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
	//fmt.Printf("%s - %s - %w", m.DeviceID(), m.DeviceType(), pl)
	Mutex.Lock()
	Data[m.DeviceID()] = pl
	Mutex.Unlock()
	/*
		var payload Payload
		json.Unmarshal([]byte(s), &payload)

		var temperature float64

		//fmt.Println(payload.G[0])
		em := Relay{}
		re := []Relay{}
		in := Input{}
		inp := []Input{}
		ro := Roller{}
		rol := []Roller{}
		has_roller := false
		for _, el := range payload.G {

			if el[1].(float64) == 1101 || el[1].(float64) == 1201 || el[1].(float64) == 1301 || el[1].(float64) == 1401 {
				em.IsOn = int(el[2].(float64)) != 0
				re = append(re, em)
			}
			if el[1].(float64) == 2101 || el[1].(float64) == 2201 || el[1].(float64) == 2301 || el[1].(float64) == 2401 {
				in.IsOn = int(el[2].(float64)) != 0
				inp = append(inp, in)
			}
			if el[1].(float64) == 1102 {
				ro.Direction = el[2].(string)
				has_roller = true
			}
			if el[1].(float64) == 1103 {
				ro.Position = int(el[2].(float64))
			}
			if el[1].(float64) == 3104 {
				temperature = el[2].(float64)
			}

		}
		if has_roller {
			rol = append(rol, ro)
		}

		ip := strings.TrimSuffix(a.String(), ":5683")
		emp := ShellyDevice{Id: m.DeviceID(), Temperature: temperature, Relay: re, Input: inp, Roller: rol, DeviceType: m.DeviceType(), Ip: ip}

		db, _ := storm.Open("devices.db")
		defer db.Close()

		err := db.Save(&emp)
		if err != nil {
			log.Println(err)
		}*/
	return nil
}

func main() {

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", Index)
	//r.HandleFunc("/device/{id}", ApiDeviceHandler)

	go func() {
		fmt.Println("CoIoT Listener started")
		log.Fatal(http.ListenAndServe(":8000", r))
	}()

	mux := coiot.NewServeMux()
	mux.Handle("/cit/s", coiot.FuncHandler(CoIoTHandler))
	log.Fatal(coiot.ListenAndServe("udp", "224.0.1.187:5683", mux))
}
func Index(w http.ResponseWriter, r *http.Request) {

	Mutex.Lock()
	//jsonData, err := json.Marshal(Data)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonData, _ := json.Marshal(&Data)
	Mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	/*	var shelly []ShellyDevice
		db, _ := storm.Open("devices.db")
		defer db.Close()
		err := db.All(&shelly)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(http.StatusOK)
		var jsonData []byte
		jsonData, _ = json.Marshal(shelly)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)

		/*for _, she := range shelly {
			fmt.Fprintf(w, "%s - %s - %f ", she.Id, she.DeviceType, she.Temperature)
			for idx, rel := range she.Relay {
				fmt.Fprintf(w, " - %d : %t ", idx, rel.IsOn)
			}


			fmt.Fprintf(w, "\n")
		} */
}

/* func ApiDeviceHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	id := vars["id"]
	dev := DeviceById[id]

	var payload Payload
	var temperature float64

	//fmt.Println(payload.G[0])
	em := Relay{}
	re := []Relay{}
	in := Input{}
	inp := []Input{}
	ro := Roller{}
	rol := []Roller{}
	has_roller := false
	w.Header().Set("Content-Type", "application/json")
	for _, el := range payload.G {

		if el[1].(float64) == 1101 || el[1].(float64) == 1201 || el[1].(float64) == 1301 || el[1].(float64) == 1401 {
			//fmt.Println("ID:", el[1], "=>", "Value:", el[2])
			em.IsOn = int(el[2].(float64)) != 0
			re = append(re, em)
		}
		if el[1].(float64) == 2101 || el[1].(float64) == 2201 || el[1].(float64) == 2301 || el[1].(float64) == 2401 {
			in.IsOn = int(el[2].(float64)) != 0
			inp = append(inp, in)
		}
		if el[1].(float64) == 1102 {
			ro.Direction = el[2].(string)
			has_roller = true
		}
		if el[1].(float64) == 1103 {
			ro.Position = int(el[2].(float64))
		}
		if el[1].(float64) == 3104 {
			temperature = el[2].(float64)
		}

	}
	if has_roller {
		rol = append(rol, ro)
	}

	ip := strings.TrimSuffix(dev.Ip, ":5683")
	emp := ShellyDevice{Id: id, Temperature: temperature, Relay: re, Input: inp, Roller: rol, DeviceType: dev.DeviceType, Ip: ip}
	var jsonData []byte
	jsonData, _ = json.Marshal(emp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}*/
