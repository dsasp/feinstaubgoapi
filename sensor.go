/*
  Copyright (c) 2018 Dieter Schmidt
*/

package feinstaubgoapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var httpclient = &http.Client{Timeout: 10 * time.Second}

func ReadSensor(reqUrl string) (data []APIData, e error) {

	//log.Println("Reading sensor data from ", reqUrl)
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("Error creating http request, err=", err)
		return nil, err
	}

	resp, err := httpclient.Do(req)
	if err != nil {
		log.Println("GET request failed, err:", err)
		return nil, err
	}
	defer resp.Body.Close()

	apidata := make([]APIData, 2)
	err = json.NewDecoder(resp.Body).Decode(&apidata)

	if err != nil {
		log.Println("Failed to decode response body. Error=", err)
		return nil, err
	}

	//log.Printf("JSON: %+v\n", apidata)
	return apidata, nil
}

func ShowSensorData(apidata APIData) {
	fmt.Printf("\n*** SENSOR DATA ***\n")
	fmt.Printf("*** Sensor Id:           %v\n", apidata.Sensor.Id)
	fmt.Printf("*** Sensor name:         %s\n", apidata.Sensor.Sensor_type.Name)
	fmt.Printf("*** Sensor manufacturer: %s\n", apidata.Sensor.Sensor_type.Manufacturer)
	fmt.Printf("*** Sample time:         %s\n", apidata.Timestamp)
	fmt.Printf("*** Sensor value(s)\n")

	//log.Printf("apidata: %+v\n", apidata)

	for i := 0; i < len(apidata.Sensordatavalues); i++ {

		switch apidata.Sensordatavalues[i].Value_type {

		case ValTypeTemp:
			fmt.Printf("    Temperature:         %s\n", apidata.Sensordatavalues[i].Value)
		case ValTypeHum:
			fmt.Printf("    Humidity:            %s\n", apidata.Sensordatavalues[i].Value)
		case ValTypePM10:
			fmt.Printf("    PM10:                %s\n", apidata.Sensordatavalues[i].Value)
		case ValTypePM25:
			fmt.Printf("    PM25:                %s\n", apidata.Sensordatavalues[i].Value)
		default:
			fmt.Printf("    Unknown value type:  %s\n", apidata.Sensordatavalues[i].Value)
		}
	}
}
