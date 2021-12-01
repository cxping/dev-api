package model

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

//车辆出三居
type Odom struct {
	Uuid       string `json:"uuid"`
	Timestamp  int64  `json:"timestamp"`
	Vehicle_id string `json:"vehicle_id"`
	Map_id     string `json:"map_id"`
	Task_id    string `json:"task_id"`
	Info       struct {
		Position_id   int `json:"position_id"`
		Position_info struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
			Z float64 `json:"z"`
		} `json:"position_info"`
		Gear      int     `json:"gear"`
		Power     float64 `json:"power"`
		Seteering float64 `json:"seteering"`
		///实时速度
		Average_speend string `json:"average_speend"`
	} `json:"info"`
	///车辆状态(0:自动 1 手动)
	Status int8 `json:"status"`
}

func Gen_data() Odom {
	rand.Seed(time.Now().UnixNano())
	return Odom{
		Uuid:       uuid.NewString(),
		Timestamp:  time.Now().Unix(),
		Vehicle_id: "aa495f8c-78d8-4e6f-861f-3f19ad488147",
		Map_id:     uuid.NewString(),
		Task_id:    uuid.NewString(),
		Info: struct {
			Position_id   int `json:"position_id"`
			Position_info struct {
				X float64 `json:"x"`
				Y float64 `json:"y"`
				Z float64 `json:"z"`
			} `json:"position_info"`
			Gear           int     `json:"gear"`
			Power          float64 `json:"power"`
			Seteering      float64 `json:"seteering"`
			Average_speend string  `json:"average_speend"`
		}{
			Position_id: 0,
			Position_info: struct {
				X float64 `json:"x"`
				Y float64 `json:"y"`
				Z float64 `json:"z"`
			}{
				X: rand.Float64(),
				Y: rand.Float64(),
				Z: rand.Float64(),
			},
			Gear:           0,
			Power:          24.0,
			Seteering:      0,
			Average_speend: "0.5",
		},
		Status: 0,
	}
}
