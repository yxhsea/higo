package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func main() {
	dir := "./report"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("ioutil.ReadDir err=%v\n", err)
		return
	}

	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())
		rows, err := getFileContent(filePath)
		if err != nil {
			fmt.Printf("getFileContent err=%v\n", err)
			continue
		}
		fmt.Printf("rows=>%v\n", rows[0][0])
	}
}

func getFileContent(filePath string) ([][]string, error) {
	var csvReader *csv.Reader
	byteBody, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("ioutil.ReadFile err=%v\n", err)
		return nil, err
	}
	reader := bytes.NewReader(byteBody)

	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest(byteBody)
	if err != nil {
		fmt.Printf("detector.DetectBest err=%v\n", err)
		return nil, err
	}
	charset := result.Charset
	switch charset {
	case "GB-18030":
		csvReader = csv.NewReader(transform.NewReader(reader, simplifiedchinese.GBK.NewDecoder()))
	case "ISO-8859-1":
		csvReader = csv.NewReader(transform.NewReader(reader, simplifiedchinese.GBK.NewDecoder()))
	case "windows-1252":
		csvReader = csv.NewReader(transform.NewReader(reader, simplifiedchinese.GBK.NewDecoder()))
	case "Shift_JIS":
		csvReader = csv.NewReader(transform.NewReader(reader, japanese.ShiftJIS.NewDecoder()))
	case "EUC-KR":
		csvReader = csv.NewReader(transform.NewReader(reader, korean.EUCKR.NewDecoder()))
	case "UTF-8":
		csvReader = csv.NewReader(reader)
	case "Windows 1258":
		csvReader = csv.NewReader(charmap.Windows1258.NewDecoder().Reader(reader))
	}
	csvReader.LazyQuotes = true
	csvReader.FieldsPerRecord = -1

	rows, err := csvReader.ReadAll()
	if err != nil {
		fmt.Printf("csvReader.ReadAll err=%v\n", err)
		return nil, err
	}
	return rows, nil
}
