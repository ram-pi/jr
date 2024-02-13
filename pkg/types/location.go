// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     campaign_finance.avsc
 *     click_stream.avsc
 *     click_stream_users.avsc
 *     clickstreamcodes.avsc
 *     credit_cards.avsc
 *     csv_product.avsc
 *     csv_user.avsc
 *     device_information.avsc
 *     fleet_mgmt_description.avsc
 *     fleet_mgmt_location.avsc
 *     fleet_mgmt_sensors.avsc
 *     gaming_games.avsc
 *     gaming_player_activity.avsc
 *     gaming_players.avsc
 *     insurance_customer_activity.avsc
 *     insurance_customers.avsc
 *     insurance_offers.avsc
 *     inventory.avsc
 *     map_dumb_schema.avsc
 *     net_device.avsc
 *     orders.avsc
 *     page_views.avsc
 *     payroll_bonus.avsc
 *     payroll_employee.avsc
 *     payroll_employee_location.avsc
 *     pizza_orders.avsc
 *     pizza_orders_cancelled.avsc
 *     pizza_orders_completed.avsc
 *     product.avsc
 *     purchase.avsc
 *     ratings.avsc
 *     shoe.avsc
 *     shoe_clickstream.avsc
 *     shoe_customer.avsc
 *     shoe_order.avsc
 *     siem_logs.avsc
 *     stock_trades.avsc
 *     stores.avsc
 *     syslog_logs.avsc
 *     transactions.avsc
 *     user.avsc
 *     users.avsc
 *     users_array_map.avsc
 */
package types

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Location struct {
	Latitude float64 `json:"latitude"`

	Longitude float64 `json:"longitude"`
}

const LocationAvroCRC64Fingerprint = "չ>\x15\xdfTZ\xe7"

func NewLocation() Location {
	r := Location{}
	return r
}

func DeserializeLocation(r io.Reader) (Location, error) {
	t := NewLocation()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeLocationFromSchema(r io.Reader, schema string) (Location, error) {
	t := NewLocation()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeLocation(r Location, w io.Writer) error {
	var err error
	err = vm.WriteDouble(r.Latitude, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.Longitude, w)
	if err != nil {
		return err
	}
	return err
}

func (r Location) Serialize(w io.Writer) error {
	return writeLocation(r, w)
}

func (r Location) Schema() string {
	return "{\"fields\":[{\"name\":\"latitude\",\"type\":\"double\"},{\"name\":\"longitude\",\"type\":\"double\"}],\"name\":\"fleet_mgmt.location\",\"type\":\"record\"}"
}

func (r Location) SchemaName() string {
	return "fleet_mgmt.location"
}

func (_ Location) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Location) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Location) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Location) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Location) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Location) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Location) SetString(v string)   { panic("Unsupported operation") }
func (_ Location) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Location) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Double{Target: &r.Latitude}

		return w

	case 1:
		w := types.Double{Target: &r.Longitude}

		return w

	}
	panic("Unknown field index")
}

func (r *Location) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Location) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Location) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Location) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Location) HintSize(int)                     { panic("Unsupported operation") }
func (_ Location) Finalize()                        {}

func (_ Location) AvroCRC64Fingerprint() []byte {
	return []byte(LocationAvroCRC64Fingerprint)
}

func (r Location) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["latitude"], err = json.Marshal(r.Latitude)
	if err != nil {
		return nil, err
	}
	output["longitude"], err = json.Marshal(r.Longitude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Location) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["latitude"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Latitude); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for latitude")
	}
	val = func() json.RawMessage {
		if v, ok := fields["longitude"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Longitude); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for longitude")
	}
	return nil
}
