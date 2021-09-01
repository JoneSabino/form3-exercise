# Go Exercise - Form3
![form3 Exercise](https://github.com/JoneSabino/form3-exercise/actions/workflows/go.yml/badge.svg)

**Name:** Jones Sabino

### This is my first time with Golang =)

To run the tests, clone this repositories and run:

```bash
$ docker-compose up
```

The same tests are also running on Github Actions.

If you don't have docker installed, please visit:

[https://www.docker.com/get-started](https://www.docker.com/get-started)

The documentation of the packages are in the [`docs`](https://github.com/JoneSabino/form3-exercise/tree/main/docs) directory.

Examples of usage:

```go
import (
	"encoding/json"
	"fmt"
	"github.com/JoneSabino/form3-exercise/pkg/form3"
	"github.com/JoneSabino/form3-exercise/pkg/model"
	"github.com/satori/go.uuid"
	"strconv"
)

func exampleCreate() string {
	class, ct := new(string), new(string)
	*class = "Personal"
	*ct = "GB"

	accAttrs := model.AccountAttributes{
		AccountClassification: class,
		AlternativeNames:      []string{"bibi"},
		BankID:                "100000",
		BankIDCode:            "GBDSC",
		Bic:                   "NWBKGB42",
		Name:                  []string{"Bianca Sabino"},
		Country:               ct,
	}

	accData := model.AccountData{
		Attributes:     &accAttrs,
		OrganisationID: uuid.NewV4().String(),
	}

	res, err := form3.Create(accData)
	if err != nil {
		panic("error creating account \n" + err.Error())
	} else {
		fmt.Println("Account Created! \n ID: " + res.Data.ID)
	}
	return res.Data.ID
}

func exampleFetch(id string) (string, string) {
	resp, err := form3.Fetch(id)
	if err != nil {
		panic(err)
	}
	version := resp.Data.Version

	content, _ := json.MarshalIndent(resp, "", "  ")

	return string(content), strconv.FormatInt(*version, 10)

}

func main() {
	//Create
	id := testCreate()

	//Fetch
	body, version := testFetch(id)
	fmt.Println(body, version)

	//Delete
	fmt.Println("Deleting")
	sts, err := form3.Delete(id, version)
	if err != nil {
		panic(err)
	}
	fmt.Println(sts)
}
```
