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
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayString(r []string, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = vm.WriteString(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayStringWrapper struct {
	Target *[]string
}

func (_ ArrayStringWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayStringWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayStringWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayStringWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayStringWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayStringWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayStringWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayStringWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayStringWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayStringWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayStringWrapper) Finalize()                        {}
func (_ ArrayStringWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayStringWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]string, 0, s)
	}
}
func (r ArrayStringWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayStringWrapper) AppendArray() types.Field {
	var v string

	*r.Target = append(*r.Target, v)
	return &types.String{Target: &(*r.Target)[len(*r.Target)-1]}
}
