package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/idna"
	"net/url"
	"os"
	"reflect"
	"strings"
)

var outputFormat string
var apiKey string
var myDomain string
var filterFields string

const version string = "1.0.0"
const programName string = "IP2WHOIS Command Line"

var showVer bool = false

func init() {
	// read config for API key if exist
	LoadConfig()
}

func main() {
	flag.StringVar(&outputFormat, "o", "json", "Output format: json | pretty")
	flag.StringVar(&apiKey, "k", "", "API key: Get your API key from https://ip2location.io")
	flag.StringVar(&filterFields, "f", "", `Filter fields: Field names separted by comma. E.g., "domain,domain_id,status,registrar.name,registrant.organization"`)
	flag.BoolVar(&showVer, "v", false, "Show version")

	flag.Usage = func() {
		PrintUsage()
	}
	flag.Parse()

	if showVer {
		PrintVersion()
		return
	}

	if apiKey == "" {
		apiKey = config.APIKey
	}

	var arg = flag.Arg(0)

	if arg == "config" {
		UpdateAPIKey(flag.Arg(1))
		return
	} else if arg == "normal2puny" {
		PrintNormal2Puny(flag.Arg(1))
		return
	} else if arg == "puny2normal" {
		PrintPuny2Normal(flag.Arg(1))
		return
	} else if len(arg) == 0 {
		PrintUsage()
		return
	} else {
		myDomain = arg
	}

	filterFields = strings.TrimSpace(filterFields)

	if filterFields != "" {
		PrintFiltered()
	} else {
		PrintNormal()
	}
}

func PrintVersion() {
	fmt.Printf("%s Version: %s\n", programName, version)
}

// GetPunycode will convert normal text to Punycode
func GetPunycode(domain string) (string, error) {
	puny, err := url.Parse("https://" + domain)
	if err != nil {
		return "", err
	}
	str, err := idna.ToASCII(puny.Host)
	if err != nil {
		return "", err
	}
	return str, nil
}

// GetNormalText will convert Punycode to normal text
func GetNormalText(domain string) (string, error) {
	puny, err := url.Parse("https://" + domain)
	if err != nil {
		return "", err
	}
	str, err := idna.ToUnicode(puny.Host)
	if err != nil {
		return "", err
	}
	return str, nil
}

func PrintNormal2Puny(domain string) {
	res, err := GetPunycode(domain)

	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%+v\n", res)
}

func PrintPuny2Normal(domain string) {
	res, err := GetNormalText(domain)

	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%+v\n", res)
}

func PrintFiltered() {
	ipw, err := LookUpMap(myDomain)

	if err != nil {
		fmt.Println(err)
	} else {
		var field string
		fields := strings.Split(filterFields, ",")
		for i := 0; i < len(fields); i++ {
			field = strings.TrimSpace(fields[i])
			fmt.Print(field)
			if i+1 < len(fields) {
				fmt.Print(",")
			}
		}
		fmt.Println("")
		for i := 0; i < len(fields); i++ {
			field = strings.TrimSpace(fields[i])
			subfields := strings.Split(field, ".")

			// traverse the nested map
			var subfield string
			ipwsub := ipw
			for j := 0; j < len(subfields); j++ {
				subfield = subfields[j]

				if v, exists := ipwsub[subfield]; exists {
					if v == nil {
						break
					}
					if j+1 == len(subfields) { // end of the traversal
						switch t := reflect.TypeOf(v).Kind(); t {
						case reflect.String:
							v2 := v.(string)
							v2 = strings.ReplaceAll(v2, `"`, `\"`)
							fmt.Printf(`"%s"`, v2)
						case reflect.Float64:
							v2 := v.(float64) // all numbers are converted to float
							fmt.Print(int(v2))
						case reflect.Slice:
							fmt.Printf("%v", v)
						case reflect.Bool:
							v2 := v.(bool)
							fmt.Print(v2)
						default:
							fmt.Print("")
						}
					} else { // still need to drill down the map
						ipwsub = ipwsub[subfield].(map[string]interface{})
					}
				} else {
					break
				}
			}
			if i+1 < len(fields) {
				fmt.Print(",")
			}
		}
		fmt.Println("")
	}
}

func PrintNormal() {
	json, err := LookUpJSON(myDomain)

	if err != nil {
		fmt.Println(err)
	} else {
		if outputFormat == "json" {
			fmt.Printf("%s\n", json)
			return
		}
		pretty, err := PrettyString(json)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(pretty)
		}
	}
}

func PrintUsage() {
	fmt.Printf("%s Version %s\n", programName, version)
	var usage string = `
To query domain WHOIS:

  Usage: EXE [OPTION]... <DOMAIN NAME>

    -v                   Display the version and exit

    -h                   Print this help

    -k                   Specify the IP2WHOIS API key
                         Get your API key from https://www.ip2location.io

    -o                   Specify the output format
                         Valid values: json (default) | pretty

    -f                   Filter the result fields
                         Field names separated by comma and using period for nested field
                         E.g. domain,domain_id,status,registrar.name,registrant.organization

To store the API key

  Usage: EXE config <API KEY>


Other functions:

To convert normal domain name to Punycode

  Usage: EXE normal2puny <DOMAIN NAME>

To convert Punycode to normal domain name

  Usage: EXE puny2normal <DOMAIN NAME>

`

	usage = strings.ReplaceAll(usage, "EXE", os.Args[0])
	fmt.Println(usage)
}
