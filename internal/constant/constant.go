package constant

import "time"

// Access Type
var RegisteredUser int = 1

// Token Time
var TokenExpiryTime time.Duration = 60 // in minutes

// JSON Message
var SuccessGeneralMessage string = "Success!"
var ErrorGeneralMessage string = "Error!"
var ErrorParsingMessage string = "Error parsing data!"
var ErrorUsernameTakenMessage string = "This username is already taken!"
var ErrorUsernamePasswordMessage string = "You have entered an invalid username or password!"
var ErrorUnathorizedMessage string = "Not Authorized!"
