/*
  Copyright (c) 2018 Dieter Schmidt

	API data structures

*/

package feinstaubgoapi

const UrlSensorAPI_1 = "http://api.luftdaten.info/v1/sensor/"

const ValTypeTemp = "temperature"
const ValTypeHum = "humidity"
const ValTypePM10 = "P1"
const ValTypePM25 = "P2"

type SensorDataValue struct {
	Value_type string `json: "value_type"`
	Value      string `json: "value"`
	Id         int    `json: "id"`
}

type SensorType struct {
	Id           int    `json: "id"`
	Name         string `json: "name"`
	Manufacturer string `json: "manufacturer"`
}

type SensorInfo struct {
	Id          int        `json: "id"`
	Sensor_type SensorType `json: "sensor_type"`
	Pin         string     `json: "pin"`
}

type LocationInfo struct {
	Country   string `json: "country"`
	Id        int    `json: "id"`
	Longitude string `json: "longitude"`
	Latitude  string `json: "latitude"`
}

type APIData struct {
	Timestamp        string            `json: "timestamp"`
	Id               int               `json: "id"`
	Sensordatavalues []SensorDataValue `json: "sensordatavalues"`
	Location         LocationInfo      `json: "location"`
	Sensor           SensorInfo        `json: "sensor"`
	SamplingRate     string            `json: "samplingrate,omitempty"`
}
