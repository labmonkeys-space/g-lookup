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
  -countrydb string
    	Path to the GeoIP2 or GeoLite2 Country database file (default "./GeoLite2-Country.mmdb")
  -lookup string
    	IPv4/IPv6 address used for the lookup (default "127.0.0.1")
```

A result will look like this:

```
`/g-lookup-darwin-amd64 -lookup 18.158.77.0
```

Result
```json
{
  "CountryRecord": {
    "Continent": {
      "Code": "EU",
      "GeoNameID": 6255148,
      "Names": {
        "de": "Europa",
        "en": "Europe",
        "es": "Europa",
        "fr": "Europe",
        "ja": "ヨーロッパ",
        "pt-BR": "Europa",
        "ru": "Европа",
        "zh-CN": "欧洲"
      }
    },
    "Country": {
      "GeoNameID": 2921044,
      "IsInEuropeanUnion": true,
      "IsoCode": "DE",
      "Names": {
        "de": "Deutschland",
        "en": "Germany",
        "es": "Alemania",
        "fr": "Allemagne",
        "ja": "ドイツ連邦共和国",
        "pt-BR": "Alemanha",
        "ru": "Германия",
        "zh-CN": "德国"
      }
    },
    "RegisteredCountry": {
      "GeoNameID": 6252001,
      "IsInEuropeanUnion": false,
      "IsoCode": "US",
      "Names": {
        "de": "USA",
        "en": "United States",
        "es": "Estados Unidos",
        "fr": "États-Unis",
        "ja": "アメリカ合衆国",
        "pt-BR": "Estados Unidos",
        "ru": "США",
        "zh-CN": "美国"
      }
    },
    "RepresentedCountry": {
      "GeoNameID": 0,
      "IsInEuropeanUnion": false,
      "IsoCode": "",
      "Names": null,
      "Type": ""
    },
    "Traits": {
      "IsAnonymousProxy": false,
      "IsSatelliteProvider": false
    }
  },
  "CityRecord": {
    "City": {
      "GeoNameID": 2925533,
      "Names": {
        "de": "Frankfurt am Main",
        "en": "Frankfurt am Main",
        "es": "Francfort",
        "fr": "Francfort-sur-le-Main",
        "ja": "フランクフルト・アム・マイン",
        "pt-BR": "Frankfurt am Main",
        "ru": "Франкфурт",
        "zh-CN": "法兰克福"
      }
    },
    "Continent": {
      "Code": "EU",
      "GeoNameID": 6255148,
      "Names": {
        "de": "Europa",
        "en": "Europe",
        "es": "Europa",
        "fr": "Europe",
        "ja": "ヨーロッパ",
        "pt-BR": "Europa",
        "ru": "Европа",
        "zh-CN": "欧洲"
      }
    },
    "Country": {
      "GeoNameID": 2921044,
      "IsInEuropeanUnion": true,
      "IsoCode": "DE",
      "Names": {
        "de": "Deutschland",
        "en": "Germany",
        "es": "Alemania",
        "fr": "Allemagne",
        "ja": "ドイツ連邦共和国",
        "pt-BR": "Alemanha",
        "ru": "Германия",
        "zh-CN": "德国"
      }
    },
    "Location": {
      "AccuracyRadius": 1000,
      "Latitude": 50.1188,
      "Longitude": 8.6843,
      "MetroCode": 0,
      "TimeZone": "Europe/Berlin"
    },
    "Postal": {
      "Code": "60313"
    },
    "RegisteredCountry": {
      "GeoNameID": 6252001,
      "IsInEuropeanUnion": false,
      "IsoCode": "US",
      "Names": {
        "de": "USA",
        "en": "United States",
        "es": "Estados Unidos",
        "fr": "États-Unis",
        "ja": "アメリカ合衆国",
        "pt-BR": "Estados Unidos",
        "ru": "США",
        "zh-CN": "美国"
      }
    },
    "RepresentedCountry": {
      "GeoNameID": 0,
      "IsInEuropeanUnion": false,
      "IsoCode": "",
      "Names": null,
      "Type": ""
    },
    "Subdivisions": [
      {
        "GeoNameID": 2905330,
        "IsoCode": "HE",
        "Names": {
          "de": "Hessen",
          "en": "Hesse",
          "es": "Hessen",
          "fr": "Hesse",
          "ru": "Гессен"
        }
      }
    ],
    "Traits": {
      "IsAnonymousProxy": false,
      "IsSatelliteProvider": false
    }
  },
  "AsnRecord": {
    "AutonomousSystemNumber": 16509,
    "AutonomousSystemOrganization": "AMAZON-02"
  },
  "GoogleMapsLink": "https://maps.google.com/?q=50.118800,8.684300"
}
```

# Build

If you want build from source run:

```
GOOS=linux GOARCH=amd64 go build -o g-lookup-linux-amd64
```

Support OS and architectures can be listed with `go tool dist list`

gl & hf
