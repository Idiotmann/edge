package main

import (
	"fmt"
	gs "github.com/Idiotmann/edge/algorithm"
	"github.com/Idiotmann/edge/model"
	"github.com/Idiotmann/edge/plotEva"
	"github.com/go-micro/plugins/v4/registry/consul"
	opentracing2 "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"log"
)

var GS gs.GS

func main() {

	//配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "micro/config")
	if err != nil {
		log.Fatal(err)
	}
	//注册中心
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})
	//链路追踪
	t, io, err := common.NewTracer("go.micro.service.product", "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//获取mysql配置,路径中不带前缀
	//mysql需要手动加载数据库驱动
	//mysqlConfig, err := common.GetMysqlFromConsul(consulConfig, "mysql")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//// 数据库类型：mysql，数据库用户名：root，密码：Kk1503060325，数据库名字：micro
	//db, err := gorm.Open("mysql", mysqlConfig.User+":"+mysqlConfig.Password+"@/"+mysqlConfig.Database+"?charset=utf8&parseTime=True&loc=Local")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer db.Close()
	//db.SingularTable(true) //禁用表名复数
	//
	////创建表之后，就注释掉
	////repository.NewProductRepository(db).InitTable()
	//productDataService := service2.NewProductDataService(repository.NewProductRepository(db))

	service := micro.NewService(
		micro.Name("go.micro.service.product"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8085"), //服务启动的地址
		micro.Registry(consulReg),       //注册中心
		//绑定链路追踪  服务端绑定handler,客户端绑定Client
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
	)
	//获取mysql配置,路径中不带前缀
	//mysql需要手动加载数据库驱动

	// Initialise service
	service.Init()

	// Register Handler
	pb.RegisterProductHandler(service.Server(), &handler.Product{ProductDataService: productDataService})
	if err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

	girlIp := model.GirlIP{"w1", "w2", "w3", "w4"}
	boyIp := model.BoyIP{"m1", "m2", "m3"}
	// 将数据存储按cost排序存储
	wbgList, _ := model.Compile(boyIp, girlIp) // 生成数据

	for _, v := range wbgList {
		//fmt.Printf("Graph长度是=%v \n", len(v.Graph))
		graph := v.Graph
		//创建people数据
		boyArr1, girlArr1 := gs.SortLike(graph)
		for i, k := range v.Graph {
			fmt.Printf("Graph[%v]宽度是=%v \n", i, len(k))

			for j, m := range k {
				fmt.Printf("Graph[%v][%v]=%v \n", i, j, m)

			}
		}
		index := GS.GaleShapley(boyArr1, girlArr1)
		fmt.Println("名字:\t最终匹配对象:")
		for _, v := range boyArr1 {
			fmt.Printf("%v\t%v\t\n", v.Name, v.Friend)
		}

		fmt.Println("")
		for _, v := range girlArr1 {
			fmt.Printf("%v\t%v\n", v.Name, v.Friend)
		}
		plotEva.Plot(index)
	}

}
