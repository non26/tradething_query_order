package config

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func ReadAWSAppConfig() (*Config, error) {
	region := os.Getenv("app_config")
	base_url := fmt.Sprintf("http://localhost:2772%v", region)
	res, err := http.Get(base_url)
	if err != nil {
		return nil, fmt.Errorf("can't pull config from AWS app config: %v", err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status from AWS app config is not ok: %v", res.StatusCode)
	}
	var m Config
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("readAll error %v", err.Error())
	}
	json.Unmarshal(body, &m)
	return &m, nil
}
