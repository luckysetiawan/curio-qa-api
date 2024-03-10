// Package constants provide constant values that should not change during
// program execution.
package constant

import "time"

// Access Types.
var RegisteredUser int = 1

// Token Expiration TIme.
var TokenExpirationTime time.Duration = 60 // in minutes

// JSON Message.
var SuccessGeneralMessage string = "Success!"
var ErrorGeneralMessage string = "Error!"
var ErrorParsingMessage string = "Error parsing data!"
var ErrorCacheMessage string = "Error caching data!"
var ErrorUsernameTakenMessage string = "This username is already taken!"
var ErrorUsernamePasswordMessage string = "You have entered an invalid username or password!"
var ErrorUnathorizedMessage string = "Not Authorized!"
