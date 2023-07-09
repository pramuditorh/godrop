package actions

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"

	"github.com/urfave/cli/v2"
)

type ProvisionData struct {
	Name string `json:"name"`
	Region string `json:"region"`
	Size  string `json:"size"`
	Image string `json:"image"`
	Ssh_keys []int `json:"ssh_keys"`
}


func Provision(c *cli.Context) error {
	var ssh_key []int = []int{}
	if c.Int("ssh-key-id") != 0 {
		ssh_key = append(ssh_key, c.Int("ssh-key-id"))	
	}

	data := ProvisionData{
		Name: c.String("name"),
		Region: c.String("region"),
		Size: c.String("size"),
		Image: c.String("image"),
		Ssh_keys: ssh_key,
	}

	marshalled, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("curl -X POST -H \"Content-Type: application/json\" -H \"Authorization: Bearer $DIGITALOCEAN_TOKEN\" -d %s \"https://api.digitalocean.com/v2/droplets\"", string(marshalled))
	fmt.Println(url)

	if err := checkDroplet(); err != nil {
		return err
	}
	return nil
}

func checkDroplet() error {
	// Simulate droplet is up or not
	var err error = errors.New("VM gak up bang")
	var s = rand.Intn(2)
	if s != 0 {
		return err
	}
	return nil
}