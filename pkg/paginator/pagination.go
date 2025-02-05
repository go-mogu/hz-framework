package paginator

import (
	"github.com/go-mogu/hz-framework/global"
	"github.com/go-mogu/hz-framework/pkg/util"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"math"
)

type PageBuilder[T any] struct {
	DB       *gorm.DB
	Model    interface{} // model struct
	Preloads []string    // 预加载
	Fields   []string    // 查询字段
}

type OnJoins struct {
	LeftTableField, RightTableField JoinTableField // LeftTableField：如：主表.ID  RightTableField：如：关联表.主表ID
}

type JoinTableField struct {
	Table, Field string
}

type SelectTableField struct {
	Model interface{}
	Table string
	Field []string
}

type Page[T any] struct {
	List        []T   `json:"list"`         // 查询的列表
	CurrentPage int   `json:"current_page"` // 当前页
	Total       int64 `json:"total"`        // 查询记录总数
	LastPage    int   `json:"last_page"`    // 最后一页
	PerPage     int   `json:"per_page"`     // 每页条数
}

func NewBuilder[T any]() *PageBuilder[T] {
	return &PageBuilder[T]{}
}

// WithDB db连接
func (pb *PageBuilder[T]) WithDB(db *gorm.DB) *PageBuilder[T] {
	pb.DB = db
	return pb
}

// NewDB 对接原生查询方式
func (pb *PageBuilder[T]) NewDB() *gorm.DB {
	return pb.DB
}

// WithField 查询单表的字段 和 过滤字段 不能与WithFields方法同用
func (pb *PageBuilder[T]) WithField(fields []string) *PageBuilder[T] {
	fieldList := filterFields(pb.Model, fields)
	pb.Fields = fieldList
	pb.DB.Select(pb.Fields)
	return pb
}

// WithFields 单多表字段查询字段（或过滤某些字段不查询 最后一个参数默认为select（不传或者传），如传omit为过滤前面传输的字段）
func (pb *PageBuilder[T]) WithFields(model interface{}, table string, fields []string) *PageBuilder[T] {
	fieldList := filterFields(model, fields)
	for i, _field := range fieldList {
		fieldList[i] = table + "." + _field
	}
	pb.Fields = append(pb.Fields, fieldList...)
	pb.DB.Select(pb.Fields)
	return pb
}

// filterFields 过滤查询字段
func filterFields(model interface{}, fields []string) []string {
	var fieldList []string
	if len(fields) == 1 {
		if fields[0] != "_omit" && fields[0] != "_select" {
			fieldList = fields
		}
	} else {
		switch fields[len(fields)-1] {
		case "_omit":
			fields = fields[:len(fields)-1]
			_fields, _ := util.GetStructColumnName(model, 1)
			fieldList, _ = lo.Difference[string](_fields, fields)
		case "_select":
			fieldList = fields[:len(fields)-1]
		default:
			fieldList = fields[:]
		}
	}
	return fieldList
}

// WithMultiFields 多表多字段查询
func (pb *PageBuilder[T]) WithMultiFields(fields []SelectTableField) *PageBuilder[T] {
	for _, field := range fields {
		pb.WithFields(field.Model, field.Table, field.Field)
	}
	return pb
}

// WithModel 查询的model struct
func (pb *PageBuilder[T]) WithModel(model interface{}) *PageBuilder[T] {
	pb.Model = model
	pb.DB = pb.DB.Model(&model)
	return pb
}

// WithOrderBy 排序
func (pb *PageBuilder[T]) WithOrderBy(orderBy interface{}) *PageBuilder[T] {
	pb.DB = pb.DB.Order(orderBy)
	return pb
}

// WithJoins join查询
func (pb *PageBuilder[T]) WithJoins(joinType string, joinFields []OnJoins) *PageBuilder[T] {
	var joins string
	for _, field := range joinFields {
		joins += " " + joinType + " JOIN " + field.RightTableField.Table
		joins += " ON " + field.LeftTableField.Table + "." + field.LeftTableField.Field + " = "
		joins += field.RightTableField.Table + "." + field.RightTableField.Field
	}
	pb.DB.Joins(joins)
	return pb
}

// WithPreloads 多表关联查询主动预加载（无条件）
func (pb *PageBuilder[T]) WithPreloads(queries []string) *PageBuilder[T] {
	pb.Preloads = queries
	return pb
}

// WithPreload 关联查询主动预加载（可传条件）
func (pb *PageBuilder[T]) WithPreload(query string, args ...interface{}) *PageBuilder[T] {
	pb.DB.Preload(query, args...)
	return pb
}

// WithCondition 查询条件
func (pb *PageBuilder[T]) WithCondition(query interface{}, args ...interface{}) *PageBuilder[T] {
	pb.DB.Where(query, args...)
	return pb
}

// Pagination 分页查询
func (pb *PageBuilder[T]) Pagination(currentPage, pageSize int) (Page[T], error) {
	query := pb.DB
	page := pb.ParsePage(currentPage, pageSize)
	offset := (page.CurrentPage - 1) * page.PerPage
	var dst = make([]T, 0)
	// 查询总数
	if err := query.Count(&page.Total).Error; err != nil {
		return page, err
	}
	// 预加载
	if len(pb.Preloads) > 0 {
		for _, preload := range pb.Preloads {
			query.Preload(preload)
		}
	}
	// 计算总页数
	if page.Total > int64(page.PerPage) {
		page.LastPage = int(math.Ceil(float64(page.Total) / float64(page.PerPage)))
	}
	// 判断总数跟最后一页的关系
	if page.CurrentPage <= page.LastPage {
		if err := query.Limit(page.PerPage).Offset(offset).Find(&dst).Error; err != nil {
			return page, err
		}
	}
	page.List = dst
	return page, nil
}

// ParsePage 分页超限设置和格式化
func (pb *PageBuilder[T]) ParsePage(currentPage, pageSize int) Page[T] {
	var page Page[T]
	// 返回每页数量
	page.PerPage = pageSize
	// 返回当前页码
	page.CurrentPage = currentPage

	if currentPage < 1 {
		page.CurrentPage = 1
	}
	if pageSize < 1 {
		page.PerPage = global.Cfg.Server.DefaultPageSize
	}
	if pageSize > global.Cfg.Server.MaxPageSize {
		page.PerPage = global.Cfg.Server.MaxPageSize
	}
	if page.LastPage < 1 {
		page.LastPage = 1
	}
	return page
}
