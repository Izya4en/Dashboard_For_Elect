package usecase

import (
	"bytes"
	"fmt"
	"io"
	"techdash/internal/domain"

	"github.com/xuri/excelize/v2"
)

type reportUsecase struct {
	actRepo domain.ActRepository
}

func NewReportUsecase(ar domain.ActRepository) domain.ReportUsecase {
	return &reportUsecase{actRepo: ar}
}

func (u *reportUsecase) GenerateActsExcel() (io.Reader, string, error) {
	// 1. Получаем данные из базы
	acts, err := u.actRepo.GetAll()
	if err != nil {
		return nil, "", err
	}

	f := excelize.NewFile()
	sheet := "Отчет по испытаниям"
	f.SetSheetName("Sheet1", sheet)

	// 2. Создаем стили (Жирный шрифт для шапки)
	style, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true},
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#E0E0E0"}, Pattern: 1},
	})

	// 3. Заголовки
	headers := []string{"ID", "Дата", "Организация", "БИН", "Вид услуги"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
		f.SetCellStyle(sheet, cell, cell, style)
	}

	// 4. Наполнение данными
	for i, act := range acts {
		row := i + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), act.ID)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), act.TestDate)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), act.OrganizationName)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), act.OrganizationBIN)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), act.ServiceName)
	}

	// Автоматическая ширина колонок
	f.SetColWidth(sheet, "A", "A", 5)
	f.SetColWidth(sheet, "B", "E", 30)

	// 5. Записываем файл в буфер (в память), чтобы не создавать временных файлов на диске
	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, "", err
	}

	fileName := "reports_export.xlsx"
	return &buf, fileName, nil
}
