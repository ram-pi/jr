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
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
	"io"
)

func writeMapInt(r map[string]int32, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for k, e := range r {
		err = vm.WriteString(k, w)
		if err != nil {
			return err
		}
		err = vm.WriteInt(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type MapIntWrapper struct {
	Target *map[string]int32
	keys   []string
	values []int32
}

func (_ *MapIntWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ *MapIntWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ *MapIntWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ *MapIntWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ *MapIntWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ *MapIntWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ *MapIntWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ *MapIntWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ *MapIntWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *MapIntWrapper) SetDefault(i int)      { panic("Unsupported operation") }

func (r *MapIntWrapper) HintSize(s int) {
	if r.keys == nil {
		r.keys = make([]string, 0, s)
		r.values = make([]int32, 0, s)
	}
}

func (r *MapIntWrapper) NullField(_ int) {
	panic("Unsupported operation")
}

func (r *MapIntWrapper) Finalize() {
	for i := range r.keys {
		(*r.Target)[r.keys[i]] = r.values[i]
	}
}

func (r *MapIntWrapper) AppendMap(key string) types.Field {
	r.keys = append(r.keys, key)
	var v int32
	r.values = append(r.values, v)
	return &types.Int{Target: &r.values[len(r.values)-1]}
}

func (_ *MapIntWrapper) AppendArray() types.Field { panic("Unsupported operation") }
