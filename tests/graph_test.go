package test

import (
	"context"
	"log"
	"testing"

	"github.com/AashrayAnand/tripit/graph"
	"github.com/AashrayAnand/tripit/secret"
	"googlemaps.github.io/maps"
)

var gclient, err = maps.NewClient(maps.WithAPIKey(secret.GOOG_API_KEY))
var cities = []string{"Seattle,WA", "New York,NY", "San Francisco,CA"}

func GetDistanceMatrix() (*maps.DistanceMatrixResponse, error) {
	// get distance matrix for set of n=3 origin/dest cities
	r := &maps.DistanceMatrixRequest{
		Origins:      cities,
		Destinations: cities,
	}

	// execute call to DistanceMatrix API
	matrix, err := gclient.DistanceMatrix(context.Background(), r)

	return matrix, err
}

func TestClientCreated(t *testing.T) {
	if err != nil {
		log.Fatalf("fatal error on gclient create: %s", err.Error())
	}
}

func TestGetDistanceMatrix(t *testing.T) {
	matrix, err := GetDistanceMatrix()
	if err != nil {
		log.Fatalf("fatal error: %s", err.Error())
	}

	// matrix should be nxn, distances between each pair of cities
	if len(matrix.Rows) != len(cities) || len(matrix.Rows[0].Elements) != len(cities) {
		log.Fatal("error, distance matrix size is incorrect")
	}
}

func TestCreateGraph(t *testing.T) {
	// want to test that an appropriate graph representation is created
	// from the distance matrix query
	matrix, err := GetDistanceMatrix()
	if err != nil {
		log.Fatalf("fatal error: %s", err.Error())
	}

	// only need to pass origin list, destination list is
	// equivalent, can use origin list to construct adjacency list
	TripGraph := graph.Construct(matrix)

	if TripGraph == nil {
		log.Fatal("graph construction not implemented")
	}

	if len(TripGraph.Locations) != len(cities) {
		log.Fatalf("not enough cities in trip graph, expected %d, got %d", len(cities), len(TripGraph.Locations))
	}

	if len(TripGraph.Edges[TripGraph.Locations[0]]) != len(cities) {
		log.Fatalf("not enough edges in trip graph, expected %d, got %d", len(cities), len(TripGraph.Edges[TripGraph.Locations[0]]))
	}
}
