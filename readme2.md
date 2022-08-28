
## init Postgres
```GO
func init() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%d dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	Postgresdb, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = Postgresdb.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to Database!")
}
```

## Init Mongo
```GO
func MongoDBinstance() *mongo.Client {

	MongoDB := os.Getenv("MONGODB_URL")
	if MongoDB == "" || len(MongoDB) < 1 {
		MongoDB = "mongodb://localhost:27017/bookDb"
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB")
	return client
}
```