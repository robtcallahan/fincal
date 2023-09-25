/*
Copyright © 2020 Rob Callahan <robtcallahan@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"encoding/json"
	cfg "fincal/pkg/config"
	"fincal/pkg/driver"
	"fincal/pkg/models"
	"fmt"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"fincal/pkg/handler"
	"fincal/pkg/plaid_auth"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Provides a web service function",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		server()
	},
}

var conn *driver.DB
var qHandler *handler.Query

func init() {
	rootCmd.AddCommand(serveCmd)

	// get and read config file
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	configFile := fmt.Sprintf("%s", viper.Get("CONFIG_FILE"))

	config, err = cfg.ReadConfig(configFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// get the hostname to determine the proper DB_HOST env var to use
	hostname, err := os.Hostname()
	log.Printf("hostname: %s\n", hostname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	re := regexp.MustCompile(`.*MacBook.*`)
	dbHost := fmt.Sprintf("%s", viper.Get("DB_HOST_REMOTE"))
	if re.Match([]byte(hostname)) {
		dbHost = fmt.Sprintf("%s", viper.Get("DB_HOST_LOCAL"))
	}

	conn, err = driver.ConnectSQL(&driver.ConnectParams{
		DBType: driver.DBType(config.DBType),
		Host:   dbHost,
		Port:   config.DBPort,
		DBName: config.DBName,
		User:   config.DBUsername,
		Pass:   config.DBPassword,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	qHandler = handler.NewQueryHandler(conn)

	client = getBankingClient()
}

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func server() {
	router := mux.NewRouter()

	router.HandleFunc("/api/ok", client.ok).Methods("GET")

	router.HandleFunc("/api/auth/login", client.login).Methods("POST", "OPTIONS")

	router.HandleFunc("/api/create_link_token", client.createLinkToken).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/exchange_public_token", client.exchangePublicToken).Methods("POST")
	router.HandleFunc("/api/get_accounts", client.getAccounts).Methods("GET")
	router.HandleFunc("/api/get_balance", client.getBalance).Methods("GET")

	router.HandleFunc("/api/get_register", client.getRegister).Methods("GET")

	router.HandleFunc("/api/get_merchants", client.getMerchants).Methods("GET")
	router.HandleFunc("/api/update_merchant", client.updateMerchant).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/delete_merchant", client.deleteMerchant).Methods("POST", "OPTIONS")

	router.HandleFunc("/api/get_categories", client.getCategories).Methods("GET")
	router.HandleFunc("/api/get_categories_for_select", client.getCategoriesForSelect).Methods("GET")
	router.HandleFunc("/api/create_category", client.createCategory).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/update_category", client.updateCategory).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/delete_category", client.deleteCategory).Methods("POST", "OPTIONS")

	router.HandleFunc("/api/get_transactions", client.getTransactions).Methods("GET")
	router.HandleFunc("/api/is_account_connected", isAccountConnected).Methods("GET")

	router.Use(mux.CORSMethodMiddleware(router))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8000", "http://localhost:5173", "http://localhost", "http://www.wetdogsoftware.com"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "OPTIONS", "GET", "DELETE", "PUT"},
		AllowedHeaders:   []string{"Content-Type", "Origin", "Accept", "token"},
		MaxAge:           86400,
	})
	myHandler := c.Handler(router)

	log.Println("Server will start at http://localhost:9000/")

	//log.Fatal(http.ListenAndServeTLS(":9000", config.CertFile, config.KeyFile, r))
	log.Fatal(http.ListenAndServe(":9000", myHandler))
}

func (c *Client) ok(w http.ResponseWriter, r *http.Request) {
	type okStruct struct {
		OK bool `json:"ok"`
	}
	ok := okStruct{
		OK: true,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(ok)
}

func (c *Client) login(w http.ResponseWriter, r *http.Request) {
	//defer r.Body.Close()
	//b, err := io.ReadAll(r.Body)
	//if err != nil {
	//	log.Println(err.Error())
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	http.Error(w, "Username or password is incorrect", http.StatusUnauthorized)
	return

}

func (c *Client) getRegister(w http.ResponseWriter, r *http.Request) {
	//sheetsProvider, err := sheets_provider.New(options.SpreadsheetID, config)
	//if err != nil {
	//	log.Println(err.Error())
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//sheetsService := sheets_service.New(sheetsProvider)
	//if err != nil {
	//	log.Println(err.Error())
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//err = sheetsService.NewRegisterSheet(config)
	//if err != nil {
	//	log.Println(err.Error())
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//googleSheet, err := sheetsService.ReadRegisterSheet()
	//if err != nil {
	//	log.Println(err.Error())
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//j, err := json.Marshal(googleSheet.Register)
	//if err != nil {
	//	log.Println(err.Error())
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//err = os.WriteFile("./json/register_sheet.json", j, 0644)
	//if err != nil {
	//	log.Println(err.Error())
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	j, err := os.ReadFile("./json/register_sheet.json")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//jString := string(j)

	w.Header().Set("Content-Type", "application/json")
	//_ = json.NewEncoder(w).Encode(googleSheet.Register)
	_, err = w.Write(j)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Client) getMerchants(w http.ResponseWriter, r *http.Request) {
	merchants := qHandler.GetMerchants()
	categories := qHandler.GetCategories()

	var merCols []models.MerchantColumn
	for _, m := range merchants {
		merCol := models.MerchantColumn{
			ID:       m.ID,
			Name:     m.Name,
			BankName: m.BankName,
		}
		for _, c := range categories {
			if c.ID == m.ColumnID {
				merCol.ColumnID = c.ID
				merCol.ColumnName = c.Name
				merCol.ColumnColor = c.Color
			}
		}
		merCols = append(merCols, merCol)
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(merCols)
}

func (c *Client) updateMerchant(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type updateStruct struct {
		ID     int    `json:"id"`
		Column string `json:"column"`
		Value  string `json:"value"`
	}
	var up updateStruct
	err = json.Unmarshal(b, &up)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// if column is column_name, then we need to get the column_id first

	merch := models.Merchant{ID: up.ID}
	qHandler.UpdateMerchant(&merch, up.Column, up.Value)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(up)
}

func (c *Client) deleteMerchant(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("json: %s", string(b))
	type deleteStruct struct {
		ID int `json:"id"`
	}
	var ds deleteStruct
	err = json.Unmarshal(b, &ds)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	merch := models.Merchant{ID: ds.ID}
	qHandler.DeleteMerchant(&merch)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(ds)
}

func (c *Client) getCategories(w http.ResponseWriter, r *http.Request) {
	columns := qHandler.GetCategories()

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(columns)
}

func (c *Client) getCategoriesForSelect(w http.ResponseWriter, r *http.Request) {
	columns := qHandler.GetCategoriesForSelect()

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(columns)
}

func (c *Client) createCategory(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var category models.Category
	err = json.Unmarshal(b, &category)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	category = *qHandler.CreateCategory(&category)
	log.Printf("returning category ID: %d", category.ID)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(category)
}

func (c *Client) updateCategory(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type updateStruct struct {
		ID     int         `json:"id"`
		Column string      `json:"column"`
		Value  interface{} `json:"value"`
	}
	var up updateStruct
	err = json.Unmarshal(b, &up)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	category := models.Category{ID: up.ID}
	qHandler.UpdateCategory(&category, up.Column, up.Value)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(up)
}

func (c *Client) deleteCategory(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("json: %s", string(b))
	type deleteStruct struct {
		ID int `json:"id"`
	}
	var ds deleteStruct
	err = json.Unmarshal(b, &ds)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	category := models.Category{ID: ds.ID}
	qHandler.DeleteCategory(&category)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(ds)
}

func (c *Client) createLinkToken(w http.ResponseWriter, r *http.Request) {
	log.Println("createLinkToken()")

	linkToken, err := plaid_auth.GetLinkToken(c.BankClient)
	if err != nil {
		log.Println(err.Error())
		fmt.Print(err)
	}
	w.Header().Set("Content-Type", "application/json")
	s := LinkToken{LinkToken: linkToken}
	json.NewEncoder(w).Encode(s)
}

func (c *Client) exchangePublicToken(w http.ResponseWriter, r *http.Request) {
	log.Println("exchangePublicToken()")

	var publicToken PublicToken

	err := json.NewDecoder(r.Body).Decode(&publicToken)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accessToken, err := plaid_auth.ExchangePublicToken(c.BankClient, publicToken.PublicToken, ctx)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: me thinks that bank.Source is needed for the token file name
	err = os.WriteFile(config.PlaidTokensDir+"/"+"AccessToken.txt", []byte(accessToken+"\n"), 0644)
	if err != nil {
		log.Printf("could not write access token: %s", err.Error())
	}

	session, err := store.Get(r, "register")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["AccessToken"] = accessToken
	err = sessions.Save(r, w)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Client) getAccounts(w http.ResponseWriter, r *http.Request) {
	log.Println("getAccounts()")

	session, err := store.Get(r, "register")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	accessToken := fmt.Sprintf("%s", session.Values["AccessToken"])
	if accessToken == "" {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	balances, err := c.BankClient.GetAccounts(accessToken, ctx)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	_ = json.NewEncoder(w).Encode(balances)
}

func (c *Client) getBalance(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "register")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	accessToken := fmt.Sprintf("%s", session.Values["AccessToken"])
	if accessToken == "" {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	balance, err := c.BankClient.GetBalance(accessToken, "fidelity", ctx)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	_ = json.NewEncoder(w).Encode(balance)
}

func (c *Client) getTransactions(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "register")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	accessToken := fmt.Sprintf("%s", session.Values["AccessToken"])
	if accessToken == "" {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// start from 2 weeks ago
	startDate := weeksAgo(2)
	endDate := today()
	transactions, err := client.BankClient.GetTransactions(options.BankIDs, startDate, endDate)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(transactions)
}

func isAccountConnected(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "register")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	accessToken := fmt.Sprintf("%s", session.Values["AccessToken"])
	type ConnectedResponse struct {
		Status bool `json:"status"`
	}
	if accessToken != "" {
		//_ = json.NewEncoder(w).Encode("{status: true}")
		resp := ConnectedResponse{Status: true}
		_ = json.NewEncoder(w).Encode(resp)
	} else {
		resp := ConnectedResponse{Status: false}
		_ = json.NewEncoder(w).Encode(resp)
	}
}
