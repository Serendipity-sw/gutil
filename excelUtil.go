// excel操作类
// create by gloomy 2017-4-18 09:14:26
package gutil

import (
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
)

// excel数据获取
// create by gloomy 2017-4-18 12:00:34
// sheet名称 数据内容 错误对象
func ReadExcel(excelFilePath string) (*map[string][]string, error) {
	xlFile, err := xlsx.OpenFile(excelFilePath)
	if err != nil {
		return nil, err
	}
	var (
		excelData   map[string][]string
		columnValue string
	)
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				columnValue, err = cell.String()
				if err != nil {
					fmt.Println("ReadExcel ", err.Error())
					continue
				}
				excelData[sheet.Name] = append(excelData[sheet.Name], columnValue)
			}
		}
	}
	return &excelData, nil
}

// excel保存
// create by gloomy 2017-4-18 09:44:39
func ExcelSave(saveContent *map[string][]string, saveFilePath string) error {
	var (
		file  *xlsx.File
		sheet *xlsx.Sheet
		row   *xlsx.Row
		cell  *xlsx.Cell
		err   error
	)
	if len(*saveContent) == 0 {
		return errors.New("send data length is 0!")
	}
	file = xlsx.NewFile()
	for sheetName, values := range *saveContent {
		sheet, err = file.AddSheet(sheetName)
		if err != nil {
			return err
		}
		if len(values) != 0 {
			row = sheet.AddRow()
			for _, columnValue := range values {
				cell = row.AddCell()
				cell.Value = columnValue
			}
		}
	}
	return file.Save(saveFilePath)
}
