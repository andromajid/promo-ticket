package main

/*
curl 'https://www.traveloka.com/api/v2/flight/search/oneway' \
-H 'origin: https://www.traveloka.com' \
-H 'accept-encoding: gzip, deflate, br' \
-H 'accept-language: en-US,en;q=0.9' \
-H 'cookie: G_ENABLED_IDPS=google; ajs_user_id=null; ajs_group_id=null; ajs_anonymous_id=%22501595a0-37e8-4150-9c45-34578228a4b1%22; __flash={}; tv-pnheader=1; tv-repeat-visit=true; useDateFlow=false; usePromoFinder=null; flightSourceAirport=JOG; flightDestinationAirport=BKKA; flightDepartureDate=10-10-2018; flightFlightType=ONE_WAY; flightNumOfAdults=1; flightNumChildren=0; flightNumInfants=0; flightSeatClassType=ECONOMY; flexibility=0; tvl=EKxlhlspXohmkW8WdeaZPqokIzPZewNKFoaMKK+2AIUHUhJKz+1BoXpjQ8ko3dXyYfD3vuN5L+nnLOa2yEpD73thJK3Oc8oMEFrTHDbcMMPfKeNXdemPQJYmpPjz8Le4LasD5nu7MQ8G3qLmAmDWKB7aLRB/LRo20JJSGu1t5zIp7flM3geEu7G2psj5RmUvyjDKZfnBFQySog//2+De71YEVuVl2qVC1XaaWluzES/oBeaEgEGrTuP+cwlWRBO/Rlv3hvKL2CDJjjw/ZIlNpQ==; tvs=aHiviipoLebGMEC5qT/rF7Niuv0lY8TIDgIjJJXlz4uV3HIUeuwnzqKzOsANTzg7pkmKMzen9Kefc05Qyn5MaU5UJH5/oaphwlt0lD1hvuGBUum3c0bCuTT4i2M6yV0G9k46s+GqAGIpEh9zj3d7hW9tbzSEpCwqbRrUsHnwgoFtj3mINbEK9of8CAYyMj0NH8wheJSz6vbAJYTOTpX4Q4t+iBZnoR21CMSeZTElWr4=' \
-H 'user-agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36' \
-H 'x-domain: flight' \
-H 'x-route-prefix: en-id' \
-H 'content-type: application/json' \
-H 'accept: application/json' \
-H 'referer: https://www.traveloka.com/en-id/flight/fullsearch?ap=JOG.BKKA&dt=10-10-2018.NA&ps=1.0.0&sc=ECONOMY' \
-H 'authority: www.traveloka.com' \
-H 'x-nonce: 13aa57aa-afd1-4eab-be19-136a46827614' --data-binary '{"clientInterface":"desktop","data":{"currency":"IDR","destinationAirportOrArea":"BKKA","flightDate":{"day":"10","month":"10","year":"2018"},"locale":"en_ID","newResult":true,"numSeats":{"numAdults":"1","numChildren":"0","numInfants":"0"},"seatPublishedClass":"ECONOMY","seqNo":null,"sortFilter":{"filterAirlines":[],"filterArrive":[],"filterDepart":[],"filterTransit":[],"selectedDeparture":"","sort":null},"sourceAirportOrArea":"JOG","searchId":null,"usePromoFinder":false,"useDateFlow":false},"fields":[]}' --compressed
*/

type TravelokaResult struct {
	ConnectingFlightRoutes []struct {
		DepartureAirport string `json:"departureAirport"`
		ArrivalAirport   string `json:"arrivalAirport"`
		ProviderID       string `json:"providerId"`
		DateTimeStamp    string `json:"dateTimeStamp"`
		NumDayOffset     string `json:"numDayOffset"`
		TotalNumStop     string `json:"totalNumStop"`
		FlightRefundInfo struct {
			RefundableStatus  string   `json:"refundableStatus"`
			RefundInfoSummary string   `json:"refundInfoSummary"`
			PolicyDetails     []string `json:"policyDetails"`
		} `json:"flightRefundInfo"`
		FlightRescheduleInfo struct {
			RescheduleStatus      string   `json:"rescheduleStatus"`
			RescheduleInfoSummary string   `json:"rescheduleInfoSummary"`
			PolicyDetails         []string `json:"policyDetails"`
		} `json:"flightRescheduleInfo"`
		RouteInventories []struct {
			SeatPublishedClass string `json:"seatPublishedClass"`
			SeatClass          string `json:"seatClass"`
			NumSeatLeft        string `json:"numSeatLeft"`
		} `json:"routeInventories"`
		Segments []struct {
			DepartureAirport     string `json:"departureAirport"`
			ArrivalAirport       string `json:"arrivalAirport"`
			FlightNumber         string `json:"flightNumber"`
			AirlineCode          string `json:"airlineCode"`
			BrandCode            string `json:"brandCode"`
			OperatingAirlineCode string `json:"operatingAirlineCode"`
			FareBasisCode        string `json:"fareBasisCode"`
			SegmentInventories   []struct {
				SeatClass      string `json:"seatClass"`
				NumSeatsLeft   string `json:"numSeatsLeft"`
				PublishedClass string `json:"publishedClass"`
				FareBasis      string `json:"fareBasis"`
				Refundable     string `json:"refundable"`
				SubclassInfo   string `json:"subclassInfo"`
			} `json:"segmentInventories"`
			AircraftType    string `json:"aircraftType"`
			HasMeal         bool   `json:"hasMeal"`
			FreeBaggageInfo struct {
				UnitOfMeasure string `json:"unitOfMeasure"`
				Weight        string `json:"weight"`
				Quantity      string `json:"quantity"`
			} `json:"freeBaggageInfo"`
			Facilities struct {
				Baggage struct {
					UnitOfMeasure  string `json:"unitOfMeasure"`
					Weight         string `json:"weight"`
					Quantity       string `json:"quantity"`
					AvailableToBuy bool   `json:"availableToBuy"`
				} `json:"baggage"`
				Wifi struct {
					Cost   string `json:"cost"`
					Chance string `json:"chance"`
				} `json:"wifi"`
				FreeMeal struct {
					Available bool `json:"available"`
				} `json:"freeMeal"`
				Entertainment struct {
					Available bool `json:"available"`
				} `json:"entertainment"`
				UsbAndPower struct {
					Usb    bool   `json:"usb"`
					Power  bool   `json:"power"`
					Chance string `json:"chance"`
				} `json:"usbAndPower"`
				Order []string `json:"order"`
			} `json:"facilities"`
			AircraftInformation struct {
				Aircraft struct {
					Model           string `json:"model"`
					Type            string `json:"type"`
					SeatInformation struct {
						Layout string `json:"layout"`
						Pitch  string `json:"pitch"`
						Type   string `json:"type"`
					} `json:"seatInformation"`
				} `json:"aircraft"`
				CabinBaggage struct {
					UnitOfMeasure string `json:"unitOfMeasure"`
					Weight        string `json:"weight"`
					Quantity      string `json:"quantity"`
				} `json:"cabinBaggage"`
			} `json:"aircraftInformation"`
			FlightLegInfoList       []string `json:"flightLegInfoList"`
			NumTransit              string   `json:"numTransit"`
			DurationMinute          string   `json:"durationMinute"`
			RouteNumDaysOffset      string   `json:"routeNumDaysOffset"`
			TzDepartureMinuteOffset string   `json:"tzDepartureMinuteOffset"`
			TzArrivalMinuteOffset   string   `json:"tzArrivalMinuteOffset"`
			DepartureDate           struct {
				Month string `json:"month"`
				Day   string `json:"day"`
				Year  string `json:"year"`
			} `json:"departureDate"`
			ArrivalDate struct {
				Month string `json:"month"`
				Day   string `json:"day"`
				Year  string `json:"year"`
			} `json:"arrivalDate"`
			DepartureTime struct {
				Hour   string `json:"hour"`
				Minute string `json:"minute"`
			} `json:"departureTime"`
			ArrivalTime struct {
				Hour   string `json:"hour"`
				Minute string `json:"minute"`
			} `json:"arrivalTime"`
			VisaRequired bool `json:"visaRequired"`
		} `json:"segments"`
		PromoLabels  []string `json:"promoLabels"`
		LoyaltyPoint string   `json:"loyaltyPoint"`
	} `json:"connectingFlightRoutes"`
	FlightMetadata []struct {
		Currency           string `json:"currency"`
		AirlineID          string `json:"airlineId"`
		SourceAirport      string `json:"sourceAirport"`
		DestinationAirport string `json:"destinationAirport"`
		NumAdult           string `json:"numAdult"`
		NumChild           string `json:"numChild"`
		NumInfant          string `json:"numInfant"`
	} `json:"flightMetadata"`
	TotalNumStop   string `json:"totalNumStop"`
	MobileAppDeal  bool   `json:"mobileAppDeal"`
	TripDuration   string `json:"tripDuration"`
	AdditionalInfo struct {
		SeatClassLabel string `json:"seatClassLabel"`
	} `json:"additionalInfo"`
	FlexibleTicket  bool `json:"flexibleTicket"`
	AirlineFareInfo struct {
		TotalSearchFare struct {
			Currency string `json:"currency"`
			Amount   string `json:"amount"`
		} `json:"totalSearchFare"`
		TotalAdditionalFare struct {
			Currency string `json:"currency"`
			Amount   string `json:"amount"`
		} `json:"totalAdditionalFare"`
		DetailedSearchFares []struct {
			AncillaryFare struct {
				AddOnFares []struct {
					Fare struct {
						Currency string `json:"currency"`
						Amount   string `json:"amount"`
					} `json:"fare"`
					Type string `json:"type"`
				} `json:"addOnFares"`
				ServiceFares []string `json:"serviceFares"`
			} `json:"ancillaryFare"`
			FlightRouteFares struct {
				AdultAirlineFare struct {
					Currency string `json:"currency"`
					Amount   string `json:"amount"`
				} `json:"adultAirlineFare"`
				AdultBaseFare struct {
					Currency string `json:"currency"`
					Amount   string `json:"amount"`
				} `json:"adultBaseFare"`
				ChildBaseFare struct {
					Currency string `json:"currency"`
					Amount   string `json:"amount"`
				} `json:"childBaseFare"`
				InfantBaseFare struct {
					Currency string `json:"currency"`
					Amount   string `json:"amount"`
				} `json:"infantBaseFare"`
				BaseFareTotal struct {
					Currency string `json:"currency"`
					Amount   string `json:"amount"`
				} `json:"baseFareTotal"`
				VatTotal struct {
					Currency string `json:"currency"`
					Amount   string `json:"amount"`
				} `json:"vatTotal"`
				CompulsoryInsuranceTotal struct {
					Currency string `json:"currency"`
					Amount   string `json:"amount"`
				} `json:"compulsoryInsuranceTotal"`
				PscTotal struct {
					Currency string `json:"currency"`
					Amount   string `json:"amount"`
				} `json:"pscTotal"`
				FuelSurchargeTotal struct {
					Currency string `json:"currency"`
					Amount   string `json:"amount"`
				} `json:"fuelSurchargeTotal"`
				AdminFeeTotal struct {
					Currency string `json:"currency"`
					Amount   string `json:"amount"`
				} `json:"adminFeeTotal"`
				TotalAirlineFare struct {
					Currency string `json:"currency"`
					Amount   string `json:"amount"`
				} `json:"totalAirlineFare"`
				TotalAdditionalFare struct {
					Currency string `json:"currency"`
					Amount   string `json:"amount"`
				} `json:"totalAdditionalFare"`
			} `json:"flightRouteFares"`
		} `json:"detailedSearchFares"`
	} `json:"airlineFareInfo"`
	AgentFareInfo struct {
		TotalSearchFare struct {
			Currency string `json:"currency"`
			Amount   string `json:"amount"`
		} `json:"totalSearchFare"`
		TotalAdditionalFare string `json:"totalAdditionalFare"`
		DetailedSearchFares []struct {
			AncillaryFare struct {
				AddOnFares []struct {
					Fare struct {
						Currency string `json:"currency"`
						Amount   string `json:"amount"`
					} `json:"fare"`
					Type string `json:"type"`
				} `json:"addOnFares"`
				ServiceFares []string `json:"serviceFares"`
			} `json:"ancillaryFare"`
			FlightRouteFares struct {
				AdultAgentFare struct {
					Currency string `json:"currency"`
					Amount   string `json:"amount"`
				} `json:"adultAgentFare"`
				ChildAgentFare struct {
					Currency string `json:"currency"`
					Amount   string `json:"amount"`
				} `json:"childAgentFare"`
				InfantAgentFare struct {
					Currency string `json:"currency"`
					Amount   string `json:"amount"`
				} `json:"infantAgentFare"`
				AdminFeeTotal struct {
					Currency string `json:"currency"`
					Amount   string `json:"amount"`
				} `json:"adminFeeTotal"`
				TotalFare struct {
					Currency string `json:"currency"`
					Amount   string `json:"amount"`
				} `json:"totalFare"`
			} `json:"flightRouteFares"`
		} `json:"detailedSearchFares"`
	} `json:"agentFareInfo"`
	FlightID     string `json:"flightId"`
	DesktopPrice struct {
		Currency string `json:"currency"`
		Amount   string `json:"amount"`
	} `json:"desktopPrice"`
	MAppsPrice struct {
		Currency string `json:"currency"`
		Amount   string `json:"amount"`
	} `json:"mAppsPrice"`
	DeltaPrice     string `json:"deltaPrice"`
	BundleFareInfo string `json:"bundleFareInfo"`
	LoyaltyPoint   string `json:"loyaltyPoint"`
}
