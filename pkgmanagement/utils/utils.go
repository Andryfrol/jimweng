package utils

type Output interface {
	// Write takes in group of points to be written to the Output
	Write(points *[]*[]*DBSchema) error
}

type Input interface {
	// Gather takes in an accumulator and adds the inputInfo to the Input
	Gather() (interface{}, error)
}

type DBSchema struct {
	Hostname   string
	Ptime      int64
	Name       string
	TimePeriod int
	Counter    string
	Ingress    int
	Egress     int
	KeyDepth   int
	KeyValue   []string
	Drilldown  int
}
