package constants

import "time"

const ServerPort = 80
const GracefulTimeout = time.Second * 15
const WebInterfaceURI = "/web"
const DatabaseAddress = "mongodb://localhost:27017"
const DatabaseName = "demo-api-users"
const DbAccountsSchema = "accounts"
