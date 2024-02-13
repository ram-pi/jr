// Autogenerated code. DO NOT EDIT.
package types

var address Address

var array_orderline_wrapper ArrayOrderlineWrapper

var array_string_wrapper ArrayStringWrapper

var bytes Bytes

var campaign_finance CampaignFinance

var click_stream ClickStream

var click_stream_codes ClickStreamCodes

var click_stream_users ClickStreamUsers

var clickstreamusers Clickstreamusers

var codes Codes

var credit_cards CreditCards

var csv_product CsvProduct

var csv_user CsvUser

var destination Destination

var device_information DeviceInformation

var deviceinformation Deviceinformation

var fleet_mgmt_description FleetMgmtDescription

var fleet_mgmt_location FleetMgmtLocation

var fleet_mgmt_sensors FleetMgmtSensors

var fleetmgmtdescription Fleetmgmtdescription

var fleetmgmtlocation Fleetmgmtlocation

var fleetmgmtsensors Fleetmgmtsensors

var gaming_games GamingGames

var gaming_player_activity GamingPlayerActivity

var gaming_players GamingPlayers

var gaminggames Gaminggames

var gamingplayeractivity Gamingplayeractivity

var gamingplayers Gamingplayers

var insurance_customer_activity InsuranceCustomerActivity

var insurance_customers InsuranceCustomers

var insurance_offers InsuranceOffers

var inventory Inventory

var location Location

var map_bool_wrapper MapBoolWrapper

var map_bytes_wrapper MapBytesWrapper

var map_float_wrapper MapFloatWrapper

var map_int_wrapper MapIntWrapper

var map_long_wrapper MapLongWrapper

var map_string_wrapper MapStringWrapper

var my_map_test_record MyMapTestRecord

var net_device NetDevice

var orderline Orderline

var orders Orders

var page_views PageViews

var payroll_bonus PayrollBonus

var payroll_employee PayrollEmployee

var payroll_employee_location PayrollEmployeeLocation

var pizza_orders PizzaOrders

var pizza_orders_cancelled PizzaOrdersCancelled

var pizza_orders_completed PizzaOrdersCompleted

var product Product

var purchase Purchase

var ratings Ratings

var shoe Shoe

var shoe_clickstream ShoeClickstream

var shoe_customer ShoeCustomer

var shoe_order ShoeOrder

var siem_logs SiemLogs

var source Source

var stock_trades StockTrades

var stores Stores

var syslog_logs SyslogLogs

var transactions Transactions

var user User

var users Users

var users_array_map UsersArrayMap

//gocyclo:ignore
func GetType(templateType string) interface{} {

	switch templateType {

	case "address":
		return &address

	case "array_orderline_wrapper":
		return &array_orderline_wrapper

	case "array_string_wrapper":
		return &array_string_wrapper

	case "bytes":
		return &bytes

	case "campaign_finance":
		return &campaign_finance

	case "click_stream":
		return &click_stream

	case "click_stream_codes":
		return &click_stream_codes

	case "click_stream_users":
		return &click_stream_users

	case "clickstreamusers":
		return &clickstreamusers

	case "codes":
		return &codes

	case "credit_cards":
		return &credit_cards

	case "csv_product":
		return &csv_product

	case "csv_user":
		return &csv_user

	case "destination":
		return &destination

	case "device_information":
		return &device_information

	case "deviceinformation":
		return &deviceinformation

	case "fleet_mgmt_description":
		return &fleet_mgmt_description

	case "fleet_mgmt_location":
		return &fleet_mgmt_location

	case "fleet_mgmt_sensors":
		return &fleet_mgmt_sensors

	case "fleetmgmtdescription":
		return &fleetmgmtdescription

	case "fleetmgmtlocation":
		return &fleetmgmtlocation

	case "fleetmgmtsensors":
		return &fleetmgmtsensors

	case "gaming_games":
		return &gaming_games

	case "gaming_player_activity":
		return &gaming_player_activity

	case "gaming_players":
		return &gaming_players

	case "gaminggames":
		return &gaminggames

	case "gamingplayeractivity":
		return &gamingplayeractivity

	case "gamingplayers":
		return &gamingplayers

	case "insurance_customer_activity":
		return &insurance_customer_activity

	case "insurance_customers":
		return &insurance_customers

	case "insurance_offers":
		return &insurance_offers

	case "inventory":
		return &inventory

	case "location":
		return &location

	case "map_bool_wrapper":
		return &map_bool_wrapper

	case "map_bytes_wrapper":
		return &map_bytes_wrapper

	case "map_float_wrapper":
		return &map_float_wrapper

	case "map_int_wrapper":
		return &map_int_wrapper

	case "map_long_wrapper":
		return &map_long_wrapper

	case "map_string_wrapper":
		return &map_string_wrapper

	case "my_map_test_record":
		return &my_map_test_record

	case "net_device":
		return &net_device

	case "orderline":
		return &orderline

	case "orders":
		return &orders

	case "page_views":
		return &page_views

	case "payroll_bonus":
		return &payroll_bonus

	case "payroll_employee":
		return &payroll_employee

	case "payroll_employee_location":
		return &payroll_employee_location

	case "pizza_orders":
		return &pizza_orders

	case "pizza_orders_cancelled":
		return &pizza_orders_cancelled

	case "pizza_orders_completed":
		return &pizza_orders_completed

	case "product":
		return &product

	case "purchase":
		return &purchase

	case "ratings":
		return &ratings

	case "shoe":
		return &shoe

	case "shoe_clickstream":
		return &shoe_clickstream

	case "shoe_customer":
		return &shoe_customer

	case "shoe_order":
		return &shoe_order

	case "siem_logs":
		return &siem_logs

	case "source":
		return &source

	case "stock_trades":
		return &stock_trades

	case "stores":
		return &stores

	case "syslog_logs":
		return &syslog_logs

	case "transactions":
		return &transactions

	case "user":
		return &user

	case "users":
		return &users

	case "users_array_map":
		return &users_array_map
	}
	return nil
}
