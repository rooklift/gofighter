package gofighter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func get_json_from_url(protocol string, url string, api_key string, postdata * RawOrder, unmarshaltarget interface{})  error {

	bodybytes, _ := json.Marshal(postdata)				// Don't dereference postdata as it might be nil
	body := bytes.NewBufferString(string(bodybytes))

	client := &http.Client{}
	req, err := http.NewRequest(protocol, url, body)
	if err != nil {
		return fmt.Errorf("error calling http.NewRequest: %s", err)
	}
	req.Header.Add("X-Starfighter-Authorization", api_key)
	api_cookie_text := fmt.Sprintf("api_key=%s", api_key)
	req.Header.Add("Cookie", api_cookie_text)

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

func CheckAPI(base_url string, api_key string)  (Heartbeat, error) {
	var ret Heartbeat
	url := base_url + "/heartbeat"
	err := get_json_from_url("GET", url, api_key, nil, &ret)
	return ret, err
}

func CheckVenue(base_url string, api_key string, venue string)  (VenueHeartbeat, error) {
	var ret VenueHeartbeat
	url := base_url + "/venues/" + venue + "/heartbeat"
	err := get_json_from_url("GET", url, api_key, nil, &ret)
	return ret, err
}

func GetVenueList(base_url string, api_key string)  (VenueList, error) {
	var ret VenueList
	url := base_url + "/venues"
	err := get_json_from_url("GET", url, api_key, nil, &ret)
	return ret, err
}

func GetStockList(base_url string, api_key string, venue string)  (StockList, error) {
	var ret StockList
	url := base_url + "/venues/" + venue + "/stocks"
	err := get_json_from_url("GET", url, api_key, nil, &ret)
	return ret, err
}

func GetOrderbook(base_url string, api_key string, venue string, symbol string)  (OrderBook, error) {
	var ret OrderBook
	url := base_url + "/venues/" + venue + "/stocks/" + symbol
	err := get_json_from_url("GET", url, api_key, nil, &ret)
	return ret, err
}

func GetQuote(base_url string, api_key string, venue string, symbol string)  (Quote, error) {
	var ret Quote
	url := base_url + "/venues/" + venue + "/stocks/" + symbol + "/quote"
	err := get_json_from_url("GET", url, api_key, nil, &ret)
	return ret, err
}

func GetStatus(base_url string, api_key string, venue string, symbol string, id int)  (Order, error) {
	var ret Order
	url := base_url + "/venues/" + venue + "/stocks/" + symbol + "/orders/" + strconv.Itoa(id)
	err := get_json_from_url("GET", url, api_key, nil, &ret)
	return ret, err
}

func Cancel(base_url string, api_key string, venue string, symbol string, id int)  (Order, error) {
	var ret Order
	url := base_url + "/venues/" + venue + "/stocks/" + symbol + "/orders/" + strconv.Itoa(id)
	err := get_json_from_url("DELETE", url, api_key, nil, &ret)
	return ret, err
}

func StatusAllOrders(base_url string, api_key string, venue string, account string)  (OrderList, error) {
	var ret OrderList
	url := base_url + "/venues/" + venue + "/accounts/" + account + "/orders"
	err := get_json_from_url("GET", url, api_key, nil, &ret)
	return ret, err
}

func StatusAllOrdersOneStock(base_url string, api_key string, venue string, account string, symbol string)  (OrderList, error) {
	var ret OrderList
	url := base_url + "/venues/" + venue + "/accounts/" + account + "/stocks/" + symbol + "/orders"
	err := get_json_from_url("GET", url, api_key, nil, &ret)
	return ret, err
}

func Execute(base_url string, api_key string, postdata RawOrder, result_chan chan Order)  (Order, error) {
	venue := postdata.Venue
	symbol := postdata.Symbol

	var ret Order
	url := base_url + "/venues/" + venue + "/stocks/" + symbol + "/orders"
	err := get_json_from_url("POST", url, api_key, &postdata, &ret)
	if err != nil {
		return ret, err
	}
	if result_chan != nil {
		result_chan <- ret
	}
	return ret, nil
}
