package main

import (
	"fmt"
	//"io/ioutil"
	"os"
	"reflect"
	"strconv"

	"github.com/linkedin/goavro/v2"
)

var (
	codec *goavro.Codec
)

func init() {

	// schema, err := ioutil.ReadFile("nested_schema.avsc")
	// if err != nil {
	// 	panic(err)
	// }

	// Определение схемы Avro
	schema := `{
		"namespace": "my.namespace.com",
		"type":	"record",
		"name": "indentity",
		"fields": [
			{ "name": "FirstName", "type": "string"},
			{ "name": "LastName", "type": "string"},
			{ "name": "Errors", "type": ["null", {"type":"array", "items":"string"}], "default": null },
			{ "name": "Address", "type": ["null",{
				"namespace": "my.namespace.com",
				"type":	"record",
				"name": "address",
				"fields": [
					{ "name": "Address1", "type": "string" },
					{ "name": "Address2", "type": ["null", "string"], "default": null },
					{ "name": "City", "type": "string" },
					{ "name": "State", "type": "string" },
					{ "name": "Zip", "type": "int" }
				]
			}],"default":null}
		]
	}`

	//Create Schema Once
	//codec, err = goavro.NewCodec(string(schema))
	codec1, err := goavro.NewCodec(schema)
	if err != nil {
		panic(err)
	}
	codec = codec1
}

func main() {

	//Sample Data
	user := &User{
		FirstName: "John",
		LastName:  "Snow",
		Address: &Address{
			Address1: "1106 Pennsylvania Avenue",
			City:     "Wilmington",
			State:    "DE",
			Zip:      19806,
		},
	}

	fmt.Printf("User in=%+v\n", user)
	fmt.Printf("Address in=%+v\n", user.Address)

	///Convert Binary From Native
	binary, err := codec.BinaryFromNative(nil, user.ToStringMap())
	if err != nil {
		panic(err)
	}

	///Convert Native from Binary
	native, _, err := codec.NativeFromBinary(binary)
	if err != nil {
		panic(err)
	}

	// decodedUser := native.(map[string]interface{})
	// fmt.Println("Decoded:", decodedUser)
	// fmt.Println("FirstName:", decodedUser["FirstName"].(string))
	// fmt.Println("LastName:", decodedUser["LastName"].(string))
	
	// address := decodedUser["Address"].(map[string]interface{})["my.namespace.com.address"].(map[string]interface{})
	// fmt.Println("Address:", address)
	// fmt.Println("Address.City:", address["City"].(string))
	// fmt.Println("Address.State:", address["State"].(string))
	// fmt.Println("Address.Zip:", address["Zip"].(int32))



	//Convert it back tp Native
	userOut := StringMapToUser(native.(map[string]interface{}))
	fmt.Printf("User out=%+v\n", userOut)
	fmt.Printf("Address out=%+v\n", userOut.Address)

	if ok := reflect.DeepEqual(user, userOut); !ok {
		fmt.Fprintf(os.Stderr, "struct Compare Failed ok=%t\n", ok)
	} else {
		fmt.Println("user and userOut are deeply equal")
	  }
}

// User holds information about a user.
type User struct {
	FirstName string
	LastName  string
	Errors    []string
	Address   *Address
}

// Address holds information about an address.
type Address struct {
	Address1 string
	Address2 string
	City     string
	State    string
	Zip      int
}

// ToStringMap returns a map representation of the User.
func (u *User) ToStringMap() map[string]interface{} {
	datumIn := map[string]interface{}{
		"FirstName": string(u.FirstName),
		"LastName":  string(u.LastName),
	}

	if len(u.Errors) > 0 {
		datumIn["Errors"] = goavro.Union("array", u.Errors)
	} else {
		datumIn["Errors"] = goavro.Union("null", nil)
	}

	if u.Address != nil {
		addDatum := map[string]interface{}{
			"Address1": string(u.Address.Address1),
			"City":     string(u.Address.City),
			"State":    string(u.Address.State),
			"Zip":      int(u.Address.Zip),
		}
		if u.Address.Address2 != "" {
			addDatum["Address2"] = goavro.Union("string", u.Address.Address2)
		} else {
			addDatum["Address2"] = goavro.Union("null", nil)
		}

		//important need namespace and record name
		datumIn["Address"] = goavro.Union("my.namespace.com.address", addDatum)

	} else {
		datumIn["Address"] = goavro.Union("null", nil)
	}
	return datumIn
}

// StringMapToUser returns a User from a map representation of the User.
func StringMapToUser(data map[string]interface{}) *User {

	ind := &User{}
	for k, v := range data {
		switch k {
		case "FirstName":
			if value, ok := v.(string); ok {
				ind.FirstName = value
			}
		case "LastName":
			if value, ok := v.(string); ok {
				ind.LastName = value
			}
		case "Errors":
			if value, ok := v.(map[string]interface{}); ok {
				for _, item := range value["array"].([]interface{}) {
					ind.Errors = append(ind.Errors, item.(string))
				}
			}
		case "Address":
			if vmap, ok := v.(map[string]interface{}); ok {
				//important need namespace and record name
				if cookieSMap, ok := vmap["my.namespace.com.address"].(map[string]interface{}); ok {
					add := &Address{}
					for k, v := range cookieSMap {
						switch k {
						case "Address1":
							if value, ok := v.(string); ok {
								add.Address1 = value
							}
						case "Address2":
							if value, ok := v.(string); ok {
								add.Address2 = value
							}
						case "City":
							if value, ok := v.(string); ok {
								add.City = value
							}
						case "State":
							if value, ok := v.(string); ok {
								add.State = value
							}
						case "Zip":
							switch value := v.(type) {
								case string:
									zip, err := strconv.Atoi(value)
									if err != nil {
										fmt.Printf("Error converting ZIP to integer: %v\n", err)
									} else {
										add.Zip = zip
									}
								case int:
									add.Zip = value
								case int32:
									add.Zip = int(value)
								default:
									fmt.Printf("Unexpected type for ZIP: %T\n", v)
							}
						}
					}
					ind.Address = add
				}
			}
		}

	}
	return ind
}