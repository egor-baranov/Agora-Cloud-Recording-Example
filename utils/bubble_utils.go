package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/raysandeep/Estimator-App/schemas"
	"github.com/spf13/viper"
)

func PatchRecord(id string, payload map[string]interface{}) {
}

func FetchRecord(shortId string) (*schemas.BubbleRecord, error) {
}
