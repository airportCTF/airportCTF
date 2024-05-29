local jwt = require "resty.jwt"
local function decode_jwt(encoded_jwt)
    local jwt_obj = jwt:load_jwt(encoded_jwt)
    if not jwt_obj.valid then
        return nil
    end
    if not jwt_obj.payload.login then
        return nil
    end
    ngx.log(ngx.ERR, jwt_obj.payload.exp)
    ngx.log(ngx.ERR, os.time())
    return jwt_obj.payload.login
end
local jwt_cookie = ngx.var.cookie_session
local jwt_payload = decode_jwt(jwt_cookie)
if jwt_payload then
    ngx.req.set_header("X-Data-Login", jwt_payload)
    ngx.req.set_header("X-Data-Auth", "True")
else
    ngx.req.set_header("X-Data-Auth", "False")
end
if jwt_cookie and not jwt_payload then
    ngx.redirect("/api/auth/v1/logout", ngx.HTTP_MOVED_TEMPORARILY)
    return
end
ngx.exec("@ticket");
