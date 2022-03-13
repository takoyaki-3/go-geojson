package gogeojson

type Geometry struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates"`
	Properties map[string]string `json:"properties"`
}
type Feature struct {
	Type       string            `json:"type"`
	Geometry   Geometry          `json:"geometry"`
	Properties map[string]string `json:"properties"`
}
type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
	Properties map[string]string `json:"properties"`
}
func NewFeatureCollection()*FeatureCollection{
	fc := new(FeatureCollection)
	fc.Type = "FeatureCollection"
	return fc
}
func NewLineString(coordinates [][]float64,props map[string]string)*Geometry{
	geom := new(Geometry)
	geom.Type = "LineString"
	geom.Coordinates = coordinates
	geom.Properties = props
	return geom
}
