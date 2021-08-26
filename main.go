package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
	"os"
)

var geoIpCountryDb, geoIpCityDb, geoIpAsnDb, ipString string

type GeoIpLookup struct {
	CountryRecord geoip2.Country
	CityRecord geoip2.City
	AsnRecord geoip2.ASN
}

func main() {
	log.SetOutput(os.Stdout)
	flag.StringVar(&geoIpCountryDb, "countrydb", "./GeoLite2-Country.mmdb", "Path to the GeoIP2 or GeoLite2 Country database file")
	flag.StringVar(&geoIpCityDb, "citydb", "./GeoLite2-City.mmdb", "Path to the GeoIP2 or GeoLite2 City database file")
	flag.StringVar(&geoIpAsnDb, "asndb", "./GeoLite2-ASN.mmdb", "Path to the GeoIP2 or GeoLite2 ASN database file")
	flag.StringVar(&ipString, "lookup", "127.0.0.1", "IPv4/IPv6 address used for the lookup")
	flag.Parse()

	countryDb, err := geoip2.Open(geoIpCountryDb)
	if err != nil {
		log.Fatal(err)
	}
	defer countryDb.Close()

	cityDb, err := geoip2.Open(geoIpCityDb)
	if err != nil {
		log.Fatal(err)
	}
	defer cityDb.Close()

	asnDb, err := geoip2.Open(geoIpAsnDb)
	if err != nil {
		log.Fatal(err)
	}
	defer asnDb.Close()

	ip := net.ParseIP(ipString)
	if ip == nil {
		log.Fatal("Error parsing the IP address " + ipString)
	}

	countryRecord, err := countryDb.Country(ip)
	if err != nil {
		log.Fatal(err)
	}

	cityRecord, err := cityDb.City(ip)
	if err != nil {
		log.Fatal(err)
	}

	asnRecord, err := asnDb.ASN(ip)
	if err != nil {
		log.Fatal(err)
	}

	r := GeoIpLookup{
		*countryRecord,
		*cityRecord,
		*asnRecord,
	}
	resultJson, err := json.MarshalIndent(r, "", "  ")

	fmt.Printf("%v\n", string(resultJson))
}

