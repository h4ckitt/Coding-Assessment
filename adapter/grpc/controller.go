package grpc

import (
	"assessment/adapter/grpc/grpc_proto"
	"assessment/domain"
	"assessment/usecases"
	"context"
	"strconv"

	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CarServiceStruct struct {
	Service usecases.CarUseCase
	grpc_proto.UnimplementedCarServiceServer
}

func NewGRPCController(service usecases.CarUseCase) *CarServiceStruct {
	return &CarServiceStruct{
		Service: service,
	}
}

func (controller *CarServiceStruct) Register(ctx context.Context, car *grpc_proto.Car) (*grpc_proto.Car, error) {
	savedCar := domain.Car{
		Name:       car.GetName(),
		Color:      car.GetColor(),
		Type:       car.GetType(),
		SpeedRange: int(car.GetSpeedRange()),
		Features:   car.GetFeatures(),
	}

	err := controller.Service.Register(savedCar)

	if err != nil {
		//	controller.Logger.LogError("%s", err)
		return &grpc_proto.Car{}, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &grpc_proto.Car{}, nil
}

func (controller *CarServiceStruct) ViewCarDetails(ctx context.Context, id *wrappers.Int32Value) (*grpc_proto.Car, error) {
	carId := strconv.Itoa(int(id.Value))

	car, err := controller.Service.ViewDetails(carId)

	if err != nil {
		//	controller.Logger.LogError("%s\n", err)
		return &grpc_proto.Car{}, status.Errorf(codes.NotFound, "The Requested Resource Was Not Found")
	}

	result := &grpc_proto.Car{
		Name:       car.Name,
		Color:      car.Color,
		Type:       car.Type,
		SpeedRange: int32(car.SpeedRange),
		Features:   car.Features,
	}

	return result, nil
}

func (controller *CarServiceStruct) GetCarsByColorOrType(ctx context.Context, filter *grpc_proto.Filter) (*grpc_proto.Cars, error) {
	if color := filter.GetColor(); color != "" {
		cars, err := controller.Service.GetCarsByColor(color)

		if err != nil {
			return &grpc_proto.Cars{}, status.Errorf(codes.NotFound, "The Requested Resource Was Not Found")
		}

		/*if len(cars) == 0 {
			return &pb.Cars{}, status.Errorf(codes.NotFound, "The Requested Resource Was Not Found")
		}*/
		var result []*grpc_proto.Car
		for _, car := range cars {
			resultCar := &grpc_proto.Car{
				Name:       car.Name,
				Color:      car.Color,
				Type:       car.Type,
				SpeedRange: int32(car.SpeedRange),
				Features:   car.Features,
			}

			result = append(result, resultCar)
		}

		return &grpc_proto.Cars{Cars: result}, nil
	} else if carType := filter.GetType(); carType != "" {
		cars, err := controller.Service.GetCarsByType(carType)

		if err != nil {
			return &grpc_proto.Cars{}, status.Errorf(codes.NotFound, "The Requested Resource Was Not Found")
		}

		/*if len(cars) == 0 {
			return &pb.Cars{}, status.Errorf(codes.NotFound, "The Requested Resource Was Not Found")
		}*/
		var result []*grpc_proto.Car
		for _, car := range cars {
			resultCar := &grpc_proto.Car{
				Name:       car.Name,
				Color:      car.Color,
				Type:       car.Type,
				SpeedRange: int32(car.SpeedRange),
				Features:   car.Features,
			}

			result = append(result, resultCar)
		}

		return &grpc_proto.Cars{Cars: result}, nil
	}

	return &grpc_proto.Cars{}, status.Errorf(codes.InvalidArgument, "")
}
