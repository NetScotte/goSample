package iniconfig

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// 将配置数据写入到文件中
func MarshalFile(s interface{}, filename string) (err error) {
	data, err := Marshal(s)
	if err != nil {
		err = fmt.Errorf("MarshalFile failed: %v", err)
	}
	err = ioutil.WriteFile(filename, data, 0755)
	if err != nil {
		err = fmt.Errorf("Write file failed, %v", err)
	}
	return
}

// 从文件中读取配置数据
func UnMarshalFile(filename string, s interface{}) (err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		err = fmt.Errorf("Failed read file %v: %v",filename, err)
	}
	err = UnMarshal(data, s)
	if err != nil {
		err = fmt.Errorf("Failed UnMarshal: %v", err)
	}
	return
}


// 将结构体序列化为文本字符
func Marshal(s interface{}) (data []byte, err error) {
	objType := reflect.TypeOf(s)
	objValue := reflect.ValueOf(s)
	if objType.Kind() != reflect.Struct {
		err = fmt.Errorf("not a struct")
		return
	}
	var content []string
	for i :=0; i < objType.NumField(); i++ {
		sectionType := objType.Field(i)
		sectionValue := objValue.Field(i)
		if sectionType.Type.Kind() != reflect.Struct {
			continue
		}
		sectionName := sectionType.Tag.Get("ini")
		if sectionName == "" {
			sectionName = strings.ToLower(sectionType.Name)
		}
		line := fmt.Sprintf("[%s]\n", sectionName)
		content = append(content, line)
		for j := 0; j < sectionValue.NumField(); j++ {
			optionValue := sectionValue.Field(j)
			optionType := sectionValue.Type().Field(j)
			optionName := optionType.Tag.Get("ini")
			if optionName == "" {
				optionName = strings.ToLower(optionType.Name)
			}
			line := fmt.Sprintf("%v=%v\n", optionName, optionValue.Interface())
			content = append(content, line)
		}
		content = append(content, "\n")
	}

	for _, value := range content {
		lineByte := []byte(value)
		data  = append(data, lineByte...)
	}

	return
}


// 将文本字符反序列为结构体
func UnMarshal(data []byte, s interface{}) (err error) {
	lineArray := strings.Split(string(data), "\n")
	var sectionName string
	typeInfo := reflect.TypeOf(s)
	if typeInfo.Kind() != reflect.Ptr {
		err = fmt.Errorf("please input address")
		return
	}

	typeInfo2 := typeInfo.Elem()
	if typeInfo2.Kind() != reflect.Struct {
		err = fmt.Errorf("please input a struct")
		return
	}

	for index, line := range lineArray {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if line[0] == ';' || line[0] == '#' {
			continue
		}

		if line[0] == '[' && line[len(line) -1] == ']' {
			sectionName, err = parserSection(line, typeInfo2)
			if err != nil {
				err = fmt.Errorf("lineno(%v): %v", index, err)
			}
			continue
		}

		if sectionName != "" {
			sectionStruct := reflect.ValueOf(s).Elem().FieldByName(sectionName)
			err = setOption(line, sectionStruct)
		}

	}
	return
}

// 从结构体中获取sectionName,
func parserSection(line string, typeStruct reflect.Type) (sectionName string, err error) {
	name := line[1:len(line)-1]
	for i := 0; i < typeStruct.NumField(); i++ {
		filed := typeStruct.Field(i)
		tagValue := filed.Tag.Get("ini")
		if tagValue == name {
			sectionName = filed.Name
			return
		}
	}
	return
}

// 为结构体设置值
func setOption(line string, sectionStruct reflect.Value) (err error) {
	index := strings.Index(line, "=")
	if index == -1 {
		err = fmt.Errorf("not found '=' in line")
		return
	}

	key := strings.TrimSpace(line[0:index])
	value := strings.TrimSpace(line[index+1:])

	for i := 0; i < sectionStruct.NumField(); i++ {
		valueField := sectionStruct.Field(i)
		typefiled := sectionStruct.Type().Field(i)
		tagValue := typefiled.Tag.Get("ini")
		if tagValue == key {
			//fmt.Printf("key: %v ------> structKey: %v will be set: %v \n", key, typefiled.Name, value)
			switch typefiled.Type.Kind() {
			case reflect.Int:
				intValue, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					break
				}
				valueField.SetInt(intValue)
			case reflect.Float64:
				floatValue, err := strconv.ParseFloat(value, 64)
				if err != nil {
					break
				}
				valueField.SetFloat(floatValue)
			case reflect.String:
				valueField.SetString(value)
			default:
				err = fmt.Errorf("unsupport struct type: %v", typefiled.Type.Kind())
			}
		}
	}
	return
}