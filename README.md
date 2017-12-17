# scURL
cURL replica providing session. User can specify headers/cookies and store them. Later requests will use that data
stored in local "session"

### Usage

1. Start a new session using `-n` flag
2. Add a cookie using `-ac` followed by `cookie-name=cookie-value`
3. Remove cookie using `-rc` followed by `cookie-name`
4. Add header using `-ah` followed by `header-name=header-value`
5. Remove header using `-rh` followed by `header-name`
6. Send a HTTP request using `-x` followed by request type (ie. POST, GET) `-b` followed by JSON body 