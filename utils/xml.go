package utils

import (
	"encoding/xml"
	"io/ioutil"
)

func ToXML(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

func FromXML(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}

func ReadXMLFile(filename string, v interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return FromXML(data, v)
}
