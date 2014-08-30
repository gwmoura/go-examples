package main

import (
  "flag"
  "html/template"
  "io/ioutil"
  "log"
  "net"
  "net/http"
  "regexp"
)

type Page struct{
  Title string
  Body  []byte
}
var funcMap = template.FuncMap{
        "getLinks": getLinks,
}
var (addr = flag.Bool("addr",false,"find open address and print to final-port.txt"))
var templates = template.Must(template.New("").Funcs(funcMap).ParseFiles("tmpl/index.html","tmpl/edit.html", "tmpl/view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func (p *Page) save() error {
  filename := p.Title+".txt"
  return ioutil.WriteFile("data/"+filename,p.Body,0600)
}

func loadPage(title string) (*Page,error) {
  filename := title+".txt"
  body, err := ioutil.ReadFile("data/"+filename)
  if err != nil{
    return nil,err
  }
  return &Page{Title:title,Body:body},nil
}

func indexHandler(w http.ResponseWriter, r *http.Request, title string ) {
  var body =  r.URL.Path[1:]
  
  p := &Page{Title: title, Body: []byte(body)}

  renderTemplate(w,"index",p)
}

func getLinks() string{
  links := ""
  pages, err := ioutil.ReadDir("data/")
  if err == nil{
    for i := range pages{
      page := pages[i].Name()[:4]
      links += "<a href=\"/view/"+page+"\">"+page+"</a><br/>"
    } 
  }else{
    links = "Não temos páginas!"
  }
  return links
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.save()
    if err != nil{
      http.Error(w,err.Error(),http.StatusInternalServerError)
      return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
  p, err := loadPage(title)
  if err != nil{
    http.Redirect(w,r,"/edit/"+title,http.StatusFound)
    return
  }
  renderTemplate(w,"view",p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
  p,err := loadPage(title)
  if err != nil{
    p = &Page{Title:title}
  }
  renderTemplate(w,"edit",p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
  err := templates.ExecuteTemplate(w, tmpl+".html", p)
  if err != nil{
    http.Error(w,err.Error(), http.StatusInternalServerError)
    return
  }
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            //http.NotFound(w, r)
            //return
            fn(w, r, "/")
        }else{
          fn(w, r, m[2])  
        }
        
    }
}

func main() {
    flag.Parse()
    http.HandleFunc("/", makeHandler(indexHandler))
    http.HandleFunc("/view/", makeHandler(viewHandler))
    http.HandleFunc("/edit/", makeHandler(editHandler))
    http.HandleFunc("/save/", makeHandler(saveHandler))

    if *addr {
        l, err := net.Listen("tcp", "127.0.0.1:0")
        if err != nil {
            log.Fatal(err)
        }
        err = ioutil.WriteFile("final-port.txt", []byte(l.Addr().String()), 0644)
        if err != nil {
            log.Fatal(err)
        }
        s := &http.Server{}
        s.Serve(l)
        return
    }

    http.ListenAndServe(":8080", nil)
}

