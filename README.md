# teamcity-rest-client-go
Teamcity rest client in Go

Nothing fancy, just a go-swagger generate api from Teamcity's swagger.

Basic useage:

```go
package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	apiclient "github.com/cldmnky/teamcity-rest-client-go/client"
	"github.com/cldmnky/teamcity-rest-client-go/client/project"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

var projectName = "API-Test"

//DefaultSchemes:  Use https
var defaultSchemes = []string{"https"}

func main() {
	// reada yaml config
	c := Config{}

	err := yaml.Unmarshal([]byte(config), &c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	// create the API client
	client := apiclient.New(httptransport.New(os.Getenv("TC_HOST"), "", defaultSchemes), strfmt.Default)

	// make the authenticated request to get all items
	bearerTokenAuth := httptransport.BearerToken(os.Getenv("TC_ACCESS_TOKEN"))

	// get project features from packer
	pp := project.NewServeProjectParams()
	pp.ProjectLocator = projectName
	project, err := client.Project.ServeProject(pp, bearerTokenAuth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ProjectFeatures: %#v\n", project.Payload.ProjectFeatures)
	for _, p := range projectFeatures.Payload.ProjectFeatures.ProjectFeature {
		for _, f := range p.Properties.Property {
			fmt.Printf("ProjectFeature: %#v -> %#v\n", f.Name, f.Value)
		}
	}
}
```