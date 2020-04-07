package getsubcommands

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func toJson(data interface{}, prettified bool) ([]byte, error) {
	var err error
	var jsonData []byte

	if prettified {
		jsonData, err = json.MarshalIndent(data, "", "  ")
		if err != nil {
			return nil, err
		}
	} else {
		jsonData, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
	}

	return jsonData, nil
}

func toXml(data interface{}, prettified bool) ([]byte, error) {
	var err error
	var xmlData []byte

	if reflect.ValueOf(data).Kind() == reflect.Map {
		byteData := data.(*[]uint8)
		xmlData, err = xml.MarshalIndent(byteData, "", "  ")
		if err != nil {
			return nil, err
		}
	}

	if prettified {
		xmlData, err = xml.MarshalIndent(data, "", "  ")
		if err != nil {
			return nil, err
		}
	} else {
		xmlData, err = xml.Marshal(data)
		if err != nil {
			return nil, err
		}
	}

	return xmlData, nil
}

func toHumanReadable(data interface{}, limit int) string {
	//the standard limit is 3
	output := superReflect(reflect.ValueOf(data), 0, limit)

	return output
}

/*
Custom function to properly generate human readable outputs
*/
func superReflect(reflectedData reflect.Value, depth int, limit int) string {
	var outputString string

	kind := reflectedData.Type().Kind()
	if kind != reflect.String && kind != reflect.Int && kind != reflect.Float64 && kind != reflect.Ptr {
		if depth > 0 {
			outputString += "\n"
		}
		outputString += strings.Repeat("  ", depth)
		outputString += reflectedData.Type().Name()
	}
	depth++

	switch kind {
	case reflect.Struct:
		outputString += "\n"
		for i := 0; i < reflectedData.NumField(); i++ {
			if reflectedData.Type().Field(i).Type.Kind() != reflect.Slice {
				outputString += strings.Repeat("  ", depth)
				outputString += reflectedData.Type().Field(i).Name + ": "
			}
			outputString += superReflect(reflectedData.Field(i), depth, limit)
		}
	case reflect.Slice:
		outputString += "(" + strconv.Itoa(reflectedData.Len()) + ") \n"
		if reflectedData.Len() == 0 {
			outputString += strings.Repeat("  ", depth)
			outputString += "/"
		}
		for j := 0; j < reflectedData.Len(); j++ {
			if depth < limit {
				outputString += superReflect(reflectedData.Index(j), depth, limit)
			}
		}
		outputString += "\n"
	case reflect.Map:
		outputString += "(" + strconv.Itoa(reflectedData.Len()) + ") \n"
		for _, key := range reflectedData.MapKeys() {
			outputString += strings.Repeat("  ", depth)
			outputString += key.String() + ": "
			if depth < limit {
				outputString += superReflect(reflectedData.MapIndex(key), depth, limit)
			}
		}
	case reflect.Int:
		fieldValue := strconv.Itoa(int(reflectedData.Int()))
		outputString += fieldValue + "\n"
	case reflect.String:
		fieldValue := reflectedData.String()
		outputString += fieldValue + "\n"
	case reflect.Float64:
		fieldValue := strconv.FormatFloat(reflectedData.Float(), 'f', -1, 64)
		outputString += fieldValue + "\n"
	case reflect.Ptr:
		indirect := reflect.Indirect(reflectedData)
		if indirect.Kind() == reflect.Invalid {
			outputString += "/\n"
		} else {
			outputString += strconv.Itoa(int(indirect.Int())) + "\n"
		}
	default:
		log.Debug().
			Msg("Could not reflect " + reflectedData.Type().Kind().String())
	}

	depth--
	return outputString
}

func parsePersistentFlags(cmd *cobra.Command) (string, int, bool) {
	format := cmd.Flag("format").Value.String()
	if !validateFormat(format) {
		log.Error().
			Msg("Invalid format")
		os.Exit(1)
	}

	depth, err := strconv.Atoi(cmd.Flag("depth").Value.String())
	if err != nil {
		log.Error().
			Msg("Error during conversion of 'depth' flag")
		os.Exit(1)
	}

	prettified, err := strconv.ParseBool(cmd.Flag("pretty").Value.String())
	if err != nil {
		log.Error().
			Msg("Error during conversion of 'pretty' flag")
		os.Exit(1)
	}

	return format, depth, prettified
}

// PrintData prints the interface 'data' that is handed in
func PrintData(format string, rawData interface{}, prettified bool, depth int) error {
	switch format {
	case "xml":
		data, err := toXml(rawData, prettified)
		if err != nil {
			return err
		}
		fmt.Println(string(data))
		break
	case "json":
		data, err := toJson(rawData, prettified)
		if err != nil {
			return err
		}
		fmt.Println(string(data))
		break
	case "human-readable":
		data := toHumanReadable(rawData, depth)
		fmt.Println(data)
		break
	default:
		break
	}

	return nil
}

func validateFormat(format string) bool {
	for _, allowedFormat := range []string{"xml", "json", "human-readable"} {
		if format == allowedFormat {
			return true
		}
	}

	return false
}
