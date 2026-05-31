package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

/*
SelectAllClients() []Client
SelectAllTasks() []Tasks
*/

func SelectAllClients() []Client {
	db, err := sql.Open("sqlite", "lesson22.db")
	if err != nil {
		log.Println(err)
		return []Client{}
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("ping db:", err)
	}

	rows, err := db.Query("SELECT * FROM clients")
	if err != nil {
		log.Println(err)
		return []Client{}
	}

	defer rows.Close()
	clients := []Client{}

	for rows.Next() {
		ph := Client{}
		err := rows.Scan(&ph.Id, &ph.Name, &ph.Created, &ph.WasOnline)
		if err != nil {
			log.Println(err)
			continue
		}
		clients = append(clients, ph)
	}
	return clients
}

func SelectAllTasks() []Task {
	db, err := sql.Open("sqlite", "lesson22.db")
	if err != nil {
		log.Println(err)
		return []Task{}
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("ping db:", err)
	}

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Println(err)
		return []Task{}
	}

	defer rows.Close()
	tasks := []Task{}

	for rows.Next() {
		ph := Task{}
		err := rows.Scan(&ph.Id, &ph.ClientId, &ph.Text, &ph.Created, &ph.Answer, &ph.Done)
		if err != nil {
			log.Println(err)
			continue
		}
		tasks = append(tasks, ph)
	}
	return tasks
}
