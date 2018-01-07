# GO wrapper code for accessing APIs created by the "luftdaten.info" project.

Sample code (datamodel, functions) accesses APIs (http://api.luftdaten.info/..) to get fine dust / particulate pollution data collected by a network of sensors.

Full details and all background information is available at [luftdaten.info](http://luftdaten.info) .

See [API info](https://github.com/opendata-stuttgart/meta/wiki/APIs) for details about APIs being used by Go sample code.

Sample code is available for the following subset of APIs:

- http://api.luftdaten.info/v1/sensor/{apiID}/ - all measurements of the last 5 minutes for one given sensor
- more to come

## Usage

- ensure you have Go installed and GOPATH is set 
- install package `go get github.com/dsasp/feinstaubgoapi`
- copy `main.go_sample` to your Go /src working directory and rename file to `main.go`
- run sample `go run main.go --sensor=<a_sensor_id>` . Open http://deutschland.maps.luftdaten.info/#6/51.165/10.455  and click on a hexagon on the map. This will show id for the related sensor.
