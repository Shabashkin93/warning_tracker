package warning

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://192.168.3.182:8090/v1/warning", nil)
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

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
