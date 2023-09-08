package main

import (
	"flag"
	"log"
	"net/url"

	"github.com/mazeyqian/go-gin-gee/internal/pkg/constants"
	"github.com/szyhf/go-excel"
)

// func (this Standard) GetXLSXSheetName() string {
// 	return "Some other sheet name if need"
// }

// type Temp struct {
// 	Foo string
// }

// self define a unmarshal interface to unmarshal string.
// func (this *Temp) UnmarshalBinary(d []byte) error {
// 	return json.Unmarshal(d, this)
// }

// func simpleUsage() {
// 	// will assume the sheet name as "Standard" from the struct name.
// 	var stdList []constants.Standard
// 	err := excel.UnmarshalXLSX("./testdata/simple.xlsx", &stdList)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func defaultUsage(filePath string, sheetNamer interface{}) {
	conn := excel.NewConnecter()
	err := conn.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Generate an new reader of a sheet
	// sheetNamer: if sheetNamer is string, will use sheet as sheet name.
	//             if sheetNamer is int, will i'th sheet in the workbook, be careful the hidden sheet is counted. i ∈ [1,+inf]
	//             if sheetNamer is a object implements `GetXLSXSheetName()string`, the return value will be used.
	//             otherwise, will use sheetNamer as struct and reflect for it's name.
	// 			   if sheetNamer is a slice, the type of element will be used to infer like before.
	rd, err := conn.NewReader(sheetNamer)
	if err != nil {
		panic(err)
	}
	defer rd.Close()

	exampleNameMap := make(map[string]int)
	exampleIdMap := map[string]int{}

	for rd.Next() {
		var s constants.Standard
		// Read a row into a struct.
		err := rd.Read(&s)
		if err != nil {
			panic(err)
		}
		urlStr := s.URL
		uu, err := url.Parse(urlStr)
		if err != nil {
			panic(err)
		}
		exampleName := uu.Query().Get("exampleName")
		if exampleName != "" {
			exampleNameMap[exampleName]++
			log.Printf("exampleName: %s", exampleName)
		}
		exampleId := uu.Query().Get("exampleId")
		if exampleId != "" {
			exampleIdMap[exampleId]++
			log.Printf("exampleId: %s", exampleId)
		}
	}
	log.Printf("exampleNameMap: %+v", exampleNameMap)
	for k, v := range exampleNameMap {
		log.Printf("name: %s, count: %d", k, v)
	}
	log.Printf("exampleIdMap: %+v", exampleIdMap)
	for k, v := range exampleIdMap {
		log.Printf("id: %s, count: %d", k, v)
	}

	// Read all is also supported.
	// var stdList []constants.Standard
	// err = rd.ReadAll(&stdList)
	// if err != nil {
	// 	panic(err)
	// 	// return
	// }
	// log.Printf("%+v", stdList)

	// map with string key is support, too.
	// if value is not string
	// will try to unmarshal to target type
	// but will skip if unmarshal failed.

	// var stdSliceList [][]string
	// err = rd.ReadAll(&stdSliceList)
	// if err != nil {
	//   panic(err)
	//	 return
	// }
	// fmt.Printf("%+v",stdSliceList)

	// var stdMapList []map[string]string
	// err = rd.ReadAll(&stdMapList)
	// if err != nil {
	//   panic(err)
	//	 return
	// }
	// fmt.Printf("%+v",stdMapList)

	// Using binary instead of file.
	// xlsxData, err := ioutil.ReadFile(filePath)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// conn := excel.NewConnecter()
	// err = conn.OpenBinary(xlsxData)
}

// Example: go run scripts/parse-excel-xlsx/main.go -file="data/example_name.xlsx" -sheet="example_name"
// https://github.com/szyhf/go-excel
func main() {
	// Define command-line flags
	filePath := flag.String("file", "", "xlsx file path")
	sheetNamer := flag.String("sheet", "", "sheet name")
	flag.Parse()
	defaultUsage(*filePath, *sheetNamer)
}
