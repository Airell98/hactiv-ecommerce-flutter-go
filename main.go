package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/AirellJordan98/hacktiv_ecommerce/api"
	db "github.com/AirellJordan98/hacktiv_ecommerce/db/sqlc"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("PORT")

	fmt.Println("HOST:", os.Getenv("DB_HOST"))
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require TimeZone=Asia/Shanghai",
		host, dbPort, user, password, dbName)
	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
	// 	host, dbPort, user, password, dbName)

	// DBSource := "postgresql://sofyan:postgres@localhost:5432/hacktiv-ecommerce?sslmode=disable"
	conn, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalln("error connecting to db", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store, conn)

	err = server.Start(":" + port)

	if err != nil {
		log.Fatalln("error while starting server", err)
	}

}

// type Person struct {
// 	Name string `required:"true" json:"name"`
// 	Age  int    `required:"true"`
// }

// func CatchErr() {
// 	if r := recover(); r != nil {
// 		fmt.Println("error brow")
// 	}
// }

// func isValid(data interface{}) error {
// 	r := reflect.TypeOf(data)
// 	v := reflect.ValueOf(data)
// 	for i := 0; i < r.NumField(); i++ {
// 		t := r.Field(i)
// 		if t.Tag.Get("required") == "true" {
// 			if v.Field(i).String() == "" {
// 				tName := t.Tag.Get("json")
// 				return fmt.Errorf("field %v is required", tName)
// 			}
// 		}
// 	}

// 	return nil
// }

// func main() {
// urls := []string{"https://golang.org", "https://www.google.com"}

// var wg sync.WaitGroup

// wg.Add(len(urls))

// for _, url := range urls {
// 	go CheckAndSaveBody(url, &wg)
// 	fmt.Println(strings.Repeat("#", 20))
// }

// fmt.Println("No. of GoRoutines:", runtime.NumGoroutine())

// wg.Wait()

// const gr = 10

// var wg sync.WaitGroup

// var n int = 0

// var m sync.Mutex

// for j := 0; j < 1; j++ {
// 	wg.Add(gr * 2)
// 	func() {
// 		for i := 0; i < gr; i++ {
// 			go func() {
// 				time.Sleep(time.Second / 10)
// 				m.Lock()
// 				n++
// 				m.Unlock()
// 				wg.Done()
// 				fmt.Println("Increment:", n)
// 			}()
// 			go func() {
// 				time.Sleep(time.Second / 10)
// 				m.Lock()
// 				n--
// 				m.Unlock()
// 				wg.Done()
// 				fmt.Println("Decrement:", n)
// 			}()
// 		}
// 	}()
// 	wg.Wait()
// 	fmt.Printf("{%d},Result: %d\n", j, n)
// 	n = 0
// }

// c := make(chan int)

// for i := 0; i <= 0; i++ {
// 	go factorial(i, c)

// 	fmt.Println("Selesai:", <-c)
// }

// fmt.Println("Diantara 2 looping")

// for i := 1; i <= 1; i++ {
// 	go factorial(i, c)

// 	fmt.Println("Selesai:", <-c)
// }

// defer close(c)
// urls := []string{"https://golang.org", "https://www.google.com", "https://proandroiddev.com"}

// ch := make(chan string)

// for _, url := range urls {
// 	go CheckAndSaveBodyChan(url, ch)

// }

// for i := 0; i < len(urls); i++ {
// 	fmt.Println("Waiting response from:", urls[i], ", index", i)
// 	u := <-ch
// 	fmt.Println(u)
// 	fmt.Println("Response has been recieved from:", strings.Split(u, ":")[1], ", index", i)
// }
// }

// func factorial(n int, c chan int) {
// 	fmt.Println("Recieved n:", n)
// 	c <- n
// }
