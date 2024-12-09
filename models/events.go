package models

import (
	"Rest-API/db"
	"fmt"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

var events []Event

/*func (event Event) Save() error {

	query := `insert into events (name,description,location,datetime,user_id) values (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserId)
	defer stmt.Close()

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	fmt.Print("id:", id)

	event.ID = id
	return err
}*/

func (event *Event) Save() error {
	query := `INSERT INTO events (name, description, location, datetime, user_id) VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserId)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	id, err := result.LastInsertId()

	//to remove  start

	query2 := "select * from events where id = ?"

	row := db.DB.QueryRow(query2, id)

	err = row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		return err
	}

	fmt.Println("event here: ", event)

	//to remove  stop

	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}

	event.ID = id
	fmt.Printf("Inserted event with ID: %d\n", id)
	return nil
}

/*func GetAllEvents() ([]Event, error) {

	query := "select * from events"
	rows, err := db.DB.Query(query)

	if err != nil {
		fmt.Print("GetAllEvents", err)
		return nil, err
	}

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

		if err != nil {
			fmt.Print("GetAllEvents !", err)
			return nil, err
		}
	}

	return events, nil
}*/

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Print("GetAllEvents Error: ", err)
		return nil, err
	}
	defer rows.Close() // Ensure rows are properly closed after processing

	// Create a local slice to hold events
	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			fmt.Print("GetAllEvents Scan Error: ", err)
			return nil, err
		}
		// Append the event to the slice
		events = append(events, event)
	}

	// Check for errors after the iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "select * from events where id = ?"

	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `update events set name = ?, description = ?, location = ?, datetime = ?, user_id = ? where id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserId, event.ID)

	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}

func (event Event) Delete() error {
	query := "delete from events where id = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)

	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}

func (e Event) Register(userId int64) error {
	query := "insert into registrations(event_id,user_id) values (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := "delete from registrations where event_id = ? and user_id = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}
