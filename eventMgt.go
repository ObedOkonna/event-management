package main

import (
	"errors"
	"fmt"
	"time"
)

// Participant struct holds participant details
type Participant struct {
	Name  string
	Email string
}

// Event struct represents an event with event details and participants
type Event struct {
	EventID      int
	EventName    string
	Date         time.Time
	Participants []Participant
}

// Global slices to store events and participants
var events []Event

// Function to add an event to the event list
func addEvent(eventID int, eventName string, date time.Time) error {
	for _, e := range events {
		if e.EventID == eventID {
			return errors.New("event ID already exists")
		}
	}
	newEvent := Event{
		EventID:   eventID,
		EventName: eventName,
		Date:      date,
	}
	events = append(events, newEvent)
	return nil
}

// Function to register a participant for an event
func registerParticipant(eventID int, participant Participant) error {
	for i, e := range events {
		if e.EventID == eventID {
			if len(e.Participants) >= 10 {
				return errors.New("event is already full")
			}
			events[i].Participants = append(events[i].Participants, participant)
			return nil
		}
	}
	return errors.New("event not found")
}

// Function to display event information
func displayEventInfo(eventID int) {
	for _, e := range events {
		if e.EventID == eventID {
			fmt.Printf("Event ID: %d\nEvent Name: %s\nDate: %s\nParticipants:\n", e.EventID, e.EventName, e.Date.Format("2006-01-02"))
			for _, p := range e.Participants {
				fmt.Printf("- Name: %s, Email: %s\n", p.Name, p.Email)
			}
			return
		}
	}
	fmt.Println("Event not found")

}

func main() {
	var choice int
	for {
		fmt.Println("Hello! You are welcome to our Event Management System")
		fmt.Println("1. Add Event")
		fmt.Println("2. Register Participant")
		fmt.Println("3. Display Event Information")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var eventID int
			var eventName string
			var dateStr string

			fmt.Print("Enter Event ID: ")
			fmt.Scanln(&eventID)

			fmt.Print("Enter Event Name: ")
			fmt.Scanln(&eventName)

			fmt.Print("Enter Date (YYYY-MM-DD): ")
			fmt.Scanln(&dateStr)

			date, err := time.Parse("2006-01-02", dateStr)
			if err != nil {
				fmt.Println("Invalid date format. Please enter date in YYYY-MM-DD format.")
				break
			}

			err = addEvent(eventID, eventName, date)
			if err != nil {
				fmt.Println("Error adding event:", err)
			} else {
				fmt.Println("Event added successfully!")
			}
		case 2:
			var eventID int
			var participantName string
			var participantEmail string

			fmt.Print("Enter Event ID to register participant: ")
			fmt.Scanln(&eventID)

			fmt.Print("Enter Participant Name: ")
			fmt.Scanln(&participantName)

			fmt.Print("Enter Participant Email: ")
			fmt.Scanln(&participantEmail)

			participant := Participant{
				Name:  participantName,
				Email: participantEmail,
			}

			err := registerParticipant(eventID, participant)
			if err != nil {
				fmt.Println("Error registering participant:", err)
			} else {
				fmt.Println("Participant registered successfully!")
			}
		case 3:
			var eventID int
			fmt.Print("Enter Event ID to display information: ")
			fmt.Scanln(&eventID)

			displayEventInfo(eventID)
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
	}
}
