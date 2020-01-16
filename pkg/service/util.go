package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"reflect"
	"strings"
	"time"
)

func stringFromPointer(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func structToMapClaims(i interface{}) (jwt.MapClaims, error) {
	now := time.Now()
	keyValues := jwt.MapClaims{
		"exp": now.Add(tokenExpirationTime).Unix(),
		"iat": now.Unix(),
	}

	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	typElem := reflect.TypeOf(i).Elem()

	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		structFieldName := typ.Field(i).Name
		typField, found := typElem.FieldByName(structFieldName)
		if !found {
			return keyValues, fmt.Errorf("expected field not found in jwt claims srtuct")
		}
		name := strings.Split(typField.Tag.Get("jwt"), ",")[0]
		if name == "-" {
			continue
		}

		var v interface{}
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = f.Int()
		case uint, uint8, uint16, uint32, uint64:
			v = f.Uint()
		case float32, float64:
			v = f.Float()
		case []byte:
			v = f.Bytes()
		case string:
			v = f.String()
			if v == "" {
				v = nil
			}
		}
		if v != nil {
			keyValues[name] = v
		}
	}

	return keyValues, nil
}
