package grpc

import (
	"assessment/domain"
	pb "assessment/infrastructure/grpc/grpc_proto"
	"assessment/usecases"
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strconv"
)

type CarServiceStruct struct {
	Service usecases.CarUseCase
	Logger  usecases.Logger
	pb.UnimplementedCarServiceServer
}

func NewGRPCController(service usecases.CarUseCase, logger usecases.Logger) *CarServiceStruct {
	return &CarServiceStruct{
		Service: service,
		Logger:  logger,
	}
}

func (controller *CarServiceStruct) Register(ctx context.Context, car *pb.Car) (*pb.Car, error) {
	log.Printf("Processing Vehicle: %v\n", car.GetName())

	savedCar := domain.Car{
		Name:       car.GetName(),
		Color:      car.GetColor(),
		Type:       car.GetType(),
		SpeedRange: int(car.GetSpeedRange()),
		Features:   car.GetFeatures(),
	}

	err := controller.Service.Register(savedCar)

	if err != nil {
		controller.Logger.LogError("%s", err)
		return &pb.Car{}, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &pb.Car{}, nil
}

func (controller *CarServiceStruct) GetCar(ctx context.Context, id *wrappers.Int32Value) (*pb.Car, error) {
	carId := strconv.Itoa(int(id.Value))

	car, err := controller.Service.ViewDetails(carId)

	if err != nil {
		controller.Logger.LogError("%s\n", err)
		return &pb.Car{}, status.Errorf(codes.NotFound, err.Error())
	}

	result := &pb.Car{
		Name:       car.Name,
		Color:      car.Color,
		Type:       car.Type,
		SpeedRange: int32(car.SpeedRange),
		Features:   car.Features,
	}

	return result, nil
}

func (controller *CarServiceStruct) GetCars(ctx context.Context, filter *pb.Filter) (*pb.Cars, error) {
	if color := filter.GetColor(); color != "" {
		cars, err := controller.Service.GetCarsByColor(color)

		if err != nil {
			return &pb.Cars{}, status.Errorf(codes.Internal, "Your Request Could Not Be Completed At This Time, Please Try Again Later")
		}

		if len(cars) == 0 {
			return &pb.Cars{}, status.Errorf(codes.NotFound, "The Requested Resource Was Not Found")
		}
		var result []*pb.Car
		for _, car := range cars {
			resultCar := &pb.Car{
				Name:       car.Name,
				Color:      car.Color,
				Type:       car.Type,
				SpeedRange: int32(car.SpeedRange),
				Features:   car.Features,
			}

			result = append(result, resultCar)
		}

		return &pb.Cars{Cars: result}, nil
	}

	return &pb.Cars{}, status.Errorf(codes.InvalidArgument, "")
}
