package gormgen

import "text/template"

func parseTemplateOrPanic(t string) *template.Template {
	tpl, err := template.New("output_template").Parse(t)
	if err != nil {
		panic(err)
	}
	return tpl
}

var outputTemplate = parseTemplateOrPanic(`
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

	// Add add one record
	func (t *{{.StructName}}) Add(db *gorm.DB)(err error) {
		if err = db.Create(t).Error;err!=nil{
			{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
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
		if err = db.Where("id = ?",t.ID).Updates(m).Error;err!=nil{
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
	{{range .Fields}}
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
`)
