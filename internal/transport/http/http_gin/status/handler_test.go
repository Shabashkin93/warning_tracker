package status

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://192.168.3.182:8090/v1/status", nil)
	if err != nil {
		fmt.Println("error request:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	req, err = http.NewRequest("GET", "http://192.168.3.182:8090/v1/statuses", nil)
	if err != nil {
		fmt.Println("error request:", err)
		return
	}

	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
