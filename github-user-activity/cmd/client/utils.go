package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetPublicToken(envFile string) string {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("A GitHub Personal Access Token was not found in file ", envFile)
		os.Exit(1)
	}
	token := os.Getenv("GITHUB_PUBLIC_API_TOKEN")
	return token
}

func (x *LabelElement) UnmarshalJSON(data []byte) error {
	x.LabelClass = nil
	var c LabelClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.LabelClass = &c
	}
	return nil
}

func (x *LabelElement) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.LabelClass != nil, x.LabelClass, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("unparsable number")
		}
		return false, errors.New("union does not contain number")
	case float64:
		return false, errors.New("decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("union does not contain array")
		}
		return false, errors.New("cannot handle delimiter")
	}
	return false, errors.New("cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("union must not be null")
}

func UnmarshalPublicUserEvents(data []byte) (PublicUserEvents, error) {
	var r PublicUserEvents
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PublicUserEvents) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
