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

type Orders struct {
	Ordertime int64 `json:"ordertime"`

	Orderid int32 `json:"orderid"`

	Itemid string `json:"itemid"`

	Orderunits float64 `json:"orderunits"`

	Address Address `json:"address"`
}

const OrdersAvroCRC64Fingerprint = "\xadEЛ4@\x0e\x1d"

func NewOrders() Orders {
	r := Orders{}
	r.Address = NewAddress()

	return r
}

func DeserializeOrders(r io.Reader) (Orders, error) {
	t := NewOrders()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeOrdersFromSchema(r io.Reader, schema string) (Orders, error) {
	t := NewOrders()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeOrders(r Orders, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.Ordertime, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Orderid, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Itemid, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.Orderunits, w)
	if err != nil {
		return err
	}
	err = writeAddress(r.Address, w)
	if err != nil {
		return err
	}
	return err
}

func (r Orders) Serialize(w io.Writer) error {
	return writeOrders(r, w)
}

func (r Orders) Schema() string {
	return "{\"fields\":[{\"name\":\"ordertime\",\"type\":{\"arg.properties\":{\"range\":{\"max\":1519273364600,\"min\":1487715775521}},\"type\":\"long\"}},{\"name\":\"orderid\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":0}},\"type\":\"int\"}},{\"name\":\"itemid\",\"type\":{\"arg.properties\":{\"regex\":\"Item_[1-9][0-9]{0,2}\"},\"type\":\"string\"}},{\"name\":\"orderunits\",\"type\":{\"arg.properties\":{\"range\":{\"max\":10,\"min\":0.1}},\"type\":\"double\"}},{\"name\":\"address\",\"type\":{\"fields\":[{\"name\":\"city\",\"type\":{\"arg.properties\":{\"regex\":\"City_[1-9]{0,2}\"},\"type\":\"string\"}},{\"name\":\"state\",\"type\":{\"arg.properties\":{\"regex\":\"State_[1-9]{0,2}\"},\"type\":\"string\"}},{\"name\":\"zipcode\",\"type\":{\"arg.properties\":{\"range\":{\"max\":99999,\"min\":10000}},\"type\":\"long\"}}],\"name\":\"address\",\"type\":\"record\"}}],\"name\":\"ksql.Orders\",\"type\":\"record\"}"
}

func (r Orders) SchemaName() string {
	return "ksql.Orders"
}

func (_ Orders) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Orders) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Orders) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Orders) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Orders) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Orders) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Orders) SetString(v string)   { panic("Unsupported operation") }
func (_ Orders) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Orders) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.Ordertime}

		return w

	case 1:
		w := types.Int{Target: &r.Orderid}

		return w

	case 2:
		w := types.String{Target: &r.Itemid}

		return w

	case 3:
		w := types.Double{Target: &r.Orderunits}

		return w

	case 4:
		r.Address = NewAddress()

		w := types.Record{Target: &r.Address}

		return w

	}
	panic("Unknown field index")
}

func (r *Orders) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Orders) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Orders) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Orders) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Orders) HintSize(int)                     { panic("Unsupported operation") }
func (_ Orders) Finalize()                        {}

func (_ Orders) AvroCRC64Fingerprint() []byte {
	return []byte(OrdersAvroCRC64Fingerprint)
}

func (r Orders) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["ordertime"], err = json.Marshal(r.Ordertime)
	if err != nil {
		return nil, err
	}
	output["orderid"], err = json.Marshal(r.Orderid)
	if err != nil {
		return nil, err
	}
	output["itemid"], err = json.Marshal(r.Itemid)
	if err != nil {
		return nil, err
	}
	output["orderunits"], err = json.Marshal(r.Orderunits)
	if err != nil {
		return nil, err
	}
	output["address"], err = json.Marshal(r.Address)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Orders) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["ordertime"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ordertime); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ordertime")
	}
	val = func() json.RawMessage {
		if v, ok := fields["orderid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Orderid); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for orderid")
	}
	val = func() json.RawMessage {
		if v, ok := fields["itemid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Itemid); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for itemid")
	}
	val = func() json.RawMessage {
		if v, ok := fields["orderunits"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Orderunits); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for orderunits")
	}
	val = func() json.RawMessage {
		if v, ok := fields["address"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Address); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for address")
	}
	return nil
}
