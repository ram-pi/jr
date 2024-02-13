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

type FleetMgmtSensors struct {
	Vehicle_id int32 `json:"vehicle_id"`

	Engine_temperature int32 `json:"engine_temperature"`

	Average_rpm int32 `json:"average_rpm"`
}

const FleetMgmtSensorsAvroCRC64Fingerprint = "$ \xae)\xcb%LV"

func NewFleetMgmtSensors() FleetMgmtSensors {
	r := FleetMgmtSensors{}
	return r
}

func DeserializeFleetMgmtSensors(r io.Reader) (FleetMgmtSensors, error) {
	t := NewFleetMgmtSensors()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeFleetMgmtSensorsFromSchema(r io.Reader, schema string) (FleetMgmtSensors, error) {
	t := NewFleetMgmtSensors()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeFleetMgmtSensors(r FleetMgmtSensors, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Vehicle_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Engine_temperature, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Average_rpm, w)
	if err != nil {
		return err
	}
	return err
}

func (r FleetMgmtSensors) Serialize(w io.Writer) error {
	return writeFleetMgmtSensors(r, w)
}

func (r FleetMgmtSensors) Schema() string {
	return "{\"fields\":[{\"name\":\"vehicle_id\",\"type\":{\"arg.properties\":{\"range\":{\"max\":9999,\"min\":1000}},\"type\":\"int\"}},{\"name\":\"engine_temperature\",\"type\":{\"arg.properties\":{\"range\":{\"max\":250,\"min\":150}},\"type\":\"int\"}},{\"name\":\"average_rpm\",\"type\":{\"arg.properties\":{\"range\":{\"max\":5000,\"min\":1800}},\"type\":\"int\"}}],\"name\":\"fleet_mgmt.FleetMgmtSensors\",\"type\":\"record\"}"
}

func (r FleetMgmtSensors) SchemaName() string {
	return "fleet_mgmt.FleetMgmtSensors"
}

func (_ FleetMgmtSensors) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ FleetMgmtSensors) SetInt(v int32)       { panic("Unsupported operation") }
func (_ FleetMgmtSensors) SetLong(v int64)      { panic("Unsupported operation") }
func (_ FleetMgmtSensors) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ FleetMgmtSensors) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ FleetMgmtSensors) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ FleetMgmtSensors) SetString(v string)   { panic("Unsupported operation") }
func (_ FleetMgmtSensors) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *FleetMgmtSensors) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Vehicle_id}

		return w

	case 1:
		w := types.Int{Target: &r.Engine_temperature}

		return w

	case 2:
		w := types.Int{Target: &r.Average_rpm}

		return w

	}
	panic("Unknown field index")
}

func (r *FleetMgmtSensors) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *FleetMgmtSensors) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ FleetMgmtSensors) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ FleetMgmtSensors) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ FleetMgmtSensors) HintSize(int)                     { panic("Unsupported operation") }
func (_ FleetMgmtSensors) Finalize()                        {}

func (_ FleetMgmtSensors) AvroCRC64Fingerprint() []byte {
	return []byte(FleetMgmtSensorsAvroCRC64Fingerprint)
}

func (r FleetMgmtSensors) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["vehicle_id"], err = json.Marshal(r.Vehicle_id)
	if err != nil {
		return nil, err
	}
	output["engine_temperature"], err = json.Marshal(r.Engine_temperature)
	if err != nil {
		return nil, err
	}
	output["average_rpm"], err = json.Marshal(r.Average_rpm)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *FleetMgmtSensors) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["vehicle_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Vehicle_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for vehicle_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["engine_temperature"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Engine_temperature); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for engine_temperature")
	}
	val = func() json.RawMessage {
		if v, ok := fields["average_rpm"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Average_rpm); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for average_rpm")
	}
	return nil
}
