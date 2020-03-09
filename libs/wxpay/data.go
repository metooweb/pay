package wxpay

import (
	"code.metooweb.com/payment/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type DataBase struct {
	vals map[string]interface{}
}

func (data *DataBase) Get(name string) interface{} {

	if val, ok := data.vals[name]; !ok {
		return nil
	} else {
		return val
	}
}

func (data *DataBase) GetString(name string) string {

	if val, ok := data.vals[name]; !ok {
		return ""
	} else {
		return val.(string)
	}

}

func (data *DataBase) GetInt(name string) int {

	if val, ok := data.vals[name]; !ok {
		return 0
	} else {
		switch val.(type) {
		case int:
			return val.(int)
		case string:
			var (
				res int
				err error
			)
			if res, err = strconv.Atoi(val.(string)); err != nil {
				panic(err)
			}
			return res
		default:
			return 0
		}
	}

}

func (data *DataBase) beforeSet() {

	if data.vals == nil {
		data.vals = make(map[string]interface{})
	}

}

func (data *DataBase) SetString(name, val string) {
	data.beforeSet()
	data.vals[name] = val
}

func (data *DataBase) SetInt(name string, val int) {
	data.beforeSet()
	data.vals[name] = val
}

func (data *DataBase) SetInt64(name string, val int64) {
	data.beforeSet()
	data.vals[name] = val
}

func (data *DataBase) MakeSign() string {

	var (
		params []string
		format string
		fields []string
	)

	for field := range data.vals {
		fields = append(fields, field)
	}

	sort.Strings(fields)

	for _, field := range fields {

		if field == "key" || field == "sign" {
			continue
		}
		val := data.Get(field)
		switch val.(type) {
		case string:
			if val == "" {
				continue
			}
			format = `%s=%s`
		case int8, int16, int32, int64, int:
			if val == 0 {
				continue
			}
			format = `%s=%d`
		}
		params = append(params, fmt.Sprintf(format, field, val))
	}

	params = append(params, "key="+data.GetString("key"))

	return utils.MD5(strings.Join(params, "&"), true)
	//return fmt.Sprintf("%X", md5.Sum([]byte( strings.Join(params, "&") )))
}

func (data *DataBase) ToXml() (xml string) {

	xml += `<xml>`

	for key, val := range data.vals {
		if key == "key" {
			continue
		}
		switch val.(type) {
		case string:
			xml += fmt.Sprintf(`<%s><![CDATA[%s]]></%s>`, key, val, key)
		case int:
			xml += fmt.Sprintf(`<%s><![CDATA[%d]]></%s>`, key, val, key)
		}
	}
	xml += `</xml>`

	return
}

func (data *DataBase) SetKey(val string) {
	data.SetString("key", val)
}

func (data *DataBase) GetVals() map[string]interface{} {
	return data.vals
}

func (data *DataBase) Init(vals map[string]interface{}) {

	data.vals = vals
}

func NewDataBase() *DataBase {
	return &DataBase{}
}
