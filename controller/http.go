package controller


type PortID = string
	const PORT_NUMBER PortID = ":10000" // port number

	chiRouter := chi.NewRouter() // initiate chi router

	chiRouter.Get("/", BlogPage)
	chiRouter.Get("/view", Home)
	chiRouter.Get("/delete/{SerialN}", DeleteBlog)
	chiRouter.Get("/blog", ViewBlog)
	chiRouter.Post("/blogs", AddBlog)
	chiRouter.Post("/update/{SerialN}", UpdateBlog)
	chiRouter.Get("/edit/{SerialN}", EditBlog)
	chiRouter.Get("/cap/{SerialN}", Capitalize)
	chiRouter.Get("/uncap/{SerialN}", UnCapitalize)
	chiRouter.Get("/readmore/{SerialN}", ReadMore)
	chiRouter.Get("/home/{SerialN}", ViewBlog)

	fmt.Printf("Server Running at Port %v \n %v\n", PORT_NUMBER, time.Now())
	//Listening
	PupulateSliceFromDataBase()
	e := http.ListenAndServe(PORT_NUMBER, chiRouter) // listening to requests from connections
	H.HandleErr(e)
