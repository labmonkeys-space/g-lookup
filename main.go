package main

import (
	"flag"
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
	"os"
)

var geoIpCityDb, geoIpAsnDb, ipString string

func main() {
	log.SetOutput(os.Stdout)
	flag.StringVar(&geoIpCityDb, "citydb", "./GeoLite2-City.mmdb", "Path to the GeoIP2 or GeoLite2 City database file")
	flag.StringVar(&geoIpAsnDb, "asndb", "./GeoLite2-ASN.mmdb", "Path to the GeoIP2 or GeoLite2 ASN database file")
	flag.StringVar(&ipString, "lookup", "127.0.0.1", "IPv4/IPv6 address used for the lookup")
	flag.Parse()
	cityDb, err := geoip2.Open(geoIpCityDb)
	if err != nil {
		log.Fatal(err)
	}
	defer cityDb.Close()

	asnDb, err := geoip2.Open(geoIpAsnDb)
	if err != nil {
		log.Fatal(err)
	}

	ip := net.ParseIP(ipString)
	if ip == nil {
		log.Fatal("Error parsing the IP address " + ipString)
	}

	cityRecord, err := cityDb.City(ip)
	if err != nil {
		log.Fatal(err)
	}

	asnRecord, err := asnDb.ASN(ip)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("English city name: %v\n", cityRecord.City.Names["en"])
	if len(cityRecord.Subdivisions) > 0 {
		fmt.Printf("English subdivision name: %v\n", cityRecord.Subdivisions[0].Names["en"])
	}
	fmt.Printf("ISO country code: %v\n", cityRecord.Country.IsoCode)
	fmt.Printf("Time zone: %v\n", cityRecord.Location.TimeZone)
	fmt.Printf("Latitude: %v\n", cityRecord.Location.Latitude)
	fmt.Printf("Longitude: %v\n", cityRecord.Location.Longitude)
	fmt.Printf("ASN: %v\n", asnRecord.AutonomousSystemNumber)
	fmt.Printf("ASN organisation: %v\n", asnRecord.AutonomousSystemOrganization)
}

