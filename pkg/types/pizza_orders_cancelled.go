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

type PizzaOrdersCancelled struct {
	Store_id int32 `json:"store_id"`

	Store_order_id int32 `json:"store_order_id"`

	Date int32 `json:"date"`

	Status string `json:"status"`
}

const PizzaOrdersCancelledAvroCRC64Fingerprint = "e\xbd^\x84\v\xb28\xac"

func NewPizzaOrdersCancelled() PizzaOrdersCancelled {
	r := PizzaOrdersCancelled{}
	return r
}

func DeserializePizzaOrdersCancelled(r io.Reader) (PizzaOrdersCancelled, error) {
	t := NewPizzaOrdersCancelled()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePizzaOrdersCancelledFromSchema(r io.Reader, schema string) (PizzaOrdersCancelled, error) {
	t := NewPizzaOrdersCancelled()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePizzaOrdersCancelled(r PizzaOrdersCancelled, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Store_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Store_order_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Date, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Status, w)
	if err != nil {
		return err
	}
	return err
}

func (r PizzaOrdersCancelled) Serialize(w io.Writer) error {
	return writePizzaOrdersCancelled(r, w)
}

func (r PizzaOrdersCancelled) Schema() string {
	return "{\"fields\":[{\"name\":\"store_id\",\"type\":{\"arg.properties\":{\"range\":{\"max\":10,\"min\":1}},\"type\":\"int\"}},{\"name\":\"store_order_id\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":1001,\"step\":2}},\"type\":\"int\"}},{\"name\":\"date\",\"type\":{\"arg.properties\":{\"range\":{\"max\":19000,\"min\":18000}},\"logicalType\":\"date\",\"type\":\"int\"}},{\"name\":\"status\",\"type\":{\"arg.properties\":{\"options\":[\"cancelled\"]},\"type\":\"string\"}}],\"name\":\"pizza_orders.PizzaOrdersCancelled\",\"type\":\"record\"}"
}

func (r PizzaOrdersCancelled) SchemaName() string {
	return "pizza_orders.PizzaOrdersCancelled"
}

func (_ PizzaOrdersCancelled) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ PizzaOrdersCancelled) SetInt(v int32)       { panic("Unsupported operation") }
func (_ PizzaOrdersCancelled) SetLong(v int64)      { panic("Unsupported operation") }
func (_ PizzaOrdersCancelled) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ PizzaOrdersCancelled) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ PizzaOrdersCancelled) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ PizzaOrdersCancelled) SetString(v string)   { panic("Unsupported operation") }
func (_ PizzaOrdersCancelled) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *PizzaOrdersCancelled) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Store_id}

		return w

	case 1:
		w := types.Int{Target: &r.Store_order_id}

		return w

	case 2:
		w := types.Int{Target: &r.Date}

		return w

	case 3:
		w := types.String{Target: &r.Status}

		return w

	}
	panic("Unknown field index")
}

func (r *PizzaOrdersCancelled) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *PizzaOrdersCancelled) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ PizzaOrdersCancelled) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ PizzaOrdersCancelled) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ PizzaOrdersCancelled) HintSize(int)                     { panic("Unsupported operation") }
func (_ PizzaOrdersCancelled) Finalize()                        {}

func (_ PizzaOrdersCancelled) AvroCRC64Fingerprint() []byte {
	return []byte(PizzaOrdersCancelledAvroCRC64Fingerprint)
}

func (r PizzaOrdersCancelled) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["store_id"], err = json.Marshal(r.Store_id)
	if err != nil {
		return nil, err
	}
	output["store_order_id"], err = json.Marshal(r.Store_order_id)
	if err != nil {
		return nil, err
	}
	output["date"], err = json.Marshal(r.Date)
	if err != nil {
		return nil, err
	}
	output["status"], err = json.Marshal(r.Status)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *PizzaOrdersCancelled) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["store_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Store_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for store_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["store_order_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Store_order_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for store_order_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["date"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Date); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for date")
	}
	val = func() json.RawMessage {
		if v, ok := fields["status"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Status); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for status")
	}
	return nil
}
