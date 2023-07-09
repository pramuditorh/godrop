package actions

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/pramuditorh/godrop/helpers"
	"github.com/urfave/cli/v2"
)

type NetworkV4 struct {
	IPAddress string `json:"ip_address"`
	Netmask   string `json:"netmask"`
	Gateway   string `json:"gateway"`
	Type      string `json:"type"`
}

type Network struct {
	V4 []NetworkV4 `json:"v4"`
}

type Droplet struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Memory int `json:"memory"`
	Vcpus int `json:"vcpus"`
	Disk int `json:"disk"`
	Networks Network `json:"networks"`
}

type Region struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Features []string `json:"features"`
	Available bool `json:"available"`
	Sizes []string `json:"sizes"`
}

type Link struct {}

type Meta struct {
	Total int `json:"total"`
}

type DropletResponse struct {
	Droplets []Droplet `json:"droplets"`
	Links Link `json:"links"`
	Meta Meta `json:"meta"`
}

type RegionResponse struct {
	Regions []Region `json:"regions"`
	Links Link `json:"links"`
	Meta Meta `json:"meta"`
}

func ListDroplets(c *cli.Context) error {
	resp, err := helpers.DoReq(http.MethodGet, "https://api.digitalocean.com/v2/droplets", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    var droplet DropletResponse
    errUn := json.Unmarshal(bodyBytes, &droplet)
		if errUn != nil {
			log.Fatal(errUn)
		}
		w := tabwriter.NewWriter(os.Stdout, 10, 1, 5, ' ', 0)
		fsHeader := "%s\t%s\t%s\t%s\t%s\n"
		fsValue := "%d\t%s\t%d\t%d\t%s\n"
		fmt.Println(string(bodyBytes))
		fmt.Fprintf(w, fsHeader, "ID", "NAME", "MEMORY", "VCPU", "NETWORK")
    for _, drp := range droplet.Droplets {
			fmt.Fprintf(w, fsValue, drp.ID, drp.Name, drp.Memory, drp.Vcpus, drp.Networks.V4[0].IPAddress)
		}
		w.Flush()
	}
	return nil
}

func ListRegions(c *cli.Context) error {
	resp, err := helpers.DoReq(http.MethodGet, "https://api.digitalocean.com/v2/regions", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

		var region RegionResponse
    errUn := json.Unmarshal(bodyBytes, &region)
		if errUn != nil {
			log.Fatal(errUn)
		}
		w := tabwriter.NewWriter(os.Stdout, 10, 1, 5, ' ', 0)
		fsHeader := "%s\t%s\t%s\n"
		fsValue := "%s\t%s\t%t\n"
		fmt.Fprintf(w, fsHeader, "REGION", "SLUG", "AVAILABILITY")
    for _, reg := range region.Regions {
			fmt.Fprintf(w, fsValue, reg.Name, reg.Slug, reg.Available)
		}
		w.Flush()
	}
	return nil
}