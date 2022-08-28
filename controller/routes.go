package controller

func PupulateSliceFromDataBase() {

	rows, err := Postgresdb.Query("SELECT * FROM gists")
	if err != nil {
		//http.Error(writer, http.StatusText(500), 500)
		log.Println(err)
		return
	}
	defer rows.Close()
	blg := H.Blog{}

	for rows.Next() {

		err := rows.Scan(&blg.SerialN, &blg.Title, &blg.Body, &blg.Date, &blg.Edit, &blg.Delete) // order matters
		if err != nil {
			//http.Error(writer, http.StatusText(500), 500)
			log.Println("rows.next err ", err)
			return
		}

		BlogsQueue.Append(blg)
	}
	if err = rows.Err(); err != nil {
		//http.Error(writer, http.StatusText(500), 500)
		fmt.Println("http StatusText(500")
		return
	}
}

func AddBlog(writer http.ResponseWriter, req *http.Request) {

	fmt.Printf("%v", req.Body)
	e := req.ParseForm() // populate r.postForm with data from form fields

	H.HandleErr(e)
	timeNow := time.Now()
	date := timeNow.Format("Mon, 02 Jan 2006 15:04:05")
	title := req.PostForm.Get("title")
	title = strings.TrimSpace(title)
	body := req.PostForm.Get("body")
	body = strings.TrimSpace(body)
	if body == "" || title == "" {
		//http.Error(writer, http.StatusText(400), http.StatusBadRequest)
		parsedTemplate, err := template.ParseFiles("./tmpl/badrequeested.html")

		if err != nil {
			H.Println(err)
		}

		parsedTemplate.Execute(writer, "BAD REQUEST GATEWAY")

	} else {
		num, _ := uuid.NewUUID()
		Uid := num.String()
		// read data from database
		_, err := Postgresdb.Exec("INSERT INTO gists VALUES ($1,$2,$3,$4)", Uid, title, body, date)
		H.HandleErr(err)

		//insert to database
		temp := H.Blog{
			SerialN: Uid,
			Title:   title,
			Body:    body,
			Date:    date,
			Edit:    "edit",
			Delete:  "delete",
		}

		BlogsQueue.Append(temp)
		parsedTemplate, err := template.ParseFiles("./tmpl/basetemplate.gohtml")

		if err != nil {
			H.Println(err)
		}

		parsedTemplate.Execute(writer, BlogsQueue.Blogs)

		//RenderPage(writer, "./tmpl/basetemplate.gohtml", Blgs1.Blogs)
		//http.Redirect(writer, req, "./tmpl/basetemplate.gohtml", http.StatusSeeOther)
		// http.Redirect(writer, req, "/blog", http.StatusPermanentRedirect)

	}

}

func DeleteBlog(writer http.ResponseWriter, req *http.Request) {

	SN := chi.URLParam(req, "SerialN")

	for _, blog := range BlogsQueue.Blogs {
		if blog.SerialN == SN {
			BlogsQueue.Delete(blog.SerialN)

		}

		_, err := Postgresdb.Exec("DELETE FROM gists WHERE serialn=$1;", SN)
		if err != nil {
			http.Error(writer, http.StatusText(500), http.StatusInternalServerError)
			return
		}

	}
	parsedTemplate, err := template.ParseFiles("./tmpl/basetemplate.gohtml")

	if err != nil {
		H.Println(err)
	}

	parsedTemplate.Execute(writer, BlogsQueue.Blogs)

}

func UpdateBlog(writer http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "SerialN")
	e := req.ParseForm() // populate r.postForm with data from form fields

	H.HandleErr(e)
	title := req.PostForm.Get("title")
	body := req.PostForm.Get("body")
	timeNow := time.Now()
	date := timeNow.Format("Mon, 02 Jan 2006 15:04:05")

	if strings.TrimSpace(body) == "" || strings.TrimSpace(title) == "" {
		parsedTemplate, err := template.ParseFiles("./tmpl/badrequeested.html")

		if err != nil {
			H.Println(err)
		}

		parsedTemplate.Execute(writer, "BAD REQUEST GATEWAY")
		//http.Error(writer, http.StatusText(400), http.StatusBadRequest)

	} else {
		_, err := Postgresdb.Exec("update gists set title = $1,body = $2, date = $3 where serialn = $4;", title, body, date, id)
		if err != nil {
			http.Error(writer, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		BlogsQueue.EditBlog(id, body, title)
		parsedTemplate, err := template.ParseFiles("./tmpl/basetemplate.gohtml")

		if err != nil {
			H.Println(err)
		}
		parsedTemplate.Execute(writer, BlogsQueue.Blogs)
	}
}

func Capitalize(writer http.ResponseWriter, req *http.Request) {
	SN := chi.URLParam(req, "SerialN")

	for ind, blog := range BlogsQueue.Blogs {
		if blog.SerialN == SN {

			strH := blog.Title

			strH2 := strings.ToUpper(strH)

			BlogsQueue.Blogs[ind].Title = strH2
			fmt.Println(strH2)
		}
	}
	http.Redirect(writer, req, "/", 302)

	//var dbs = Blogger{} // database
}

func UnCapitalize(writer http.ResponseWriter, req *http.Request) {
	SN := chi.URLParam(req, "SerialN")

	for ind, blog := range BlogsQueue.Blogs {
		if blog.SerialN == SN {
			str := blog.Body
			strH := blog.Title
			str2 := strings.ToLower(str)
			strH2 := strings.ToLower(strH)
			BlogsQueue.Blogs[ind].Body = str2
			BlogsQueue.Blogs[ind].Title = strH2
		}
	}

	RenderPage(writer, "./tmpl/basetemplate.gohtml", BlogsQueue.Blogs)

}
func ReadMore(writer http.ResponseWriter, req *http.Request) {
	SN := chi.URLParam(req, "SerialN")
	//template.parse and execute

	parsedTemplate, err := template.ParseFiles("./tmpl/readmore.gohtml")

	if err != nil {
		H.Println(err)
	}

	for _, blog := range BlogsQueue.Blogs {
		if SN == blog.SerialN {
			parsedTemplate.Execute(writer, blog)
		}

	}

}

func EditBlog(writer http.ResponseWriter, req *http.Request) {
	SN := chi.URLParam(req, "SerialN")
	for _, blog := range BlogsQueue.Blogs {
		if blog.SerialN == SN {
			//This points to the html location
			TemplateFile, err1 := template.ParseFiles("./tmpl/index.gohtml")
			H.HandleErr(err1)

			err2 := TemplateFile.Execute(writer, blog)
			H.HandleErr(err2)

		}
	}
}
