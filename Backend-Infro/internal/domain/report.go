package domain

import "io"

type ReportUsecase interface {
	GenerateActsExcel() (io.Reader, string, error)
}
