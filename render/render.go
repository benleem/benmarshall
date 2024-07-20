package render

import (
	"fmt"
	"io"
	"path/filepath"
	"text/template"
)

type TemplateMap map[string]*template.Template

type Templates struct {
	Tmpls TemplateMap
}

func (t *Templates) Render(w io.Writer, name string, data interface{}) error {
	tmpl, ok := t.Tmpls[name]
	if !ok {
		return fmt.Errorf("The template %s does not exist.", name)
	}
	return tmpl.ExecuteTemplate(w, "layout.html", data)
	// return tmpl.ExecuteTemplate(w, "layout.html", data)
}

func NewTemplateRender(basePath string) (*Templates, error) {
	tmpls, err := parseTemplates(basePath, nil)
	if err != nil {
		return nil, err
	}
	return &Templates{tmpls}, nil
}

func parseTemplates(basePath string, funcMap template.FuncMap) (map[string]*template.Template, error) {
	pagesPath := filepath.Join(basePath, "pages")
	compPath := filepath.Join(basePath, "components")
	// Base layout
	layoutPath := filepath.Join(basePath, "layout.html")
	// Pages
	homePath := filepath.Join(pagesPath, "index.html")
	workPath := filepath.Join(pagesPath, "work.html")
	// Components
	headerPath := filepath.Join(compPath, "header.html")
	footerPath := filepath.Join(compPath, "footer.html")

	layoutTmpl := template.New("layout.html").Funcs(funcMap)
	template.Must(layoutTmpl.ParseFiles(layoutPath, headerPath, footerPath))
	tmplMap := TemplateMap{
		"home": template.Must(layoutTmpl.Clone()),
		"work": template.Must(layoutTmpl.Clone()),
	}
	template.Must(tmplMap["home"].ParseFiles(homePath))
	template.Must(tmplMap["work"].ParseFiles(workPath))

	// cleanRoot := filepath.Clean(basePath)
	// pfx := len(cleanRoot) + 1
	// // fmt.Println("TEMPLATE INFO:")
	// err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, err error) error {
	// 	if !info.IsDir() && strings.HasSuffix(path, ".html") {
	// 		// fmt.Println(path)
	// 		if err != nil {
	// 			return err
	// 		}

	// 		fileBytes, err := os.ReadFile(path)
	// 		if err != nil {
	// 			return err
	// 		}

	// 		name := path[pfx:]
	// 		tmpl := template.Must(template.New(name).Funcs(funcMap).Parse(string(fileBytes)))
	// 		if strings.Contains(name, "/") {
	// 			lastIndex := strings.LastIndex(name, "/") + 1
	// 			name = name[lastIndex:]
	// 		}
	// 		// fmt.Println(tmpl.DefinedTemplates())
	// 		tmplMap[name] = tmpl
	// 	}
	// 	return nil
	// })

	return tmplMap, nil
}
