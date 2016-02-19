package gofighter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func get_json_from_url(method string, url string, api_key string, postdata * RawOrder, unmarshaltarget interface{})  error {

	bodybytes, _ := json.Marshal(postdata)				// Don't dereference postdata as it might be nil
	body := bytes.NewBufferString(string(bodybytes))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return fmt.Errorf("error calling http.NewRequest: %s", err)
	}
	req.Header.Add("X-Starfighter-Authorization", api_key)
	api_cookie_text := fmt.Sprintf("api_key=%s", api_key)
	req.Header.Add("Cookie", api_cookie_text)

	if method == "POST" && postdata != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error calling client.Do: %s", err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("error calling ioutil.ReadAll: %s", err)
	}

	err = json.Unmarshal(b, unmarshaltarget)
	if err != nil {
		return fmt.Errorf("error calling json.Unmarshal: %s", err)
	}

	return nil
}

func CheckAPI(info TradingInfo)  (Heartbeat, error) {
	var ret Heartbeat
	url := info.BaseURL + "/heartbeat"
	err := get_json_from_url("GET", url, info.ApiKey, nil, &ret)
	return ret, err
}

func CheckVenue(info TradingInfo)  (VenueHeartbeat, error) {
	var ret VenueHeartbeat
	url := info.BaseURL + "/venues/" + info.Venue + "/heartbeat"
	err := get_json_from_url("GET", url, info.ApiKey, nil, &ret)
	return ret, err
}

func GetVenueList(info TradingInfo)  (VenueList, error) {
	var ret VenueList
	url := info.BaseURL + "/venues"
	err := get_json_from_url("GET", url, info.ApiKey, nil, &ret)
	return ret, err
}

func GetStockList(info TradingInfo)  (StockList, error) {
	var ret StockList
	url := info.BaseURL + "/venues/" + info.Venue + "/stocks"
	err := get_json_from_url("GET", url, info.ApiKey, nil, &ret)
	return ret, err
}

func GetOrderbook(info TradingInfo)  (OrderBook, error) {
	var ret OrderBook
	url := info.BaseURL + "/venues/" + info.Venue + "/stocks/" + info.Symbol
	err := get_json_from_url("GET", url, info.ApiKey, nil, &ret)
	return ret, err
}

func GetQuote(info TradingInfo)  (Quote, error) {
	var ret Quote
	url := info.BaseURL + "/venues/" + info.Venue + "/stocks/" + info.Symbol + "/quote"
	err := get_json_from_url("GET", url, info.ApiKey, nil, &ret)
	return ret, err
}

func GetStatus(info TradingInfo, id int)  (Order, error) {
	var ret Order
	url := info.BaseURL + "/venues/" + info.Venue + "/stocks/" + info.Symbol + "/orders/" + strconv.Itoa(id)
	err := get_json_from_url("GET", url, info.ApiKey, nil, &ret)
	return ret, err
}

func Cancel(info TradingInfo, id int)  (Order, error) {
	var ret Order
	url := info.BaseURL + "/venues/" + info.Venue + "/stocks/" + info.Symbol + "/orders/" + strconv.Itoa(id)
	err := get_json_from_url("DELETE", url, info.ApiKey, nil, &ret)
	return ret, err
}

func StatusAllOrders(info TradingInfo)  (OrderList, error) {
	var ret OrderList
	url := info.BaseURL + "/venues/" + info.Venue + "/accounts/" + info.Account + "/orders"
	err := get_json_from_url("GET", url, info.ApiKey, nil, &ret)
	return ret, err
}

func StatusAllOrdersOneStock(info TradingInfo)  (OrderList, error) {
	var ret OrderList
	url := info.BaseURL + "/venues/" + info.Venue + "/accounts/" + info.Account + "/stocks/" + info.Symbol + "/orders"
	err := get_json_from_url("GET", url, info.ApiKey, nil, &ret)
	return ret, err
}

func Execute(info TradingInfo, orderinfo ShortOrderer, result_chan chan Order)  (Order, error) {

	// Accepts either RawOrder or ShortOrder as the second argument type.
	// If it's a RawOrder and the account, venue or symbol differ from that
	// in the TradingInfo struct, the TradingInfo struct prevails.

	shortorder := orderinfo.MakeShortOrder()
	postdata := RawOrder{
		Account:    info.Account,
		Venue:      info.Venue,
		Symbol:     info.Symbol,
		Direction:  shortorder.Direction,
		OrderType:  shortorder.OrderType,
		Qty:        shortorder.Qty,
		Price:      shortorder.Price,
	}

	var ret Order
	url := info.BaseURL + "/venues/" + info.Venue + "/stocks/" + info.Symbol + "/orders"
	err := get_json_from_url("POST", url, info.ApiKey, &postdata, &ret)
	if err != nil {
		return ret, err
	}
	if result_chan != nil {
		result_chan <- ret
	}
	return ret, nil
}
