package main

import "bnqkl/chain-cms/app"

// @title			           Chain Cms Doc
// @version		               1.0
// @description	               This is a complex Chain Cms.

// @securityDefinitions.apikey Bearer
// @in                         header
// @name                       Authorization
// @bearerFormat               Bearer
// @description                Add prefix "Bearer ", such as "Bearer token".

// @license.name	           Apache 2.0
// @license.url	               http://www.apache.org/licenses/LICENSE-2.0.html

// @host	                   localhost:3000
func main() {
	app.Run()
}
