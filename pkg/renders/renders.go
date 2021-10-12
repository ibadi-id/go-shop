package renders

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ibadi-id/gostart/pkg/config"
	"github.com/ibadi-id/gostart/pkg/models"
)


var functions = template.FuncMap{

}

var app *config.AppConfig



// NewTemplate fungsi untuk mengambil template cache berasal dari main.go
func NewTemplate(a *config.AppConfig){
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData{
	return td
}

// RenderTemplate fungsi untuk menampilkan data setiap halaman
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache{
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	

	t, ok := tc[tmpl] 
	if !ok {
		log.Fatal("Tidak bisa mendapatkan template dari main")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println("Gagal menampilkan template", err)
	}

	// parsedTemplate, _ := template.ParseFiles("./template/" + tmpl)
	// err := parsedTemplate.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println("Error parsig template", err)
	// 	return
	// }

}

// CreateTemplateCache fungsi untuk menggabungkan base layout dengan content pada template
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./template/*.page.html")
	if err != nil {
		fmt.Println("Tidak bisa menemukan template", err)
		return myCache, err
	}

	for _, page := range pages {

		name := filepath.Base(page)
		
		tp, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			fmt.Println("Tidak bisa menemukan layout", err)
			return myCache, err
		}

		matches, err := filepath.Glob("./template/*.layout.html")
		if err != nil {
			fmt.Println("Tidak bisa menemukan layout", err)
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err := tp.ParseGlob("./template/*.layout.html")
			if err != nil {
				fmt.Println("Tidak bisa menemukan layout", err)
				return myCache, err
			}
			myCache[name] = ts
		}

	}

	return myCache, nil
	
}