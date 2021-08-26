# GeoIP lookup tool

Simple CLI tool to lookup long/lat and ASN information from IPv4/IPv6 addresses using the [MaxMind GeoIP2 databases](https://www.maxmind.com/en/accounts/411272/geoip/downloads).
If you want to keep it regularly updated, create a [free account](https://dev.maxmind.com/geoip/geolite2-free-geolocation-data?lang=en) and use the account ID and a generated license key.

The docker-compose.yaml file in this repository is a background service which can handle the updates for you in a given time interval, the default is set to 48h.
You need to change the `GEOIPUPDATE_ACCOUNT_ID`, `GEOIPUPDATE_LICENSE_KEY`.
As soon you have the GeoIP databases downloaded you can run the g-lookup tool like this:

```
Usage of ./g-lookup-linux-amd64:
  -asndb string
    	Path to the GeoIP2 or GeoLite2 ASN database file (default "./GeoLite2-ASN.mmdb")
  -citydb string
    	Path to the GeoIP2 or GeoLite2 City database file (default "./GeoLite2-City.mmdb")
  -lookup string
    	IPv4/IPv6 address used for the lookup (default "127.0.0.1")
```


A result will look like this:

```
`/g-lookup-darwin-amd64 -lookup 18.158.77.0
English city name: Frankfurt am Main
English subdivision name: Hesse
ISO country code: DE
Time zone: Europe/Berlin
Latitude: 50.1188
Longitude: 8.6843
ASN: 16509
ASN organisation: AMAZON-02
```

# Build

If you want build from source run:

```
GOOS=linux GOARCH=amd64 go build -o g-lookup-linux-amd64
```

Support OS and architectures can be listed with `go tool dist list`

gl & hf
