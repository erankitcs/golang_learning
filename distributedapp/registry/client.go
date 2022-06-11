package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterService(r Registration) error {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(r)

	if err != nil {
		//fmt.Println(err)
		return err
	}

	res, err := http.Post(ServiceURL, "application/json", buf)

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to register service. Registry Service "+
			"responded with code %v:", res.StatusCode)
	}
	return nil

}
