package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(s Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.CreateAccount(acc); err != nil {
		log.Fatal(err)

	}

	fmt.Println("new account +> ", acc.Number)
	return acc

}

func seedAccounts(s Storage){
	seedAccount(s, "antony", "dadd","test")
}


func main() {
	seed:=flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}


	if *seed{
		fmt.Println("seeding database")
		seedAccounts(store)
	}
	// seed stuf
	



	server := NewAPIServer(":3000", store)
	server.Run()

}

//ders ikiden 0.35de qaldim
//create accountu elave tdim postgresqle datani gonderdim
