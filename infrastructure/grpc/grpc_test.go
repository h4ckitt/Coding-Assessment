package grpc_test

import (
	"assessment/domain"
	"assessment/infrastructure/db/inmemoryteststore"
	rpc "assessment/infrastructure/grpc"
	pb "assessment/infrastructure/grpc/grpc_proto"
	"assessment/usecases"
	"context"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net"
	"testing"
)

type GRPCTestSuite struct {
	suite.Suite
	pb.UnimplementedCarServiceServer
	client pb.CarServiceClient
	closer func()
}

func (g *GRPCTestSuite) SetupSuite() {
	repo := inmemoryteststore.NewMemoryStore()

	cars := []domain.Car{
		{
			Name:       "Toyota Tundra",
			Color:      "red",
			Type:       "suv",
			SpeedRange: 180,
			Features:   []string{"panorama", "surround-system"},
		},
		{
			Name:       "Tesla Model S Plaid",
			Color:      "blue",
			Type:       "sedan",
			SpeedRange: 240,
			Features:   []string{"auto-parking", "panorama", "surround-system"},
		},
		{
			Name:       "Ducati 99R",
			Color:      "green",
			Type:       "motor-bike",
			SpeedRange: 200,
			Features:   []string{"surround-system"},
		},
		{
			Name:       "Ford ES500",
			Color:      "blue",
			Type:       "van",
			SpeedRange: 80,
			Features:   []string{"surround-system", "sunroof", "surround-system"},
		},
	}

	for _, car := range cars {
		err := repo.Store(car)

		require.NoError(g.T(), err)
	}

	service := usecases.NewService(repo)

	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	pb.RegisterCarServiceServer(server, rpc.NewGRPCController(service))

	go func() {
		if err := server.Serve(listener); err != nil {
			g.T().Fatalf("Error Starting Server: %v\n", err)
		}
	}()

	ctx := context.TODO()

	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}), grpc.WithInsecure())

	if err != nil {
		g.T().Fatalf("Error Creating A Connection: %v\n", err)
	}

	g.client = pb.NewCarServiceClient(conn)

	g.closer = func() {
		err := listener.Close()
		if err != nil {
			g.T().Fatal(err)
		}
		server.Stop()
	}

}

func (g *GRPCTestSuite) TearDownSuite() {
	g.closer()
}

func TestGRPCSuite(t *testing.T) {
	suite.Run(t, new(GRPCTestSuite))
}

func (g *GRPCTestSuite) TestRegister() {
	ctx := context.TODO()

	car := &pb.Car{
		Name:       "Toyota Tundra",
		Color:      "Blue",
		Type:       "SUV",
		SpeedRange: 240,
		Features:   []string{"sunroof", "panorama"},
	}

	_, err := g.client.Register(ctx, car)

	require.NoError(g.T(), err)

	for _, color := range []string{"", "brown", "black", "pink"} {
		car.Color = color

		_, err = g.client.Register(ctx, car)

		require.Errorf(g.T(), err, "Expected Error, Got: %v\n", err)
	}

	car.Color = "blue"

	for _, types := range []string{"bike", "unicycle", "sandals"} {
		car.Type = types

		_, err = g.client.Register(ctx, car)

		require.Errorf(g.T(), err, "Expected Error, Got: %v\n", err)
	}

	car.Type = "suv"

	for _, speed := range []int32{-1, 250, 30000} {
		car.SpeedRange = speed

		_, err = g.client.Register(ctx, car)

		require.Errorf(g.T(), err, "Expected Error, Got: %v\n", err)
	}

	car.SpeedRange = 200

	car.Features = []string{"high-suspension", "climate-control", "moonrooof"}

	_, err = g.client.Register(ctx, car)

	require.Errorf(g.T(), err, "Expected Error, Got: %v\n", err)

}

func (g *GRPCTestSuite) TestViewCarDetails() {
	ctx := context.TODO()
	id := wrapperspb.Int32(4)
	_, err := g.client.ViewCarDetails(ctx, id)

	require.NoError(g.T(), err)

	for _, id := range []int{-1, 0, 2000} {
		d := wrapperspb.Int32(int32(id))
		_, err := g.client.ViewCarDetails(ctx, d)

		require.Errorf(g.T(), err, "Expected An Error, Got: %v\n", err)
	}
}

func (g *GRPCTestSuite) TestGetCarByColorOrType() {
	ctx := context.TODO()

	filter := &pb.Filter{
		FilterType: &pb.Filter_Color{
			Color: "blue",
		},
	}

	cars, err := g.client.GetCarsByColorOrType(ctx, filter)

	require.NoError(g.T(), err)

	for _, car := range cars.GetCars() {
		require.Equal(g.T(), "blue", car.GetColor())
	}

	filter = &pb.Filter{
		FilterType: &pb.Filter_Color{
			Color: "maroon",
		},
	}

	_, err = g.client.GetCarsByColorOrType(ctx, filter)

	require.Errorf(g.T(), err, "Expected An Error, Got: %v\n", err)

	filter = &pb.Filter{
		FilterType: &pb.Filter_Type{
			Type: "sedan",
		},
	}

	cars, err = g.client.GetCarsByColorOrType(ctx, filter)

	require.NoError(g.T(), err)

	for _, car := range cars.GetCars() {
		require.Equal(g.T(), "sedan", car.GetType())
	}

	filter = &pb.Filter{
		FilterType: &pb.Filter_Type{
			Type: "carriage",
		},
	}

	_, err = g.client.GetCarsByColorOrType(ctx, filter)

	require.Errorf(g.T(), err, "Expected An Error, Got: %v\n", err)

	filter = &pb.Filter{
		FilterType: &pb.Filter_Type{},
	}

	_, err = g.client.GetCarsByColorOrType(ctx, filter)

	require.Errorf(g.T(), err, "Expected An Error, Got: %v\n", err)
}
