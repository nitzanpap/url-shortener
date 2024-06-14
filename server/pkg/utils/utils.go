package utils

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func DoesContainEmptyStrings(values []interface{}, v reflect.Value) (bool, []string) {
	var emptyFields []string
	for i := 0; i < len(values); i++ {
		if values[i] == "" {
			emptyEntry := v.Type().Field(i).Name
			emptyFields = append(emptyFields, emptyEntry)
		} else if reflect.ValueOf(values[i]).Kind() == reflect.Struct {
			// if the value is a struct, recursively check for empty values
			innerValues := make([]interface{}, v.Field(i).NumField())
			for j := 0; j < v.Field(i).NumField(); j++ {
				innerValues[j] = v.Field(i).Field(j).Interface()
			}
			isInvalidConfig, errStringArr := DoesContainEmptyStrings(innerValues, v.Field(i))
			if isInvalidConfig {
				emptyFields = append(emptyFields, errStringArr...)
			}
		}
	}
	return len(emptyFields) > 0, emptyFields
}

// OkHandler accepts a version number that is an integer or nil.
func OkHandler(c *gin.Context, version *int) {
	if version == nil {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "version": version})
}
