package main

import (
	"errors"
	"fmt"
	"reflect"
)

type user struct {
	Id       string
	Name     string
	Gender   string
	Age      int
	State    string
	District string
}

func (u user) addUser(users []user) []user {
	users = append(users, u)
	return users
}

type vaccinationCenter struct {
	Id       string
	State    string
	District string
}

func (v vaccinationCenter) addVC(vaccinationCentres []vaccinationCenter) []vaccinationCenter {
	vaccinationCentres = append(vaccinationCentres, v)
	return vaccinationCentres
}

type capacity struct {
	VcId     string
	Day      int
	Capacity int
}

func addCapacity(VcId string, Day, Capacity int, capacities []capacity) []capacity {
	capacities = append(capacities, capacity{VcId: VcId, Day: Day, Capacity: Capacity})
	return capacities
}

func listVcs(district string, capacities []capacity, vcs []vaccinationCenter) {
	for i := 0; i < len(vcs); i++ {
		if vcs[i].District == district {
			fmt.Printf("Vaccination Centre Id\tDay\tCapacity\n")
			for j := 0; j < len(capacities); j++ {
				if capacities[j].VcId == vcs[i].Id {
					fmt.Printf("%s\t%d\t%d\n", vcs[i].Id, capacities[j].Day, capacities[j].Capacity)
				}
			}
		}
	}
}

type Booking struct {
	Id     string
	VcId   string
	Day    int
	UserId string
}

func addBooking(bookingId, vcId string, day int, userId string, bookings []Booking, capacities []capacity, users []user) ([]Booking, []capacity, error) {
	var availability int
	for i := 0; i < len(capacities); i++ {
		if vcId == capacities[i].VcId {
			availability = capacities[i].Capacity
			break
		}
	}

	var age int
	for i := 0; i < len(users); i++ {
		if userId == users[i].Id {
			age = users[i].Age
			break
		}
	}

	if availability > 0 && age >= 18 {
		bookings = append(bookings, Booking{Id: bookingId, VcId: vcId, Day: day, UserId: userId})
		for i := 0; i < len(capacities); i++ {
			if vcId == capacities[i].VcId {
				capacities[i].Capacity--
				break
			}
		}
	} else {
		return bookings, capacities, errors.New("booking criteria not satisfied")
	}

	return bookings, capacities, nil
}

func cancelBooking(bookingId, vcId, userId string, bookings []Booking, capacities []capacity) ([]Booking, []capacity, error) {
	var booking Booking
	for i := 0; i < len(bookings); i++ {
		if bookingId == bookings[i].Id && userId == bookings[i].UserId {
			booking = bookings[i]
			bookings = append(bookings[:i], bookings[i+1:]...)
		}
	}
	if reflect.ValueOf(booking).IsZero() {
		return bookings, capacities, errors.New("no booking with the given entry")
	} else {
		for i := 0; i < len(capacities); i++ {
			if vcId == capacities[i].VcId && booking.Day == capacities[i].Day {
				capacities[i].Capacity++
				break
			}
		}
	}
	return bookings, capacities, nil
}

func searchVaccinationCentres(day int, district string, vaccinationCentres []vaccinationCenter, capacities []capacity) {
	var vcIds []string
	for i := 0; i < len(vaccinationCentres); i++ {
		if district == vaccinationCentres[i].District {
			vcIds = append(vcIds, vaccinationCentres[i].Id)
		}
	}

	for j := 0; j < len(vcIds); j++ {
		for k := 0; k < len(capacities); k++ {
			if capacities[k].VcId == vcIds[j] && day == capacities[k].Day {
				fmt.Printf("%s\t%d\n", vcIds[j], capacities[k].Capacity)
			}
		}
	}
}

func listAllBookings(vcId string, day int, bookings []Booking) {
	for i := 0; i < len(bookings); i++ {
		if vcId == bookings[i].VcId && day == bookings[i].Day {
			fmt.Println(bookings[i].Id)
		}
	}
}

func main() {
	var users []user
	u1 := user{Id: "U1", Name: "Harry", Gender: "Male", State: "Karnataka", District: "Bangalore", Age: 21}
	users = u1.addUser(users)

	var vaccinationCentres []vaccinationCenter
	vc1 := vaccinationCenter{Id: "vc1", State: "Karnataka", District: "Bangalore"}
	vaccinationCentres = vc1.addVC(vaccinationCentres)

	var capacities []capacity
	capacities = addCapacity("vc1", 5, 5, capacities)

	listVcs("Bangalore", capacities, vaccinationCentres)

	var bookings []Booking
	bookings, capacities, err := addBooking("bk123", "vc1", 5, "U1", bookings, capacities, users)
	if err != nil {
		fmt.Println(err.Error())
	}

	bookings, capacities, err = cancelBooking("bk123", "vc1", "U1", bookings, capacities)
	if err != nil {
		fmt.Println(err.Error())
	}

	listAllBookings("vc1", 5, bookings)

	searchVaccinationCentres(5, "Bangalore", vaccinationCentres, capacities)
}
