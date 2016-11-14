package google

import (
	"context"
	"log"

	"strconv"
	"strings"

	"googlemaps.github.io/maps"
)

type googleGeo struct {
	client *maps.Client
}

func NewGoogleGeo(apiKey string) *googleGeo {
	g := new(googleGeo)
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	g.client = client
	return g
}

func (g *googleGeo) ResolveLocation(req *GeoReq) (*GeoResp, error) {

	gcr := &maps.GeocodingRequest{
		Address:  req.Postcode,
		Language: req.Lang,
		Region:   req.Region,
	}
	parseComponents(req.Components, gcr)
	parseBounds(req.Bounds, gcr)
	parseLatLng(req.LatLng, gcr)
	parseResultType(req.ResultType, gcr)
	parseLocationType(req.LocationType, gcr)

	googleResp, err := g.client.Geocode(context.Background(), gcr)
	if err != nil {
		return nil, err
	}
	return &GeoResp{
		Lat:    googleResp[0].Geometry.Location.Lat,
		Long:   googleResp[0].Geometry.Location.Lng,
		FmtAdd: googleResp[0].FormattedAddress,
	}, nil
}

func parseComponents(components string, r *maps.GeocodingRequest) {
	if components != "" {
		c := strings.Split(components, "|")
		for _, cf := range c {
			i := strings.Split(cf, ":")
			switch i[0] {
			case "route":
				r.Components[maps.ComponentRoute] = i[1]
			case "locality":
				r.Components[maps.ComponentLocality] = i[1]
			case "administrative_area":
				r.Components[maps.ComponentAdministrativeArea] = i[1]
			case "postal_code":
				r.Components[maps.ComponentPostalCode] = i[1]
			case "country":
				r.Components[maps.ComponentCountry] = i[1]
			}
		}
	}
}

func parseBounds(bounds string, r *maps.GeocodingRequest) {
	if bounds != "" {
		b := strings.Split(bounds, "|")
		sw := strings.Split(b[0], ",")
		ne := strings.Split(b[1], ",")

		swLat, err := strconv.ParseFloat(sw[0], 64)
		if err != nil {
			log.Fatalf("Couldn't parse bounds: %#v", err)
		}
		swLng, err := strconv.ParseFloat(sw[1], 64)
		if err != nil {
			log.Fatalf("Couldn't parse bounds: %#v", err)
		}
		neLat, err := strconv.ParseFloat(ne[0], 64)
		if err != nil {
			log.Fatalf("Couldn't parse bounds: %#v", err)
		}
		neLng, err := strconv.ParseFloat(ne[1], 64)
		if err != nil {
			log.Fatalf("Couldn't parse bounds: %#v", err)
		}

		r.Bounds = &maps.LatLngBounds{
			NorthEast: maps.LatLng{Lat: neLat, Lng: neLng},
			SouthWest: maps.LatLng{Lat: swLat, Lng: swLng},
		}
	}
}

func parseLatLng(latlng string, r *maps.GeocodingRequest) {
	if latlng != "" {
		l := strings.Split(latlng, ",")
		lat, err := strconv.ParseFloat(l[0], 64)
		if err != nil {
			log.Fatalf("Couldn't parse latlng: %#v", err)
		}
		lng, err := strconv.ParseFloat(l[1], 64)
		if err != nil {
			log.Fatalf("Couldn't parse latlng: %#v", err)
		}
		r.LatLng = &maps.LatLng{
			Lat: lat,
			Lng: lng,
		}
	}
}

func parseResultType(resultType string, r *maps.GeocodingRequest) {
	if resultType != "" {
		r.ResultType = strings.Split(resultType, "|")
	}
}

func parseLocationType(locationType string, r *maps.GeocodingRequest) {
	if locationType != "" {
		for _, l := range strings.Split(locationType, "|") {
			switch l {
			case "ROOFTOP":
				r.LocationType = append(r.LocationType, maps.GeocodeAccuracyRooftop)
			case "RANGE_INTERPOLATED":
				r.LocationType = append(r.LocationType, maps.GeocodeAccuracyRangeInterpolated)
			case "GEOMETRIC_CENTER":
				r.LocationType = append(r.LocationType, maps.GeocodeAccuracyGeometricCenter)
			case "APPROXIMATE":
				r.LocationType = append(r.LocationType, maps.GeocodeAccuracyApproximate)
			}
		}

	}
}
