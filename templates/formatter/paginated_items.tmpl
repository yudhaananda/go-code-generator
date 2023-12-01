package formatter

type PaginatedItems struct {
	Data      interface{} `json:"data"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
	DataCount int         `json:"dataCount"`
	PageCount int         `json:"pageCount"`
}

func (f *PaginatedItems) Format(pageIndex int, pageSize, count float64, take float64, data interface{}) {
	f.Data = data
	f.PageIndex = pageIndex
	f.DataCount = int(count)
	f.PageSize = int(pageSize)
	if pageSize > 0 {
		temp := count / take
		if int(count)%int(take) != 0 {
			temp++
		}
		f.PageCount = int(temp)
	}
	if take < 0 {
		f.PageCount = 1
	}
}
