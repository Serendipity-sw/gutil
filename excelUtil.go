// excel操作类
// create by gloomy 2017-4-18 09:14:26
package gutil

import (
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
func ExcelSave(sheetName string, columnName *[]string, saveContent *[][]string, saveFilePath string) error {
	var (
		file  *xlsx.File
		sheet *xlsx.Sheet
		row   *xlsx.Row
		cell  *xlsx.Cell
		err   error
	)
	file = xlsx.NewFile()
	sheet, err = file.AddSheet(sheetName)
	if err != nil {
		return err
	}
	if columnName != nil && len(*columnName) != 0 {
		row = sheet.AddRow()
		for _, value := range *columnName {
			cell = row.AddCell()
			cell.Value = value
		}
	}
	if saveContent != nil && len(*saveContent) != 0 {
		for _, rows := range *saveContent {
			if len(rows) == 0 {
				continue
			}
			row = sheet.AddRow()
			for _, value := range rows {
				cell = row.AddCell()
				cell.Value = value
			}
		}
	}
	return file.Save(saveFilePath)
}
