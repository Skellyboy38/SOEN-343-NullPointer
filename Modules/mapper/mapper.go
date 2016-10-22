package mapper

import (
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Modules/dB"
	_ "github.com/lib/pq"
	"reflect"
	"strconv"
	"strings"
)

func Insert(object interface{}) error {
	query := "INSERT INTO " + strings.TrimPrefix(reflect.TypeOf(object).String(), "*main.") + " ("
	val := reflect.Indirect(reflect.ValueOf(object))
	numfields := val.NumField()

	// name the fields to insert
	for x := 0; x < numfields; x++ {

		if (x + 1) == val.NumField() {
			query += val.Type().Field(x).Name + ")"
		} else {
			query += val.Type().Field(x).Name + ", "
		}
	}
	fmt.Println(query)
	query += " VALUES ("

	// insert the values of the fields
	for x := 0; x < numfields; x++ {

		valType := val.Field(x).Type().String()

		if valType == "int" {

			current := strconv.FormatInt(val.Field(x).Int(), 10)

			if (x + 1) == val.NumField() {
				query += current + ")"
			} else {
				query += current + ", "
			}

		} else {

			current := val.Field(x).String()

			if (x + 1) == val.NumField() {
				query += "'" + current + "')"
			} else {
				query += "'" + current + "', "
			}
		}
	}

	fmt.Println(query)

	return dB.Insert(query)
}

func GetAll(object interface{}) ([][]string, error) {
	// val := reflect.Indirect(reflect.ValueOf(object))
	query := "SELECT * FROM " + strings.TrimPrefix(reflect.TypeOf(object).String(), "*main.") + " ;"

	return dB.GetAll(query)
}
