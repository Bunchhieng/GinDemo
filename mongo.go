package main

import (
	"gopkg.in/mgo.v2"
	"log"
)

type Person struct {
	Name  string
	Phone string
}

func InsertData() {
	query := func(c *mgo.Collection) (err error) {
		err = c.Insert(&Person{"Bun", "4389247982374"}, &Person{"Test", "324245"})
		if err != nil {
			log.Fatal(err)
		}
		return err
	}

	insert := func() error {
		return withCollection("test", query)
	}

	err := insert()
	if err != nil {
		log.Fatal(err)
	}
}

func SearchPerson(q interface{}, skip int, limit int) (searchResults []Person, searchErr string) {
	searchErr = ""
	searchResults = []Person{}
	query := func(c *mgo.Collection) error {
		fn := c.Find(q).Skip(skip).Limit(limit).All(&searchResults)
		if limit < 0 {
			fn = c.Find(q).Skip(skip).All(&searchResults)
		}
		return fn
	}
	search := func() error {
		return withCollection("person", query)
	}
	err := search()
	if err != nil {
		searchErr = "Database Error"
	}
	return
}
