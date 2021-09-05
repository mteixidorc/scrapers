# Booking rooms scraper

This is a very basic scraper that finds the cheapest hotel room on a HTML generated by this web request:
[booking request](https://www.booking.com/hotel/no/spitsbergen.en-gb.html?no_rooms=1&checkin=2021-12-03&checkout=2021-12-07&group_adults=2&group_children=0&req_adults=2&req_children=0&dist=0&type=total&selected_currency=NOK)

Actually this URL is hardcoded, sorry.


## Clone the project

First things first, clone the project:


## Requirements

You will need GO compiler to build and/or execute the application 
[Golang](https://golang.org)


## Build / Execute 
It's pretty simple, build an executable first if you haven't:
```
$ go build
```

And run it:
```
$ ./scrapers 
```

Alternatively you could execute project directly without a previous compilation

```
$ go run main.go 
```