package gormgen

import (
	"fmt"
	"text/template"
)

func parseTemplateOrPanic(t string) *template.Template {
	tpl, err := template.New("output_template").Parse(t)
	if err != nil {
		panic(err)
	}
	return tpl
}

var commonTemplate = parseTemplateOrPanic(fmt.Sprintf(`
package {{.PkgName}}
type FieldData struct {
		Value interface{} %sjson:"value" form:"value"%s
		Symbol string %sjson:"symbol" form:"symbol"%s
	
}
`, "`", "`", "`", "`"))

var outputTemplate = parseTemplateOrPanic(fmt.Sprintf(`
package {{.PkgName}}
{{$TransformErr :=.TransformErr}}
import (
{{if $TransformErr}} "errors" {{end}}
	{{range .ImportPkgs}}
	"{{.Pkg}}"
	{{end}}
)


{{$LogName := .LogName}}
	
 {{if $TransformErr}} var(
		 ErrCreate{{.StructName}} = errors.New("create {{.StructName}} failed") 
		 ErrDelete{{.StructName}} = errors.New("delete {{.StructName}} failed") 
		 ErrGet{{.StructName}} = errors.New("get {{.StructName}} failed") 
		 ErrUpdate{{.StructName}} = errors.New("update {{.StructName}} failed") 
	)
{{end}}

	// New{{.StructName}} new
	func New{{.StructName}}()*{{.StructName}}{
		return new({{.StructName}})
	}
	// Add add one record
	func (t *{{.StructName}}) Add(db *gorm.DB)(err error) {
		if err = db.Create(t).Error;err!=nil{
			{{if $LogName}} {{ $LogName}}.Errorln(err){{end}}
			{{if $TransformErr}} err = ErrCreate{{.StructName}}{{end}}
			return
		}
		return 
	}

	// Delete delete record
	func (t *{{.StructName}}) Delete(db *gorm.DB)(err error) {
		if err =  db.Delete(t).Error;err!=nil{
			{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
			{{if $TransformErr}} err = ErrDelete{{.StructName}} {{end}}
			return
		}
		return
	}
	
	// Updates update record
	func (t *{{.StructName}}) Updates(db *gorm.DB, m map[string]interface{})(err error) {
		if err = db.Model(&{{.StructName}}{}).Where("id = ?",t.ID).Updates(m).Error;err!=nil{
			{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
			{{if $TransformErr}} err = ErrUpdate{{.StructName}} {{end}}
			return
		}
		return 
	}

	// Get{{.StructName}}All get all record
	func Get{{.StructName}}All(db *gorm.DB)(ret []*{{.StructName}},err error){
		if err = db.Find(&ret).Error;err!=nil{
			{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
			{{if $TransformErr}} err = ErrGet{{.StructName}} {{end}}
			return
		}
		return
	}
	
	// Get{{.StructName}}Count get count
	func Get{{.StructName}}Count(db *gorm.DB)(ret int64){
		db.Model(&{{.StructName}}{}).Count(&ret)
		return
	}

	{{$StructName := .StructName}}
	type Query{{$StructName}}Form struct{
	{{range .OptionFields}} {{.FieldName}} *FieldData %sjson:"{{.HumpName}}" form:"{{.HumpName}}"%s; {{end}}
		Order []string %sjson:"order" form:"order"%s
		PageNum int %sjson:"pageNum" form:"pageNum"%s
		PageSize int %sjson:"pageSize" form:"pageSize"%s
		}

	//  Get{{$StructName}}List get {{$StructName}} list some field value or some condition
	func Get{{$StructName}}List(q *Query{{$StructName}}Form, db *gorm.DB)(ret []*{{$StructName}},err error){
		// order
		if len(q.Order)>0{
			for _,v:=range q.Order {
				db = db.Order(v)
			}
		}
		// pageSize
		if q.PageSize!=0{
			db = db.Limit(q.PageSize)
		}
		// pageNum
		if q.PageNum!=0{
			q.PageNum = (q.PageNum - 1) * q.PageSize
			db = db.Offset(q.PageNum)
		}
	{{range .OptionFields}} 
		// {{.FieldName}}
		if q.{{.FieldName}}!=nil{
			db = db.Where("{{.ColumnName}}" +q.{{.FieldName}}.Symbol +"?",q.{{.FieldName}}.Value)
		}  ; {{end}}
		if err = db.Find(&ret).Error;err!=nil{
			return	
		}
		return 
	}
	{{range .OnlyFields}}
		// QueryBy{{.FieldName}} query cond by {{.FieldName}}
		func (t *{{$StructName}}) SetQueryBy{{.FieldName}}({{.ColumnName}} {{.FieldType}})*{{$StructName}} {
			t.{{.FieldName}} = {{.ColumnName}}
			return  t
		}
		// GetBy{{.FieldName}} get one record by {{.FieldName}}
		func (t *{{$StructName}})GetBy{{.FieldName}}(db *gorm.DB)(err error){
			if err = db.First(t,"{{.ColumnName}} = ?",t.{{.FieldName}}).Error;err!=nil{
				{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
				{{if $TransformErr}} err = ErrGet{{$StructName}} {{end}}
				return
			}
			return
		}
		// DeleteBy{{.FieldName}} delete record by {{.FieldName}}
		func (t *{{$StructName}}) DeleteBy{{.FieldName}}(db *gorm.DB)(err error) {
			if err= db.Delete(t,"{{.ColumnName}} = ?",t.{{.FieldName}}).Error;err!=nil{
				{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
				{{if $TransformErr}} err = ErrDelete{{$StructName}} {{end}}
				return
				}
			return
		}
	{{end}}
`, "`", "`", "`", "`", "`", "`", "`", "`"))
