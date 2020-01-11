package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/vladfin1/grpc/internal/connection"
	epb "github.com/vladfin1/grpc/proto/service/employeepb"
	pb "github.com/vladfin1/grpc/proto/service/unitpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct{}

func New() *Server {
	return &Server{}
}

func streamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	t := time.Now()
	fmt.Println("Method: " + info.FullMethod + ", Time: " + time.Since(t).String())
	return handler(srv, stream)
}

func (g *Server) Start() error {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer(grpc.StreamInterceptor(streamInterceptor))
	pb.RegisterUnitServiceServer(grpcServer, g)
	epb.RegisterEmployeeServiceServer(grpcServer, g)
	grpcServer.Serve(lis)

	return nil

}

func main() {
	g := New()
	g.Start()
}

type UnitItem struct {
	ID    string `json:"u_id"`
	Title string `json:"title"`
}

type EmployeeItem struct {
	ID    string `json:"e_id"`
	Name  string `json:"name"`
	Lname string `json:"lname"`
	ID_U  string `json:"id_u"`
}

func (g *Server) CreateUnit(ctx context.Context, req *pb.CreateUnitReq) (*pb.CreateUnitRes, error) {
	unit := req.GetUnit()
	data := UnitItem{
		Title: unit.GetTitle(),
	}
	db := connection.GetConnection()
	db.Exec("INSERT INTO unit(name) VALUES('" + data.Title + "')")
	defer db.Close()
	return &pb.CreateUnitRes{Unit: unit}, nil
}

func (g *Server) ReadUnit(ctx context.Context, req *pb.ReadUnitReq) (*pb.ReadUnitRes, error) {
	var response *pb.ReadUnitRes
	uid := req.GetId()
	db := connection.GetConnection()
	rows, err := db.Query("SELECT* FROM unit WHERE unit_id  = " + uid)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		unit := pb.Unit{}
		err = rows.Scan(&unit.UId, &unit.Title)
		if err != nil {
			panic(err)
		}
		response = &pb.ReadUnitRes{
			Unit: &pb.Unit{
				UId:   unit.UId,
				Title: unit.Title,
			},
		}
		return response, nil
	}
	defer rows.Close()
	defer db.Close()
	return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find unit with Id %s: %v", req.GetId(), err))
}

func (g *Server) DeleteUnit(ctx context.Context, req *pb.DeleteUnitReq) (*pb.DeleteUnitRes, error) {
	uid := req.GetId()
	db := connection.GetConnection()
	db.Exec("DELETE FROM unit WHERE unit_id  = " + uid)
	defer db.Close()
	return &pb.DeleteUnitRes{
		Success: true,
	}, nil
}

func (g *Server) UpdateUnit(ctx context.Context, req *pb.UpdateUnitReq) (*pb.UpdateUnitRes, error) {
	unit := req.GetUnit()
	db := connection.GetConnection()
	db.Exec("UPDATE unit SET name = '" + unit.Title + "' WHERE unit_id = " + unit.UId)
	defer db.Close()
	return &pb.UpdateUnitRes{Unit: unit}, nil
}

func (g *Server) ListUnits(req *pb.ListUnitReq, stream pb.UnitService_ListUnitsServer) error {
	db := connection.GetConnection()
	rows, err := db.Query("SELECT* FROM unit limit 11")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		unit := pb.Unit{}
		err = rows.Scan(&unit.UId, &unit.Title)
		if err != nil {
			panic(err)
		}
		stream.Send(&pb.ListUnitRes{
			Unit: &pb.Unit{
				UId:   unit.UId,
				Title: unit.Title,
			},
		})
	}
	defer db.Close()
	return nil
}

func (g *Server) CreateEmployee(ctx context.Context, req *epb.CreateEmployeeReq) (*epb.CreateEmployeeRes, error) {
	empl := req.GetEmployee()
	data := EmployeeItem{
		Name:  empl.GetName(),
		Lname: empl.GetLname(),
		ID_U:  empl.GetIdU(),
	}
	db := connection.GetConnection()
	db.Exec("INSERT INTO employee(name, last_name, uid) VALUES('" + data.Name + "' , '" + data.Lname + "', '" + data.ID_U + "')")
	defer db.Close()
	return &epb.CreateEmployeeRes{Employee: empl}, nil
}

func (g *Server) ReadEmployee(ctx context.Context, req *epb.ReadEmployeeReq) (*epb.ReadEmployeeRes, error) {
	var response *epb.ReadEmployeeRes
	eid := req.GetId()
	db := connection.GetConnection()
	rows, err := db.Query("SELECT* FROM employee WHERE emp_id  = " + eid)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		empl := epb.Employee{}
		err = rows.Scan(&empl.EId, &empl.Name, &empl.Lname, &empl.IdU)
		if err != nil {
			panic(err)
		}
		response = &epb.ReadEmployeeRes{
			Employee: &epb.Employee{
				Name:  empl.GetName(),
				Lname: empl.GetLname(),
				IdU:   empl.GetIdU(),
			},
		}
		return response, nil
	}
	defer rows.Close()
	defer db.Close()
	return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find employee with Id %s: %v", req.GetId(), err))
}

func (g *Server) DeleteEmployee(ctx context.Context, req *epb.DeleteEmployeeReq) (*epb.DeleteEmployeeRes, error) {
	eid := req.GetId()
	db := connection.GetConnection()
	db.Exec("DELETE FROM employee WHERE emp_id  = " + eid)
	defer db.Close()
	return &epb.DeleteEmployeeRes{
		Success: true,
	}, nil
}

func (g *Server) UpdateEmployee(ctx context.Context, req *epb.UpdateEmployeeReq) (*epb.UpdateEmployeeRes, error) {
	empl := req.GetEmployee()
	db := connection.GetConnection()
	db.Exec("UPDATE employee SET name = '" + empl.Name + "', last_name = '" + empl.Lname + "', uid = '" + empl.IdU + "' WHERE emp_id = " + empl.EId)
	defer db.Close()
	return &epb.UpdateEmployeeRes{Employee: empl}, nil
}

func (g *Server) ListEmployees(req *epb.ListEmployeeReq, stream epb.EmployeeService_ListEmployeesServer) error {
	db := connection.GetConnection()
	rows, err := db.Query("SELECT* FROM employee limit 11")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		empl := epb.Employee{}
		err = rows.Scan(&empl.EId, &empl.Name, &empl.Lname, &empl.IdU)
		if err != nil {
			panic(err)
		}
		stream.Send(&epb.ListEmployeeRes{
			Employee: &epb.Employee{
				Name:  empl.GetName(),
				Lname: empl.GetLname(),
				IdU:   empl.GetIdU(),
			},
		})
	}
	defer db.Close()
	return nil
}
