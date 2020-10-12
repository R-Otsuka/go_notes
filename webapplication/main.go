package main
import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Page struct{
	Title string
	Body []byte
}

func (p *Page)save()error{
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body,0600)
}

func loadPage(title string) (*Page, error){
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil{
		return nil, err
	}
	return &Page{Title:title,Body:body},nil
}

var templates = template.Must(template.ParseFiles("edit.html","view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page){
	//t,_ := template.ParseFiles(tmpl + ".html")
	////htmlページに対してpを埋め込んでwに渡すようなイメージ
	//t.Execute(w,p)
	//キャッシングしたファイルにw,ファイル名,pを渡す
	err := templates.ExecuteTemplate(w, tmpl+".html",p)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter,r *http.Request, title string){
	p, err := loadPage(title)
	if err != nil{
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	//fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>",p.Title,p.Body)
	renderTemplate(w,"view",p)
}

func editHandler(w http.ResponseWriter,r *http.Request, title string){
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w,"edit",p)
}

//存在しない場合はtextの新規作成。存在する場合は上書き
func saveHandler(w http.ResponseWriter,r *http.Request, title string){
	body := r.FormValue("text")
	p := &Page{Title:title, Body:[]byte(body)}
	err := p.save()
	if err !=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc{
	return func(w http.ResponseWriter, r*http.Request){
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil{
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}
func main(){
	http.HandleFunc("/view/",makeHandler(viewHandler))
	http.HandleFunc("/edit/",makeHandler(editHandler))
	http.HandleFunc("/save/",makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
