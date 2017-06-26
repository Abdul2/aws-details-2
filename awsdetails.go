//uses drone to build
package main

import (
	"fmt"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/mgutz/logxi/v1"
	"os"
	_ "github.com/mattn/go-sqlite3"
)


type Instance struct {
	InstanceId   string `json:"instanceId"`
	InstanceType string `json:"instanceType"`
}

//holds json file data
var Listofinstances = make([]Instance, 0)


func GetData() {

	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("eu-west-1")})

	resp, err := svc.DescribeInstances(nil)

	if err != nil {

		panic(err)

	}

	for idx, res := range resp.Reservations {

		fmt.Println("  > Number of Reservations: ", len(res.Instances))

		var i Instance

		for _, inst := range resp.Reservations[idx].Instances {

			i.InstanceId = fmt.Sprintf("%s", *inst.InstanceId)
			i.InstanceType = fmt.Sprintf("%s", *inst.InstanceType)
		}

		Listofinstances = append(Listofinstances, i)

	}

}


func Writejsonfile() {

	content, _ := json.Marshal(Listofinstances)

	cont_string := fmt.Sprintf("%s", content)

	fmt.Printf("%s", cont_string)

	f, err := os.Create("output.json")

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err2 := f.WriteString(cont_string)

	if err2 != nil {
		log.Fatal(err2.Error())
	}

}

func main() {

	GetData()

	Writejsonfile()
}
