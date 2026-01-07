package pagination

import (
	"math"
	"news-web/internal/core/domain/entity"
)

const DefaultPerPage = 10

type PaginationInterface interface {
	AddPagination(totalData, page, perPage int) (*entity.Page, error)
}

type PaginationService struct{}

// AddPagination implements PaginationInterface.
func (s *PaginationService) AddPagination(totalData, page, perPage int) (*entity.Page, error) {
	if page < 1 {
		return nil, ErrorPage
	}

	limitData := DefaultPerPage
	if perPage > 0 {
		limitData = perPage
	}

	// Kalau tidak ada data sama sekali
	if totalData == 0 {
		// Hanya halaman 1 yang valid
		if page != 1 {
			return nil, ErrorMaxPage
		}

		zeroPage := &entity.Page{
			Page:       page,
			Perpage:    limitData,
			PageCount:  1,
			TotalCount: 0,
			First:      0,
			Last:       0,
		}
		return zeroPage, nil
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limitData)))

	if page > totalPage {
		return nil, ErrorMaxPage
	}

	last := page * limitData
	if last > totalData {
		last = totalData
	}

	first := last - limitData
	if first < 0 {
		first = 0
	}

	pages := &entity.Page{
		Page:       page,
		Perpage:    limitData,
		PageCount:  totalPage,
		TotalCount: totalData,
		First:      first,
		Last:       last,
	}

	return pages, nil
}

func NewPagination() PaginationInterface {
	return &PaginationService{}
}
